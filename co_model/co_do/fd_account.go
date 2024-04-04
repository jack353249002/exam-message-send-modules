// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FdAccount is the golang structure of table co_fd_account for DAO operations like Where/Data.
type FdAccount struct {
	g.Meta             `orm:"table:co_fd_account, do:true"`
	Id                 interface{} // ID
	Name               interface{} // 账户名称
	UnionLicenseId     interface{} // 关联资质ID，大于0时必须保值与 union_user_id 关联得上
	UnionUserId        interface{} // 关联用户ID
	CurrencyCode       interface{} // 货币代码
	IsEnabled          interface{} // 是否启用：1启用，0禁用
	LimitState         interface{} // 限制状态：0不限制，1限制支出、2限制收入
	PrecisionOfBalance interface{} // 货币单位精度：1:元，10:角，100:分，1000:厘，10000:毫，……
	Balance            interface{} // 当前余额，必须要与账单最后一笔交易余额对应得上
	Version            interface{} // 乐观锁所需数据版本字段
	CreatedAt          *gtime.Time //
	CreatedBy          interface{} //
	UpdatedAt          *gtime.Time //
	UpdatedBy          interface{} //
	DeletedAt          *gtime.Time //
	DeletedBy          interface{} //
	SceneType          interface{} // 场景类型：0不限、1充电佣金收入
	AccountType        interface{} // 账户类型：1系统账户、2银行卡、3支付宝、4微信、5云闪付、6翼支付
	AccountNumber      interface{} // 账户编号，例如银行卡号、支付宝账号、微信账号等对应账户类型的编号
	UnionMainId        interface{} // 关联主体ID，与union_license_id 中的union_main_id 一致
	AllowExceed        interface{} // 是否允许存在负余额: 0禁止、1允许
}