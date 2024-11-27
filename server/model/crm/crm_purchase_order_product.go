// 自动生成模板CrmPurchaseOrderProduct
package crm

// crmPurchaseOrderProduct表 结构体  CrmPurchaseOrderProduct
type CrmPurchaseOrderProduct struct {
	ProductId       *int     `json:"productId" form:"productId" gorm:"column:product_id;comment:;size:20;"`                    //productId字段
	PurchaseOrderId *int     `json:"purchaseOrderId" form:"purchaseOrderId" gorm:"column:purchase_order_id;comment:;size:20;"` //purchaseOrderId字段
	Quantity        *int     `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;"`                                //quantity字段
	Specifications  string   `json:"specifications" form:"specifications" gorm:"column:specifications;comment:;size:10;"`      //specifications字段
	Price           *float64 `json:"price" form:"price" gorm:"column:price;comment:产品单价;"`                                     //产品单价
	PurchaseOrder   CrmPurchaseOrder
	Product         CrmProduct
}

// TableName crmPurchaseOrderProduct表 CrmPurchaseOrderProduct自定义表名 crm_purchase_order_product
func (CrmPurchaseOrderProduct) TableName() string {
	return "crm_purchase_order_product"
}
