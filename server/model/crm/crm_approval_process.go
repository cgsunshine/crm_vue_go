// 自动生成模板CrmApprovalProcess
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 审批流程 结构体  CrmApprovalProcess
type CrmApprovalProcess struct {
 global.GVA_MODEL 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:审批流程描述;size:191;"`  //审批流程描述 
      ProcessName  string `json:"processName" form:"processName" gorm:"column:processName;comment:审批流程名称，如“财务报销流程”;size:255;"`  //审批流程名称，如“财务报销流程” 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:创建者ID;size:10;"`  //创建者ID 
}


// TableName 审批流程 CrmApprovalProcess自定义表名 crm_approval_process
func (CrmApprovalProcess) TableName() string {
  return "crm_approval_process"
}

