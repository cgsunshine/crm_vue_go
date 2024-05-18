// 自动生成模板CrmApprovalRecord
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmApprovalRecord表 结构体  CrmApprovalRecord
type CrmApprovalRecord struct {
	global.GVA_MODEL
	Code           string     `json:"code" form:"code" gorm:"column:code;comment:审批流程代码;size:64;"`                              //审批流程代码
	ModuleId       string     `json:"moduleId" form:"moduleId" gorm:"column:moduleId;comment:审批事项ID或模块ID;size:64;"`             //审批事项ID或模块ID
	Message        string     `json:"message" form:"message" gorm:"column:message;comment:审批相关信息或备注;size:191;"`                 //审批相关信息或备注
	ApplicantId    string     `json:"applicantId" form:"applicantId" gorm:"column:applicantId;comment:发起人ID;size:64;"`          //发起人ID
	ApplyTime      *time.Time `json:"applyTime" form:"applyTime" gorm:"column:applyTime;comment:申请时间;"`                         //申请时间
	Status         string     `json:"status" form:"status" gorm:"column:status;comment:审批状态;"`                                  //审批状态
	CurrentNodeId  string     `json:"currentNodeId" form:"currentNodeId" gorm:"column:currentNodeId;comment:当前审批节点ID;size:64;"` //当前审批节点ID
	ApproverId     string     `json:"approverId" form:"approverId" gorm:"column:approverId;comment:当前审批人ID;size:64;"`           //当前审批人ID
	ApproveTime    *time.Time `json:"approveTime" form:"approveTime" gorm:"column:approveTime;comment:审批时间;"`                   //审批时间
	ApproveOpinion string     `json:"approveOpinion" form:"approveOpinion" gorm:"column:approveOpinion;comment:审批意见;size:191;"` //审批意见
	FinalResult    string     `json:"finalResult" form:"finalResult" gorm:"column:finalResult;comment:最终审批结果;size:191;"`        //最终审批结果
	CloseTime      *time.Time `json:"closeTime" form:"closeTime" gorm:"column:closeTime;comment:流程关闭时间;"`                       //流程关闭时间
	Creator        string     `json:"creator" form:"creator" gorm:"column:creator;comment:记录创建者;size:64;"`                      //记录创建者
	CreateTime     *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`                     //创建时间
	Updator        string     `json:"updator" form:"updator" gorm:"column:updator;comment:记录最后更新者;size:64;"`                    //记录最后更新者
	UpdateTime     *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`                     //更新时间
	IsDeleted      *time.Time `json:"isDeleted" form:"isDeleted" gorm:"column:isDeleted;comment:逻辑删除标志;"`                       //逻辑删除标志
}

// TableName crmApprovalRecord表 CrmApprovalRecord自定义表名 crm_approval_record
func (CrmApprovalRecord) TableName() string {
	return "crm_approval_record"
}
