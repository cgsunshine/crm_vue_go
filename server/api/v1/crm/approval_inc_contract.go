package crm

import "github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"

type ApprovalIncContract struct {
}

func NewApprovalIncContract() *ApprovalIncContract {
	return &ApprovalIncContract{}
}

func (a ApprovalIncContract) ApprovalProcessSuccess(id *int) error {
	return crmContractService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Pass,
	})
}

func (a ApprovalIncContract) ApprovalProcessFail(id *int) error {
	//合同审批
	//有人拒绝，修改合同状态
	return crmContractService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Rafuse,
	})
}
