package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmApprovalRecordSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      ApprovalType  *int `json:"approvalType" form:"approvalType" `
                      ApproverId  *int `json:"approverId" form:"approverId" `
                      AssociatedId  *int `json:"associatedId" form:"associatedId" `
                      Status  string `json:"status" form:"status" `
    request.PageInfo
}
