package main

import (
	_ "service/app/search/internal/packed"
	"service/app/search/utility/elasticsearch"

	"github.com/gogf/gf/v2/os/gctx"

	"service/app/search/internal/cmd"
)

func main() {
	ctx := gctx.New()

	if err := elasticsearch.Init(ctx); err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
