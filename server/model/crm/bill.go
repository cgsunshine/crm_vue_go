// 自动生成模板CrmBill
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmBill表 结构体  CrmBill
type CrmPageBill struct {
	global.GVA_MODEL
	Amount            *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:账单金额 关联订单ID金额;size:22;"`                   //账单金额 关联订单ID金额
	Currency          string     `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:10;"`                        //币种
	CustomerId        *int       `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;"`                       //客户ID
	Description       string     `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`              //备注
	ExpirationTime    *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:到期时间;"`           //到期时间
	OrderId           *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID;"`                                //订单ID
	PaymentCollention *int       `json:"paymentCollention" form:"paymentCollention" gorm:"column:payment_collention;comment:回款单ID;"` //回款单ID
	PaymentStatus     string     `json:"paymentStatus" form:"paymentStatus" gorm:"column:payment_status;comment:付款状态;size:191;"`     //付款状态
	PaymentTime       *time.Time `json:"paymentTime" form:"paymentTime" gorm:"column:payment_time;comment:付款时间;"`                    //付款时间
	PaymentType       string     `json:"paymentType" form:"paymentType" gorm:"column:payment_type;comment:付款方式;size:10;"`            //付款方式
	UserId            *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;"`                              //管理ID 销售代表
	//以下是联表查询字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	//crm_customers 表
	CustomerName string `json:"customerName" form:"customerName" gorm:"column:customer_name;comment:客户名;size:191;"` //客户名
	//crm_order 表
	OrderName string `json:"orderName" form:"orderName" gorm:"column:order_name;comment:订单名称;size:10;" binding:"required"` //订单名称
}
