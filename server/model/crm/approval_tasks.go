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
	//crm_contract 表 合同审批状态
	ContractName string `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合同名称;size:191;"` //合同名称
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
	//crm_business_opportunity 表 合同审批状态
	BusinessOpportunityName string `json:"businessOpportunityName" form:"businessOpportunityName" gorm:"column:business_opportunity_name;comment:商机名称;size:191;" binding:"required"` //商机名称
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
	//crm_payment_collention 表 合同审批状态
	PaymentCollentionName string `json:"paymentCollentionName" form:"paymentCollentionName" gorm:"column:payment_collention_name;comment:凭证;size:191;"` //回款名称
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
	//crm_order 表 合同审批状态
	OrderName string `json:"orderName" form:"orderName" gorm:"column:order_name;comment:订单名称;size:10;" binding:"required"` //订单名称
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
	//crm_business_opportunity 表 合同审批状态
	DepositsName string `json:"depositsName" form:"depositsName" gorm:"column:deposits_name;comment:押金名称;size:191;"` //押金名称
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
	//crm_statement_account 表 合同审批状态
	StatementAccountName string `json:"statementAccountName" form:"statementAccountName" gorm:"column:statement_account_name;comment:对账单名称;size:191;"` //对账单名称
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
	//crm_payment 表 合同审批状态
	PaymentName string `json:"paymentName" form:"paymentName" gorm:"column:payment_name;comment:对账单名称;size:191;"` //付款名称
}
