package co_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/jack353249002/exam-message-send-modules/api/co_message_send_v1"
	"github.com/jack353249002/exam-message-send-modules/co_model"
)

type CreateSendReq struct {
	g.Meta ` method:"post" summary:"设置发送" tags:"消息推送系统-推送"`
	co_message_send_v1.CreateSendReq
}
type SetSendActionReq struct {
	g.Meta ` method:"post" summary:"设置发送指令" tags:"消息推送系统-推送"`
	co_message_send_v1.SetSendActionReq
}
type SetMessageInfoReq struct {
	g.Meta ` method:"post" summary:"设置消息信息" tags:"消息推送系统-消息设置"`
	co_message_send_v1.SetMessageInfoReq
}
type GetMessageListReq struct {
	g.Meta ` method:"post" summary:"获取消息列表" tags:"消息推送系统-消息设置"`
	co_message_send_v1.GetMessageListReq
}
type GetSendInfoListReq struct {
	g.Meta ` method:"post" summary:"获取推送规则列表" tags:"消息推送系统-推送"`
	co_message_send_v1.GetSendInfoListReq
}
type SendReq struct {
	Title        string `json:"title"    description:"标题"  default:"" `
	MessageId    int    `json:"message_id"    description:"消息id"  default:""`
	SendServerId string `json:"send_server_id"    description:"发送服务器id"  default:""`
	Receive      string `json:"receive"    description:"接收账号"  default:""`
	co_model.CoSend
}
type AddSendReq struct {
	g.Meta `path:"/addSend" method:"post" summary:"设置发送" tags:"消息推送系统"`
	*SendReq
}
type UpdateSendReq struct {
	g.Meta `path:"/updateSend" method:"post" summary:"设置发送" tags:"消息推送系统"`
	*SendReq
	Id int64 `json:"id"    description:"id"  default:"" `
}
