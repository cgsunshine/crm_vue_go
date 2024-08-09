package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
)

// 订单审批
type ApprovalIncOrder struct{}

func NewApprovalIncOrder() *ApprovalIncOrder {
	return &ApprovalIncOrder{}
}

func (a ApprovalIncOrder) ApprovalProcessSuccess(id *int) error {
	//订单审批
	//所有人都同意，修改订单状态
	return crmOrderService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Pass,
	})
}

func (a ApprovalIncOrder) ApprovalProcessFail(id *int) error {
	//合同审批
	//有人拒绝，修改合同状态
	return crmOrderService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Rafuse,
	})
}
