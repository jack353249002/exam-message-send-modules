package co_send

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/jack353249002/exam-message-send-modules/co_interface"
	"github.com/jack353249002/exam-message-send-modules/co_model"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_dao"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_entity"
	"strings"
)

type sSend[
	TR co_model.IMessageRes,
	TSR co_model.ISendRes,
] struct {
	//base_hook.ResponseFactoryHook[TR]
	modules co_interface.IModules[TR, TSR]
	dao     co_dao.XDao
	//makeMoreFunc func(ctx context.Context, data co_model.ICompanyRes, employeeModule co_interface.IEmployee[co_model.IEmployeeRes]) co_model.ICompanyRes
}

// FactoryMakeResponseInstance 响应实例工厂方法
/*func (s *sMessage[TR]) FactoryMakeResponseInstance() TR {
	var ret co_model.IMessageRes
	ret = &co_model.MessageRes{Message: &co_entity.Message{}}
	return ret.(TR)
}*/
func NewSend[
	TR co_model.IMessageRes,
	TSR co_model.ISendRes,
](modules co_interface.IModules[TR, TSR]) co_interface.ISend[TSR] {
	result := &sSend[TR, TSR]{
		modules: modules,
	}

	//result.makeMoreFunc = MakeMore

	//result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)
	return result
}

// 添加发送规则
func (s *sSend[TR, TSR]) CreateSend(ctx context.Context, title string, messageId int, sendServerId string, receive string) (bool, error) {
	receives := strings.Split(receive, ",")
	err := co_dao.Send.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		send := &co_entity.Send{
			Title:         title,
			MessageId:     messageId,
			SendServerId:  sendServerId,
			CreatedAt:     gtime.Now(),
			UpdatedAt:     gtime.Now(),
			SendListCount: len(receives),
		}
		res, err := co_dao.Send.Ctx(ctx).OmitEmpty().OmitNilData().Data(send).Insert()
		lastInsertId, err := res.LastInsertId()
		var sendList []map[string]interface{}
		for _, val := range receives {
			dataRow := map[string]interface{}{"send_id": lastInsertId, "receive": val, "status": 0, "send_server_id": 0}
			sendList = append(sendList, dataRow)
		}
		if len(sendList) > 0 {
			res, err = co_dao.SendList.Ctx(ctx).OmitEmpty().OmitNilData().Data(sendList).Insert()
		}
		return err
	})
	if err == nil {
		return true, nil
	} else {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, err.Error()), "", co_dao.SendList.Table())
	}
	return true, nil
}
func (s *sSend[TR, TSR]) SetSendInfoAction(ctx context.Context, sendId int, status int8) (bool, error) {
	return true, nil
}
