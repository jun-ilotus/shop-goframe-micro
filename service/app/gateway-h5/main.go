package main

import (
	_ "service/app/gateway-h5/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"service/app/gateway-h5/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
