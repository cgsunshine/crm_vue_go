// 自动生成模板CrmOrderProduct
package crm

// crmOrderProduct表 结构体  CrmOrderProduct
type CrmOrderProduct struct {
	OrderId   uint `json:"orderId" form:"orderId" gorm:"column:order_id;comment:;size:10;"`       //orderId字段
	ProductId uint `json:"productId" form:"productId" gorm:"column:product_id;comment:;size:10;"` //productId字段
}

// TableName crmOrderProduct表 CrmOrderProduct自定义表名 crm_order_product
func (CrmOrderProduct) TableName() string {
	return "crm_order_product"
}
