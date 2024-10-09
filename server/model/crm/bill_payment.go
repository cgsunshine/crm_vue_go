// 自动生成模板CrmBillPayment
package crm

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// 应付账单 结构体  CrmBillPayment
type CrmPageBillPayment struct {
	global.GVA_MODEL
	Amount             *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:金额;size:22;" binding:"required"`                       //金额
	BillPaymentName    string   `json:"billPaymentName" form:"billPaymentName" gorm:"column:bill_payment_name;comment:应付账单名称;size:191;"`        //应付账单名称
	BillPaymentNumber  string   `json:"billPaymentNumber" form:"billPaymentNumber" gorm:"column:bill_payment_number;comment:应付账单名称;size:191;"`  //应付账单编号
	Currency           string   `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:11;"`                                    //币种
	CustomerId         *int     `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;"`                                   //客户ID
	PaymentStatus      string   `json:"paymentStatus" form:"paymentStatus" gorm:"column:payment_status;comment:付款状态;size:191;"`                 //付款状态
	StatementAccountId *int     `json:"statementAccountId" form:"statementAccountId" gorm:"column:statement_account_id;comment:对账单ID;size:19;"` //对账单ID
	UserId             *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:19;"`                                       //用户id
	PaymentId          *int     `json:"paymentId" form:"paymentId" gorm:"column:payment_id;comment:对账单ID;size:19;"`                             //付款管理ID
	//联表查询字段 crm_statement_account
	StatementAccountName   string `json:"statementAccountName" form:"statementAccountName" gorm:"column:statement_account_name;comment:对账单名称;size:191;"`       //对账单名称
	StatementAccountNumber string `json:"statementAccountNumber" form:"statementAccountNumber" gorm:"column:statement_account_number;comment:对账单编号;size:191;"` //对账单编号

}
