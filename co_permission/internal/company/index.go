package company

import (
	"github.com/jack353249002/exam-message-send-modules/co_interface"
)

type company struct {
	PermissionType func(modules co_interface.IConfig) *permissionType[co_interface.IConfig]
}

var Company = company{
	PermissionType: PermissionType,
}
