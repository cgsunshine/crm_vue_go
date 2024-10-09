// 自动生成模板CrmRefundTasks
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 退款管理 结构体  CrmRefundTasks
type CrmPageRefundTasks struct {
	global.GVA_MODEL
	AssigneeId   *int   `json:"assigneeId" form:"assigneeId" gorm:"column:assignee_id;comment:指派审批人id;"`                    //指派审批人id
	RefundStatus string `json:"refundStatus" form:"refundStatus" gorm:"column:refund_status;comment:退款状态;size:191;"`         //退款状态
	Comments     string `json:"comments" form:"comments" gorm:"column:comments;comment:补充意见;size:191;"`                      //补充意见
	Valid        *int   `json:"valid" form:"valid" gorm:"column:valid;comment:审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）;"` //审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）
	AssociatedId *int   `json:"associatedId" form:"associatedId" gorm:"column:associated_id;comment:关联id 合同 商机 回款;"`     //关联id 合同 商机 回款 押金
	RefundType   *int   `json:"refundType" form:"refundType" gorm:"column:refund_type;comment:退款类型 1押金;"`                  //退款类型 1押金
	//crm_deposits 联表查询
	DepositsName   string `json:"depositsName" form:"depositsName" gorm:"column:deposits_name;comment:押金名称;size:191;"`        //押金名称
	DepositsNumber string `json:"depositsNumber" form:"depositsNumbere" gorm:"column:deposits_number;comment:押金编号;size:191;"` //押金编号

}
