// 自动生成模板CrmApprovalTasks
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 审批任务 结构体  CrmApprovalTasks
type CrmApprovalTasks struct {
 global.GVA_MODEL 
      ApprovalStatus  string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`  //审批状态 
      ApprovalType  *int `json:"approvalType" form:"approvalType" gorm:"column:approval_type;comment:审批类型 1合同 2商机 3回款;size:10;"`  //审批类型 1合同 2商机 3回款 
      AssigneeId  *int `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;size:19;"`  //指派审批人id 
      AssociatedId  *int `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;size:10;"`  //关联id 合同 商机 回款 
      Comments  string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`  //审批意见 
      RequestId  *int `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`  //关联的审批请求ID 
      StepId  *int `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`  //当前审批步骤ID 
      Valid  *int `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:19;"`  //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝） 
}


// TableName 审批任务 CrmApprovalTasks自定义表名 crm_approval_tasks
func (CrmApprovalTasks) TableName() string {
  return "crm_approval_tasks"
}

