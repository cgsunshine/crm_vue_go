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
	BusinessOpportunity int `json:"businessOpportunity"` //待审批商机
	Order               int `json:"order"`               //待审批订单
	SalesContract       int `json:"salesContract"`       //待审批销售合同
	PaymentCollection   int `json:"paymentCollection"`   //待审批回款
	Deposit             int `json:"deposit"`             //待审批押金
	Payment             int `json:"payment"`             //待审批付款
}

// 新增业务数量
type AddBusiness struct {
	BusinessOpportunity         int `json:"businessOpportunity"`         //商机
	BusinessOpportunityRelation int `json:"businessOpportunityRelation"` //商机关单
	Order                       int `json:"order"`                       //订单
}

// 客户
type CustomerInfo struct {
	TransactionCustomer            int `json:"transactionCustomer"`            //成交客户
	PublicPoolCustomers            int `json:"publicPoolCustomers"`            //公共池客户
	PublicPoolTransactionCustomers int `json:"publicPoolTransactionCustomers"` //公共池成交客户
}

// 合同信息
type ContactInfo struct {
	AddContact        int `json:"addContact"`        //新增合同
	PendingRenewal    int `json:"pendingRenewal"`    //待续约合同
	ExpiredNotRenewed int `json:"expiredNotRenewed"` //已过期未续约合同
}

// 业务跟进
type BusinessFollowUp struct {
	Customer          int `json:"customer"`          //已跟进客户
	PaymentCollection int `json:"paymentCollection"` //回款
}
