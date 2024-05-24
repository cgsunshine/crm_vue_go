package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateCrmContactFileUploadAndDownloads 更新crmContactFileUploadAndDownloads表
// @Tags CrmContactFileUploadAndDownloads
// @Summary 更新crmContactFileUploadAndDownloads表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactFileUploadAndDownloads true "更新crmContactFileUploadAndDownloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactFileUploadAndDownloads/updateCrmContactFileUploadAndDownloads [put]
func (crmContactFileUploadAndDownloadsApi *CrmContactFileUploadAndDownloadsApi) UpdateCrmContactFileSort(c *gin.Context) {

	sorts, ok := c.GetPostFormArray("sorts")
	if ok {
		response.FailWithMessage("更新失败", c)
		return
	}
	for i, sort := range sorts {
		if err := crmContactFileUploadAndDownloadsService.Sort(i, sort); err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			response.FailWithMessage("更新失败", c)
			return
		}
	}
	response.OkWithMessage("更新成功", c)

}
