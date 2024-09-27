package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	"go.uber.org/zap"
)

// 对账单管理
type ApprovalIncStatementAccount struct{}

func NewApprovalIncStatementAccount() *ApprovalIncStatementAccount {
	return &ApprovalIncStatementAccount{}
}

func (a ApprovalIncStatementAccount) ApprovalProcessSuccess(id *int) error {

	crmStatementAccount, err := crmStatementAccountService.GetCrmStatementAccountId(id)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		return err
	}
	crmBillPayment := crm.CrmBillPayment{
		Amount:             crmStatementAccount.Amount,
		BillPaymentName:    "",
		Currency:           crmStatementAccount.Currency, //缺个币种
		PaymentStatus:      comm.PaymentStatusUnpaid,
		StatementAccountId: id,
		UserId:             crmStatementAccount.UserId,
		BillPaymentNumber:  comm.GetBusinessNumber(comm.BillPaymentNumberPrefix, crmBillPaymentService.GetCrmBillPaymentTodayCount()),
	}

	if err := crmBillPaymentService.CreateCrmBillPayment(&crmBillPayment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		return err
	}

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
