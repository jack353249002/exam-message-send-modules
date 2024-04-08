package internal

import (
	"github.com/jack353249002/exam-message-send-modules/co_interface"
	"github.com/jack353249002/exam-message-send-modules/co_model"
	"github.com/jack353249002/exam-message-send-modules/co_module"
	"github.com/kysion/base-library/utility/base_permission"
)

type Global struct {
	co_interface.IModules[*co_model.MessageRes, *co_model.SendRes]

	PermissionTree []base_permission.IPermission

	// FinancialPermissionTree 财务服务权限树 (可选)
	FinancialPermissionTree []base_permission.IPermission
}

var global *Global

func Modules() *Global {
	if global != nil {
		return global
	}

	global = &Global{
		IModules: co_module.NewModules[*co_model.MessageRes, *co_model.SendRes](nil, nil),
	}

	return global
}
