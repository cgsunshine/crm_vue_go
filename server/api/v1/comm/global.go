package comm

const (
	Approval_Status_Under   = "1" //待审批
	Approval_Status_Pending = "2" //审批中
	Approval_Status_Pass    = "3" //审批通过
	Approval_Status_Rafuse  = "4" //审批拒绝
)

const (
	BusinessOpportunityStatus = "2" //商机状态
)

const (
	Contact_Approval_Tasks_valid_Effective = 1 //审批任务有效
	Contact_Approval_Tasks_Valid_Invalid   = 2 //审批任务无效
)

const (
	ContractApproval            = "合同审批"
	BusinessOpportunityApproval = "商机审批"
	PaymentCollentionApproval   = "回款审批"
	OrderApproval               = "订单审批"
	DepositsApproval            = "押金审批"
	StatementAccountApproval    = "对账单审批"
	PaymentApproval             = "付款审批"
	PurchaseOrderApproval       = "订购单审批"
	ProcurementContractApproval = "订购合同审批"
)

const (
	ContractApprovalType            = 1
	BusinessOpportunityApprovalType = 2
	PaymentCollentionApprovalType   = 3
	OrderApprovalType               = 4
	DepositsApprovalType            = 5
	StatementAccountApprovalType    = 6
	PaymentApprovalType             = 7
	PurchaseOrderApprovalType       = 8
	ProcurementContractApprovalType = 9
)

const (
	DepositsRefundType = 1 //押金
)

const (
	PaymentStatusUnpaid = "1" //未付款
	PaymentStatusPaid   = "2" //已付款
)

const (
	Deposits_Processing_Status_Unprocessed = "1" //未处理
	Deposits_Processing_Status_Processed   = "2" //已处理
)

// 审批类型对应的名称
var ApprovalConfigToType = map[int]string{
	ContractApprovalType:            ContractApproval,
	BusinessOpportunityApprovalType: BusinessOpportunityApproval,
	PaymentCollentionApprovalType:   PaymentCollentionApproval,
	OrderApprovalType:               OrderApproval,
	DepositsApprovalType:            DepositsApproval,
	StatementAccountApprovalType:    StatementAccountApproval,
	PaymentApprovalType:             PaymentApproval,
	PurchaseOrderApprovalType:       PurchaseOrderApproval,
	ProcurementContractApprovalType: ProcurementContractApproval,
}

const (
	Deposits_Status_Under   = "1" //待处理
	Deposits_Status_Payment = "2" //已支付
	Deposits_Status_Refund  = "3" //已退款
)

// 默认审批通过人数
const NumberApprovedPersonnel = 1

const (
	Cash_Payment   = "1" //现付款
	Pre_Deposit    = "2" //预存款
	Advance_charge = "3" //预付款
)

const (
	RUnsubmitted_Refund_Status = "1" //未提交
	Submitted_Refund_Status    = "2" //已提交
	Processed_Refund_Status    = "3" //已处理
)

const (
	Deposits_Processing_Status_Tickets_Unprocessed = "1" //未处理
	Deposits_Processing_Status_Tickets_Processing  = "2" //处理中
	Deposits_Processing_Status_Tickets_Processed   = "3" //处理完成
)

const (
	BusinessOpportunityRelationType  = 1 //未关单
	BusinessOpportunityUnrelatedType = 2 //已关单
)

// 合同状态
const (
	ProcurementContractStatusEffectiveType = "1" //有效
	ProcurementContractStatusInvalidType   = "2" //无效
)

// 编号前缀
const (
	PaymentNumberPrefix             = "ASF"
	BillNumberPrefix                = "ASZ"
	PaymentCollentionNumberPrefix   = "ASH"
	BillPaymentNumberPrefix         = "ASY"
	ContractNumberPrefix            = "ASC"
	BusinessOpportunityNumberPrefix = "ASSJ"
	DepositsNumberPrefix            = "ASYj"
	StatementAccountNumberPrefix    = "ASDZD"
	OrderNumberPrefix               = "ASDD"
	PurchaseOrderNumberPrefix       = "ASDGD"
	ProcurementContractNumberPrefix = "ASDGH"
)

// ascending descending
func OrderHandle(o string) string {
	if o == "ascending" {
		return "ASC"
	}
	return "DESC"
}
