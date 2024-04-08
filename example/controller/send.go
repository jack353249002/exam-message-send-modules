package controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/jack353249002/exam-message-send-modules/api/co_v1"
	"github.com/jack353249002/exam-message-send-modules/co_controller"
	"github.com/jack353249002/exam-message-send-modules/co_interface"
	"github.com/jack353249002/exam-message-send-modules/co_interface/i_controller"
	"github.com/jack353249002/exam-message-send-modules/co_model"
)

type SendController[
	TIRes co_model.IMessageRes,
	TISRes co_model.ISendRes,
] struct {
	i_controller.ISend[TISRes]
	//modules co_interface.IModules[TIRes]
}

func Send[
	TIRes co_model.IMessageRes,
	TISRes co_model.ISendRes,
](modules co_interface.IModules[TIRes, TISRes]) *SendController[TIRes, TISRes] {
	return &SendController[TIRes, TISRes]{ISend: co_controller.Send(modules)}
}

func (c *SendController[TIRes, TISRes]) CreateSend(ctx context.Context, req *co_v1.CreateSendReq) (api_v1.BoolRes, error) {
	ok, err := c.ISend.CreateSend(ctx, &req.CreateSendReq)
	return ok, err
}

// 消息列表
func (c *SendController[TIRes, TISRes]) SetSendInfoAction(ctx context.Context, req *co_v1.SetSendActionReq) (api_v1.BoolRes, error) {
	ok, err := c.ISend.SetSendInfoAction(ctx, &req.SetSendActionReq)
	return ok, err
}
