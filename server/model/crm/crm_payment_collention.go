// 自动生成模板CrmPaymentCollention
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmPaymentCollention表 结构体  CrmPaymentCollention
type CrmPaymentCollention struct {
	global.GVA_MODEL
	ApprovedById          string     `json:"approvedById" form:"approvedById" gorm:"column:approved_by_id;comment:审批人;size:191;"`                             //审批人
	ReviewStatus          string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`                             //审核状态
	AuditingTime          *time.Time `json:"auditingTime" form:"auditingTime" gorm:"column:auditing_time;comment:审核时间;"`                                      //审核时间
	OrderId               *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID;"`                                                     //订单ID
	Currency              string     `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:10;"`                                             //币种
	Amount                *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`                                                           //金额
	CustomerId            *int       `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;"`                                            //客户ID
	Description           string     `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`                                   //备注
	Proof                 string     `json:"proof" form:"proof" gorm:"column:proof;comment:凭证;size:191;"`                                                     //凭证
	UserId                *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;"`                                                   //管理ID 销售代表
	PaymentCollentionName string     `json:"paymentCollentionName" form:"paymentCollentionName" gorm:"column:payment_collention_name;comment:回款名称;size:191;"` //回款名称
	BusinessOpportunityId *int       `json:"businessOpportunityId" form:"businessOpportunityId" gorm:"column:business_opportunity_id;comment:商机id;"`          //商机id
	PaymentTime           *time.Time `json:"paymentTime" form:"paymentTime" gorm:"column:payment_time;comment:付款时间;"`                                         //回款时间
	PaymentType           string     `json:"paymentType" form:"paymentType" gorm:"column:payment_type;comment:付款方式;size:10;"`                                 //回款方式
	BillId                *int       `json:"billId" form:"billId" gorm:"column:bill_id;comment:账单id;"`                                                        //账单id
}

// TableName crmPaymentCollention表 CrmPaymentCollention自定义表名 crm_payment_collention
func (CrmPaymentCollention) TableName() string {
	return "crm_payment_collention"
}
