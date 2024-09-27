// 自动生成模板CrmPayment
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmPayment表 结构体  CrmPayment
type CrmPagePayment struct {
	global.GVA_MODEL
	BillPaymentId         *int       `json:"billPaymentId" form:"billPaymentId" gorm:"column:bill_paymentId;comment:对账单id;size:19;"`                          //应付账单id
	StatementAccountId    *int       `json:"statementAccountId" form:"statementAccountId" gorm:"column:statement_account_id;comment:对账单id;size:19;"`          //对账单id
	PaymentAmount         *float64   `json:"paymentAmount" form:"paymentAmount" gorm:"column:payment_amount;comment:付款金额;" binding:"required"`                //付款金额
	PaymentTime           *time.Time `json:"paymentTime" form:"paymentTime" gorm:"column:payment_time;comment:付款时间;"`                                         //付款时间
	PaymentVoucher        string     `json:"paymentVoucher" form:"paymentVoucher" gorm:"column:payment_voucher;comment:付款凭证;size:191;"`                       //付款凭证
	UserId                *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:销售ID;"`                                                        //销售ID
	PaymentType           string     `json:"paymentType" form:"paymentType" gorm:"column:payment_type;comment:付款方式;size:10;"`                                 //付款方式
	PaymentStatus         string     `json:"paymentStatus" form:"paymentStatus" gorm:"column:payment_status;comment:付款状态;size:191;"`                          //付款状态
	ReviewStatus          string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`                             //审核状态
	PaymentName           string     `json:"paymentName" form:"paymentName" gorm:"column:payment_name;comment:对账单名称;size:191;"`                               //付款名称
	PaymentBusinessNumber string     `json:"paymentBusinessNumber" form:"paymentBusinessNumber" gorm:"column:payment_business_number;comment:付款编号;size:191;"` //付款编号
	//以下是联表查询字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	//crm_bill_payment 表
	BusinessNumber *int `json:"businessNumber" form:"businessNumber" gorm:"column:business_number;comment:对账单名称;size:191;"` //应付账单编号
}
