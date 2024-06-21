// 自动生成模板CrmApprovalRecord
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 审批记录 结构体  CrmApprovalRecord
type CrmApprovalRecord struct {
 global.GVA_MODEL 
      ApprovalType  *int `json:"approvalType" form:"approvalType" gorm:"column:approval_type;comment:审批类型 1合同 2商机 3回款;size:10;"`  //审批类型 1合同 2商机 3回款 
      ApproveOpinion  string `json:"approveOpinion" form:"approveOpinion" gorm:"column:approve_opinion;comment:审批意见;size:191;"`  //审批意见 
      ApproverId  *int `json:"approverId" form:"approverId" gorm:"column:approver_id;comment:当前审批人ID;"`  //当前审批人ID 
      AssociatedId  *int `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;size:10;"`  //关联id 合同 商机 回款 
      Status  string `json:"status" form:"status" gorm:"column:status;comment:审批状态;size:191;"`  //审批状态 
}


// TableName 审批记录 CrmApprovalRecord自定义表名 crm_approval_record
func (CrmApprovalRecord) TableName() string {
  return "crm_approval_record"
}

