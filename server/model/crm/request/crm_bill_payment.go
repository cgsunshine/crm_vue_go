package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmBillPaymentSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Amount  *float64 `json:"amount" form:"amount" `
                      BillPaymentName  string `json:"billPaymentName" form:"billPaymentName" `
                      Currency  string `json:"currency" form:"currency" `
                      CustomerId  *int `json:"customerId" form:"customerId" `
                      PaymentStatus  string `json:"paymentStatus" form:"paymentStatus" `
                      StatementAccountId  *int `json:"statementAccountId" form:"statementAccountId" `
                      UserId  *int `json:"userId" form:"userId" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
