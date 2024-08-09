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

type CrmConfigApi struct {
}

var crmConfigService = service.ServiceGroupApp.CrmServiceGroup.CrmConfigService


// CreateCrmConfig 创建系统配置
// @Tags CrmConfig
// @Summary 创建系统配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmConfig true "创建系统配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmConfig/createCrmConfig [post]
func (crmConfigApi *CrmConfigApi) CreateCrmConfig(c *gin.Context) {
	var crmConfig crm.CrmConfig
	err := c.ShouldBindJSON(&crmConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmConfigService.CreateCrmConfig(&crmConfig); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmConfig 删除系统配置
// @Tags CrmConfig
// @Summary 删除系统配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmConfig true "删除系统配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmConfig/deleteCrmConfig [delete]
func (crmConfigApi *CrmConfigApi) DeleteCrmConfig(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmConfigService.DeleteCrmConfig(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmConfigByIds 批量删除系统配置
// @Tags CrmConfig
// @Summary 批量删除系统配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmConfig/deleteCrmConfigByIds [delete]
func (crmConfigApi *CrmConfigApi) DeleteCrmConfigByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmConfigService.DeleteCrmConfigByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmConfig 更新系统配置
// @Tags CrmConfig
// @Summary 更新系统配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmConfig true "更新系统配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmConfig/updateCrmConfig [put]
func (crmConfigApi *CrmConfigApi) UpdateCrmConfig(c *gin.Context) {
	var crmConfig crm.CrmConfig
	err := c.ShouldBindJSON(&crmConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmConfigService.UpdateCrmConfig(crmConfig); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmConfig 用id查询系统配置
// @Tags CrmConfig
// @Summary 用id查询系统配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmConfig true "用id查询系统配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmConfig/findCrmConfig [get]
func (crmConfigApi *CrmConfigApi) FindCrmConfig(c *gin.Context) {
	ID := c.Query("ID")
	if recrmConfig, err := crmConfigService.GetCrmConfig(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmConfig": recrmConfig}, c)
	}
}

// GetCrmConfigList 分页获取系统配置列表
// @Tags CrmConfig
// @Summary 分页获取系统配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmConfigSearch true "分页获取系统配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmConfig/getCrmConfigList [get]
func (crmConfigApi *CrmConfigApi) GetCrmConfigList(c *gin.Context) {
	var pageInfo crmReq.CrmConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmConfigService.GetCrmConfigInfoList(pageInfo); err != nil {
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

// GetCrmConfigPublic 不需要鉴权的系统配置接口
// @Tags CrmConfig
// @Summary 不需要鉴权的系统配置接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmConfigSearch true "分页获取系统配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmConfig/getCrmConfigList [get]
func (crmConfigApi *CrmConfigApi) GetCrmConfigPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的系统配置接口信息",
    }, "获取成功", c)
}
