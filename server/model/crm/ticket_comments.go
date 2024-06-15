// 自动生成模板CrmTicketComments
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 共单回复 结构体  CrmTicketComments
type CrmPageTicketComments struct {
	global.GVA_MODEL
	TicketId *int   `json:"ticketId" form:"ticketId" gorm:"column:ticket_id;comment:关联的工单ID;"`   //关联的工单ID
	UserId   *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:评论用户ID;"`          //评论用户ID
	Comment  string `json:"comment" form:"comment" gorm:"column:comment;comment:回复内容;size:191;"` //回复内容
	//以下是联表查询字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
}
