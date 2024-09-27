// 自动生成模板CrmPaymentCollention
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// CrmPagePaymentCollention 表 结构体  CrmPaymentCollention
type CrmPagePaymentCollention struct {
	global.GVA_MODEL

	OrderId                 *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID;"`                                                           //订单ID
	CustomerId              *int       `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;size:19;"`                                          //客户ID
	UserId                  *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;size:19;"`                                                 //管理ID 销售代表
	Description             string     `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`                                         //备注
	Currency                string     `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:10;"`                                                   //币种
	Amount                  *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`                                                                 //金额
	Proof                   string     `json:"proof" form:"proof" gorm:"column:proof;comment:凭证;size:191;"`                                                           //凭证
	ReviewStatus            string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`                                   //审核状态
	ApprovedById            string     `json:"approvedById" form:"approvedById" gorm:"column:approved_by_id;comment:审批人;size:191;"`                                   //审批人
	AuditingTime            *time.Time `json:"auditingTime" form:"auditingTime" gorm:"column:auditing_time;comment:审核时间;"`                                            //审核时间
	PaymentCollentionName   string     `json:"paymentCollentionName" form:"paymentCollentionName" gorm:"column:payment_collention_name;comment:回款名称;size:191;"`       //回款名称
	PaymentCollentionNumber string     `json:"paymentCollentionNumber" form:"paymentCollentionNumber" gorm:"column:payment_collention_number;comment:回款名称;size:191;"` //回款编号
	BusinessOpportunityId   *int       `json:"businessOpportunityId" form:"businessOpportunityId" gorm:"column:business_opportunity_id;comment:商机id;"`                //商机id
	CurrencyId              *int       `json:"currencyId" form:"currencyId" gorm:"column:currency_id;comment:币种id;size:11;"`                                          //币种ID
	PaymentTime             *time.Time `json:"paymentTime" form:"paymentTime" gorm:"column:payment_time;comment:付款时间;"`                                               //回款时间
	PaymentType             string     `json:"paymentType" form:"paymentType" gorm:"column:payment_type;comment:付款方式;size:10;"`                                       //回款方式
	BillId                  *int       `json:"billId" form:"billId" gorm:"column:bill_id;comment:账单id;"`                                                              //账单id
	PaymentCollentionType   string     `json:"paymentCollentionType" form:"paymentCollentionType" gorm:"column:payment_collention_type;comment:审核状态;size:191;"`       //回款类型
	//以下是联表查询字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	//crm_customers 表
	CustomerName string `json:"customerName" form:"customerName" gorm:"column:customer_name;comment:客户名;size:191;"` //客户名
	//crm_order
	Price     *float64 `json:"price" form:"price" gorm:"column:price;comment:产品原价;"`                                         //产品原价
	OrderName string   `json:"orderName" form:"orderName" gorm:"column:order_name;comment:订单名称;size:10;" binding:"required"` //订单名称
	//crm_business_opportunity 表
	BusinessOpportunityName string `json:"businessOpportunityName" form:"businessOpportunityName" gorm:"column:business_opportunity_name;comment:商机名称;size:191;" binding:"required"` //商机名称
	//crm_currency 表 后面替换去掉注释
	CurrencyName string `json:"currencyName" form:"currencyName" gorm:"column:currency_name;comment:币种;size:10;"` //币种名称
	//crm_bill
	BillName string `json:"billName" form:"billName" gorm:"column:bill_name;comment:账单名称;size:191;"` //账单名称
}
