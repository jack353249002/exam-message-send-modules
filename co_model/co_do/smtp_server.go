// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SmtpServer is the golang structure of table co_smtp_server for DAO operations like Where/Data.
type SmtpServer struct {
	g.Meta         `orm:"table:co_smtp_server, do:true"`
	Id             interface{} //
	SmtpServer     interface{} //
	SmtpPassword   interface{} //
	Title          interface{} //
	SmtpSendEmail  interface{} //
	Port           interface{} //
	MaxConcurrency interface{} //
}
