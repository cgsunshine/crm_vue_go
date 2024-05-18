// 自动生成模板CrmUser
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// crmUser表 结构体  CrmUser
type CrmUser struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:255;"`  //name字段 
}


// TableName crmUser表 CrmUser自定义表名 crm_user
func (CrmUser) TableName() string {
  return "crm_user"
}

