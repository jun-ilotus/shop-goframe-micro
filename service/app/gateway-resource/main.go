package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	_ "service/app/gateway-resource/internal/packed"
	"service/utility/middleware"

	"github.com/gogf/gf/v2/os/gctx"

	"service/app/gateway-resource/internal/cmd"
)

func main() {
	var ctx = gctx.New()
	conf, err := g.Cfg().Get(ctx, "etcd.address")
	if err != nil {
		panic(err)
	}

	var address = conf.String()
	grpcx.Resolver.Register(etcd.New(address))

	// 创建 HTTP 服务
	s := g.Server()

	// 设置 CORS 头
	s.Use(middleware.MiddlewareCORS)
	cmd.Main.Run(ctx)
}
