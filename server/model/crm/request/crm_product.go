package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmProductSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      ProductName  string `json:"productName" form:"productName" `
                      ProductGroupId  *int `json:"productGroupId" form:"productGroupId" `
                      ProductTypeId  *int `json:"productTypeId" form:"productTypeId" `
                      Inventory  *int `json:"inventory" form:"inventory" `
                      Price  *float64 `json:"price" form:"price" `
                      DiscountPrice  *float64 `json:"discountPrice" form:"discountPrice" `
                      SalesPrice  *float64 `json:"salesPrice" form:"salesPrice" `
                      ResourceId  *int `json:"resourceId" form:"resourceId" `
    request.PageInfo
}
