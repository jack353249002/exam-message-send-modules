package sendunit

import (
	"context"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_entity"
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
	AddSendList(*co_entity.SendList)
	Init(int, *co_entity.Send, *[]co_entity.SendList, context.Context)
	Stop()
	Pause()
	Start()
}
