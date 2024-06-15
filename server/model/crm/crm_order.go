// 自动生成模板CrmOrder
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// crmOrder表 结构体  CrmOrder
type CrmOrder struct {
	global.GVA_MODEL
	Currency     string   `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:10;" binding:"required"`       //币种
	CustomerId   *int     `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;"`                         //客户ID
	Description  string   `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`                //备注
	DiscountRate *int     `json:"discountRate" form:"discountRate" gorm:"column:discount_rate;comment:折扣率;" binding:"required"` //折扣率
	Price        *float64 `json:"price" form:"price" gorm:"column:price;comment:产品原价;"`                                         //产品原价
	ProductId    string   `json:"productId" form:"productId" gorm:"column:product_id;comment:产品id;size:191;"`                   //产品id
	SalesPrice   *float64 `json:"salesPrice" form:"salesPrice" gorm:"column:sales_price;comment:产品销售价格;"`                       //产品销售价格
	UserId       *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;"`                                //管理ID 销售代表
}

// TableName crmOrder表 CrmOrder自定义表名 crm_order
func (CrmOrder) TableName() string {
	return "crm_order"
}
