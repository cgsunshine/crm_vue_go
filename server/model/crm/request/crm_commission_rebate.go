package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmCommissionRebateSearch struct {
	baseSearchReq
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	OrderId       *int     `json:"orderId" form:"orderId" `
	UserId        *int     `json:"userId" form:"userId" `
	Payee         string   `json:"payee" form:"payee" `
	PaymentMethod string   `json:"paymentMethod" form:"paymentMethod" `
	Account       string   `json:"account" form:"account" `
	Amount        *float64 `json:"amount" form:"amount" `
	request.PageInfo
}
