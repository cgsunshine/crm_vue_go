package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetCrmSupplierList 分页获取crmSupplier表列表
// @Tags CrmSupplier
// @Summary 分页获取crmSupplier表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmSupplierSearch true "分页获取crmSupplier表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmSupplier/getCrmSupplierList [get]
func (crmSupplierApi *CrmSupplierApi) GetCrmPageSupplierList(c *gin.Context) {
	var pageInfo crmReq.CrmSupplierSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	pageInfo.UserId = userService.FindUserDataStatusById(&userID)

	if list, total, err := crmSupplierService.GetCrmPageSupplierInfoList(pageInfo); err != nil {
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

// FindCrmPageSupplier 用id查询crmSupplier表
// @Tags CrmSupplier
// @Summary 用id查询crmSupplier表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmSupplier true "用id查询crmSupplier表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmSupplier/findCrmSupplier [get]
func (crmSupplierApi *CrmSupplierApi) FindCrmPageSupplier(c *gin.Context) {
	ID := c.Query("ID")
	if recrmSupplier, err := crmSupplierService.GetCrmPageSupplier(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmSupplier": recrmSupplier}, c)
	}
}
