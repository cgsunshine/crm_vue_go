package request

import (
	"time"
)

type CrmAdminHomeSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	UserId         *int       `json:"userId" form:"userId" `
}
