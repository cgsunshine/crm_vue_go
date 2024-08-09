// 自动生成模板CrmBill
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// crmBill表 结构体  CrmBill
type CrmBill struct {
 global.GVA_MODEL 
      OrderId  *int `json:"orderId" form:"orderId" gorm:"column:order_id;comment:产品名称;size:10;"`  //产品名称 
      CustomerId  *int `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:产品ID;size:10;"`  //产品ID 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;size:10;"`  //管理ID 销售代表 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`  //备注 
      Money  *float64 `json:"money" form:"money" gorm:"column:money;comment:账单金额 关联订单ID金额;size:10;"`  //账单金额 关联订单ID金额 
      Currency  string `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:10;"`  //币种 
      ExpirationTime  *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:到期时间;"`  //到期时间 
      PaymentStatus  string `json:"paymentStatus" form:"paymentStatus" gorm:"column:payment_status;comment:付款状态;size:255;"`  //付款状态 
      PaymentType  string `json:"paymentType" form:"paymentType" gorm:"column:payment_type;comment:付款方式;size:255;"`  //付款方式 
      PaymentTime  string `json:"paymentTime" form:"paymentTime" gorm:"column:payment_time;comment:付款时间;size:255;"`  //付款时间 
      PaymentCollention  *int `json:"paymentCollention" form:"paymentCollention" gorm:"column:payment_collention;comment:回款单ID;size:10;"`  //回款单ID 
}


// TableName crmBill表 CrmBill自定义表名 crm_bill
func (CrmBill) TableName() string {
  return "crm_bill"
}

