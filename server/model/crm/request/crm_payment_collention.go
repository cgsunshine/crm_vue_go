package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmPaymentCollentionSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Amount  *float64 `json:"amount" form:"amount" `
                      ApprovedById  string `json:"approvedById" form:"approvedById" `
                      Currency  string `json:"currency" form:"currency" `
                      CustomerId  *int `json:"customerId" form:"customerId" `
                StartPaymentTime  *time.Time  `json:"startPaymentTime" form:"startPaymentTime"`
                EndPaymentTime  *time.Time  `json:"endPaymentTime" form:"endPaymentTime"`
                      ReviewStatus  string `json:"reviewStatus" form:"reviewStatus" `
                      UserId  *int `json:"userId" form:"userId" `
    request.PageInfo
}
