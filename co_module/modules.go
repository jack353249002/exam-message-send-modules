package co_module

import (
	"github.com/jack353249002/exam-message-send-modules/co_interface"
	"github.com/jack353249002/exam-message-send-modules/co_model"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_dao"
	"github.com/jack353249002/exam-message-send-modules/internal/logic/co_message"
	"github.com/jack353249002/exam-message-send-modules/internal/logic/co_send"
)

type Modules[
	TMessageRes co_model.IMessageRes,
	TSendRes co_model.ISendRes,
] struct {
	co_interface.IConfig
	message co_interface.IMessage[TMessageRes]
	send    co_interface.ISend[TSendRes]
	xDao    *co_dao.MessageDao
}

func (m *Modules[TMessageRes, TSendRes]) Message() co_interface.IMessage[TMessageRes] {
	return m.message
}
func (m *Modules[TMessageRes, TSendRes]) Send() co_interface.ISend[TSendRes] {
	return m.send
}

// 这里需要添加绑定
func NewModules[
	ITMessageRes co_model.IMessageRes,
	ITSendRes co_model.ISendRes,
](conf *co_model.Config,
	xDao *co_dao.MessageDao) (response co_interface.IModules[ITMessageRes, ITSendRes]) {
	module := &Modules[ITMessageRes, ITSendRes]{
		xDao: xDao,
	}
	response = module
	module.message = co_message.NewMessage(response)
	module.send = co_send.NewSend(response)
	// 权限树追加权限
	//co_consts.PermissionTree = append(co_consts.PermissionTree, boot.InitPermission(response)...)
	//co_consts.FinancialPermissionTree = append(co_consts.FinancialPermissionTree, boot.InitFinancialPermission(response)...)

	return module
}
