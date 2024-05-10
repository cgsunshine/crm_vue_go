// 自动生成模板CrmLoginLog
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 登录日志 结构体  CrmLoginLog
type CrmLoginLog struct {
 global.GVA_MODEL 
      SysUserName  string `json:"sysUserName" form:"sysUserName" gorm:"column:sys_user_name;comment:登录人用户名称;size:191;"`  //登录人用户名称 
      SysUserId  *int `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:登录用户ID;size:10;"`  //登录用户ID 
}


// TableName 登录日志 CrmLoginLog自定义表名 crm_login_log
func (CrmLoginLog) TableName() string {
  return "crm_login_log"
}

