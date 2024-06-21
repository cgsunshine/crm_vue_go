// 自动生成模板CrmApprovalNode
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 审批节点 结构体  CrmApprovalNode
type CrmApprovalNode struct {
 global.GVA_MODEL 
      ConditionExpression  string `json:"conditionExpression" form:"conditionExpression" gorm:"column:conditionExpression;comment:预留条件;size:191;"`  //预留条件 
      NodeName  string `json:"nodeName" form:"nodeName" gorm:"column:nodeName;comment:流程名称;size:255;"`  //流程名称 
      NodeOrder  *int `json:"nodeOrder" form:"nodeOrder" gorm:"column:nodeOrder;comment:节点顺序;size:19;"`  //节点顺序 
      NumberApprovedPersonnel  *int `json:"numberApprovedPersonnel" form:"numberApprovedPersonnel" gorm:"column:number_approved_personnel;comment:审批通过的人数，需要几人同意才能通过审批;"`  //审批通过的人数，需要几人同意才能通过审批 
      ProcessId  *int `json:"processId" form:"processId" gorm:"column:processId;comment:所属审批流程ID;size:19;"`  //所属审批流程ID 
      RoleIds  string `json:"roleIds" form:"roleIds" gorm:"column:roleIds;comment:角色ID;size:255;"`  //角色ID 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:添加记录的用户;"`  //添加记录的用户 
      UserIds  string `json:"userIds" form:"userIds" gorm:"column:user_ids;comment:特定审批人ID;size:255;"`  //特定审批人ID 
}


// TableName 审批节点 CrmApprovalNode自定义表名 crm_approval_node
func (CrmApprovalNode) TableName() string {
  return "crm_approval_node"
}

