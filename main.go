package main

import (
	_ "GoTab/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"GoTab/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
