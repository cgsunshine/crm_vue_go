package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmProcurementContractSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      ContractAmount  *float64 `json:"contractAmount" form:"contractAmount" `
                      ContractName  string `json:"contractName" form:"contractName" `
                      ContractStatus  string `json:"contractStatus" form:"contractStatus" `
                StartCreationTime  *time.Time  `json:"startCreationTime" form:"startCreationTime"`
                EndCreationTime  *time.Time  `json:"endCreationTime" form:"endCreationTime"`
                StartExpirationTime  *time.Time  `json:"startExpirationTime" form:"startExpirationTime"`
                EndExpirationTime  *time.Time  `json:"endExpirationTime" form:"endExpirationTime"`
                      UserId  *int `json:"userId" form:"userId" `
    request.PageInfo
}
