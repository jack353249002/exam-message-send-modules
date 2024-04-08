package co_message

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/jack353249002/exam-message-send-modules/co_interface"
	"github.com/jack353249002/exam-message-send-modules/co_model"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_dao"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
)

type sMessage[
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
func NewMessage[
	TR co_model.IMessageRes,
	TSR co_model.ISendRes,
](modules co_interface.IModules[TR, TSR]) co_interface.IMessage[TR] {
	result := &sMessage[TR, TSR]{
		modules: modules,
	}

	//result.makeMoreFunc = MakeMore

	//result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)
	return result
}
func (s *sMessage[TR, TSR]) CreateMessage(ctx context.Context, title string, body string) (bool, error) {
	data := &co_entity.Message{
		Title:     title,
		Body:      body,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}
	_, err := co_dao.Message.Ctx(ctx).OmitEmpty().OmitNilData().Data(data).Insert()
	if err == nil {
		return true, nil
	} else {
		return false, sys_service.SysLogs().ErrorSimple(ctx, gerror.NewCode(gcode.CodeDbOperationError, err.Error()), "", co_dao.Message.Table())
	}
	return true, nil
}

// 消息列表
func (s *sMessage[TR, TSR]) QueryMessageList(ctx context.Context, filter *base_model.SearchParams) (*base_model.CollectRes[TR], error) {
	result, err := daoctl.Query[TR](co_dao.Message.Ctx(ctx), filter, false)

	return result, err

}
