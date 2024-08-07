package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetCrmPurchaseOrderList 分页获取crmPurchaseOrder表列表
// @Tags CrmPurchaseOrder
// @Summary 分页获取crmPurchaseOrder表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPurchaseOrderSearch true "分页获取crmPurchaseOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPurchaseOrder/getCrmPurchaseOrderList [get]
func (crmPurchaseOrderApi *CrmPurchaseOrderApi) GetCrmPagePurchaseOrderList(c *gin.Context) {
	var pageInfo crmReq.CrmPurchaseOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)

	if list, total, err := crmPurchaseOrderService.GetCrmPagePurchaseOrderInfoList(pageInfo); err != nil {
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

// FindCrmPagePurchaseOrder 用id查询订购单管理
// @Tags CrmPurchaseOrder
// @Summary 用id查询订购单管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmPurchaseOrder true "用id查询订购单管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPurchaseOrder/findCrmPurchaseOrder [get]
func (crmPurchaseOrderApi *CrmPurchaseOrderApi) FindCrmPagePurchaseOrder(c *gin.Context) {
	ID := c.Query("ID")
	if recrmPurchaseOrder, err := crmPurchaseOrderService.GetCrmPagePurchaseOrder(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmPurchaseOrder": recrmPurchaseOrder}, c)
	}
}
