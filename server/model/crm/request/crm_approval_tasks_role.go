package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmApprovalTasksRoleSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      ApprovalStatus  string `json:"approvalStatus" form:"approvalStatus" `
                      ApprovalType  *int `json:"approvalType" form:"approvalType" `
                      AssigneeId  *int `json:"assigneeId" form:"assigneeId" `
                      AssociatedId  *int `json:"associatedId" form:"associatedId" `
                      Comments  string `json:"comments" form:"comments" `
                      RequestId  *int `json:"requestId" form:"requestId" `
                      RoleId  *int `json:"roleId" form:"roleId" `
    request.PageInfo
}
