package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmRefundTasksSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      AssigneeId  *int `json:"assigneeId" form:"assigneeId" `
                      RefundStatus  string `json:"refundStatus" form:"refundStatus" `
                      Valid  *int `json:"valid" form:"valid" `
                      AssociatedId  *int `json:"associatedId" form:"associatedId" `
                      RefundType  *int `json:"refundType" form:"refundType" `
    request.PageInfo
}
