package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmStatementAccountSearch struct {
	baseSearchReq
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	PurchaseOrderId        *int       `json:"purchaseOrderId" form:"purchaseOrderId" `
	Amount                 *float64   `json:"amount" form:"amount" `
	StartCreationTime      *time.Time `json:"startCreationTime" form:"startCreationTime"`
	EndCreationTime        *time.Time `json:"endCreationTime" form:"endCreationTime"`
	StartPayableTime       *time.Time `json:"startPayableTime" form:"startPayableTime"`
	EndPayableTime         *time.Time `json:"endPayableTime" form:"endPayableTime"`
	PaymentStatus          string     `json:"paymentStatus" form:"paymentStatus" `
	UserId                 *int       `json:"userId" form:"userId" `
	StatementAccountNumber string     `json:"statementAccountNumber" form:"statementAccountNumber"` //对账单编号
	request.PageInfo
}
