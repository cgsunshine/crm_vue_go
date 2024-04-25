// 自动生成模板CrmTest
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// crmTest表 结构体  CrmTest
type CrmTest struct {
 global.GVA_MODEL 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:255;"`  //name字段 
      Age  *int `json:"age" form:"age" gorm:"column:age;comment:;size:10;"`  //age字段 
}


// TableName crmTest表 CrmTest自定义表名 crm_test
func (CrmTest) TableName() string {
  return "crm_test"
}

