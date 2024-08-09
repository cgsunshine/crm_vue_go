// 自动生成模板CrmPurchaseOrder
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// 订购单 结构体  CrmPurchaseOrder
type CrmPurchaseOrder struct {
 global.GVA_MODEL 
      Amount  *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:金额;size:22;"`  //金额 
      ContractId  *int `json:"contractId" form:"contractId" gorm:"column:contract_id;comment:合同ID;size:19;"`  //合同ID 
      CreationTime  *time.Time `json:"creationTime" form:"creationTime" gorm:"column:creation_time;comment:创建时间;"`  //创建时间 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`  //备注 
      ExpirationTime  *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:到期时间;"`  //到期时间 
      ProductId  *int `json:"productId" form:"productId" gorm:"column:product_id;comment:产品ID;size:19;"`  //产品ID 
      Quantity  *int `json:"quantity" form:"quantity" gorm:"column:quantity;comment:数量;size:19;"`  //数量 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:负责人ID;size:19;"`  //负责人ID 
}


// TableName 订购单 CrmPurchaseOrder自定义表名 crm_purchase_order
func (CrmPurchaseOrder) TableName() string {
  return "crm_purchase_order"
}

