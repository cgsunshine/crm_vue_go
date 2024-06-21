// 自动生成模板CrmApprovalProcess
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// crmApprovalProcess表 结构体  CrmApprovalProcess
type CrmApprovalProcess struct {
 global.GVA_MODEL 
      ProcessName  string `json:"processName" form:"processName" gorm:"column:processName;comment:审批流程名称，如“财务报销流程”;size:255;"`  //路程名称 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:审批流程描述;"`  //流程描述 
      CreatedBy  string `json:"createdBy" form:"createdBy" gorm:"column:createdBy;comment:创建者ID;size:64;"`  //创建者ID 
}


// TableName crmApprovalProcess表 CrmApprovalProcess自定义表名 crm_approval_process
func (CrmApprovalProcess) TableName() string {
  return "crm_approval_process"
}

