// 自动生成模板CrmContactApprovalTasks
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 合同审批表 结构体  CrmContactApprovalTasks
type CrmContactApprovalTasks struct {
 global.GVA_MODEL 
      RequestId  *int `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`  //关联的审批请求ID 
      StepId  *int `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`  //当前审批步骤ID 
      Assignee  string `json:"assignee" form:"assignee" gorm:"column:assignee;comment:指派审批人;size:50;"`  //指派审批人 
      ApprovalStatus  string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`  //审批状态 
      Comments  string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`  //审批意见 
      ContactId  *int `json:"contactId" form:"contactId" gorm:"column:contact_id;comment:合同id;size:10;"`  //合同id 
      Valid  *int `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:10;"`  //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝） 
}


// TableName 合同审批表 CrmContactApprovalTasks自定义表名 crm_contact_approval_tasks
func (CrmContactApprovalTasks) TableName() string {
  return "crm_contact_approval_tasks"
}

