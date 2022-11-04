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
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net"
	"opentelemetry-example/grpc/api/userauth"
	"opentelemetry-example/provider"
	"os"
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

const (
	port = ":8888"
)

var tracer = otel.Tracer("grpc-example-2")
var zlog *otelzap.Logger

// server is used to implement api.HelloServiceServer.
type GeminiAuth struct {
	geminiuserauth.UnimplementedGeminiAuthServer
}

func (s *GeminiAuth) UserLogin(ctx context.Context, in *geminiuserauth.UserLoginReq) (*geminiuserauth.UserLoginReply, error) {
	InitLog()
	log.Printf("Received: %v\n", in)

	span := trace.SpanFromContext(ctx)
	traceId := span.SpanContext().TraceID().String()
	log.Printf("traceId: %v\n", traceId)

	time.Sleep(50 * time.Millisecond)
	//span := trace.SpanFromContext(ctx)
	//defer span.End()
	zlog.InfoContext(ctx, "UserLogin",
		zap.String("username", in.Username),
		zap.String("foo", "bar"))
	//_, span := tracer.Start(ctx, "UserLogin",
	//	trace.WithAttributes(attribute.String("extra.key", "extra.value")))
	//defer span.End()
	//traceId := span.SpanContext().TraceID().String()
	//fmt.Println("traceId:", traceId)
	return &geminiuserauth.UserLoginReply{
		Code: 0,
		Msg:  "",
		Data: &geminiuserauth.UserLoginReplyData{
			Token:        "",
			RefreshToken: "",
			UserId:       0,
			Username:     "daijun",
			DisplayName:  "代军",
			PhoneNum:     "",
			Email:        "",
		},
	}, nil
}

func (s *GeminiAuth) GetUserInfoById(ctx context.Context, in *geminiuserauth.GetUserInfoByIdRequest) (*geminiuserauth.GetUserInfoByIdReply, error) {
	InitLog()
	log.Printf("Received: %v\n", in)

	span := trace.SpanFromContext(ctx)
	traceId := span.SpanContext().TraceID().String()
	log.Printf("traceId: %v\n", traceId)
	time.Sleep(50 * time.Millisecond)
	return &geminiuserauth.GetUserInfoByIdReply{
		UserInfo: &geminiuserauth.UserInfo{
			UserId:   in.UserId,
			UserName: "chenjie",
		},
		IsSuccess: false,
		Err:       "",
	}, nil
}

func main() {
	tp, err := provider.JaegerProvider("userAuth")
	//tp, err := provider.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	geminiuserauth.RegisterGeminiAuthServer(s, &GeminiAuth{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func InitLog() {
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(DefaultEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)
	// 初始化 zap Logger
	zapLogger := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	zlog = otelzap.New(zapLogger, otelzap.WithTraceIDField(true), otelzap.WithMinLevel(zap.InfoLevel))
}

func DefaultEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     timeEncoder,                    // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格式为绝对路径
	}
}
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
