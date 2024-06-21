// 自动生成模板CrmApprovalTasks
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// crmApprovalTasks表 结构体  CrmApprovalTasks
type CrmApprovalTasks struct {
	global.GVA_MODEL
	RequestId *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;size:10;"` //关联的审批请求ID
	StepId    *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;size:10;"`           //当前审批步骤ID
	Assignee  string `json:"assignee" form:"assignee" gorm:"column:assignee;comment:指派审批人;size:50;"`         //指派审批人
	Status    string `json:"status" form:"status" gorm:"column:status;comment:任务状态;"`                        //任务状态
	Comments  string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;"`                  //审批意见
}

// TableName crmApprovalTasks表 CrmApprovalTasks自定义表名 crm_approval_tasks
func (CrmApprovalTasks) TableName() string {
	return "crm_approval_tasks"
}
