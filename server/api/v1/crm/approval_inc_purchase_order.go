package crm

import "github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"

// 对账单管理
type ApprovalIncPurchaseOrder struct{}

func NewApprovalIncPurchaseOrder() *ApprovalIncPurchaseOrder {
	return &ApprovalIncPurchaseOrder{}
}

func (a ApprovalIncPurchaseOrder) ApprovalProcessSuccess(id *int) error {

	return crmPurchaseOrderService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Pass,
	})
}

func (a ApprovalIncPurchaseOrder) ApprovalProcessFail(id *int) error {
	//合同审批
	//有人拒绝，修改合同状态
	return crmPurchaseOrderService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Rafuse,
	})
}
