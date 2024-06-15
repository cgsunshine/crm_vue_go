package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
)

type CrmFileResponse struct {
	File []crm.CrmContactFile `json:"file"`
}

type CrmContactFileResponse struct {
	File crm.CrmContactFile `json:"file"`
}

type CrmStatementAccountResponse struct {
	File crm.CrmStatementAccountFile `json:"file"`
}

type CrmBusinessOpportunityFileResponse struct {
	File crm.CrmBusinessOpportunityFile `json:"file"`
}
