package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/jack353249002/exam-message-send-modules/co_interface"
	"github.com/jack353249002/exam-message-send-modules/co_model"
	"github.com/jack353249002/exam-message-send-modules/example/controller"
)

func ModulesGroup[
	ITMessageRes co_model.IMessageRes,
	ITSendRes co_model.ISendRes,
](modules co_interface.IModules[ITMessageRes, ITSendRes], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	MessageGroup(modules, group)
	SendGroup(modules, group)
	return group
}
func MessageGroup[
	ITMessage co_model.IMessageRes,
	ITSend co_model.ISendRes,
](modules co_interface.IModules[ITMessage, ITSend], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := "/" + "message"
	controllerObj := controller.Message(modules)

	group.POST(routePrefix+"/createMessage", controllerObj.CreateMessage)
	group.POST(routePrefix+"/queryMessageList", controllerObj.QueryMessageList)
	return group
}
func SendGroup[
	ITMessage co_model.IMessageRes,
	ITSend co_model.ISendRes,
](modules co_interface.IModules[ITMessage, ITSend], group *ghttp.RouterGroup) *ghttp.RouterGroup {
	routePrefix := "/" + "send"
	controller := controller.Send(modules)

	group.POST(routePrefix+"/createSend", controller.CreateSend)
	group.POST(routePrefix+"/setSendAction", controller.SetSendInfoAction)
	group.POST(routePrefix+"/querySendInfoList", controller.QuerySendInfoList)
	return group
}
