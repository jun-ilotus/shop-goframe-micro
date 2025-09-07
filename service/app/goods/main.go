package main

import (
	_ "service/app/goods/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"service/app/goods/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
