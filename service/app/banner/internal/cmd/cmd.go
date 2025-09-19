package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"google.golang.org/grpc"
	"service/app/banner/internal/controller/position_info"
	"service/app/banner/internal/controller/rotation_info"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "banner grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			position_info.Register(s)
			rotation_info.Register(s)
			s.Run()
			return nil
		},
	}
)
