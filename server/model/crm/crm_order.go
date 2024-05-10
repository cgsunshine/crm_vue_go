// 自动生成模板CrmOrder
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// 订单管理 结构体  CrmOrder
type CrmOrder struct {
 global.GVA_MODEL 
      ProductId  string `json:"productId" form:"productId" gorm:"column:product_id;comment:产品名称;size:191;"`  //产品名称 
      CustomerId  *int `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:产品ID;size:10;"`  //产品ID 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;size:10;"`  //管理ID 销售代表 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`  //备注 
      Price  *float64 `json:"price" form:"price" gorm:"column:price;comment:产品原价;size:10;"`  //产品原价 
      SalesPrice  *float64 `json:"salesPrice" form:"salesPrice" gorm:"column:sales_price;comment:产品折扣价;size:10;"`  //产品折扣价 
      Currency  string `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:10;"`  //币种 
      DiscountRate  *int `json:"discountRate" form:"discountRate" gorm:"column:discount_rate;comment:折扣率;size:10;"`  //折扣率 
      CreatedTime  *time.Time `json:"createdTime" form:"createdTime" gorm:"column:created_time;comment:创建时间;"`  //创建时间 
}


// TableName 订单管理 CrmOrder自定义表名 crm_order
func (CrmOrder) TableName() string {
  return "crm_order"
}

