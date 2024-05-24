// 自动生成模板CrmCommissionRebate
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 返佣管理 结构体  CrmCommissionRebate
type CrmCommissionRebate struct {
 global.GVA_MODEL 
      Age  *int `json:"age" form:"age" gorm:"column:age;comment:;"`  //age字段 
      OrderId  *int `json:"orderId" form:"orderId" gorm:"column:order_id;comment:账单ID;size:10;"`  //账单ID 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:员工ID 负责人;size:10;"`  //员工ID 负责人 
      Payee  string `json:"payee" form:"payee" gorm:"column:payee;comment:收款人;size:191;"`  //收款人 
      PaymentMethod  string `json:"paymentMethod" form:"paymentMethod" gorm:"column:payment_method;comment:收款方式;size:191;"`  //收款方式 
      Account  string `json:"account" form:"account" gorm:"column:account;comment:账户;size:191;"`  //账户 
      Amount  *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:金额;size:10;"`  //金额 
      RebateDetails  string `json:"rebateDetails" form:"rebateDetails" gorm:"column:rebate_details;comment:返利详情;size:191;"`  //返利详情 
}


// TableName 返佣管理 CrmCommissionRebate自定义表名 crm_commission_rebate
func (CrmCommissionRebate) TableName() string {
  return "crm_commission_rebate"
}

