package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "service/app/user/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"service/app/user/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
