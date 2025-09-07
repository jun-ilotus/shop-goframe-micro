package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gctx"

	"service/app/goods/internal/cmd"
)

func main() {
	var ctx = gctx.New()
	conf, err := g.Cfg().Get(ctx, "etcd.address")
	if err != nil {
		panic(err)
	}

	var address = conf.String()
	grpcx.Resolver.Register(etcd.New(address))

	cmd.Main.Run(gctx.GetInitCtx())
}
