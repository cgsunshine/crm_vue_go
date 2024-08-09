package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmPaymentSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	StatementAccountId *int       `json:"statementAccountId" form:"statementAccountId" `
	PaymentAmount      *float64   `json:"paymentAmount" form:"paymentAmount" `
	StartPaymentTime   *time.Time `json:"startPaymentTime" form:"startPaymentTime"`
	EndPaymentTime     *time.Time `json:"endPaymentTime" form:"endPaymentTime"`
	PaymentVoucher     string     `json:"paymentVoucher" form:"paymentVoucher" `
	UserId             *int       `json:"userId" form:"userId" `
	request.PageInfo
}
