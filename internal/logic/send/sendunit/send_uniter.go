package sendunit

import (
	"context"
	"github.com/jack353249002/exam-message-send/sys_model/sys_entity"
)

func SendFactory(unitType int) SendUniter {
	switch unitType {
	case 0:
		return &EmailSend{}
	default:
		return nil
	}
}

type SendUniter interface {
	AddSendList(*sys_entity.CoSendList)
	Init(int, *sys_entity.CoSend, *[]sys_entity.CoSendList, context.Context)
	Stop()
	Pause()
	Start()
}
