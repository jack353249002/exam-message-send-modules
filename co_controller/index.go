package co_controller

import (
	"github.com/jack353249002/exam-message-send-modules/co_controller/internal"
	"github.com/jack353249002/exam-message-send-modules/co_interface"
	"github.com/jack353249002/exam-message-send-modules/co_interface/i_controller"
	"github.com/jack353249002/exam-message-send-modules/co_model"
)

func Message[
	TIRes co_model.IMessageRes,
	TISRes co_model.ISendRes,
](modules co_interface.IModules[TIRes, TISRes]) i_controller.IMessage[TIRes] {
	return internal.Message[TIRes](modules)
}

func Send[
	TIRes co_model.IMessageRes,
	TISRes co_model.ISendRes,
](modules co_interface.IModules[TIRes, TISRes]) i_controller.ISend[TISRes] {
	return internal.Send[TIRes, TISRes](modules)
}
