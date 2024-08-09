package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
)

// 回款管理
type ApprovalIncPaymentColletion struct{}

func NewApprovalIncPaymentColletion() *ApprovalIncPaymentColletion {
	return &ApprovalIncPaymentColletion{}
}

func (a ApprovalIncPaymentColletion) ApprovalProcessSuccess(id *int) error {
	//回款审批
	//所有人都同意，修改商机状态
	err := crmPaymentCollentionService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Pass,
	})
	if err != nil {
		return err
	}

	pc, err := crmPaymentCollentionService.GetCrmPaymentIdCollention(*id)
	if err != nil {
		return err
	}

	//通过审批以后，商机更新为已关单
	err = crmBusinessOpportunityService.UpdApprovalStatus(pc.BusinessOpportunityId, map[string]interface{}{
		"status": comm.BusinessOpportunityStatus,
	})
	if err != nil {
		return err
	}

	order, err := crmOrderService.GetCrmOrderId(pc.OrderId)
	if err != nil {
		return err
	}

	//通过审批以后，商机更新为已关单
	err = crmBusinessOpportunityService.UpdApprovalStatus(order.BusinessOpportunityId, map[string]interface{}{
		"status": comm.BusinessOpportunityStatus,
	})

	if err != nil {
		return err
	}

	//跟新账单状态 付款状态
	err = crmBillService.UpdApprovalStatus(pc.BillId, map[string]interface{}{
		"payment_status": comm.PaymentStatusPaid,
	})
	if err != nil {
		return err
	}

	if pc.PaymentCollentionType == comm.Pre_Deposit || pc.PaymentCollentionType == comm.Advance_charge {
		//更新用户余额 包含特定条件更新
		err = crmCustomersService.UpdateBalanceIncrease(*pc.CustomerId, *pc.Amount)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a ApprovalIncPaymentColletion) ApprovalProcessFail(id *int) error {
	//合同审批
	//有人拒绝，修改合同状态
	return crmPaymentCollentionService.UpdApprovalStatus(id, map[string]interface{}{
		"review_status": comm.Approval_Status_Rafuse,
	})
}
