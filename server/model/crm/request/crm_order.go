package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmOrderSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Currency  string `json:"currency" form:"currency" `
                      CustomerId  *int `json:"customerId" form:"customerId" `
                      DiscountRate  *int `json:"discountRate" form:"discountRate" `
                      Price  *float64 `json:"price" form:"price" `
                      ProductId  string `json:"productId" form:"productId" `
                      SalesPrice  *float64 `json:"salesPrice" form:"salesPrice" `
                      UserId  *int `json:"userId" form:"userId" `
    request.PageInfo
}
