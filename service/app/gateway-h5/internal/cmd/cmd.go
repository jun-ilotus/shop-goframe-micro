package cmd

import (
	"context"
	"service/app/gateway-h5/internal/controller/interaction"
	"service/utility/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"service/app/gateway-h5/internal/controller/user"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http gateway-h5 server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/frontend", func(group *ghttp.RouterGroup) {
					group.Bind(
						user.NewV1(),
					)
				})
				// 需要JWT验证的路由
				group.Group("/frontend", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTAuth)
					group.Bind(
						// 需要认证的接口
						interaction.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
