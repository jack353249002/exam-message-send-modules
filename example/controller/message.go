package controller

import (
	"context"
	"fmt"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/jack353249002/exam-message-send-modules/api/co_v1"
	"github.com/jack353249002/exam-message-send-modules/co_controller"
	"github.com/jack353249002/exam-message-send-modules/co_interface"
	"github.com/jack353249002/exam-message-send-modules/co_interface/i_controller"
	"github.com/jack353249002/exam-message-send-modules/co_model"
	"github.com/kysion/base-library/utility/kconv"
)

type MessageController[
	TIRes co_model.IMessageRes,
] struct {
	i_controller.IMessage[TIRes]
	//modules co_interface.IModules[TIRes]
}

func Message[
	TIRes co_model.IMessageRes,
	TISRes co_model.ISendRes,
](modules co_interface.IModules[TIRes, TISRes]) *MessageController[TIRes] {
	return &MessageController[TIRes]{IMessage: co_controller.Message(modules)}
}

// CreateMessage 创建消息
func (c *MessageController[TIRes]) CreateMessage(ctx context.Context, req *co_v1.SetMessageInfoReq) (api_v1.BoolRes, error) {
	fmt.Println("show", c.IMessage)
	ret, err := c.IMessage.CreateMessage(ctx, &req.SetMessageInfoReq)
	return ret == true, err
}

// 消息列表
func (c *MessageController[TIRes]) QueryMessageList(ctx context.Context, req *co_v1.GetMessageListReq) (*co_model.MessageListRes, error) {
	ret, err := c.IMessage.QueryMessageList(ctx, &req.GetMessageListReq)
	return kconv.Struct(ret, &co_model.MessageListRes{}), err
}
