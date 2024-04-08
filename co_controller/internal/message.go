package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/jack353249002/exam-message-send-modules/api/co_message_send_v1"
	"github.com/jack353249002/exam-message-send-modules/co_interface"
	"github.com/jack353249002/exam-message-send-modules/co_interface/i_controller"
	"github.com/jack353249002/exam-message-send-modules/co_model"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_dao"
	"github.com/kysion/base-library/base_model"
)

type MessageController[
	TIRes co_model.IMessageRes,
	TISRes co_model.ISendRes,
] struct {
	i_controller.IMessage[TIRes]
	modules co_interface.IModules[TIRes, TISRes]
	dao     co_dao.MessageDao
}

func Message[
	TIRes co_model.IMessageRes,
	TISRes co_model.ISendRes,
](modules co_interface.IModules[TIRes, TISRes]) i_controller.IMessage[TIRes] {
	return &MessageController[TIRes, TISRes]{
		modules: modules,
	}
}

// GetCompanyById 通过ID获取公司信息
func (c *MessageController[TIRes, TISRes]) CreateMessage(ctx context.Context, req *co_message_send_v1.SetMessageInfoReq) (api_v1.BoolRes, error) {
	ok, err := c.modules.Message().CreateMessage(ctx, *req.Title, *req.Body)
	return ok == true, err
}

// 消息列表
func (c *MessageController[TIRes, TISRes]) QueryMessageList(ctx context.Context, req *co_message_send_v1.GetMessageListReq) (*base_model.CollectRes[TIRes], error) {
	return c.modules.Message().QueryMessageList(ctx, &req.SearchParams)
}
