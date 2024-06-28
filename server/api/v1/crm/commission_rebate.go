package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetCrmCommissionRebateList 分页获取crmCommissionRebate表列表
// @Tags CrmCommissionRebate
// @Summary 分页获取crmCommissionRebate表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmCommissionRebateSearch true "分页获取crmCommissionRebate表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCommissionRebate/getCrmCommissionRebateList [get]
func (crmCommissionRebateApi *CrmCommissionRebateApi) GetCrmPageCommissionRebateList(c *gin.Context) {
	var pageInfo crmReq.CrmCommissionRebateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)
	if list, total, err := crmCommissionRebateService.GetCrmPageCommissionRebateInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// FindCrmPageCommissionRebate 用id查询crmCommissionRebate表
// @Tags CrmCommissionRebate
// @Summary 用id查询crmCommissionRebate表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmCommissionRebate true "用id查询crmCommissionRebate表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmCommissionRebate/findCrmCommissionRebate [get]
func (crmCommissionRebateApi *CrmCommissionRebateApi) FindCrmPageCommissionRebate(c *gin.Context) {
	ID := c.Query("ID")
	if recrmCommissionRebate, err := crmCommissionRebateService.GetCrmPageCommissionRebate(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmCommissionRebate": recrmCommissionRebate}, c)
	}
}
