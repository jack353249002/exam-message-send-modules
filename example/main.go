package main

import (
	_ "github.com/SupenBysz/gf-admin-community"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/jack353249002/exam-message-send-modules/example/internal/boot"

	_ "github.com/jack353249002/exam-message-send-modules/example/internal/consts"
	_ "github.com/jack353249002/exam-message-send-modules/internal/logic"
)

func main() {
	boot.Main.Run(gctx.New())
}
