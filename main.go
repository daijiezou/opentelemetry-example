package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"opentelemetry-example/grpc/api"
	"time"
)

//var tracer oteltrace.Tracer
var tracer = otel.Tracer("gin-server-new")

const (
	service     = "gin-server"
	environment = "production"
)

func main() {
	tp, err := JaegerProvider("http://127.0.0.1:14268/api/traces")
	//tp, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	r := gin.New()
	r.Use(otelgin.Middleware("my-server"))
	r.GET("/users/:id", helloHandler)

	if err := r.Run("localhost:9999"); err != nil {
		log.Print(err)
	}
}

func helloHandler(c *gin.Context) {
	id := c.Param("id")
	/*	_, span := tracer.Start(c.Request.Context(), "helloHandler", oteltrace.WithAttributes(attribute.String("id", id)))
		defer span.End()
		span.SetAttributes(attribute.KeyValue{
			Key:   "name",
			Value: attribute.StringValue("daijun"),
		})
		span.SetAttributes(attribute.KeyValue{
			Key:   "current time",
			Value: attribute.StringValue(time.Now().Format("2006-01-02 15:04:05")),
		})*/

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
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
	}
	log.Printf("Response from server: %s", response.Reply)
	span := oteltrace.SpanFromContext(ctx)
	//traceId := span.SpanContext().TraceID().String()
	c.Header("trace_id", span.SpanContext().TraceID().String())
	c.Header("span-id", span.SpanContext().SpanID().String())
	c.JSON(200, gin.H{
		"reply":   response.Reply,
		"traceId": span.SpanContext().TraceID().String(),
	})
}

func JaegerProvider(url string) (*sdktrace.TracerProvider, error) {
	fmt.Println("init traceProvider")
	// 创建jaeger provider
	// 可以直接连collector也可以连agent
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
