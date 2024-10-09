// 自动生成模板CrmStatementAccount
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmStatementAccount表 结构体  CrmStatementAccount
type CrmStatementAccount struct {
	global.GVA_MODEL
	PurchaseOrderId *int     `json:"purchaseOrderId" form:"purchaseOrderId" gorm:"column:purchase_order_id;comment:订购单ID;"` //订购单ID
	Amount          *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:金额;" binding:"required"`              //金额
	BillFile        string   `json:"billFile" form:"billFile" gorm:"column:bill_file;comment:账单文件;size:191;"`               //账单文件
	//CreationTime         *time.Time `json:"creationTime" form:"creationTime" gorm:"column:creation_time;comment:创建时间;"`                  //创建时间
	PayableTime            *time.Time `json:"payableTime" form:"payableTime" gorm:"column:payable_time;comment:应付时间;"`                                             //应付时间
	PaymentStatus          string     `json:"paymentStatus" form:"paymentStatus" gorm:"column:payment_status;comment:付款状态;size:191;"`                              //付款状态
	UserId                 *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:负责人;"`                                                             //负责人
	ReviewStatus           string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`                                 //审核状态
	StatementAccountName   string     `json:"statementAccountName" form:"statementAccountName" gorm:"column:statement_account_name;comment:对账单名称;size:191;"`       //对账单名称
	StatementAccountNumber string     `json:"statementAccountNumber" form:"statementAccountNumber" gorm:"column:statement_account_number;comment:对账单编号;size:191;"` //对账单编号
	Currency               string     `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:11;"`                                                 //币种
	Invoice                string     `json:"invoice" form:"invoice" gorm:"column:invoice;comment:发票;size:11;"`                                                    //发票
}

// TableName crmStatementAccount表 CrmStatementAccount自定义表名 crm_statement_account
func (CrmStatementAccount) TableName() string {
	return "crm_statement_account"
}
