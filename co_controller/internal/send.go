package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/jack353249002/exam-message-send-modules/api/co_message_send_v1"
	"github.com/jack353249002/exam-message-send-modules/co_interface"
	"github.com/jack353249002/exam-message-send-modules/co_interface/i_controller"
	"github.com/jack353249002/exam-message-send-modules/co_model"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_dao"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_do"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_entity"
	"github.com/jack353249002/exam-message-send-modules/internal/logic/send/sendunit"
	"math"
	"strconv"
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
	var coSend co_entity.Send
	var coSendList []co_entity.SendList
	sendId := req.Id
	co_dao.Send.Ctx(ctx).Where(co_do.Send{Id: req.Id}).Scan(&coSend)
	status := req.Status
	if coSend.Id == 0 {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, "发送消息规则不存在"), "", co_dao.Send.Table())
	} else {
		if status == 1 || status == -1 || status == -2 {
			key := strconv.Itoa(sendId)
			emailSend, sendUnitHave := sendunit.GetSendEmailUnitPool(key)
			switch status {
			case 1:
				if sendUnitHave {
					emailSend.Start()
				} else {
					co_dao.SendList.Ctx(ctx).Where("send_id=? AND status=0", sendId).Scan(&coSendList)
					num := math.Ceil(float64(len(coSendList)) / 2.0)
					sendUnit := sendunit.SendFactory(coSend.SendModel)
					if sendUnit == nil {
						return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, "发送模型不存在"), "", co_dao.Send.Table())
					}
					sendUnit.Init(int(num), &coSend, &coSendList, context.Background())
				}
			case -2:
				if sendUnitHave {
					emailSend.Pause()
				} else {
					return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, "控制单元不存在"), "", co_dao.Send.Table())
				}
			case -1:
				if sendUnitHave {
					emailSend.Stop()
				} else {
					return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, "控制单元不存在"), "", co_dao.Send.Table())
				}
			}
		} else {
			return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, "控制单元不存在"), "", co_dao.Send.Table())
		}
	}
	return true, nil
}
