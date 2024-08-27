package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmDepositsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Amount           *float64   `json:"amount" form:"amount" `
	Currency         string     `json:"currency" form:"currency" `
	CustomerId       *int       `json:"customerId" form:"customerId" `
	OrderId          *int       `json:"orderId" form:"orderId" `
	StartPaymentDate *time.Time `json:"startPaymentDate" form:"startPaymentDate"`
	EndPaymentDate   *time.Time `json:"endPaymentDate" form:"endPaymentDate"`
	StartRefundDate  *time.Time `json:"startRefundDate" form:"startRefundDate"`
	EndRefundDate    *time.Time `json:"endRefundDate" form:"endRefundDate"`
	Status           string     `json:"status" form:"status" `
	UserId           *int       `json:"userId" form:"userId" `
	ReviewStatus     string     `json:"reviewStatus" form:"reviewStatus"`
	RefundStatus     string     `json:"refundStatus" form:"refundStatus"`
	request.PageInfo
}
