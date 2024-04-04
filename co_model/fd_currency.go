package co_model

import "github.com/jack353249002/exam-message-send-modules/co_model/co_entity"

type FdCurrencyRes struct {
	co_entity.FdCurrency
}

func (m *FdCurrencyRes) Data() *FdCurrencyRes {
	return m
}

type IFdCurrencyRes interface {
	Data() *FdCurrencyRes
}
