package main

import (
	_ "service/app/gateway-admin/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"service/app/gateway-admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
