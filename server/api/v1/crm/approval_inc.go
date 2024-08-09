package crm

import "github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"

type ApprovalInc interface {
	ApprovalProcessSuccess(id *int) error //审批成功
	ApprovalProcessFail(id *int) error    //审批失败
}

var ApprovalEntrance = map[int]ApprovalInc{
	comm.ContractApprovalType:            NewApprovalIncContract(),
	comm.BusinessOpportunityApprovalType: NewApprovalIncBusinessOpportunity(),
	comm.PaymentCollentionApprovalType:   NewApprovalIncPaymentColletion(),
	comm.OrderApprovalType:               NewApprovalIncOrder(),
	comm.DepositsApprovalType:            NewApprovalIncDeposits(),
	comm.StatementAccountApprovalType:    NewApprovalIncStatementAccount(),
	comm.PaymentApprovalType:             NewApprovalIncPayment(),
}
