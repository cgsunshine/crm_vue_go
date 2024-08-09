package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmPurchaseOrderSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Amount              *float64   `json:"amount" form:"amount" `
	ContractId          *int       `json:"contractId" form:"contractId" `
	StartCreationTime   *time.Time `json:"startCreationTime" form:"startCreationTime"`
	EndCreationTime     *time.Time `json:"endCreationTime" form:"endCreationTime"`
	StartExpirationTime *time.Time `json:"startExpirationTime" form:"startExpirationTime"`
	EndExpirationTime   *time.Time `json:"endExpirationTime" form:"endExpirationTime"`
	ProductId           *int       `json:"productId" form:"productId" `
	Quantity            *int       `json:"quantity" form:"quantity" `
	UserId              *int       `json:"userId" form:"userId" `
	PurchaseOrderName   string     `json:"purchaseOrderName" form:"purchaseOrderName"` //订购单名称
	request.PageInfo
}
