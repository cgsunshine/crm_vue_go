package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmContactApprovalTasksSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Assignee  string `json:"assignee" form:"assignee" `
                      ApprovalStatus  string `json:"approvalStatus" form:"approvalStatus" `
    request.PageInfo
}
