package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmPurchaseOrderSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      ContractId  *int `json:"contractId" form:"contractId" `
                      ProductId  *int `json:"productId" form:"productId" `
                      Quantity  *int `json:"quantity" form:"quantity" `
                      Amount  *float64 `json:"amount" form:"amount" `
                      UserId  *int `json:"userId" form:"userId" `
                StartCreationTime  *time.Time  `json:"startCreationTime" form:"startCreationTime"`
                EndCreationTime  *time.Time  `json:"endCreationTime" form:"endCreationTime"`
                StartExpirationTime  *time.Time  `json:"startExpirationTime" form:"startExpirationTime"`
                EndExpirationTime  *time.Time  `json:"endExpirationTime" form:"endExpirationTime"`
    request.PageInfo
}
