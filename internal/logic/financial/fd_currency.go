package financial

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/jack353249002/exam-message-send-modules/co_interface"
	"github.com/jack353249002/exam-message-send-modules/co_model"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_dao"
	"github.com/jack353249002/exam-message-send-modules/co_model/co_do"
	"github.com/kysion/base-library/base_hook"
)

// 货币类型管理
type sFdCurrency[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	TR co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	base_hook.ResponseFactoryHook[TR]
	modules co_interface.IModules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		TR,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]
	dao *co_dao.XDao
}

func NewFdCurrency[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	TR co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TR,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) co_interface.IFdCurrency[TR] {
	result := &sFdCurrency[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		TR,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]{
		modules: modules,
		dao:     modules.Dao(),
	}

	result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)

	return result
}

// FactoryMakeResponseInstance 响应实例工厂方法
func (s *sFdCurrency[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TR,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) FactoryMakeResponseInstance() TR {
	var ret co_model.IFdCurrencyRes
	ret = &co_model.FdCurrencyRes{}
	return ret.(TR)
}

// GetCurrencyByCurrencyCode 根据货币代码查找货币(主键)
func (s *sFdCurrency[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TR,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetCurrencyByCurrencyCode(ctx context.Context, currencyCode string) (response TR, err error) {
	if currencyCode == "" {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_CurrencyCode_NotNull"), s.dao.FdCurrency.Table())
	}

	result := s.FactoryMakeResponseInstance()

	err = s.dao.FdCurrency.Ctx(ctx).Where(co_do.FdCurrency{CurrencyCode: currencyCode}).Scan(result.Data())

	if err != nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#Currency}{#error_Data_Get_Failed}"), s.dao.FdCurrency.Table())
	}

	return result, nil
}

// GetCurrencyByCnName 根据国家查找货币信息
func (s *sFdCurrency[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	TR,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetCurrencyByCnName(ctx context.Context, cnName string) (response TR, err error) {
	if cnName == "" {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_CurrencyCnName_NotNull"), s.dao.FdCurrency.Table())
	}

	result := s.FactoryMakeResponseInstance()

	err = s.dao.FdCurrency.Ctx(ctx).Where(co_do.FdCurrency{CnName: cnName}).Scan(result.Data())
	if err != nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#Currency}{#error_Data_Get_Failed}"), s.dao.FdCurrency.Table())
	}

	return result, nil
}