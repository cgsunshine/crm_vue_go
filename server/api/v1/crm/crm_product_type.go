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

type CrmProductTypeApi struct {
}

var crmProductTypeService = service.ServiceGroupApp.CrmServiceGroup.CrmProductTypeService


// CreateCrmProductType 创建产品类型
// @Tags CrmProductType
// @Summary 创建产品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmProductType true "创建产品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmProductType/createCrmProductType [post]
func (crmProductTypeApi *CrmProductTypeApi) CreateCrmProductType(c *gin.Context) {
	var crmProductType crm.CrmProductType
	err := c.ShouldBindJSON(&crmProductType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmProductTypeService.CreateCrmProductType(&crmProductType); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmProductType 删除产品类型
// @Tags CrmProductType
// @Summary 删除产品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmProductType true "删除产品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmProductType/deleteCrmProductType [delete]
func (crmProductTypeApi *CrmProductTypeApi) DeleteCrmProductType(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmProductTypeService.DeleteCrmProductType(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmProductTypeByIds 批量删除产品类型
// @Tags CrmProductType
// @Summary 批量删除产品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmProductType/deleteCrmProductTypeByIds [delete]
func (crmProductTypeApi *CrmProductTypeApi) DeleteCrmProductTypeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmProductTypeService.DeleteCrmProductTypeByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmProductType 更新产品类型
// @Tags CrmProductType
// @Summary 更新产品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmProductType true "更新产品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmProductType/updateCrmProductType [put]
func (crmProductTypeApi *CrmProductTypeApi) UpdateCrmProductType(c *gin.Context) {
	var crmProductType crm.CrmProductType
	err := c.ShouldBindJSON(&crmProductType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmProductTypeService.UpdateCrmProductType(crmProductType); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmProductType 用id查询产品类型
// @Tags CrmProductType
// @Summary 用id查询产品类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmProductType true "用id查询产品类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmProductType/findCrmProductType [get]
func (crmProductTypeApi *CrmProductTypeApi) FindCrmProductType(c *gin.Context) {
	ID := c.Query("ID")
	if recrmProductType, err := crmProductTypeService.GetCrmProductType(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmProductType": recrmProductType}, c)
	}
}

// GetCrmProductTypeList 分页获取产品类型列表
// @Tags CrmProductType
// @Summary 分页获取产品类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmProductTypeSearch true "分页获取产品类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProductType/getCrmProductTypeList [get]
func (crmProductTypeApi *CrmProductTypeApi) GetCrmProductTypeList(c *gin.Context) {
	var pageInfo crmReq.CrmProductTypeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmProductTypeService.GetCrmProductTypeInfoList(pageInfo); err != nil {
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

// GetCrmProductTypePublic 不需要鉴权的产品类型接口
// @Tags CrmProductType
// @Summary 不需要鉴权的产品类型接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmProductTypeSearch true "分页获取产品类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProductType/getCrmProductTypeList [get]
func (crmProductTypeApi *CrmProductTypeApi) GetCrmProductTypePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的产品类型接口信息",
    }, "获取成功", c)
}
