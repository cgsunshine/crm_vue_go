package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmOrderProductSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	OrderId   *int `json:"orderId" form:"orderId" `
	ProductId *int `json:"productId" form:"productId" `
	request.PageInfo
}

// Modify  user's auth structure
type SetOrderProduct struct {
	ID         uint
	ProductIds []uint `json:"productIds"` // 产品ID
}
