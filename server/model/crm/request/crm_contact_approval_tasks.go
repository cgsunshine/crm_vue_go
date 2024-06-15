package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmContactApprovalTasksSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      AssigneeId  *int `json:"assigneeId" form:"assigneeId" `
                      ApprovalStatus  string `json:"approvalStatus" form:"approvalStatus" `
                      ContactId  *int `json:"contactId" form:"contactId" `
    request.PageInfo
}
