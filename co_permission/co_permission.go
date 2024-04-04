package co_permission

import (
	"github.com/jack353249002/exam-message-send-modules/co_permission/internal/company"
	"github.com/jack353249002/exam-message-send-modules/co_permission/internal/employee"
	"github.com/jack353249002/exam-message-send-modules/co_permission/internal/financial"
	"github.com/jack353249002/exam-message-send-modules/co_permission/internal/liceense"
	"github.com/jack353249002/exam-message-send-modules/co_permission/internal/team"
)

type (
	CompanyPermissionType   = company.Permission
	EmployeePermissionType  = employee.Permission
	TeamPermissionType      = team.Permission
	FinancialPermissionType = financial.Permission

	LicensePermissionType = liceense.PermissionTypeEnum
)

var (
	Company   = company.Company
	Employee  = employee.Employee
	Team      = team.Team
	Financial = financial.Financial

	License = liceense.License
)
