package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmBillSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Amount  *float64 `json:"amount" form:"amount" `
                      Currency  string `json:"currency" form:"currency" `
                      CustomerId  *int `json:"customerId" form:"customerId" `
                StartExpirationTime  *time.Time  `json:"startExpirationTime" form:"startExpirationTime"`
                EndExpirationTime  *time.Time  `json:"endExpirationTime" form:"endExpirationTime"`
                      OrderId  *int `json:"orderId" form:"orderId" `
                      PaymentCollention  *int `json:"paymentCollention" form:"paymentCollention" `
                      PaymentStatus  string `json:"paymentStatus" form:"paymentStatus" `
                StartPaymentTime  *time.Time  `json:"startPaymentTime" form:"startPaymentTime"`
                EndPaymentTime  *time.Time  `json:"endPaymentTime" form:"endPaymentTime"`
                      PaymentType  string `json:"paymentType" form:"paymentType" `
                      UserId  *int `json:"userId" form:"userId" `
    request.PageInfo
}
