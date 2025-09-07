package cmd

import (
	"context"
	"service/app/gateway-admin/internal/controller/admin"
	"service/app/gateway-admin/internal/controller/file"
	"service/app/gateway-admin/internal/controller/goods"
	"service/utility/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/backend", func(group *ghttp.RouterGroup) {
					group.Bind(
						admin.NewV1(),
					)
				})
				group.Group("/backend", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.JWTAuth)
					group.Bind(
						file.NewV1(),
						goods.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
