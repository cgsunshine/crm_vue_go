package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetCrmOrderList 分页获取crmOrder表列表
// @Tags CrmOrder
// @Summary 分页获取crmOrder表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmOrderSearch true "分页获取crmOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmOrder/getCrmOrderList [get]
func (crmOrderApi *CrmOrderApi) GetCrmPageOrderList(c *gin.Context) {
	var pageInfo crmReq.CrmOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	pageInfo.UserId = userService.FindUserDataStatusById(&userID)

	if list, total, err := crmOrderService.GetCrmPageOrderInfoList(pageInfo); err != nil {
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

// FindCrmPageOrder 用id查询crmOrder表
// @Tags CrmOrder
// @Summary 用id查询crmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmOrder true "用id查询crmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmOrder/findCrmOrder [get]
func (crmOrderApi *CrmOrderApi) FindCrmPageOrder(c *gin.Context) {
	ID := c.Query("ID")
	if recrmOrder, err := crmOrderService.GetCrmPageOrder(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmOrder": recrmOrder}, c)
	}
}
