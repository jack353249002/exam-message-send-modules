package team

import (
	"github.com/jack353249002/exam-message-send-modules/co_interface"
)

type team struct {
	PermissionType func(modules co_interface.IConfig) *permissionType[co_interface.IConfig]
}

var Team = team{
	PermissionType: PermissionType,
}
