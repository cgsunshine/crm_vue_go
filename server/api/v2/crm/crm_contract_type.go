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

type CrmContractTypeApi struct {
}

var crmContractTypeService = service.ServiceGroupApp.CrmServiceGroup.CrmContractTypeService


// CreateCrmContractType 创建合同类型
// @Tags CrmContractType
// @Summary 创建合同类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContractType true "创建合同类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContractType/createCrmContractType [post]
func (crmContractTypeApi *CrmContractTypeApi) CreateCrmContractType(c *gin.Context) {
	var crmContractType crm.CrmContractType
	err := c.ShouldBindJSON(&crmContractType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmContractTypeService.CreateCrmContractType(&crmContractType); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmContractType 删除合同类型
// @Tags CrmContractType
// @Summary 删除合同类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContractType true "删除合同类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContractType/deleteCrmContractType [delete]
func (crmContractTypeApi *CrmContractTypeApi) DeleteCrmContractType(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmContractTypeService.DeleteCrmContractType(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmContractTypeByIds 批量删除合同类型
// @Tags CrmContractType
// @Summary 批量删除合同类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmContractType/deleteCrmContractTypeByIds [delete]
func (crmContractTypeApi *CrmContractTypeApi) DeleteCrmContractTypeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmContractTypeService.DeleteCrmContractTypeByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmContractType 更新合同类型
// @Tags CrmContractType
// @Summary 更新合同类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContractType true "更新合同类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContractType/updateCrmContractType [put]
func (crmContractTypeApi *CrmContractTypeApi) UpdateCrmContractType(c *gin.Context) {
	var crmContractType crm.CrmContractType
	err := c.ShouldBindJSON(&crmContractType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmContractTypeService.UpdateCrmContractType(crmContractType); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmContractType 用id查询合同类型
// @Tags CrmContractType
// @Summary 用id查询合同类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContractType true "用id查询合同类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContractType/findCrmContractType [get]
func (crmContractTypeApi *CrmContractTypeApi) FindCrmContractType(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContractType, err := crmContractTypeService.GetCrmContractType(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContractType": recrmContractType}, c)
	}
}

// GetCrmContractTypeList 分页获取合同类型列表
// @Tags CrmContractType
// @Summary 分页获取合同类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContractTypeSearch true "分页获取合同类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContractType/getCrmContractTypeList [get]
func (crmContractTypeApi *CrmContractTypeApi) GetCrmContractTypeList(c *gin.Context) {
	var pageInfo crmReq.CrmContractTypeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmContractTypeService.GetCrmContractTypeInfoList(pageInfo); err != nil {
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

// GetCrmContractTypePublic 不需要鉴权的合同类型接口
// @Tags CrmContractType
// @Summary 不需要鉴权的合同类型接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContractTypeSearch true "分页获取合同类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContractType/getCrmContractTypeList [get]
func (crmContractTypeApi *CrmContractTypeApi) GetCrmContractTypePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的合同类型接口信息",
    }, "获取成功", c)
}
