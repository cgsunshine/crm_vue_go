// 自动生成模板CrmApprovalNode
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// crmApprovalNode表 结构体  CrmApprovalNode
type CrmApprovalNode struct {
 global.GVA_MODEL 
      ProcessId  string `json:"processId" form:"processId" gorm:"column:processId;comment:所属审批流程ID;size:64;"`  //所属审批流程ID 
      NodeName  string `json:"nodeName" form:"nodeName" gorm:"column:nodeName;comment:流程名称;size:255;"`  //流程名称 
      NodeOrder  *int `json:"nodeOrder" form:"nodeOrder" gorm:"column:nodeOrder;comment:节点顺序;size:10;"`  //节点顺序，用于确定审批步骤 
      RoleIds  string `json:"roleIds" form:"roleIds" gorm:"column:roleIds;comment:角色ID;size:255;"`  //角色ID 多个逗号分隔 
      UserId  string `json:"userId" form:"userId" gorm:"column:userId;comment:特定审批人ID;size:255;"`  //特定审批人ID 可以多个逗号分隔 
      ConditionExpression  string `json:"conditionExpression" form:"conditionExpression" gorm:"column:conditionExpression;comment:预留条件;"`  //条件表达式，用于动态路由到不同节点，可为空 
}


// TableName crmApprovalNode表 CrmApprovalNode自定义表名 crm_approval_node
func (CrmApprovalNode) TableName() string {
  return "crm_approval_node"
}

