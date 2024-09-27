// 自动生成模板CrmPayment
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmPayment表 结构体  CrmPaymentb
type CrmPayment struct {
	global.GVA_MODEL
	BillPaymentId      *int       `json:"billPaymentId" form:"billPaymentId" gorm:"column:bill_payment_id;comment:对账单id;size:19;"`                //应付账单id
	StatementAccountId *int       `json:"statementAccountId" form:"statementAccountId" gorm:"column:statement_account_id;comment:对账单id;size:19;"` //对账单id
	PaymentAmount      *float64   `json:"paymentAmount" form:"paymentAmount" gorm:"column:payment_amount;comment:付款金额;"`                          //付款金额
	PaymentTime        *time.Time `json:"paymentTime" form:"paymentTime" gorm:"column:payment_time;comment:付款时间;"`                                //付款时间
	PaymentVoucher     string     `json:"paymentVoucher" form:"paymentVoucher" gorm:"column:payment_voucher;comment:付款凭证;size:191;"`              //付款凭证
	UserId             *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:销售ID;"`                                               //销售ID
	PaymentType        string     `json:"paymentType" form:"paymentType" gorm:"column:payment_type;comment:付款方式;size:10;"`                        //付款方式
	PaymentStatus      string     `json:"paymentStatus" form:"paymentStatus" gorm:"column:payment_status;comment:付款状态;size:191;"`                 //付款状态
	ReviewStatus       string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`                    //审核状态
	PaymentName        string     `json:"paymentName" form:"paymentName" gorm:"column:payment_name;comment:对账单名称;size:191;"`                      //付款名称
	PaymentNumber      string     `json:"paymentNumber" form:"paymentNumber" gorm:"column:payment_number;comment:付款编号;size:191;"`                 //付款编号
}

// TableName crmPayment表 CrmPayment自定义表名 crm_payment
func (CrmPayment) TableName() string {
	return "crm_payment"
}
