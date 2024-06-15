package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetCrmBusinessOpportunityList 分页获取crmBusinessOpportunity表列表
// @Tags CrmBusinessOpportunity
// @Summary 分页获取crmBusinessOpportunity表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunitySearch true "分页获取crmBusinessOpportunity表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunity/getPageCrmBusinessOpportunityList [get]
func (crmBusinessOpportunityApi *CrmBusinessOpportunityApi) GetCrmPageBusinessOpportunityList(c *gin.Context) {
	var pageInfo crmReq.CrmBusinessOpportunitySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	pageInfo.UserId = userService.FindUserDataStatusById(&userID)
	if list, total, err := crmBusinessOpportunityService.GetPageCrmBusinessOpportunityInfoList(pageInfo); err != nil {
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

// FindCrmPageBusinessOpportunity 用id查询商机管理
// @Tags CrmBusinessOpportunity
// @Summary 用id查询商机管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBusinessOpportunity true "用id查询商机管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunity/findCrmBusinessOpportunity [get]
func (crmBusinessOpportunityApi *CrmBusinessOpportunityApi) FindCrmPageBusinessOpportunity(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBusinessOpportunity, err := crmBusinessOpportunityService.GetCrmPageBusinessOpportunity(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmBusinessOpportunity": recrmBusinessOpportunity}, c)
	}
}
