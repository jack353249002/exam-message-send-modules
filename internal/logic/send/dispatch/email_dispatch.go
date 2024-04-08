package dispatch

import (
	"context"
	"fmt"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_entity"
	"sync"
	"time"
)

func NewEmailDispatchPollingFactory() EmailDispatcher {
	return &EmailDispatchPolling{}
}

type SmtpServerInfo struct {
	SmtpServer            co_entity.SmtpServer
	ErrorCount            int
	Index                 int
	Have                  bool
	ErrorCheckLastOutTime time.Time //错误检测最后释放时间
	ErrorCheckCoolTime    int       //冷却时间
}

func (s *SmtpServerInfo) IsInCoolTime() bool {
	return s.ErrorCheckLastOutTime.Unix() > 0 && s.ErrorCheckLastOutTime.Unix()+int64(s.ErrorCheckCoolTime) > time.Now().Unix()
}

type SignalInfo struct {
	ServerInfoResponse chan SmtpServerInfo
}

// 邮箱发送轮询法
type EmailDispatchPolling struct {
	Index                    int
	SmtpServer               []SmtpServerInfo
	SmtpServerReady          []*SmtpServerInfo
	SmtpServerSendErrorCheck []*SmtpServerInfo //异常的smtp服务器
	SmtpServerReadyIndex     map[int]int
	Length                   int //总长度
	LengthReady              int //可用长度
	Lock                     sync.Mutex
	ErrorMax                 int //错误最大数
	Context                  context.Context
	Signal                   chan struct{}   //通知是否有可用的服务器
	SignalGroup              chan SignalInfo //用于阻塞模式下等待到有服务器后通知所有获取者
	IsWaitServer             bool            //是否等待有服务器
	IsWaitServerLock         sync.RWMutex
}

func (e *EmailDispatchPolling) SetIsWaitServer() {
	e.IsWaitServerLock.Lock()
	e.IsWaitServer = true
	e.IsWaitServerLock.Unlock()
}
func (e *EmailDispatchPolling) SetNotWaitServer() {
	e.IsWaitServerLock.Lock()
	e.IsWaitServer = false
	e.IsWaitServerLock.Unlock()
	select {
	case e.Signal <- struct{}{}:

	default:
	}
}
func (e *EmailDispatchPolling) GetIsWaitServer() bool {
	e.IsWaitServerLock.RLock()
	lock := e.IsWaitServer
	e.IsWaitServerLock.RUnlock()
	return lock
}

func (e *EmailDispatchPolling) MoveIndex() {
	e.Lock.Lock()
	if e.Index > e.LengthReady-1 {
		e.Index = 0
	}
	e.Index++
	e.Lock.Unlock()
	return
}
func (e *EmailDispatchPolling) RemoveServer(index int, id int, ruleType string) {
	e.Lock.Lock()
	switch ruleType {
	//根据错误次数进行回收，达到次数后直接释放掉
	case "ErrorCount-1":
		if index < len(e.SmtpServerReady) && e.SmtpServerReady[index].SmtpServer.Id == id && len(e.SmtpServerReady) > 0 {
			if e.SmtpServerReady[index].ErrorCount >= e.ErrorMax-1 {
				e.LengthReady--
				e.SmtpServerReady = append(e.SmtpServerReady[:index], e.SmtpServerReady[index+1:]...)
			} else {
				e.SmtpServerReady[index].ErrorCount++
			}
		} else if index < len(e.SmtpServerReady) && e.SmtpServerReady[index].SmtpServer.Id != id {
			indextemp := 0
			have := false
			for _, val := range e.SmtpServerReady {
				if val.SmtpServer.Id == id {
					have = true
					break
				}
				indextemp++
			}
			if have {
				//e.SmtpServerReady = append(e.SmtpServerReady[:indextemp], e.SmtpServerReady[indextemp+1:]...)
				if e.SmtpServerReady[index].ErrorCount >= e.ErrorMax-1 {
					e.LengthReady--
					e.SmtpServerReady = append(e.SmtpServerReady[:indextemp], e.SmtpServerReady[indextemp+1:]...)
				} else {
					e.SmtpServerReady[indextemp].ErrorCount++
				}
			}
		}
		//存放到检测区
	case "ErrorCount-2":
		if index < len(e.SmtpServerReady) && e.SmtpServerReady[index].SmtpServer.Id == id && len(e.SmtpServerReady) > 0 {
			smtpServerInfo := e.SmtpServerReady[index]
			if !smtpServerInfo.IsInCoolTime() {
				e.LengthReady--
				e.SmtpServerReady = append(e.SmtpServerReady[:index], e.SmtpServerReady[index+1:]...)
				go e.errorCheck(smtpServerInfo)
			}

		} else if index < len(e.SmtpServerReady) && e.SmtpServerReady[index].SmtpServer.Id != id {
			indextemp := 0
			have := false
			for readyIndexTemp, val := range e.SmtpServerReady {
				if val.SmtpServer.Id == id {
					have = true
					indextemp = readyIndexTemp
					break
				}
				//indextemp++
			}
			if have {
				smtpServerInfo := e.SmtpServerReady[indextemp]
				//e.SmtpServerReady = append(e.SmtpServerReady[:indextemp], e.SmtpServerReady[indextemp+1:]...)
				if !smtpServerInfo.IsInCoolTime() {
					e.LengthReady--
					e.SmtpServerReady = append(e.SmtpServerReady[:indextemp], e.SmtpServerReady[indextemp+1:]...)
					go e.errorCheck(smtpServerInfo)
				}
			}
		}

	}
	e.Lock.Unlock()
}
func (e *EmailDispatchPolling) errorCheck(smtpServerInfo *SmtpServerInfo) {
	//timestr := time.Now().String()
	//warnmessage := fmt.Sprintf("时间:%s 有错误的smtp服务器进入检查:[id:%d,地址:%s]", timestr, smtpServerInfo.SmtpServer.Id, smtpServerInfo.SmtpServer.SmtpServer)
	defer func() {
		//timestr = time.Now().String()
		//warnmessage = fmt.Sprintf("时间:%s 有错误的smtp服务器放出检查:[id:%d,地址:%s]", timestr, smtpServerInfo.SmtpServer.ID, smtpServerInfo.SmtpServer.SmtpServer)
	}()
	for {
		if true {
			e.Lock.Lock()
			smtpServerInfo.ErrorCheckLastOutTime = time.Now()
			e.SmtpServerReady = append(e.SmtpServerReady, smtpServerInfo)
			e.LengthReady++
			e.Lock.Unlock()
			select {
			case e.Signal <- struct{}{}:
				fmt.Println(smtpServerInfo.SmtpServer.SmtpServer + "signal success")
			default:
				fmt.Println(smtpServerInfo.SmtpServer.SmtpServer + "signal default")
			}
			fmt.Println(smtpServerInfo.SmtpServer.SmtpServer + "成功2")
			break
		} else {
			fmt.Println(smtpServerInfo.SmtpServer.SmtpServer + "失败2")
			//fmt.Println("错误信息:", err.Error())
		}
		time.Sleep(5 * time.Second)
		select {
		case <-e.Context.Done():
			fmt.Println("结束循环检测")
			return
		default:

		}
	}
}
func (e *EmailDispatchPolling) getSignal() {
	var getServerFunc func() (SmtpServerInfo, int, bool)
	getServerFunc = func() (SmtpServerInfo, int, bool) {
		server, index, have := e.GetServer()
		return server, index, have
	}
	for signalServer := range e.SignalGroup {
		server, index, have := getServerFunc()
		var smtpserverInfo SmtpServerInfo
		smtpserverInfo.ErrorCheckCoolTime = ErrorCheckCoolTime
		if have {
			smtpserverInfo.SmtpServer = server.SmtpServer
			smtpserverInfo.Index = index
			signalServer.ServerInfoResponse <- smtpserverInfo
		} else {
			//如果需要执行指令，需要解除等待。直接强制发送已有新服务器的信号
			isWait := e.GetIsWaitServer()
			fmt.Println("need_wait:", isWait)
			if isWait {
				<-e.Signal
				server, index, have = getServerFunc()
				smtpserverInfo.SmtpServer = server.SmtpServer
				smtpserverInfo.Index = index
				smtpserverInfo.Have = have
				signalServer.ServerInfoResponse <- smtpserverInfo
			} else {
				smtpserverInfo.SmtpServer = server.SmtpServer
				smtpserverInfo.Index = index
				smtpserverInfo.Have = have
				signalServer.ServerInfoResponse <- smtpserverInfo
			}
		}
	}
}
func (e *EmailDispatchPolling) Init(errorMax int, context context.Context) {
	e.ErrorMax = errorMax
	e.Context = context
	e.Signal = make(chan struct{})
	e.SignalGroup = make(chan SignalInfo)
	e.SetIsWaitServer()
	go e.getSignal()
}
func (e *EmailDispatchPolling) FillServer(smtps []co_entity.SmtpServer) {
	e.Length = len(smtps)
	e.LengthReady = e.Length
	//e.SmtpServer=smtps
	for _, val := range smtps {
		var smtpinfo SmtpServerInfo
		smtpinfo.ErrorCheckCoolTime = ErrorCheckCoolTime
		smtpinfo.SmtpServer = val
		e.SmtpServer = append(e.SmtpServer, smtpinfo)
		e.SmtpServerReady = append(e.SmtpServerReady, &smtpinfo)
	}
}

func (e *EmailDispatchPolling) GetServer() (res SmtpServerInfo, index int, have bool) {
	e.Lock.Lock()
	have = true
	if e.Index > e.LengthReady-1 {
		e.Index = 0
	}
	if e.LengthReady <= 0 {
		index = 0
		have = false
		e.Lock.Unlock()
		return
	} else {
		resinfo := e.SmtpServerReady[e.Index]
		res = *resinfo
		index = e.Index
		e.Index++
	}
	e.Lock.Unlock()
	return
}

// listenHaveServer:如果当前没有可用服务器，可以进行阻塞并待有数据后再触发
func (e *EmailDispatchPolling) GetServerListen() (res co_entity.SmtpServer, index int, have bool, inErrCheckCoolTime bool) {
	smtpInfo, index, have := e.GetServer()
	if have {
		res = smtpInfo.SmtpServer
		inErrCheckCoolTime = smtpInfo.IsInCoolTime()
		return
	} else {
		var signalInfo SignalInfo
		var smtpserver chan SmtpServerInfo
		smtpserver = make(chan SmtpServerInfo)
		signalInfo.ServerInfoResponse = smtpserver
		e.SignalGroup <- signalInfo
		server := <-smtpserver
		res = server.SmtpServer
		index = server.Index
		have = server.Have
		inErrCheckCoolTime = server.IsInCoolTime()
		fmt.Println(res.SmtpServer + "异步获取成功")
		return
	}
}
