// 自动生成模板CrmBusinessOpportunityApprovalRecord
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 商机审批记录 结构体  CrmBusinessOpportunityApprovalRecord
type CrmBusinessOpportunityApprovalRecord struct {
 global.GVA_MODEL 
      ApproveOpinion  string `json:"approveOpinion" form:"approveOpinion" gorm:"column:approve_opinion;comment:审批意见;size:191;"`  //审批意见 
      ApproverId  *int `json:"approverId" form:"approverId" gorm:"column:approver_id;comment:当前审批人ID;"`  //当前审批人ID 
      BusinessOpportunityId  *int `json:"businessOpportunityId" form:"businessOpportunityId" gorm:"column:business_opportunity_id;comment:商机id;size:19;"`  //商机id 
      Status  string `json:"status" form:"status" gorm:"column:status;comment:审批状态;size:191;"`  //审批状态 
}


// TableName 商机审批记录 CrmBusinessOpportunityApprovalRecord自定义表名 crm_business_opportunity_approval_record
func (CrmBusinessOpportunityApprovalRecord) TableName() string {
  return "crm_business_opportunity_approval_record"
}

