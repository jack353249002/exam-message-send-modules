package i_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/jack353249002/exam-message-send-modules/api/co_message_send_v1"
	"github.com/jack353249002/exam-message-send-modules/co_model"
)

type ISend[
	TIRes co_model.ISendRes,
] interface {
	// 添加发送规则
	CreateSend(ctx context.Context, req *co_message_send_v1.CreateSendReq) (api_v1.BoolRes, error)
	// 设置发送消息
	SetSendInfoAction(ctx context.Context, req *co_message_send_v1.SetSendActionReq) (api_v1.BoolRes, error)
}
