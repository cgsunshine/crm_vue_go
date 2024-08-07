package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
)

// 回款管理
type ApprovalIncPayment struct{}

func NewApprovalIncPayment() *ApprovalIncPayment {
	return &ApprovalIncPayment{}
}

func (a ApprovalIncPayment) ApprovalProcessSuccess(id *int) error {
	//回款审批
	//所有人都同意，修改商机状态
	err := crmPaymentService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Pass,
	})
	if err != nil {
		return err
	}

	//pc, err := crmPaymentService.GetCrmPaymentIdInfo(id)
	//if err != nil {
	//	return err
	//}

	////通过审批以后，应付账单更新为已付款 暂时手动修改
	//err = crmBillPaymentService.UpdApprovalStatus(pc.StatementAccountId, map[string]interface{}{
	//	"status": comm.BusinessOpportunityStatus,
	//})
	//if err != nil {
	//	return err
	//}

	return nil
}

func (a ApprovalIncPayment) ApprovalProcessFail(id *int) error {
	//付款审批
	//有人拒绝，修改付款状态
	return crmPaymentService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Rafuse,
	})
}
