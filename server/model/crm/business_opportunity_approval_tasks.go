// 自动生成模板CrmBusinessOpportunityApprovalTasks
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 商机审批 结构体  CrmBusinessOpportunityApprovalTasks
type CrmPageBusinessOpportunityApprovalTasks struct {
	global.GVA_MODEL
	ApprovalStatus        string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"`              //审批状态
	AssigneeId            *int   `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;"`                                //指派审批人id
	BusinessOpportunityId *int   `json:"businessOpportunityId" form:"businessOpportunityId" gorm:"column:business_opportunity_id;comment:商机id;"` //商机id
	Comments              string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`                                 //审批意见
	RequestId             *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;size:19;"`                         //关联的审批请求ID
	StepId                *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;size:19;"`                                   //当前审批步骤ID
	Valid                 *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;"`                           //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
	//联表查询
	//crm_contract 表 合同审批状态
	ContractName string `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合同名称;size:191;"` //合同名称
}
