// 自动生成模板CrmApprovalTasks
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 审批任务 合同 结构体  CrmApprovalTasks
type CrmContactInfoApprovalTasks struct {
	global.GVA_MODEL
	ApprovalStatus string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`      //审批状态
	ApprovalType   *int   `json:"approvalType" form:"approvalType" gorm:"column:approval_type;comment:审批类型 1合同 2商机 3回款;size:10;"` //审批类型 1合同 2商机 3回款
	AssigneeId     *int   `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;size:19;"`                //指派审批人id
	AssociatedId   *int   `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;size:10;"`    //关联id 合同 商机 回款
	Comments       string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`                         //审批意见
	RequestId      *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`                         //关联的审批请求ID
	StepId         *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`                                   //当前审批步骤ID
	Valid          *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:19;"`           //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
	//联表查询
	//crm_contract 表
	ContractName   string `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合同名称;size:191;"`       //合同名称
	ContractNumber string `json:"contractNumber" form:"contractNumber" gorm:"column:contract_number;comment:合同编号;size:191;"` //合同编号
}

// 审批任务 商机 结构体  CrmApprovalTasks
type CrmBusinessOpportunityInfoApprovalTasks struct {
	global.GVA_MODEL
	ApprovalStatus string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`      //审批状态
	ApprovalType   *int   `json:"approvalType" form:"approvalType" gorm:"column:approval_type;comment:审批类型 1合同 2商机 3回款;size:10;"` //审批类型 1合同 2商机 3回款
	AssigneeId     *int   `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;size:19;"`                //指派审批人id
	AssociatedId   *int   `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;size:10;"`    //关联id 合同 商机 回款
	Comments       string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`                         //审批意见
	RequestId      *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`                         //关联的审批请求ID
	StepId         *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`                                   //当前审批步骤ID
	Valid          *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:19;"`           //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
	//联表查询
	//crm_business_opportunity 表
	BusinessOpportunityName   string `json:"businessOpportunityName" form:"businessOpportunityName" gorm:"column:business_opportunity_name;comment:商机名称;size:191;" binding:"required"` //商机名称
	BusinessOpportunityNumber string `json:"businessOpportunityNumber" form:"businessOpportunityNumber" gorm:"column:business_opportunity_number;comment:商机编号;size:191;"`              //商机编号
}

// 审批任务 回款 结构体  CrmApprovalTasks
type CrmPaymentCollentionInfoApprovalTasks struct {
	global.GVA_MODEL
	ApprovalStatus string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`      //审批状态
	ApprovalType   *int   `json:"approvalType" form:"approvalType" gorm:"column:approval_type;comment:审批类型 1合同 2商机 3回款;size:10;"` //审批类型 1合同 2商机 3回款
	AssigneeId     *int   `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;size:19;"`                //指派审批人id
	AssociatedId   *int   `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;size:10;"`    //关联id 合同 商机 回款
	Comments       string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`                         //审批意见
	RequestId      *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`                         //关联的审批请求ID
	StepId         *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`                                   //当前审批步骤ID
	Valid          *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:19;"`           //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
	//联表查询
	//crm_payment_collention 表
	PaymentCollentionName   string `json:"paymentCollentionName" form:"paymentCollentionName" gorm:"column:payment_collention_name;comment:凭证;size:191;"`         //回款名称
	PaymentCollentionNumber string `json:"paymentCollentionNumber" form:"paymentCollentionNumber" gorm:"column:payment_collention_number;comment:回款名称;size:191;"` //回款编号
}

// CrmOrderInfoApprovalTasks 审批任务 订单 结构体
type CrmOrderInfoApprovalTasks struct {
	global.GVA_MODEL
	ApprovalStatus string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`      //审批状态
	ApprovalType   *int   `json:"approvalType" form:"approvalType" gorm:"column:approval_type;comment:审批类型 1合同 2商机 3回款;size:10;"` //审批类型 1合同 2商机 3回款
	AssigneeId     *int   `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;size:19;"`                //指派审批人id
	AssociatedId   *int   `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;size:10;"`    //关联id 合同 商机 回款
	Comments       string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`                         //审批意见
	RequestId      *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`                         //关联的审批请求ID
	StepId         *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`                                   //当前审批步骤ID
	Valid          *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:19;"`           //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
	//联表查询
	//crm_order 表
	OrderName   string `json:"orderName" form:"orderName" gorm:"column:order_name;comment:订单名称;size:10;" binding:"required"` //订单名称
	OrderNumber string `json:"orderNumber" form:"orderNumber" gorm:"column:order_number;comment:订单编号;size:191;"`             //订单编号
}

// CrmDepositsInfoApprovalTasks 审批任务 押金 结构体
type CrmDepositsInfoApprovalTasks struct {
	global.GVA_MODEL
	ApprovalStatus string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`      //审批状态
	ApprovalType   *int   `json:"approvalType" form:"approvalType" gorm:"column:approval_type;comment:审批类型 1合同 2商机 3回款;size:10;"` //审批类型 1合同 2商机 3回款
	AssigneeId     *int   `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;size:19;"`                //指派审批人id
	AssociatedId   *int   `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;size:10;"`    //关联id 合同 商机 回款
	Comments       string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`                         //审批意见
	RequestId      *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`                         //关联的审批请求ID
	StepId         *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`                                   //当前审批步骤ID
	Valid          *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:19;"`           //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
	//联表查询
	//crm_business_opportunity 表
	DepositsName   string `json:"depositsName" form:"depositsName" gorm:"column:deposits_name;comment:押金名称;size:191;"`        //押金名称
	DepositsNumber string `json:"depositsNumber" form:"depositsNumbere" gorm:"column:deposits_number;comment:押金编号;size:191;"` //押金编号
}

// CrmDepositsInfoApprovalTasks 审批任务 对账单 结构体
type CrmStatementAccountInfoApprovalTasks struct {
	global.GVA_MODEL
	ApprovalStatus string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`      //审批状态
	ApprovalType   *int   `json:"approvalType" form:"approvalType" gorm:"column:approval_type;comment:审批类型 1合同 2商机 3回款;size:10;"` //审批类型 1合同 2商机 3回款
	AssigneeId     *int   `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;size:19;"`                //指派审批人id
	AssociatedId   *int   `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;size:10;"`    //关联id 合同 商机 回款
	Comments       string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`                         //审批意见
	RequestId      *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`                         //关联的审批请求ID
	StepId         *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`                                   //当前审批步骤ID
	Valid          *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:19;"`           //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
	//联表查询
	//crm_statement_account 表
	StatementAccountName   string `json:"statementAccountName" form:"statementAccountName" gorm:"column:statement_account_name;comment:对账单名称;size:191;"`       //对账单名称
	StatementAccountNumber string `json:"statementAccountNumber" form:"statementAccountNumber" gorm:"column:statement_account_number;comment:对账单编号;size:191;"` //对账单编号
}

// CrmDepositsInfoApprovalTasks 审批任务 付款 结构体
type CrmPaymentInfoApprovalTasks struct {
	global.GVA_MODEL
	ApprovalStatus string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`      //审批状态
	ApprovalType   *int   `json:"approvalType" form:"approvalType" gorm:"column:approval_type;comment:审批类型 1合同 2商机 3回款;size:10;"` //审批类型 1合同 2商机 3回款
	AssigneeId     *int   `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;size:19;"`                //指派审批人id
	AssociatedId   *int   `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;size:10;"`    //关联id 合同 商机 回款
	Comments       string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`                         //审批意见
	RequestId      *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`                         //关联的审批请求ID
	StepId         *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`                                   //当前审批步骤ID
	Valid          *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:19;"`           //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
	//联表查询
	//crm_payment 表
	PaymentName   string `json:"paymentName" form:"paymentName" gorm:"column:payment_name;comment:对账单名称;size:191;"`      //付款名称
	PaymentNumber string `json:"paymentNumber" form:"paymentNumber" gorm:"column:payment_number;comment:付款编号;size:191;"` //付款编号
}

// 订购单管理 结构体  CrmPurchaseOrder
type CrmPurchaseOrderInfoApprovalTasks struct {
	global.GVA_MODEL
	ApprovalStatus string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`      //审批状态
	ApprovalType   *int   `json:"approvalType" form:"approvalType" gorm:"column:approval_type;comment:审批类型 1合同 2商机 3回款;size:10;"` //审批类型 1合同 2商机 3回款
	AssigneeId     *int   `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;size:19;"`                //指派审批人id
	AssociatedId   *int   `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;size:10;"`    //关联id 合同 商机 回款
	Comments       string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`                         //审批意见
	RequestId      *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`                         //关联的审批请求ID
	StepId         *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`                                   //当前审批步骤ID
	Valid          *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:19;"`           //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
	//联表查询
	//crm_purchase_order 表
	PurchaseOrderName   string `json:"purchaseOrderName" form:"purchaseOrderName" gorm:"column:purchase_order_name;comment:订购单名称;size:191;"`       //订购单名称
	PurchaseOrderNumber string `json:"purchaseOrderNumber" form:"purchaseOrderNumber" gorm:"column:purchase_order_number;comment:订购单编号;size:191;"` //订购单编号
}

// 订购合同 结构体  CrmProcurementContract
type CrmProcurementContractInfoApprovalTasks struct {
	global.GVA_MODEL
	ApprovalStatus string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`      //审批状态
	ApprovalType   *int   `json:"approvalType" form:"approvalType" gorm:"column:approval_type;comment:审批类型 1合同 2商机 3回款;size:10;"` //审批类型 1合同 2商机 3回款
	AssigneeId     *int   `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;size:19;"`                //指派审批人id
	AssociatedId   *int   `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;size:10;"`    //关联id 合同 商机 回款
	Comments       string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`                         //审批意见
	RequestId      *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`                         //关联的审批请求ID
	StepId         *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`                                   //当前审批步骤ID
	Valid          *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:19;"`           //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
	//联表查询
	//crm_procurement_contract 表
	ContractName              string `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合同名称;size:191;" binding:"required"`                        //合同名称
	ProcurementContractNumber string `json:"procurementContractNumber" form:"procurementContractNumber" gorm:"column:procurement_contract_number;comment:订购合同编号;size:191;"` //订购合同编号
}
