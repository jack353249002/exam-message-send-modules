package internal

import (
	"github.com/jack353249002/exam-message-send-modules/internal/boot"
)

func init() {
	// 注册自定义参数校验规则
	boot.InitCustomRules()
}
