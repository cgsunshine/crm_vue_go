// 自动生成模板CrmTickets
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmTickets表 结构体  CrmTickets
type CrmTickets struct {
	global.GVA_MODEL
	Title                   string     `json:"title" form:"title" gorm:"column:title;comment:工单标题;size:255;"`                                            //工单标题
	Description             string     `json:"description" form:"description" gorm:"column:description;comment:工单描述;"`                                   //工单描述
	SubmitterId             *int       `json:"submitterId" form:"submitterId" gorm:"column:submitter_id;comment:提交人ID;size:10;"`                         //提交人ID
	AssigneeId              *int       `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派人ID;size:10;"`                            //指派人ID
	CategoryId              *int       `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:工单类别ID;size:10;"`                           //工单类别ID
	Status                  string     `json:"status" form:"status" gorm:"column:status;comment:工单状态;"`                                                  //工单状态
	Priority                string     `json:"priority" form:"priority" gorm:"column:priority;comment:工单优先级;"`                                           //工单优先级
	EstimatedResolutionTime *time.Time `json:"estimatedResolutionTime" form:"estimatedResolutionTime" gorm:"column:estimated_resolution_time;comment:;"` //estimatedResolutionTime字段
	ActualResolutionTime    *time.Time `json:"actualResolutionTime" form:"actualResolutionTime" gorm:"column:actual_resolution_time;comment:;"`          //actualResolutionTime字段
}

// TableName crmTickets表 CrmTickets自定义表名 crm_tickets
func (CrmTickets) TableName() string {
	return "crm_tickets"
}
