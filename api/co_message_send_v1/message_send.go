package co_message_send_v1

import (
	"github.com/jack353249002/exam-message-send-modules/co_model"
	"github.com/kysion/base-library/base_model"
)

type SetMessageInfoReq struct {
	co_model.CoMessage
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}
type GetMessageListReq struct {
	base_model.SearchParams
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}
type CreateSendReq struct {
	co_model.CoSend
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}
type SetSendActionReq struct {
	Id      string   `json:"id" v:"required#id不能为空"  dc:"id"`
	Status  int8     `json:"status" v:"required#状态不能为空"  dc:"status"`
	Include []string `json:"include" dc:"需要附加数据的返回值字段集，如果没有填写，默认不附加数据"`
}
