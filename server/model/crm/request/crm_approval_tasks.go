package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmApprovalTasksSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      ApprovalStatus  string `json:"approvalStatus" form:"approvalStatus" `
                      ApprovalType  *int `json:"approvalType" form:"approvalType" `
                      AssigneeId  *int `json:"assigneeId" form:"assigneeId" `
                      AssociatedId  *int `json:"associatedId" form:"associatedId" `
                      RequestId  *int `json:"requestId" form:"requestId" `
    request.PageInfo
}
