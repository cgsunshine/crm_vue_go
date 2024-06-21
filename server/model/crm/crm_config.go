// 自动生成模板CrmConfig
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 系统配置 结构体  CrmConfig
type CrmConfig struct {
 global.GVA_MODEL 
      ConfigName  string `json:"configName" form:"configName" gorm:"column:config_name;comment:配置名称;size:10;"`  //配置名称 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:描述;size:199;"`  //描述 
      AssociatedId  *int `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 可为空;size:10;"`  //关联id 可为空 
      ConfigContent  string `json:"configContent" form:"configContent" gorm:"column:config_content;comment:配置内容 可为空;size:255;"`  //配置内容 可为空 
}


// TableName 系统配置 CrmConfig自定义表名 crm_config
func (CrmConfig) TableName() string {
  return "crm_config"
}

