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
)

const (
	ContractApprovalType            = 1
	BusinessOpportunityApprovalType = 2
	PaymentCollentionApprovalType   = 3
	OrderApprovalType               = 4
)

const (
	PaymentStatusUnpaid = "1" //未付款
	PaymentStatusPaid   = "2" //已付款
)

// 审批类型对应的名称
var ApprovalConfigToType = map[int]string{
	ContractApprovalType:            ContractApproval,
	BusinessOpportunityApprovalType: BusinessOpportunityApproval,
	PaymentCollentionApprovalType:   PaymentCollentionApproval,
	OrderApprovalType:               OrderApproval,
}

// 默认审批通过人数
const NumberApprovedPersonnel = 1
