package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetCrmPaymentCollentionList 分页获取crmPaymentCollention表列表
// @Tags CrmPaymentCollention
// @Summary 分页获取crmPaymentCollention表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPaymentCollentionSearch true "分页获取crmPaymentCollention表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPaymentCollention/getCrmPaymentCollentionList [get]
func (crmPaymentCollentionApi *CrmPaymentCollentionApi) GetCrmPagePaymentCollentionInfoList(c *gin.Context) {
	var pageInfo crmReq.CrmPaymentCollentionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	pageInfo.UserId = userService.FindUserDataStatusById(&userID)

	if list, total, err := crmPaymentCollentionService.GetCrmPagePaymentCollentionInfoList(pageInfo); err != nil {
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

// FindCrmPagePaymentCollention 用id查询crmPaymentCollention表
// @Tags CrmPaymentCollention
// @Summary 用id查询crmPaymentCollention表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmPaymentCollention true "用id查询crmPaymentCollention表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPaymentCollention/findCrmPaymentCollention [get]
func (crmPaymentCollentionApi *CrmPaymentCollentionApi) FindCrmPagePaymentCollention(c *gin.Context) {
	ID := c.Query("ID")
	if recrmPaymentCollention, err := crmPaymentCollentionService.GetCrmPagePaymentCollention(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmPaymentCollention": recrmPaymentCollention}, c)
	}
}
