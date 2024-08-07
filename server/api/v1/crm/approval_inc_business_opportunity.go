package crm

import "github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"

// 商机审批
type ApprovalIncBusinessOpportunity struct{}

func NewApprovalIncBusinessOpportunity() *ApprovalIncBusinessOpportunity {
	return &ApprovalIncBusinessOpportunity{}
}

func (a ApprovalIncBusinessOpportunity) ApprovalProcessSuccess(id *int) error {
	return crmBusinessOpportunityService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Pass,
	})
}

func (a ApprovalIncBusinessOpportunity) ApprovalProcessFail(id *int) error {
	//合同审批
	//有人拒绝，修改合同状态
	return crmBusinessOpportunityService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Rafuse,
	})
}
