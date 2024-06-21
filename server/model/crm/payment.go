// 自动生成模板CrmPayment
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmPayment表 结构体  CrmPayment
type CrmPagePayment struct {
	global.GVA_MODEL
	OrderId        *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:账单ID;"`                               //账单ID
	PaymentAmount  *float64   `json:"paymentAmount" form:"paymentAmount" gorm:"column:payment_amount;comment:付款金额;"`             //付款金额
	PaymentTime    *time.Time `json:"paymentTime" form:"paymentTime" gorm:"column:payment_time;comment:付款时间;"`                   //付款时间
	PaymentVoucher string     `json:"paymentVoucher" form:"paymentVoucher" gorm:"column:payment_voucher;comment:付款凭证;size:191;"` //付款凭证
	UserId         *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:销售ID;"`                                  //销售ID
	//以下是联表查询字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	//crm_order 表
	OrderName string `json:"orderName" form:"orderName" gorm:"column:order_name;comment:订单名称;size:10;" binding:"required"` //订单名称
}
