// 自动生成模板CrmPaymentCollentionApprovalRecord
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 回款审批记录 结构体  CrmPaymentCollentionApprovalRecord
type CrmPaymentCollentionApprovalRecord struct {
 global.GVA_MODEL 
      ApproveOpinion  string `json:"approveOpinion" form:"approveOpinion" gorm:"column:approve_opinion;comment:审批意见;size:191;"`  //审批意见 
      ApproverId  *int `json:"approverId" form:"approverId" gorm:"column:approver_id;comment:当前审批人ID;"`  //当前审批人ID 
      PaymentCollentionId  *int `json:"paymentCollentionId" form:"paymentCollentionId" gorm:"column:payment_collention_id;comment:回款id;size:19;"`  //回款id 
      Status  string `json:"status" form:"status" gorm:"column:status;comment:审批状态;size:191;"`  //审批状态 
}


// TableName 回款审批记录 CrmPaymentCollentionApprovalRecord自定义表名 crm_payment_collention_approval_record
func (CrmPaymentCollentionApprovalRecord) TableName() string {
  return "crm_payment_collention_approval_record"
}

