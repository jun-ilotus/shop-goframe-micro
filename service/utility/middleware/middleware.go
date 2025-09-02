package middleware

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

func GrpcClientTimeout(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")
	r.Response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	r.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// 处理预检请求
	if r.Method == http.MethodOptions {
		r.Response.WriteHeader(204) // No Content
		return
	}
	r.Middleware.Next()
}
