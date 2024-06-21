package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmPaymentCollentionApprovalRecordSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      ApproverId  *int `json:"approverId" form:"approverId" `
                      PaymentCollentionId  *int `json:"paymentCollentionId" form:"paymentCollentionId" `
    request.PageInfo
}
