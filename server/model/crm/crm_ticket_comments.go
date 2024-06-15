// 自动生成模板CrmTicketComments
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 共单回复 结构体  CrmTicketComments
type CrmTicketComments struct {
 global.GVA_MODEL 
      TicketId  *int `json:"ticketId" form:"ticketId" gorm:"column:ticket_id;comment:关联的工单ID;"`  //关联的工单ID 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:评论用户ID;"`  //评论用户ID 
      Comment  string `json:"comment" form:"comment" gorm:"column:comment;comment:回复内容;size:191;"`  //回复内容 
}


// TableName 共单回复 CrmTicketComments自定义表名 crm_ticket_comments
func (CrmTicketComments) TableName() string {
  return "crm_ticket_comments"
}

