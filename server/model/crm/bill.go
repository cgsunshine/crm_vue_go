// 自动生成模板CrmBill
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmBill表 结构体  CrmBill
type CrmPageBill struct {
	global.GVA_MODEL

	OrderId             *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID;"`                                       //订单ID
	PaymentCollentionId *int       `json:"paymentCollentionId" form:"paymentCollentionId" gorm:"column:payment_collention_id;comment:回款单ID;"` //回款单ID
	PaymentId           *int       `json:"PaymentId" form:"PaymentId" gorm:"column:payment_id;comment:付款ID;"`                                 //付款ID
	Amount              *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`                                             //金额
	BalanceConsumption  *float64   `json:"balanceConsumption" form:"balanceConsumption" gorm:"column:balance_consumption;comment:余额消费;"`      //余额消费
	Currency            string     `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:11;"`                               //币种
	UserId              *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;size:19;"`                             //管理ID 销售代表
	BillNumber          string     `json:"billNumber" form:"billNumber" gorm:"column:bill_number;comment:账单编号;size:191;"`                     //账单编号
	PaymentStatus       string     `json:"paymentStatus" form:"paymentStatus" gorm:"column:payment_status;comment:付款状态;size:191;"`            //付款状态
	ExpirationTime      *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:到期时间;"`                  //到期时间
	BillName            string     `json:"billName" form:"billName" gorm:"column:bill_name;comment:账单名称;size:191;"`                           //账单名称
	CustomerId          *int       `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;size:19;"`                      //客户ID

	//以下是联表查询字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	//crm_customers 表
	CustomerName    string `json:"customerName" form:"customerName" gorm:"column:customer_name;comment:客户名;size:191;"`           //客户名
	CustomerCompany string `json:"customerCompany" form:"customerCompany" gorm:"column:customer_company;comment:客户公司;size:191;"` //客户公司
	CustomerAddress string `json:"customerAddress" form:"customerAddress" gorm:"column:customer_address;comment:客户地址;size:191;"` //客户地址
	//crm_order 表
	OrderName        string     `json:"orderName" form:"orderName" gorm:"column:order_name;comment:订单名称;size:10;"`               //订单名称
	BillingStartTime *time.Time `json:"billingStartTime" form:"billingStartTime" gorm:"column:billing_start_time;comment:到期时间;"` //计费开始时间
	BillingEndTime   *time.Time `json:"billingEndTime" form:"billingEndTime" gorm:"column:billing_end_time;comment:到期时间;"`       //计费结束时间
	ProductId        string     `json:"productId" form:"productId" gorm:"column:product_id;comment:产品id;size:191;"`              //产品id
	DiscountRate     *float64   `json:"discountRate" form:"discountRate" gorm:"column:discount_rate;comment:折扣率;"`               //折扣率
	OrderNumber      string     `json:"orderNumber" form:"orderNumber" gorm:"column:order_number;comment:订单编号;size:191;"`        //订单编号
	//crm_payment_collention 表
	PaymentCollentionName string     `json:"paymentCollentionName" form:"paymentCollentionName" gorm:"column:payment_collention_name;comment:回款名称;size:191;"` //回款名称
	PaymentType           string     `json:"paymentType" form:"paymentType" gorm:"column:payment_type;comment:付款方式;size:10;"`                                 //回款方式
	PaymentTime           *time.Time `json:"paymentTime" form:"paymentTime" gorm:"column:payment_time;comment:付款时间;"`                                         //回款时间
}
