// 自动生成模板CrmProduct
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// crmProduct表 结构体  CrmProduct
type CrmProduct struct {
	global.GVA_MODEL
	ProductName    string   `json:"productName" form:"productName" gorm:"column:product_name;comment:产品名称;size:191;" binding:"required"` //产品名称
	UserId         *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;"`                                       //管理ID 销售代表
	Description    string   `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`                       //备注
	ProductGroupId *int     `json:"productGroupId" form:"productGroupId" gorm:"column:product_group_id;comment:产品分组id;"`                 //产品分组id
	ProductTypeId  *int     `json:"productTypeId" form:"productTypeId" gorm:"column:product_type_id;comment:产品类型id;"`                    //产品类型id
	Inventory      *int     `json:"inventory" form:"inventory" gorm:"column:inventory;comment:库存;"`                                      //库存
	Price          *float64 `json:"price" form:"price" gorm:"column:price;comment:产品原价;"`                                                //产品原价
	DiscountPrice  *float64 `json:"discountPrice" form:"discountPrice" gorm:"column:discount_price;comment:产品折扣价;"`                      //产品折扣价
	SalesPrice     *float64 `json:"salesPrice" form:"salesPrice" gorm:"column:sales_price;comment:产品销售价格;"`                              //产品销售价格
	ResourceId     *int     `json:"resourceId" form:"resourceId" gorm:"column:resource_id;comment:资源id;size:10;"`                        //资源id
	//OrderProducts  []CrmOrderProduct `gorm:"foreignKey:ProductID"`
	//OrderProducts []CrmOrderProduct
}

// TableName crmProduct表 CrmProduct自定义表名 crm_product
func (CrmProduct) TableName() string {
	return "crm_product"
}
