package i_controller

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/jack353249002/exam-message-send-modules/api/co_message_send_v1"
	"github.com/jack353249002/exam-message-send-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type IMessage[
	TIRes co_model.IMessageRes,
] interface {
	CreateMessage(ctx context.Context, req *co_message_send_v1.SetMessageInfoReq) (api_v1.BoolRes, error)
	QueryMessageList(ctx context.Context, info *co_message_send_v1.GetMessageListReq) (*base_model.CollectRes[TIRes], error)
}
