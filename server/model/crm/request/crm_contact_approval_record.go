package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmContactApprovalRecordSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      ApproveOpinion  string `json:"approveOpinion" form:"approveOpinion" `
                      ApproverId  *int `json:"approverId" form:"approverId" `
    request.PageInfo
}
