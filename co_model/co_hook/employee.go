package co_hook

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/jack353249002/exam-message-send-modules/co_model"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_enum"
)

type EmployeeHookFilter struct {
	InOutType     co_enum.FinancialInOutType
	TradeType     co_enum.FinancialTradeType
	InTransaction bool
}

type EmployeeHookFunc sys_model.HookFunc[EmployeeHookFilter, co_model.EmployeeRes]
