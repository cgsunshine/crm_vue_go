// 自动生成模板CrmOperationRecords
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 操作记录 结构体  CrmOperationRecords
type CrmOperationRecords struct {
 global.GVA_MODEL 
      SysUserName  string `json:"sysUserName" form:"sysUserName" gorm:"column:sys_user_name;comment:操作用户名称;size:191;"`  //操作用户名称 
      ApiName  string `json:"apiName" form:"apiName" gorm:"column:api_name;comment:操作接口名称;size:191;"`  //操作接口名称 
      SysUserId  *int `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:操作用户ID;size:10;"`  //操作用户ID 
      ApiPath  string `json:"apiPath" form:"apiPath" gorm:"column:api_path;comment:操作接口路径;size:191;"`  //操作接口路径 
}


// TableName 操作记录 CrmOperationRecords自定义表名 crm_operation_records
func (CrmOperationRecords) TableName() string {
  return "crm_operation_records"
}

