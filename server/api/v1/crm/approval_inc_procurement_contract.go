package crm

import "github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"

// 对账单管理
type ApprovalIncProcurementContract struct{}

func NewApprovalIncProcurementContract() *ApprovalIncProcurementContract {
	return &ApprovalIncProcurementContract{}
}

func (a ApprovalIncProcurementContract) ApprovalProcessSuccess(id *int) error {

	return crmProcurementContractService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Pass,
	})
}

func (a ApprovalIncProcurementContract) ApprovalProcessFail(id *int) error {
	//合同审批
	//有人拒绝，修改合同状态
	return crmProcurementContractService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Rafuse,
	})
}
