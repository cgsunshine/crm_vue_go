// 自动生成模板CrmCommissionRebate
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// crmCommissionRebate表 结构体  CrmCommissionRebate
type CrmCommissionRebate struct {
	global.GVA_MODEL
	OrderId       *int     `json:"orderId" form:"orderId" gorm:"column:order_id;comment:账单ID;size:10;" binding:"required"` //账单ID
	UserId        *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:销售ID 负责人;size:10;"`                   //销售ID 负责人
	Payee         string   `json:"payee" form:"payee" gorm:"column:payee;comment:收款人;size:191;" `                          //收款人
	PaymentMethod string   `json:"paymentMethod" form:"paymentMethod" gorm:"column:payment_method;comment:收款方式;size:191;"` //收款方式
	Account       string   `json:"account" form:"account" gorm:"column:account;comment:账户;size:191;"`                      //账户
	Amount        *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:金额;" binding:"required"`               //金额
	RebateDetails string   `json:"rebateDetails" form:"rebateDetails" gorm:"column:rebate_details;comment:返利详情;size:191;"` //返利详情
}

// TableName crmCommissionRebate表 CrmCommissionRebate自定义表名 crm_commission_rebate
func (CrmCommissionRebate) TableName() string {
	return "crm_commission_rebate"
}
