package custom

type CrmHomeData struct {
	//待审批部分
	PendingApprovalTask *PendingApprovalTask `json:"pendingApprovalTask"` //待审批商机
	AddBusiness         *AddBusiness         `json:"addBusiness"`
	CustomerInfo        *CustomerInfo        `json:"customerInfo"`
	ContactInfo         *ContactInfo         `json:"contactInfo"`
	BusinessFollowUp    *BusinessFollowUp    `json:"businessFollowUp"`
	PerformanceTarget   int                  `json:"performanceTarget"` //目标达成率
}

// 待审批数量
type PendingApprovalTask struct {
	BusinessOpportunity int64 `json:"businessOpportunity"` //待审批商机
	Order               int64 `json:"order"`               //待审批订单
	SalesContract       int64 `json:"salesContract"`       //待审批销售合同
	PaymentCollection   int64 `json:"paymentCollection"`   //待审批回款
	Deposit             int64 `json:"deposit"`             //待审批押金
	Payment             int64 `json:"payment"`             //待审批付款
}

// 新增业务数量
type AddBusiness struct {
	BusinessOpportunity         int64 `json:"businessOpportunity"`         //商机
	BusinessOpportunityRelation int64 `json:"businessOpportunityRelation"` //商机关单
	Order                       int64 `json:"order"`                       //订单
}

// 客户
type CustomerInfo struct {
	TransactionCustomer            int64 `json:"transactionCustomer"`            //成交客户
	PublicPoolCustomers            int64 `json:"publicPoolCustomers"`            //公共池客户
	PublicPoolTransactionCustomers int64 `json:"publicPoolTransactionCustomers"` //公共池成交客户
}

// 合同信息
type ContactInfo struct {
	AddContact        int64 `json:"addContact"`        //新增合同
	PendingRenewal    int64 `json:"pendingRenewal"`    //待续约合同
	ExpiredNotRenewed int64 `json:"expiredNotRenewed"` //已过期未续约合同
}

// 业务跟进
type BusinessFollowUp struct {
	Customer          int64 `json:"customer"`          //已跟进客户
	PaymentCollection int64 `json:"paymentCollection"` //回款
}
