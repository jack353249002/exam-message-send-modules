package financial

import "github.com/jack353249002/exam-message-send-modules/co_interface"

type financial struct {
	PermissionType func(modules co_interface.IConfig) *permissionType[co_interface.IConfig]
}

var Financial = financial{
	PermissionType: PermissionType,
}
