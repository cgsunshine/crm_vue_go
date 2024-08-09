// 自动生成模板CrmContactApprovalTasks
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 合同审批 结构体  CrmReqContactApprovalTasks
type CrmReqContactApprovalTasks struct {
	global.GVA_MODEL
	RequestId      *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`                    //关联的审批请求ID
	StepId         *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`                              //当前审批步骤ID
	AssigneeId     []*int `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;size:10;"`           //指派审批人id
	ApprovalStatus string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"` //审批状态
	Comments       string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`                    //审批意见
	ContactId      *int   `json:"contactId" form:"contactId" gorm:"column:contact_id;comment:合同id;size:10;"`                 //合同id
	Valid          *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:10;"`      //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
}

// 合同审批 结构体  CrmRespContactApprovalTasks
type CrmRespContactApprovalTasks struct {
	global.GVA_MODEL
	RequestId      *int   `json:"requestId" form:"requestId" gorm:"column:request_id;comment:关联的审批请求ID;"`                    //关联的审批请求ID
	StepId         *int   `json:"stepId" form:"stepId" gorm:"column:step_id;comment:当前审批步骤ID;"`                              //当前审批步骤ID
	AssigneeId     *int   `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;size:10;"`           //指派审批人id
	ApprovalStatus string `json:"approvalStatus" form:"approvalStatus" gorm:"column:approval_status;comment:审批状态;size:191;"` //审批状态
	Comments       string `json:"comments" form:"comments" gorm:"column:comments;comment:审批意见;size:191;"`                    //审批意见
	ContactId      *int   `json:"contactId" form:"contactId" gorm:"column:contact_id;comment:合同id;size:10;"`                 //合同id
	Valid          *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;size:10;"`      //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
	//联表查询
	//crm_contract 表 合同审批状态
	ContractName string `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合同名称;size:191;"` //合同名称
}
