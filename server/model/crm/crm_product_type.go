// 自动生成模板CrmProductType
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 产品类型 结构体  CrmProductType
type CrmProductType struct {
 global.GVA_MODEL 
      TypeName  string `json:"typeName" form:"typeName" gorm:"column:type_name;comment:分组名称;size:191;"`  //分组名称 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`  //备注 
      Unit  string `json:"unit" form:"unit" gorm:"column:unit;comment:单位;size:191;"`  //单位 
}


// TableName 产品类型 CrmProductType自定义表名 crm_product_type
func (CrmProductType) TableName() string {
  return "crm_product_type"
}

