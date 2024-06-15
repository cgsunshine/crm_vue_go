package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetCrmProductList 分页获取产品管理列表
// @Tags CrmProduct
// @Summary 分页获取产品管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmProductSearch true "分页获取产品管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProduct/getPageCrmProductList [get]
func (crmProductApi *CrmProductApi) GetCrmPageProductList(c *gin.Context) {
	var pageInfo crmReq.CrmProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//userID, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	//
	//pageInfo.UserId = userService.FindUserDataStatusById(&userID)

	if list, total, err := crmProductService.GetCrmPageProductInfoList(pageInfo); err != nil {
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

// FindCrmPageProduct 用id查询crmProduct表
// @Tags CrmProduct
// @Summary 用id查询crmProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmProduct true "用id查询crmProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmProduct/findCrmProduct [get]
func (crmProductApi *CrmProductApi) FindCrmPageProduct(c *gin.Context) {
	ID := c.Query("ID")
	if recrmProduct, err := crmProductService.GetCrmPageProduct(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmProduct": recrmProduct}, c)
	}
}
