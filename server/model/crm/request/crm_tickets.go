package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmTicketsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	SubmitterId               *int       `json:"submitterId" form:"submitterId"` //提交人ID
	TicketStatus              string     `json:"ticketStatus" form:"ticketStatus"`
	AssigneeId                *int       `json:"assigneeId" form:"assigneeId"` //指派人ID
	CategoryId                *int       `json:"categoryId" form:"categoryId"`
	CreatedAtOrder            string     `json:"createdAtOrder" form:"createdAtOrder" default:"DESC"`
	Title                     string     `json:"title" form:"title"`                                         //工单标题
	Priority                  string     `json:"priority" form:"priority"`                                   //工单优先级
	StartActualResolutionTime *time.Time `json:"startActualResolutionTime" form:"startActualResolutionTime"` //实际完成时间开始
	EndActualResolutionTime   *time.Time `json:"endActualResolutionTime" form:"endActualResolutionTime"`     //实际完成时间结束
	StartLastReplyTime        *time.Time `json:"startLastReplyTime" form:"startLastReplyTime"`               //最后一次回复时间开始
	EndLastReplyTime          *time.Time `json:"endLastReplyTime" form:"endLastReplyTime"`                   //最后一次回复时间结束
	request.PageInfo
}

//crm_tickets.created_at DESC
