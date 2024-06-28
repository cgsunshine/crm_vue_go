package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetCrmBillList 分页获取crmBill表列表
// @Tags CrmBill
// @Summary 分页获取crmBill表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBillSearch true "分页获取crmBill表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBill/getCrmBillList [get]
func (crmBillApi *CrmBillApi) GetCrmPageBillList(c *gin.Context) {
	var pageInfo crmReq.CrmBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)

	if list, total, err := crmBillService.GetCrmPageBillInfoList(pageInfo); err != nil {
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

// FindPageCrmBill 用id查询crmBill表
// @Tags CrmBill
// @Summary 用id查询crmBill表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBill true "用id查询crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBill/findCrmBill [get]
func (crmBillApi *CrmBillApi) FindPageCrmBill(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBill, err := crmBillService.GetCrmPageBill(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmBill": recrmBill}, c)
	}
}
