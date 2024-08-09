// 自动生成模板CrmContactApprovalRecord
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 合同审批记录 结构体  CrmContactApprovalRecord
type CrmContactApprovalRecord struct {
 global.GVA_MODEL 
      ApproveOpinion  string `json:"approveOpinion" form:"approveOpinion" gorm:"column:approve_opinion;comment:审批意见;size:191;"`  //审批意见 
      ApproverId  string `json:"approverId" form:"approverId" gorm:"column:approver_id;comment:当前审批人ID;size:64;"`  //当前审批人ID 
      ContactId  *int `json:"contactId" form:"contactId" gorm:"column:contact_id;comment:合同id;size:10;"`  //合同id 
      Status  string `json:"status" form:"status" gorm:"column:status;comment:审批状态;size:191;"`  //审批状态 
}


// TableName 合同审批记录 CrmContactApprovalRecord自定义表名 crm_contact_approval_record
func (CrmContactApprovalRecord) TableName() string {
  return "crm_contact_approval_record"
}

