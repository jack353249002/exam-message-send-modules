package co_dao

type XDao struct {
	Message    MessageDao
	Company    CompanyDao
	Employee   CompanyEmployeeDao
	Team       CompanyTeamDao
	TeamMember CompanyTeamMemberDao

	FdAccount       FdAccountDao
	FdAccountBill   FdAccountBillDao
	FdInvoice       FdInvoiceDao
	FdInvoiceDetail FdInvoiceDetailDao
	FdCurrency      FdCurrencyDao
	FdBankCard      FdBankCardDao
	FdAccountDetail FdAccountDetailDao
}
