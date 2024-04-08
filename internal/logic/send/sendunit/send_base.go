package sendunit

import "sync"

var SendUnitPool sync.Map

func GetSendEmailUnitPool(key string) (res *EmailSend, have bool) {
	val, have := SendUnitPool.Load(key)
	if have {
		res = val.(*EmailSend)
	}
	return
}
