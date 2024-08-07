package crm

import "github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"

// 对账单管理
type ApprovalIncStatementAccount struct{}

func NewApprovalIncStatementAccount() *ApprovalIncStatementAccount {
	return &ApprovalIncStatementAccount{}
}

func (a ApprovalIncStatementAccount) ApprovalProcessSuccess(id *int) error {

	return crmStatementAccountService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Pass,
	})
}

func (a ApprovalIncStatementAccount) ApprovalProcessFail(id *int) error {
	//合同审批
	//有人拒绝，修改合同状态
	return crmStatementAccountService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Rafuse,
	})
}
