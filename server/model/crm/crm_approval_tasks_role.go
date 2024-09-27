// 自动生成模板CrmApprovalTasksRole
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 角色审批表 结构体  CrmApprovalTasksRole
type CrmApprovalTasksRole struct {
 global.GVA_MODEL 
      ApprovalStatus  string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`  //审批状态 
      ApprovalType  *int `json:"approvalType" form:"approvalType" gorm:"column:approval_type;comment:审批类型 1合同 2商机 3回款;"`  //审批类型 1合同 2商机 3回款 
      AssigneeId  *int `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:审批人id;"`  //审批人id 
      AssociatedId  *int `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;"`  //关联id 合同 商机 回款 
      Comments  string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`  //审批意见 
      RequestId  *int `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;size:19;"`  //关联的审批请求ID 
      RoleId  *int `json:"roleId" form:"roleId" gorm:"column:role_id;comment:审批角色id;size:10;"`  //审批角色id 
      StepId  *int `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;size:19;"`  //当前审批步骤ID 
      Valid  *int `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;"`  //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝） 
}


// TableName 角色审批表 CrmApprovalTasksRole自定义表名 crm_approval_tasks_role
func (CrmApprovalTasksRole) TableName() string {
  return "crm_approval_tasks_role"
}

