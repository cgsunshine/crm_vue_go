// 自动生成模板CrmBill
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmBill表 结构体  CrmBill
type CrmBill struct {
	global.GVA_MODEL
	OrderId             *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID;"`                                       //订单ID
	PaymentCollentionId *int       `json:"paymentCollentionId" form:"paymentCollentionId" gorm:"column:payment_collention_id;comment:回款单ID;"` //回款单ID
	PaymentId           *int       `json:"PaymentId" form:"PaymentId" gorm:"column:payment_id;comment:付款ID;"`                                 //付款ID
	Amount              *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`                                             //金额
	Currency            string     `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:11;"`                               //币种
	UserId              *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;"`                                          //管理ID 销售代表
	BillNumber          string     `json:"billNumber" form:"billNumber" gorm:"column:bill_number;comment:账单编号;size:191;"`                     //账单编号
	PaymentStatus       string     `json:"paymentStatus" form:"paymentStatus" gorm:"column:payment_status;comment:付款状态;size:191;"`            //付款状态
	ExpirationTime      *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:到期时间;"`                  //到期时间
	BillName            string     `json:"billName" form:"billName" gorm:"column:bill_name;comment:账单名称;size:191;"`                           //账单名称
	CustomerId          *int       `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;size:19;"`                      //客户ID
}

// TableName crmBill表 CrmBill自定义表名 crm_bill
func (CrmBill) TableName() string {
	return "crm_bill"
}
