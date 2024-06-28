package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CrmOrderProductApi struct {
}

var crmOrderProductService = service.ServiceGroupApp.CrmServiceGroup.CrmOrderProductService

// CreateCrmOrderProduct 创建crmOrderProduct表
// @Tags CrmOrderProduct
// @Summary 创建crmOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmOrderProduct true "创建crmOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmOrderProduct/createCrmOrderProduct [post]
func (crmOrderProductApi *CrmOrderProductApi) CreateCrmOrderProduct(c *gin.Context) {
	var crmOrderProduct crm.CrmOrderProduct
	err := c.ShouldBindJSON(&crmOrderProduct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmOrderProductService.CreateCrmOrderProduct(&crmOrderProduct); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DelCrmOrderProduct 创建crmOrderProduct表
// @Tags CrmOrderProduct
// @Summary 创建crmOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmOrderProduct true "创建crmOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmOrderProduct/createCrmOrderProduct [post]
func (crmOrderProductApi *CrmOrderProductApi) DelCrmOrderProduct(c *gin.Context) {
	var crmOrderProduct crm.CrmOrderProduct
	err := c.ShouldBindJSON(&crmOrderProduct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmOrderProductService.DeleteUnscopedCrmOrderProduct(&crmOrderProduct); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmOrderProduct 删除crmOrderProduct表
// @Tags CrmOrderProduct
// @Summary 删除crmOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmOrderProduct true "删除crmOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmOrderProduct/deleteCrmOrderProduct [delete]
func (crmOrderProductApi *CrmOrderProductApi) DeleteCrmOrderProduct(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmOrderProductService.DeleteCrmOrderProduct(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmOrderProductByIds 批量删除crmOrderProduct表
// @Tags CrmOrderProduct
// @Summary 批量删除crmOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmOrderProduct/deleteCrmOrderProductByIds [delete]
func (crmOrderProductApi *CrmOrderProductApi) DeleteCrmOrderProductByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmOrderProductService.DeleteCrmOrderProductByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmOrderProduct 更新crmOrderProduct表
// @Tags CrmOrderProduct
// @Summary 更新crmOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmOrderProduct true "更新crmOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmOrderProduct/updateCrmOrderProduct [put]
func (crmOrderProductApi *CrmOrderProductApi) UpdateCrmOrderProduct(c *gin.Context) {
	var crmOrderProduct crm.CrmOrderProduct
	err := c.ShouldBindJSON(&crmOrderProduct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmOrderProductService.UpdateCrmOrderProduct(crmOrderProduct); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmOrderProduct 用id查询crmOrderProduct表
// @Tags CrmOrderProduct
// @Summary 用id查询crmOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmOrderProduct true "用id查询crmOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmOrderProduct/findCrmOrderProduct [get]
func (crmOrderProductApi *CrmOrderProductApi) FindCrmOrderProduct(c *gin.Context) {
	ID := c.Query("ID")
	if recrmOrderProduct, err := crmOrderProductService.GetCrmOrderProduct(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmOrderProduct": recrmOrderProduct}, c)
	}
}

// GetCrmOrderProductList 分页获取crmOrderProduct表列表
// @Tags CrmOrderProduct
// @Summary 分页获取crmOrderProduct表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmOrderProductSearch true "分页获取crmOrderProduct表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmOrderProduct/getCrmOrderProductList [get]
func (crmOrderProductApi *CrmOrderProductApi) GetCrmOrderProductList(c *gin.Context) {
	var pageInfo crmReq.CrmOrderProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmOrderProductService.GetCrmOrderProductInfoList(pageInfo); err != nil {
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

// GetCrmOrderProductPublic 不需要鉴权的crmOrderProduct表接口
// @Tags CrmOrderProduct
// @Summary 不需要鉴权的crmOrderProduct表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmOrderProductSearch true "分页获取crmOrderProduct表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmOrderProduct/getCrmOrderProductList [get]
func (crmOrderProductApi *CrmOrderProductApi) GetCrmOrderProductPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的crmOrderProduct表接口信息",
	}, "获取成功", c)
}
