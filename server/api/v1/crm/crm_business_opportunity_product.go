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

type CrmBusinessOpportunityProductApi struct {
}

var crmBusinessOpportunityProductService = service.ServiceGroupApp.CrmServiceGroup.CrmBusinessOpportunityProductService


// CreateCrmBusinessOpportunityProduct 创建crmBusinessOpportunityProduct表
// @Tags CrmBusinessOpportunityProduct
// @Summary 创建crmBusinessOpportunityProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityProduct true "创建crmBusinessOpportunityProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunityProduct/createCrmBusinessOpportunityProduct [post]
func (crmBusinessOpportunityProductApi *CrmBusinessOpportunityProductApi) CreateCrmBusinessOpportunityProduct(c *gin.Context) {
	var crmBusinessOpportunityProduct crm.CrmBusinessOpportunityProduct
	err := c.ShouldBindJSON(&crmBusinessOpportunityProduct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBusinessOpportunityProductService.CreateCrmBusinessOpportunityProduct(&crmBusinessOpportunityProduct); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmBusinessOpportunityProduct 删除crmBusinessOpportunityProduct表
// @Tags CrmBusinessOpportunityProduct
// @Summary 删除crmBusinessOpportunityProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityProduct true "删除crmBusinessOpportunityProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityProduct/deleteCrmBusinessOpportunityProduct [delete]
func (crmBusinessOpportunityProductApi *CrmBusinessOpportunityProductApi) DeleteCrmBusinessOpportunityProduct(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmBusinessOpportunityProductService.DeleteCrmBusinessOpportunityProduct(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmBusinessOpportunityProductByIds 批量删除crmBusinessOpportunityProduct表
// @Tags CrmBusinessOpportunityProduct
// @Summary 批量删除crmBusinessOpportunityProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmBusinessOpportunityProduct/deleteCrmBusinessOpportunityProductByIds [delete]
func (crmBusinessOpportunityProductApi *CrmBusinessOpportunityProductApi) DeleteCrmBusinessOpportunityProductByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmBusinessOpportunityProductService.DeleteCrmBusinessOpportunityProductByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmBusinessOpportunityProduct 更新crmBusinessOpportunityProduct表
// @Tags CrmBusinessOpportunityProduct
// @Summary 更新crmBusinessOpportunityProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityProduct true "更新crmBusinessOpportunityProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBusinessOpportunityProduct/updateCrmBusinessOpportunityProduct [put]
func (crmBusinessOpportunityProductApi *CrmBusinessOpportunityProductApi) UpdateCrmBusinessOpportunityProduct(c *gin.Context) {
	var crmBusinessOpportunityProduct crm.CrmBusinessOpportunityProduct
	err := c.ShouldBindJSON(&crmBusinessOpportunityProduct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBusinessOpportunityProductService.UpdateCrmBusinessOpportunityProduct(crmBusinessOpportunityProduct); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmBusinessOpportunityProduct 用id查询crmBusinessOpportunityProduct表
// @Tags CrmBusinessOpportunityProduct
// @Summary 用id查询crmBusinessOpportunityProduct表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBusinessOpportunityProduct true "用id查询crmBusinessOpportunityProduct表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunityProduct/findCrmBusinessOpportunityProduct [get]
func (crmBusinessOpportunityProductApi *CrmBusinessOpportunityProductApi) FindCrmBusinessOpportunityProduct(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBusinessOpportunityProduct, err := crmBusinessOpportunityProductService.GetCrmBusinessOpportunityProduct(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmBusinessOpportunityProduct": recrmBusinessOpportunityProduct}, c)
	}
}

// GetCrmBusinessOpportunityProductList 分页获取crmBusinessOpportunityProduct表列表
// @Tags CrmBusinessOpportunityProduct
// @Summary 分页获取crmBusinessOpportunityProduct表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunityProductSearch true "分页获取crmBusinessOpportunityProduct表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityProduct/getCrmBusinessOpportunityProductList [get]
func (crmBusinessOpportunityProductApi *CrmBusinessOpportunityProductApi) GetCrmBusinessOpportunityProductList(c *gin.Context) {
	var pageInfo crmReq.CrmBusinessOpportunityProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmBusinessOpportunityProductService.GetCrmBusinessOpportunityProductInfoList(pageInfo); err != nil {
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

// GetCrmBusinessOpportunityProductPublic 不需要鉴权的crmBusinessOpportunityProduct表接口
// @Tags CrmBusinessOpportunityProduct
// @Summary 不需要鉴权的crmBusinessOpportunityProduct表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunityProductSearch true "分页获取crmBusinessOpportunityProduct表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityProduct/getCrmBusinessOpportunityProductList [get]
func (crmBusinessOpportunityProductApi *CrmBusinessOpportunityProductApi) GetCrmBusinessOpportunityProductPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的crmBusinessOpportunityProduct表接口信息",
    }, "获取成功", c)
}
