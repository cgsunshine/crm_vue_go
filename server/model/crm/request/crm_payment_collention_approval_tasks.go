package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmPaymentCollentionApprovalTasksSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      ApprovalStatus  string `json:"approvalStatus" form:"approvalStatus" `
                      AssigneeId  *int `json:"assigneeId" form:"assigneeId" `
                      PaymentCollentionId  *int `json:"paymentCollentionId" form:"paymentCollentionId" `
    request.PageInfo
}
