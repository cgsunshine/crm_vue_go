// 自动生成模板CrmPaymentCollentionApprovalTasks
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// crmPaymentCollentionApprovalTasks表 结构体  CrmPaymentCollentionApprovalTasks
type CrmPaymentCollentionApprovalTasks struct {
 global.GVA_MODEL 
      ApprovalStatus  string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`  //审批状态 
      AssigneeId  *int `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;"`  //指派审批人id 
      Comments  string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`  //审批意见 
      PaymentCollentionId  *int `json:"paymentCollentionId" form:"paymentCollentionId" gorm:"column:payment_collention_id;comment:回款id;"`  //回款id 
      RequestId  *int `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;size:19;"`  //关联的审批请求ID 
      StepId  *int `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;size:19;"`  //当前审批步骤ID 
      Valid  *int `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;"`  //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝） 
}


// TableName crmPaymentCollentionApprovalTasks表 CrmPaymentCollentionApprovalTasks自定义表名 crm_payment_collention_approval_tasks
func (CrmPaymentCollentionApprovalTasks) TableName() string {
  return "crm_payment_collention_approval_tasks"
}

