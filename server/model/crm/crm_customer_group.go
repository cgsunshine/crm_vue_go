// 自动生成模板CrmCustomerGroup
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// crmCustomerGroup表 结构体  CrmCustomerGroup
type CrmCustomerGroup struct {
 global.GVA_MODEL 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:描述;size:191;"`  //描述 
      GroupName  string `json:"groupName" form:"groupName" gorm:"column:group_name;comment:分组名称;size:90;"`  //分组名称 
}


// TableName crmCustomerGroup表 CrmCustomerGroup自定义表名 crm_customer_group
func (CrmCustomerGroup) TableName() string {
  return "crm_customer_group"
}

