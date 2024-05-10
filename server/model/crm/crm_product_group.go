// 自动生成模板CrmProductGroup
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 产品分组 结构体  CrmProductGroup
type CrmProductGroup struct {
 global.GVA_MODEL 
      GroupName  string `json:"groupName" form:"groupName" gorm:"column:group_name;comment:分组名称;size:191;"`  //分组名称 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:描述;size:191;"`  //描述 
}


// TableName 产品分组 CrmProductGroup自定义表名 crm_product_group
func (CrmProductGroup) TableName() string {
  return "crm_product_group"
}

