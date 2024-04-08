package boot

import (
	"github.com/jack353249002/exam-message-send-modules/utility/co_rules"
)

// InitCustomRules 注册自定义参数校验规则
func InitCustomRules() {
	// 注册资质自定义规则
	co_rules.RequiredLicense()
}
