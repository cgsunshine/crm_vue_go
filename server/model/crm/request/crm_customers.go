package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmCustomersSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      CustomerName  string `json:"customerName" form:"customerName" `
                      CustomerPhoneData  string `json:"customerPhoneData" form:"customerPhoneData" `
                      UserId  *int `json:"userId" form:"userId" `
                      CustomerCompany  string `json:"customerCompany" form:"customerCompany" `
                      CustomerAddress  string `json:"customerAddress" form:"customerAddress" `
                      CustomerStatus  string `json:"customerStatus" form:"customerStatus" `
                      CustomerGroup  string `json:"customerGroup" form:"customerGroup" `
    request.PageInfo
}
