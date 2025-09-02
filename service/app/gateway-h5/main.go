package main

import (
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	_ "service/app/gateway-h5/internal/packed"
	"service/utility/middleware"

	"github.com/gogf/gf/v2/os/gctx"

	"service/app/gateway-h5/internal/cmd"
)

func main() {
	var ctx = gctx.New()
	conf, err := g.Cfg().Get(ctx, "etcd.address")
	if err != nil {
		panic(err)
	}

	var address = conf.String()
	grpcx.Resolver.Register(etcd.New(address))

	// 创建 http 服务
	s := g.Server()

	// 设置 CORS 头
	s.Use(middleware.MiddlewareCORS)
	cmd.Main.Run(gctx.GetInitCtx())
}
