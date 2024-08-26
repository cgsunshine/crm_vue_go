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
	ID           uint
	ProductIds   []uint            `json:"productIds"` // 产品ID
	ProductsInfo []CrmOrderProduct `json:"productsInfo" form:"productsInfo"`
}

type CrmOrderProduct struct {
	OrderId        *int   `json:"orderId" form:"orderId" gorm:"column:order_id;comment:;size:10;"`                     //orderId字段
	ProductId      *int   `json:"productId" form:"productId" gorm:"column:product_id;comment:;size:10;"`               //productId字段
	Quantity       *int   `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;size:10;"`                   //产品数量
	Specifications string `json:"specifications" form:"specifications" gorm:"column:specifications;comment:;size:10;"` //产品规格
}
