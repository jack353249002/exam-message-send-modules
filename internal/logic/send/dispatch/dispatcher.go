package dispatch

import (
	"context"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_entity"
)

const ErrorCheckCoolTime = 120

type EmailDispatcher interface {
	MoveIndex()
	FillServer([]co_entity.SmtpServer)
	GetServer() (SmtpServerInfo, int, bool)
	GetServerListen() (co_entity.SmtpServer, int, bool, bool)
	RemoveServer(int, int, string)
	Init(int, context.Context)
	SetIsWaitServer()
	SetNotWaitServer()
	GetIsWaitServer() bool
}
