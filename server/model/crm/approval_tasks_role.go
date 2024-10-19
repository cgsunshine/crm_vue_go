// 自动生成模板CrmApprovalTasksRole
package crm

// 审批任务 合同 结构体  CrmApprovalTasks
type CrmContactInfoApprovalTasksRole struct {
	CrmApprovalTasksRole
	//联表查询
	//crm_contract 表
	ContractName   string `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合同名称;size:191;"`       //合同名称
	ContractNumber string `json:"contractNumber" form:"contractNumber" gorm:"column:contract_number;comment:合同编号;size:191;"` //合同编号
	//sys_authorities 表
	AuthorityName string `json:"authorityName" gorm:"comment:角色名"` // 角色名
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名

}

// 审批任务 商机 结构体  CrmApprovalTasks
type CrmBusinessOpportunityInfoApprovalTasksRole struct {
	CrmApprovalTasksRole
	//联表查询
	//crm_business_opportunity 表
	BusinessOpportunityName   string `json:"businessOpportunityName" form:"businessOpportunityName" gorm:"column:business_opportunity_name;comment:商机名称;size:191;" binding:"required"` //商机名称
	AuthorityName             string `json:"authorityName" gorm:"comment:角色名"`                                                                                                         // 角色名
	BusinessOpportunityNumber string `json:"businessOpportunityNumber" form:"businessOpportunityNumber" gorm:"column:business_opportunity_number;comment:商机编号;size:191;"`              //商机编号
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
}

// 审批任务 回款 结构体  CrmApprovalTasks
type CrmPaymentCollentionInfoApprovalTasksRole struct {
	CrmApprovalTasksRole
	//联表查询
	//crm_payment_collention 表
	PaymentCollentionName   string `json:"paymentCollentionName" form:"paymentCollentionName" gorm:"column:payment_collention_name;comment:凭证;size:191;"`         //回款名称
	AuthorityName           string `json:"authorityName" gorm:"comment:角色名"`                                                                                      // 角色名
	PaymentCollentionNumber string `json:"paymentCollentionNumber" form:"paymentCollentionNumber" gorm:"column:payment_collention_number;comment:回款名称;size:191;"` //回款编号
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
}

// CrmOrderInfoApprovalTasks 审批任务 订单 结构体
type CrmOrderInfoApprovalTasksRole struct {
	CrmApprovalTasksRole
	//联表查询
	//crm_order 表
	OrderName     string `json:"orderName" form:"orderName" gorm:"column:order_name;comment:订单名称;size:10;" binding:"required"` //订单名称
	AuthorityName string `json:"authorityName" gorm:"comment:角色名"`                                                             // 角色名
	OrderNumber   string `json:"orderNumber" form:"orderNumber" gorm:"column:order_number;comment:订单编号;size:191;"`             //订单编号
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
}

// CrmDepositsInfoApprovalTasks 审批任务 押金 结构体
type CrmDepositsInfoApprovalTasksRole struct {
	CrmApprovalTasksRole
	//联表查询
	//crm_business_opportunity 表
	DepositsName   string `json:"depositsName" form:"depositsName" gorm:"column:deposits_name;comment:押金名称;size:191;"`        //押金名称
	AuthorityName  string `json:"authorityName" gorm:"comment:角色名"`                                                           // 角色名
	DepositsNumber string `json:"depositsNumber" form:"depositsNumbere" gorm:"column:deposits_number;comment:押金编号;size:191;"` //押金编号
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
}

// CrmDepositsInfoApprovalTasks 审批任务 对账单 结构体
type CrmStatementAccountInfoApprovalTasksRole struct {
	CrmApprovalTasksRole
	//联表查询
	//crm_statement_account 表
	StatementAccountName   string `json:"statementAccountName" form:"statementAccountName" gorm:"column:statement_account_name;comment:对账单名称;size:191;"`       //对账单名称
	AuthorityName          string `json:"authorityName" gorm:"comment:角色名"`                                                                                    // 角色名
	StatementAccountNumber string `json:"statementAccountNumber" form:"statementAccountNumber" gorm:"column:statement_account_number;comment:对账单编号;size:191;"` //对账单编号
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
}

// CrmDepositsInfoApprovalTasks 审批任务 付款 结构体
type CrmPaymentInfoApprovalTasksRole struct {
	CrmApprovalTasksRole
	//联表查询
	//crm_payment 表
	PaymentName   string `json:"paymentName" form:"paymentName" gorm:"column:payment_name;comment:对账单名称;size:191;"`      //付款名称
	AuthorityName string `json:"authorityName" gorm:"comment:角色名"`                                                       // 角色名
	PaymentNumber string `json:"paymentNumber" form:"paymentNumber" gorm:"column:payment_number;comment:付款编号;size:191;"` //付款编号
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
}

// 订购单管理 结构体  CrmPurchaseOrder
type CrmPurchaseOrderInfoApprovalTasksRole struct {
	CrmApprovalTasksRole
	//联表查询
	//crm_purchase_order 表
	PurchaseOrderName   string `json:"purchaseOrderName" form:"purchaseOrderName" gorm:"column:purchase_order_name;comment:订购单名称;size:191;"`       //订购单名称
	AuthorityName       string `json:"authorityName" gorm:"comment:角色名"`                                                                           // 角色名
	PurchaseOrderNumber string `json:"purchaseOrderNumber" form:"purchaseOrderNumber" gorm:"column:purchase_order_number;comment:订购单编号;size:191;"` //订购单编号
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
}

// 订购合同 结构体  CrmProcurementContract
type CrmProcurementContractInfoApprovalTasksRole struct {
	CrmApprovalTasksRole
	//联表查询
	//crm_procurement_contract 表
	ContractName              string `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合同名称;size:191;" binding:"required"`                        //合同名称
	AuthorityName             string `json:"authorityName" gorm:"comment:角色名"`                                                                                              // 角色名
	ProcurementContractNumber string `json:"procurementContractNumber" form:"procurementContractNumber" gorm:"column:procurement_contract_number;comment:订购合同编号;size:191;"` //订购合同编号
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
}
