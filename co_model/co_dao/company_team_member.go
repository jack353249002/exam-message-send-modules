// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package co_dao

import (
	"github.com/jack353249002/exam-message-send-modules/co_model/co_dao/internal"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
)

type CompanyTeamMemberDao = dao_interface.TIDao[internal.CompanyTeamMemberColumns]

func NewCompanyTeamMember(dao ...dao_interface.IDao) CompanyTeamMemberDao {
	return (CompanyTeamMemberDao)(internal.NewCompanyTeamMemberDao(dao...))
}

var (
	CompanyTeamMember = NewCompanyTeamMember()
)
