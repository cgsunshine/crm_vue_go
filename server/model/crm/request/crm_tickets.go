package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmTicketsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	SubmitterId    *int   `json:"submitterId" form:"submitterId"`
	ID             *int   `json:"id" form:"id"`
	TicketStatus   string `json:"ticketStatus" form:"ticketStatus"`
	AssigneeId     *int   `json:"assigneeId" form:"assigneeId" `
	CategoryId     *int   `json:"categoryId" form:"categoryId" `
	CreatedAtOrder string `json:"createdAtOrder" form:"createdAtOrder" default:"DESC"`
	request.PageInfo
}

//crm_tickets.created_at DESC
