package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
)

type CrmFileResponse struct {
	File []crm.CrmContactFileUploadAndDownloads `json:"file"`
}

type CrmContactFileResponse struct {
	File crm.CrmContactFileUploadAndDownloads `json:"file"`
}
