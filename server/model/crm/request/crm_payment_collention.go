package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmPaymentCollentionSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      BillId  *int `json:"billId" form:"billId" `
                      CustomerId  *int `json:"customerId" form:"customerId" `
                      UserId  *int `json:"userId" form:"userId" `
                      Currency  string `json:"currency" form:"currency" `
                      Proof  string `json:"proof" form:"proof" `
                      AuditingStatus  string `json:"auditingStatus" form:"auditingStatus" `
                      ApprovedById  string `json:"approvedById" form:"approvedById" `
                StartAuditingTime  *time.Time  `json:"startAuditingTime" form:"startAuditingTime"`
                EndAuditingTime  *time.Time  `json:"endAuditingTime" form:"endAuditingTime"`
    request.PageInfo
}
