package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmContactApprovalRecordSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      ApproverId  string `json:"approverId" form:"approverId" `
                      ContactId  *int `json:"contactId" form:"contactId" `
    request.PageInfo
}
