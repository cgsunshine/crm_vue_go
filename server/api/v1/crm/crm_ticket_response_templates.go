package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CrmTicketResponseTemplatesApi struct {
}

var crmTicketResponseTemplatesService = service.ServiceGroupApp.CrmServiceGroup.CrmTicketResponseTemplatesService

// CreateCrmTicketResponseTemplates 创建快捷回复模板
// @Tags CrmTicketResponseTemplates
// @Summary 创建快捷回复模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTicketResponseTemplates true "创建快捷回复模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmTicketResponseTemplates/createCrmTicketResponseTemplates [post]
func (crmTicketResponseTemplatesApi *CrmTicketResponseTemplatesApi) CreateCrmTicketResponseTemplates(c *gin.Context) {
	var crmTicketResponseTemplates crm.CrmTicketResponseTemplates
	err := c.ShouldBindJSON(&crmTicketResponseTemplates)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmTicketResponseTemplates.CreatedBy = comm.GetHeaderUserId(c)

	if err := crmTicketResponseTemplatesService.CreateCrmTicketResponseTemplates(&crmTicketResponseTemplates); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmTicketResponseTemplates 删除快捷回复模板
// @Tags CrmTicketResponseTemplates
// @Summary 删除快捷回复模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTicketResponseTemplates true "删除快捷回复模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTicketResponseTemplates/deleteCrmTicketResponseTemplates [delete]
func (crmTicketResponseTemplatesApi *CrmTicketResponseTemplatesApi) DeleteCrmTicketResponseTemplates(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmTicketResponseTemplatesService.DeleteCrmTicketResponseTemplates(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmTicketResponseTemplatesByIds 批量删除快捷回复模板
// @Tags CrmTicketResponseTemplates
// @Summary 批量删除快捷回复模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmTicketResponseTemplates/deleteCrmTicketResponseTemplatesByIds [delete]
func (crmTicketResponseTemplatesApi *CrmTicketResponseTemplatesApi) DeleteCrmTicketResponseTemplatesByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmTicketResponseTemplatesService.DeleteCrmTicketResponseTemplatesByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmTicketResponseTemplates 更新快捷回复模板
// @Tags CrmTicketResponseTemplates
// @Summary 更新快捷回复模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTicketResponseTemplates true "更新快捷回复模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmTicketResponseTemplates/updateCrmTicketResponseTemplates [put]
func (crmTicketResponseTemplatesApi *CrmTicketResponseTemplatesApi) UpdateCrmTicketResponseTemplates(c *gin.Context) {
	var crmTicketResponseTemplates crm.CrmTicketResponseTemplates
	err := c.ShouldBindJSON(&crmTicketResponseTemplates)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmTicketResponseTemplatesService.UpdateCrmTicketResponseTemplates(crmTicketResponseTemplates); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmTicketResponseTemplates 用id查询快捷回复模板
// @Tags CrmTicketResponseTemplates
// @Summary 用id查询快捷回复模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmTicketResponseTemplates true "用id查询快捷回复模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmTicketResponseTemplates/findCrmTicketResponseTemplates [get]
func (crmTicketResponseTemplatesApi *CrmTicketResponseTemplatesApi) FindCrmTicketResponseTemplates(c *gin.Context) {
	ID := c.Query("ID")
	if recrmTicketResponseTemplates, err := crmTicketResponseTemplatesService.GetCrmTicketResponseTemplates(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmTicketResponseTemplates": recrmTicketResponseTemplates}, c)
	}
}

// GetCrmTicketResponseTemplatesList 分页获取快捷回复模板列表
// @Tags CrmTicketResponseTemplates
// @Summary 分页获取快捷回复模板列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmTicketResponseTemplatesSearch true "分页获取快捷回复模板列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTicketResponseTemplates/getCrmTicketResponseTemplatesList [get]
func (crmTicketResponseTemplatesApi *CrmTicketResponseTemplatesApi) GetCrmTicketResponseTemplatesList(c *gin.Context) {
	var pageInfo crmReq.CrmTicketResponseTemplatesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmTicketResponseTemplatesService.GetCrmTicketResponseTemplatesInfoList(pageInfo); err != nil {
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

// GetCrmTicketResponseTemplatesPublic 不需要鉴权的快捷回复模板接口
// @Tags CrmTicketResponseTemplates
// @Summary 不需要鉴权的快捷回复模板接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmTicketResponseTemplatesSearch true "分页获取快捷回复模板列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTicketResponseTemplates/getCrmTicketResponseTemplatesList [get]
func (crmTicketResponseTemplatesApi *CrmTicketResponseTemplatesApi) GetCrmTicketResponseTemplatesPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的快捷回复模板接口信息",
	}, "获取成功", c)
}
