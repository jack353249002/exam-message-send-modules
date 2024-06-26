package sendunit

import (
	"container/list"
	"context"
	"fmt"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_dao"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_do"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_entity"
	dispatch2 "github.com/jack353249002/exam-message-send-modules/internal/logic/send/dispatch"
	"github.com/jack353249002/exam-message-send-modules/internal/logic/send/sendserver"
	"strconv"
	"strings"
	"sync"
)

type EmailSendResponse struct {
	Code int
}
type EmailSend struct {
	SendInfo      *co_entity.Send
	SendList      chan *co_entity.SendList
	SendListModel *[]co_entity.SendList
	Message       *co_entity.Message
	MaxGo         int //协程数量
	Response      chan *EmailSendResponse
	WithGroup     sync.WaitGroup
	Ctx           context.Context
	SmtpServer    co_entity.SmtpServer
	Dispatch      dispatch2.EmailDispatcher
	Signal        *list.List
	SignalLock    sync.RWMutex
	SignPool      []*SignalCmd
	HaveSignalCmd chan int8
	Status        int8
	StatusLock    sync.RWMutex
}
type SignalCmd struct {
	WithGroup *sync.WaitGroup
	Signal    chan int8
}

func (e *EmailSend) AddSendList(sendList *co_entity.SendList) {
	//e.WithGroup.Add(1)
	e.SendList <- sendList
}
func (e *EmailSend) Init(maxGo int, sendInfo *co_entity.Send, sendList *[]co_entity.SendList, ctx context.Context) {
	e.SendListModel = sendList
	e.Signal = list.New()
	e.HaveSignalCmd = make(chan int8)
	key := strconv.Itoa(sendInfo.Id)
	SendUnitPool.Store(key, e)
	e.Ctx = ctx
	messageInfo := co_entity.Message{}
	co_dao.Message.Ctx(ctx).Where(co_do.Message{
		Id: sendInfo.MessageId,
	}).Scan(&messageInfo)
	e.Message = &messageInfo
	e.Dispatch = dispatch2.NewEmailDispatchPollingFactory()
	serverIds := strings.Split(sendInfo.SendServerId, ",")
	var smtpServer []co_entity.SmtpServer
	co_dao.SmtpServer.Ctx(ctx).WhereIn(co_dao.SmtpServer.Columns().Id, serverIds).Scan(&smtpServer)
	e.Dispatch.Init(5, e.Ctx)
	e.Dispatch.FillServer(smtpServer)
	sendListLength := len(*sendList)
	e.WithGroup = sync.WaitGroup{}
	e.WithGroup.Add(sendListLength)
	co_dao.Send.Ctx(ctx).Data(co_do.Send{Status: 1}).Where("id", sendInfo.Id).Update()
	e.StatusLock.Lock()
	e.Status = 1
	e.StatusLock.Unlock()
	e.SendList = make(chan *co_entity.SendList, sendListLength)
	for _, val := range *sendList {
		valtemp := val
		select {
		case e.SendList <- &valtemp:
		default:
		}
	}
	e.MaxGo = maxGo
	e.SendInfo = sendInfo
	for i := 0; i < e.MaxGo; i++ {
		var signalCmd SignalCmd
		signalCmd.Signal = make(chan int8)
		e.SignPool = append(e.SignPool, &signalCmd)
		go e.StartSend(&signalCmd, i)
		fmt.Println("max_go:", i)
	}
	go func() {
		var beforeSign int8
		for sign := range e.HaveSignalCmd {
			if sign != beforeSign {
				var waitGroup sync.WaitGroup
				waitGroup.Add(e.MaxGo)
				fmt.Println("signcmd_gonum", e.MaxGo)
				if sign == 1 || sign == -2 || sign == -1 {
					e.Dispatch.SetNotWaitServer()
				}
				for _, val := range e.SignPool {
					val.WithGroup = &waitGroup
					val.Signal <- sign
				}
				fmt.Println("cmd with start")
				waitGroup.Wait()
				fmt.Println("cmd with end")
				if sign == 1 && beforeSign != -2 {
					co_dao.Send.Ctx(ctx).Data(co_do.Send{Status: 1}).Where("id", e.SendInfo.Id).Update()
					e.Status = 1
					e.StatusLock.Unlock()
				}
				if sign == 1 && beforeSign == -2 {
					co_dao.Send.Ctx(ctx).Data(co_do.Send{Status: 1}).Where("id", e.SendInfo.Id).Update()
				}
				if sign == -2 {
					co_dao.Send.Ctx(ctx).Data(co_do.Send{Status: 3}).Where("id", e.SendInfo.Id).Update()
				}
				if sign == -1 && beforeSign != -2 {
					e.Status = 0
					e.StatusLock.Unlock()
					e.Close()
					sendInfoIDKey := strconv.Itoa(e.SendInfo.Id)
					SendUnitPool.Delete(sendInfoIDKey)
					co_dao.Send.Ctx(ctx).Data(co_do.Send{Status: 2}).Where("id", e.SendInfo.Id).Update()
				}
				if sign == -1 && beforeSign == -2 {
					e.StatusLock.Lock()
					e.Status = 0
					e.StatusLock.Unlock()
					e.Close()
					sendInfoIDKey := strconv.Itoa(e.SendInfo.Id)
					SendUnitPool.Delete(sendInfoIDKey)
					co_dao.Send.Ctx(ctx).Data(co_do.Send{Status: 2}).Where("id", e.SendInfo.Id).Update()
				}
				beforeSign = sign
				fmt.Println("waitend")
			}
		}
	}()
	go func() {
		e.WithGroup.Wait()
		e.Close()
		//sql = fmt.Sprintf("send_id=%d AND status=1", e.SendInfo.ID)
		count, _ := co_dao.SendList.Ctx(ctx).Where("send_id", e.SendInfo.Id).Where("status", 1).Count()
		co_dao.Send.Ctx(ctx).Data(co_do.Send{Status: 2, SuccessCount: count}).Where("id", e.SendInfo.Id).Update()
		sendInfoKey := strconv.Itoa(sendInfo.Id)
		SendUnitPool.Delete(sendInfoKey)
	}()
}
func (e *EmailSend) Stop() {
	e.StatusLock.Lock()
	e.HaveSignalCmd <- -1
}
func (e *EmailSend) Pause() {
	//e.Signal<--2
	e.HaveSignalCmd <- -2
}
func (e *EmailSend) Start() {
	e.StatusLock.Lock()
	status := e.Status
	e.StatusLock.Unlock()
	if status != 0 {
		e.HaveSignalCmd <- 1
	} else {
		e.Init(e.MaxGo, e.SendInfo, e.SendListModel, e.Ctx)
	}
}
func (e *EmailSend) Close() {
	if e.SendList != nil {
		close(e.SendList)
	}
	if e.HaveSignalCmd != nil {
		close(e.HaveSignalCmd)
	}
	if e.Response != nil {
		close(e.Response)
	}
}
func (e *EmailSend) StartSend(signalCmd *SignalCmd, goIndex int) {
	type SingnalInfo struct {
		SignalType uint8 //0=数据,1=信号
		Signal     int8
		SendList   *co_entity.SendList
	}
	var privateSignal chan *SingnalInfo
	privateSignal = make(chan *SingnalInfo)
	//var sendSignal chan bool  //true代表可以发
	//sendSignal=make(chan bool)
	var signalContoller chan bool
	signalContoller = make(chan bool)
	go func() {
		//sendSignal<-true
	}()
	go func() {
		for signalval := range signalCmd.Signal {
			fmt.Println("signalcmd:", signalval)
			var signalInfo SingnalInfo
			signalInfo.SignalType = 1
			signalInfo.Signal = signalval
			//<-sendSignal
			privateSignal <- &signalInfo
			<-signalContoller
		}
	}()
	go func() {
		for send := range e.SendList {
			var signalInfo SingnalInfo
			signalInfo.SignalType = 0
			signalInfo.Signal = -1
			signalInfo.SendList = send
			/*sendSignal<-true
			privateSignal<-&signalInfo*/
			//<-sendSignal //取出消息
			privateSignal <- &signalInfo
			//time.Sleep(5 * time.Second)
		}
	}()
	for val := range privateSignal {
		fmt.Println("aaa:", val.Signal)
		if val.SignalType == 1 {
			switch val.Signal {
			case 1:
				signalContoller <- true
				//sendSignal<-true
				fmt.Println("open:", goIndex)
			case -1:
				signalContoller <- true
				signalCmd.WithGroup.Done()
				return
			case -2:
				fmt.Println("purse:", goIndex)
				signalCmd.WithGroup.Done()
				for signalval := range signalCmd.Signal {
					if signalval == 1 {
						isWait := e.Dispatch.GetIsWaitServer()
						if !isWait {
							e.Dispatch.SetIsWaitServer()
						}
						signalContoller <- true
						//sendSignal<-true
						fmt.Println("open2:", goIndex)
						signalCmd.WithGroup.Done()
						break
					}
					if signalval == -1 {
						<-signalContoller
						//sendSignal<-true
						fmt.Println("stop2:", goIndex)
						//停止
						return
					}
				}
				/*for signalTempPause := range privateSignal {
					if signalTempPause.SignalType == 1 && signalTempPause.Signal == 1 {
						sendSignal<-true
						fmt.Println("open2:", goIndex)
						break
					}
					if signalTempPause.SignalType == 1 && signalTempPause.Signal == -1 {
						sendSignal<-true
						fmt.Println("stop2:", goIndex)
						//停止
						return
					}
					fmt.Println("other",signalTempPause)
				}*/
			}
		} else {
			server, index, have, inCoolTime := e.Dispatch.GetServerListen()
			if have {
				err := e.Send(val.SendList, &server)
				if err != nil {
					if !inCoolTime {
						e.Dispatch.RemoveServer(index, server.Id, "ErrorCount-2")
					}
					e.AddSendList(val.SendList)
				} else {
					co_dao.SendList.Ctx(e.Ctx).Data(co_do.SendList{Status: 1, SendServerId: server.Id}).Where("id", val.SendList.Id).Update()
					e.WithGroup.Done()
				}
			} else {
				e.AddSendList(val.SendList)
			}
			//sendSignal<-true
		}
	}
}
func (e *EmailSend) Send(send *co_entity.SendList, smtpserver *co_entity.SmtpServer) (err error) {
	var smtp sendserver.Smtp
	smtp.Port = smtpserver.Port
	smtp.Server = smtpserver.SmtpServer
	smtp.SenderEmail = smtpserver.SmtpSendEmail
	smtp.SenderPassword = smtpserver.SmtpPassword
	var attch []string
	if e.Message.Attach != "" {
		attch = strings.Split(e.Message.Attach, ",")
	}
	err = smtp.Send(send.Receive, e.Message.Title, e.Message.Body, attch)
	/*fmt.Println(attch)
	fmt.Println("smtpserver:", smtpserver.Title)
	fmt.Println("receive:", send.Receive)
	fmt.Println("---------------")*/
	/*if err==nil {

	}*/
	/*if smtpserver.ID==1 || smtpserver.ID==2{
		err=fmt.Errorf("错误")
	}*/
	return
}
