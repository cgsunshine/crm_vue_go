package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmTicketResponseTemplatesSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	CreatedAtOrder string     `json:"createdAtOrder" form:"createdAtOrder" default:"DESC"`
	TemplateName   string     `json:"templateName" form:"templateName"` //回复模板名称
	request.PageInfo
}
