package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type CrmSupplierSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      CommpanyName  string `json:"commpanyName" form:"commpanyName" `
                      ContactPerson  string `json:"contactPerson" form:"contactPerson" `
                      Email  string `json:"email" form:"email" `
                      Telephone  string `json:"telephone" form:"telephone" `
                StartNoteAddTime  *time.Time  `json:"startNoteAddTime" form:"startNoteAddTime"`
                EndNoteAddTime  *time.Time  `json:"endNoteAddTime" form:"endNoteAddTime"`
                      UserId  *int `json:"userId" form:"userId" `
    request.PageInfo
}
