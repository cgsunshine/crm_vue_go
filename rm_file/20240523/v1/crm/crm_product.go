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

type CrmProductApi struct {
}

var crmProductService = service.ServiceGroupApp.CrmServiceGroup.CrmProductService


// CreateCrmProduct 创建产品管理
// @Tags CrmProduct
// @Summary 创建产品管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmProduct true "创建产品管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmProduct/createCrmProduct [post]
func (crmProductApi *CrmProductApi) CreateCrmProduct(c *gin.Context) {
	var crmProduct crm.CrmProduct
	err := c.ShouldBindJSON(&crmProduct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmProductService.CreateCrmProduct(&crmProduct); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmProduct 删除产品管理
// @Tags CrmProduct
// @Summary 删除产品管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmProduct true "删除产品管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmProduct/deleteCrmProduct [delete]
func (crmProductApi *CrmProductApi) DeleteCrmProduct(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmProductService.DeleteCrmProduct(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmProductByIds 批量删除产品管理
// @Tags CrmProduct
// @Summary 批量删除产品管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmProduct/deleteCrmProductByIds [delete]
func (crmProductApi *CrmProductApi) DeleteCrmProductByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmProductService.DeleteCrmProductByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmProduct 更新产品管理
// @Tags CrmProduct
// @Summary 更新产品管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmProduct true "更新产品管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmProduct/updateCrmProduct [put]
func (crmProductApi *CrmProductApi) UpdateCrmProduct(c *gin.Context) {
	var crmProduct crm.CrmProduct
	err := c.ShouldBindJSON(&crmProduct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmProductService.UpdateCrmProduct(crmProduct); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmProduct 用id查询产品管理
// @Tags CrmProduct
// @Summary 用id查询产品管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmProduct true "用id查询产品管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmProduct/findCrmProduct [get]
func (crmProductApi *CrmProductApi) FindCrmProduct(c *gin.Context) {
	ID := c.Query("ID")
	if recrmProduct, err := crmProductService.GetCrmProduct(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmProduct": recrmProduct}, c)
	}
}

// GetCrmProductList 分页获取产品管理列表
// @Tags CrmProduct
// @Summary 分页获取产品管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmProductSearch true "分页获取产品管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProduct/getCrmProductList [get]
func (crmProductApi *CrmProductApi) GetCrmProductList(c *gin.Context) {
	var pageInfo crmReq.CrmProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmProductService.GetCrmProductInfoList(pageInfo); err != nil {
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

// GetCrmProductPublic 不需要鉴权的产品管理接口
// @Tags CrmProduct
// @Summary 不需要鉴权的产品管理接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmProductSearch true "分页获取产品管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProduct/getCrmProductList [get]
func (crmProductApi *CrmProductApi) GetCrmProductPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的产品管理接口信息",
    }, "获取成功", c)
}
