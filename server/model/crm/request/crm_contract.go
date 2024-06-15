package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmContractSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	StartApplicationTime *time.Time `json:"startApplicationTime" form:"startApplicationTime"`
	EndApplicationTime   *time.Time `json:"endApplicationTime" form:"endApplicationTime"`
	ContractName         string     `json:"contractName" form:"contractName" `
	ContractStatus       string     `json:"contractStatus" form:"contractStatus" `
	CustomerId           *int       `json:"customerId" form:"customerId" `
	StartExpirationTime  *time.Time `json:"startExpirationTime" form:"startExpirationTime"`
	EndExpirationTime    *time.Time `json:"endExpirationTime" form:"endExpirationTime"`
	OrderId              *int       `json:"orderId" form:"orderId" `
	ReviewResult         string     `json:"reviewResult" form:"reviewResult" `
	ReviewStatus         string     `json:"reviewStatus" form:"reviewStatus" `
	UserId               *int       `json:"userId" form:"userId" `
	request.PageInfo
}
