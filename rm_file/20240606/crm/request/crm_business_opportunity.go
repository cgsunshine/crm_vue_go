package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmBusinessOpportunitySearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Amount  *float64 `json:"amount" form:"amount" `
                      BusinessOpportunityName  string `json:"businessOpportunityName" form:"businessOpportunityName" `
                      CustomerId  *int `json:"customerId" form:"customerId" `
                StartInputTime  *time.Time  `json:"startInputTime" form:"startInputTime"`
                EndInputTime  *time.Time  `json:"endInputTime" form:"endInputTime"`
                      OfferValidityPeriod  *time.Time `json:"offerValidityPeriod" form:"offerValidityPeriod" `
                      ProductId  *int `json:"productId" form:"productId" `
                      Status  string `json:"status" form:"status" `
                      UserId  *int `json:"userId" form:"userId" `
    request.PageInfo
}
