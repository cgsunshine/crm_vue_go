package crm

import "github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"

// 押金审批
type ApprovalIncDeposits struct{}

func NewApprovalIncDeposits() *ApprovalIncDeposits {
	return &ApprovalIncDeposits{}
}

func (a ApprovalIncDeposits) ApprovalProcessSuccess(id *int) error {
	//押金审批
	//所有人都同意，修改押金状态
	return crmDepositsService.UpdDepositsInfo(id, map[string]interface{}{
		"review_status":   comm.Approval_Status_Pass,
		"deposits_status": comm.Deposits_Status_Payment,
	})
}

func (a ApprovalIncDeposits) ApprovalProcessFail(id *int) error {
	//合同审批
	//有人拒绝，修改合同状态
	return crmDepositsService.UpdDepositsInfo(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Rafuse,
	})
}
