package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CrmPurchaseOrderProductApi struct {
}

var crmPurchaseOrderProductService = service.ServiceGroupApp.CrmServiceGroup.CrmPurchaseOrderProductService


// CreateCrmPurchaseOrderProduct 创建crmPurchaseOrderProduct表
// @Tags CrmPurchaseOrderProduct
// @Summary 创建crmPurchaseOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPurchaseOrderProduct true "创建crmPurchaseOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPurchaseOrderProduct/createCrmPurchaseOrderProduct [post]
func (crmPurchaseOrderProductApi *CrmPurchaseOrderProductApi) CreateCrmPurchaseOrderProduct(c *gin.Context) {
	var crmPurchaseOrderProduct crm.CrmPurchaseOrderProduct
	err := c.ShouldBindJSON(&crmPurchaseOrderProduct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmPurchaseOrderProductService.CreateCrmPurchaseOrderProduct(&crmPurchaseOrderProduct); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmPurchaseOrderProduct 删除crmPurchaseOrderProduct表
// @Tags CrmPurchaseOrderProduct
// @Summary 删除crmPurchaseOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPurchaseOrderProduct true "删除crmPurchaseOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPurchaseOrderProduct/deleteCrmPurchaseOrderProduct [delete]
func (crmPurchaseOrderProductApi *CrmPurchaseOrderProductApi) DeleteCrmPurchaseOrderProduct(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmPurchaseOrderProductService.DeleteCrmPurchaseOrderProduct(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmPurchaseOrderProductByIds 批量删除crmPurchaseOrderProduct表
// @Tags CrmPurchaseOrderProduct
// @Summary 批量删除crmPurchaseOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmPurchaseOrderProduct/deleteCrmPurchaseOrderProductByIds [delete]
func (crmPurchaseOrderProductApi *CrmPurchaseOrderProductApi) DeleteCrmPurchaseOrderProductByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmPurchaseOrderProductService.DeleteCrmPurchaseOrderProductByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmPurchaseOrderProduct 更新crmPurchaseOrderProduct表
// @Tags CrmPurchaseOrderProduct
// @Summary 更新crmPurchaseOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPurchaseOrderProduct true "更新crmPurchaseOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmPurchaseOrderProduct/updateCrmPurchaseOrderProduct [put]
func (crmPurchaseOrderProductApi *CrmPurchaseOrderProductApi) UpdateCrmPurchaseOrderProduct(c *gin.Context) {
	var crmPurchaseOrderProduct crm.CrmPurchaseOrderProduct
	err := c.ShouldBindJSON(&crmPurchaseOrderProduct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmPurchaseOrderProductService.UpdateCrmPurchaseOrderProduct(crmPurchaseOrderProduct); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmPurchaseOrderProduct 用id查询crmPurchaseOrderProduct表
// @Tags CrmPurchaseOrderProduct
// @Summary 用id查询crmPurchaseOrderProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmPurchaseOrderProduct true "用id查询crmPurchaseOrderProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPurchaseOrderProduct/findCrmPurchaseOrderProduct [get]
func (crmPurchaseOrderProductApi *CrmPurchaseOrderProductApi) FindCrmPurchaseOrderProduct(c *gin.Context) {
	ID := c.Query("ID")
	if recrmPurchaseOrderProduct, err := crmPurchaseOrderProductService.GetCrmPurchaseOrderProduct(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmPurchaseOrderProduct": recrmPurchaseOrderProduct}, c)
	}
}

// GetCrmPurchaseOrderProductList 分页获取crmPurchaseOrderProduct表列表
// @Tags CrmPurchaseOrderProduct
// @Summary 分页获取crmPurchaseOrderProduct表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPurchaseOrderProductSearch true "分页获取crmPurchaseOrderProduct表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPurchaseOrderProduct/getCrmPurchaseOrderProductList [get]
func (crmPurchaseOrderProductApi *CrmPurchaseOrderProductApi) GetCrmPurchaseOrderProductList(c *gin.Context) {
	var pageInfo crmReq.CrmPurchaseOrderProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmPurchaseOrderProductService.GetCrmPurchaseOrderProductInfoList(pageInfo); err != nil {
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

// GetCrmPurchaseOrderProductPublic 不需要鉴权的crmPurchaseOrderProduct表接口
// @Tags CrmPurchaseOrderProduct
// @Summary 不需要鉴权的crmPurchaseOrderProduct表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPurchaseOrderProductSearch true "分页获取crmPurchaseOrderProduct表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPurchaseOrderProduct/getCrmPurchaseOrderProductList [get]
func (crmPurchaseOrderProductApi *CrmPurchaseOrderProductApi) GetCrmPurchaseOrderProductPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的crmPurchaseOrderProduct表接口信息",
    }, "获取成功", c)
}
