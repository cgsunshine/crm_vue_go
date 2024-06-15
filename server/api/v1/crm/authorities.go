package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateCrmAuthorities 更新数据权限状态
// @Tags CrmAuthorities
// @Summary 更新数据权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmAuthorities true "更新数据权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmAuthorities/updateCrmAuthorities [put]
func (crmAuthoritiesApi *CrmAuthoritiesApi) UpdateCrmStatusAuthorities(c *gin.Context) {
	var crmAuthorities crm.CrmAuthorities
	err := c.ShouldBindJSON(&crmAuthorities)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmAuthoritiesService.UpdateCrmStatusAuthorities(crmAuthorities); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
