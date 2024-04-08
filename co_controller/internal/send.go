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

type SendController[
	TIRes co_model.IMessageRes,
	TISRes co_model.ISendRes,
] struct {
	i_controller.ISend[TISRes]
	modules co_interface.IModules[TIRes, TISRes]
	dao     co_dao.MessageDao
}

func Send[
	TIRes co_model.IMessageRes,
	TISRes co_model.ISendRes,
](modules co_interface.IModules[TIRes, TISRes]) i_controller.ISend[TISRes] {
	return &SendController[TIRes, TISRes]{
		modules: modules,
	}
}

func (c *SendController[TIRes, TISRes]) CreateSend(ctx context.Context, req *co_message_send_v1.CreateSendReq) (api_v1.BoolRes, error) {
	ok, err := c.modules.Send().CreateSend(ctx, *req.Title, *req.MessageId, *req.SendServerId, *req.Receive)
	return ok == true, err
}

// 消息列表
func (c *SendController[TIRes, TISRes]) SetSendInfoAction(ctx context.Context, req *co_message_send_v1.SetSendActionReq) (api_v1.BoolRes, error) {
	ok, err := c.modules.Send().SetSendInfoAction(ctx, req.Id, req.Status)
	return ok == true, err
}

// 发送规则列表
func (c *SendController[TIRes, TISRes]) QuerySendInfoList(ctx context.Context, req *co_message_send_v1.GetSendInfoListReq) (*base_model.CollectRes[TISRes], error) {
	return c.modules.Send().QuerySendInfoList(ctx, &req.SearchParams)
}
