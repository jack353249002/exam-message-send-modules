package co_model

import (
	"github.com/jack353249002/exam-message-send-modules/base_interface"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_do"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
)

type CoMessage struct {
	OverrideDo base_interface.DoModel[co_do.Message] `json:"-"`
	Title      *string                               `json:"title" v:"required#请输入标题"  dc:"标题"`
	Body       *string                               `json:"body" v:"required#请输入内容"  dc:"内容"`
}

type MessageRes struct {
	*co_entity.Message
}

type MessageListRes base_model.CollectRes[MessageRes]

func (m *MessageRes) Data() *MessageRes {
	return m
}

type IMessageRes interface {
	Data() *MessageRes
}
