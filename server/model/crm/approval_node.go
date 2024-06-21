// 自动生成模板CrmApprovalNode
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 节点审批 结构体  CrmApprovalNode
type CrmPageApprovalNode struct {
	global.GVA_MODEL
	ConditionExpression     string   `json:"conditionExpression" form:"conditionExpression" gorm:"column:conditionExpression;comment:预留条件;size:191;"`                              //预留条件
	NodeName                string   `json:"nodeName" form:"nodeName" gorm:"column:nodeName;comment:流程名称;size:255;"`                                                               //流程名称
	NodeOrder               *int     `json:"nodeOrder" form:"nodeOrder" gorm:"column:nodeOrder;comment:节点顺序;"`                                                                     //节点顺序
	NumberApprovedPersonnel *int     `json:"numberApprovedPersonnel" form:"numberApprovedPersonnel" gorm:"column:number_approved_personnel;comment:审批通过的人数，需要几人同意才能通过审批;size:10;"` //审批通过的人数，需要几人同意才能通过审批
	ProcessId               *int     `json:"processId" form:"processId" gorm:"column:processId;comment:所属审批流程ID;size:64;"`                                                         //所属审批流程ID
	RoleIds                 []string `json:"roleIds" form:"roleIds" gorm:"column:roleIds;comment:角色ID;size:255;"`                                                                  //角色ID
	UserId                  *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:添加记录的用户;size:10;"`                                                                  //添加记录的用户
	UserIds                 string   `json:"userIds" form:"userIds" gorm:"column:user_ids;comment:特定审批人ID;size:255;"`                                                              //特定审批人ID
}
