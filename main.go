package main

import (
	"context"
	"fmt"
	"log"
	"opentelemetry-example/grpc/api"
	"opentelemetry-example/grpc/api/errors"
	"opentelemetry-example/provider"
	"time"

	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/attribute"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var tracer = otel.Tracer("gin-server")

func main() {
	/*
		如果没有初始化Provider，则 otelgin.Middleware 会从全局获取到一个noopTracer
		noopTracer is an implementation of Tracer that preforms no operations.
	*/
	//tp, err := JaegerProvider("http://127.0.0.1:14268/api/traces")
	tp, err := provider.StdOutTracer()
	if err != nil {
		log.Fatal(err)
		history2
	}

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	r := gin.New()
	r.Use(otelgin.Middleware("my-server"), TokenAuth())
	r.GET("/users/:id", helloHandler)

	if err := r.Run("localhost:9999"); err != nil {
		log.Print(err)
	}
}

func helloHandler(c *gin.Context) {
	id := c.Param("id")
	_, span := tracer.Start(c.Request.Context(), "helloHandler", oteltrace.WithAttributes(attribute.String("id", id)))
	defer span.End()
	span.SetAttributes(attribute.KeyValue{
		Key:   "name",
		Value: attribute.StringValue("daijun"),
	})
	span.SetAttributes(attribute.KeyValue{
		Key:   "current time",
		Value: attribute.StringValue(time.Now().Format("2006-01-02 15:04:05")),
	})
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer func() { _ = conn.Close() }()

	cli := api.NewHelloServiceClient(conn)
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "web-api-client-us-east-1",
		"user-id", "some-test-user-id",
		"id", id,
	)
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)
	response, err := cli.SayHello(ctx, &api.HelloRequest{Greeting: "World"})
	if err != nil {
		gerr := errors.FromError(err)
		fmt.Printf("gerr:%+v", gerr)
		c.JSON(200, gin.H{
			"err": gerr.Error(),
		})
		return
	}
	log.Printf("Response from server: %s", response.Reply)
	c.JSON(200, gin.H{
		"reply":   response.Reply,
		"traceId": span.SpanContext().TraceID().String(),
	})
}

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 模拟token校验
		span := oteltrace.SpanFromContext(c.Request.Context())
		c.Header("Trace-ID", span.SpanContext().TraceID().String())
	}
}
