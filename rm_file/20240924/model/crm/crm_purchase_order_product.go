// 自动生成模板CrmPurchaseOrderProduct
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// crmPurchaseOrderProduct表 结构体  CrmPurchaseOrderProduct
type CrmPurchaseOrderProduct struct {
 global.GVA_MODEL 
      ProductId  *int `json:"productId" form:"productId" gorm:"column:product_id;comment:;size:20;"`  //productId字段 
      PurchaseOrderId  *int `json:"purchaseOrderId" form:"purchaseOrderId" gorm:"column:purchase_order_id;comment:;size:20;"`  //purchaseOrderId字段 
      Quantity  *int `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;"`  //quantity字段 
      Specifications  string `json:"specifications" form:"specifications" gorm:"column:specifications;comment:;size:10;"`  //specifications字段 
}


// TableName crmPurchaseOrderProduct表 CrmPurchaseOrderProduct自定义表名 crm_purchase_order_product
func (CrmPurchaseOrderProduct) TableName() string {
  return "crm_purchase_order_product"
}

