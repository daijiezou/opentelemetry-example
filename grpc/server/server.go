// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"opentelemetry-example/grpc/api/errors"
	geminiuserauth "opentelemetry-example/grpc/api/userauth"
	"opentelemetry-example/provider"
	"time"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/example/api"
	"go.opentelemetry.io/otel"

	"google.golang.org/grpc"
)

const (
	port = ":7777"
)

var tracer = otel.Tracer("grpc-example")

// server is used to implement api.HelloServiceServer.
type server struct {
	api.HelloServiceServer
}

// SayHello implements api.HelloServiceServer.
func (s *server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	log.Printf("Received: %v\n", in.GetGreeting())
	//bag := baggage.FromContext(ctx)
	//mem := bag.Member("gemini").Value()
	//log.Printf("gemini:%+v\n", mem)
	s.workHard(ctx)
	time.Sleep(50 * time.Millisecond)

	return &api.HelloResponse{Reply: "Hello " + in.Greeting}, nil
}

func ReturnError() error {
	gErr := errors.New(1, "db error", "say hello error")
	return gErr.WithMetadata(map[string]string{
		"name":       "daijun",
		"age":        "18",
		"sweetheart": "xuxueqin",
	}).WithCause(fmt.Errorf("cause error"))

}

func (s *server) workHard(ctx context.Context) {
	//_, span := tracer.Start(ctx, "workHard",
	//	trace.WithAttributes(attribute.String("extra.key", "extra.value")))
	//defer span.End()
	//
	//traceId := span.SpanContext().TraceID().String()
	span := trace.SpanFromContext(ctx)
	traceId := span.SpanContext().TraceID().String()
	fmt.Println("traceId:", traceId)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer func() { _ = conn.Close() }()
	cli := geminiuserauth.NewGeminiAuthClient(conn)

	response, err := cli.UserLogin(ctx, &geminiuserauth.UserLoginReq{
		Username: "daijun",
		Password: "123",
	})
	fmt.Printf("response:%+v", response)
	time.Sleep(50 * time.Millisecond)
}

func main() {
	tp, err := provider.JaegerProvider("sayHello")
	//tp, err := provider.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	GetUserById()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	api.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func GetUserById() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer func() { _ = conn.Close() }()
	cli := geminiuserauth.NewGeminiAuthClient(conn)
	spanCtx, span := tracer.Start(context.Background(), "helloServiceMain")
	defer span.End()
	log.Printf("traceId: %v\n", span.SpanContext().TraceID().String())
	response, err := cli.GetUserInfoById(spanCtx, &geminiuserauth.GetUserInfoByIdRequest{UserId: 123})
	if err != nil {
		fmt.Printf("err:%+v", err)
		return
	}
	fmt.Printf("response:%+v", response)
}
