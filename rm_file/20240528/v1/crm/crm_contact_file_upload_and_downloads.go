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

type CrmContactFileUploadAndDownloadsApi struct {
}

var crmContactFileUploadAndDownloadsService = service.ServiceGroupApp.CrmServiceGroup.CrmContactFileUploadAndDownloadsService

// CreateCrmContactFileUploadAndDownloads 创建crmContactFileUploadAndDownloads表
// @Tags CrmContactFileUploadAndDownloads
// @Summary 创建crmContactFileUploadAndDownloads表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactFileUploadAndDownloads true "创建crmContactFileUploadAndDownloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContactFileUploadAndDownloads/createCrmContactFileUploadAndDownloads [post]
func (crmContactFileUploadAndDownloadsApi *CrmContactFileUploadAndDownloadsApi) CreateCrmContactFileUploadAndDownloads(c *gin.Context) {
	var crmContactFileUploadAndDownloads crm.CrmContactFileUploadAndDownloads
	err := c.ShouldBindJSON(&crmContactFileUploadAndDownloads)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmContactFileUploadAndDownloadsService.CreateCrmContactFileUploadAndDownloads(&crmContactFileUploadAndDownloads); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmContactFileUploadAndDownloads 删除crmContactFileUploadAndDownloads表
// @Tags CrmContactFileUploadAndDownloads
// @Summary 删除crmContactFileUploadAndDownloads表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactFileUploadAndDownloads true "删除crmContactFileUploadAndDownloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContactFileUploadAndDownloads/deleteCrmContactFileUploadAndDownloads [delete]
func (crmContactFileUploadAndDownloadsApi *CrmContactFileUploadAndDownloadsApi) DeleteCrmContactFileUploadAndDownloads(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmContactFileUploadAndDownloadsService.DeleteCrmContactFileUploadAndDownloads(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmContactFileUploadAndDownloadsByIds 批量删除crmContactFileUploadAndDownloads表
// @Tags CrmContactFileUploadAndDownloads
// @Summary 批量删除crmContactFileUploadAndDownloads表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmContactFileUploadAndDownloads/deleteCrmContactFileUploadAndDownloadsByIds [delete]
func (crmContactFileUploadAndDownloadsApi *CrmContactFileUploadAndDownloadsApi) DeleteCrmContactFileUploadAndDownloadsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmContactFileUploadAndDownloadsService.DeleteCrmContactFileUploadAndDownloadsByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmContactFileUploadAndDownloads 更新crmContactFileUploadAndDownloads表
// @Tags CrmContactFileUploadAndDownloads
// @Summary 更新crmContactFileUploadAndDownloads表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactFileUploadAndDownloads true "更新crmContactFileUploadAndDownloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactFileUploadAndDownloads/updateCrmContactFileUploadAndDownloads [put]
func (crmContactFileUploadAndDownloadsApi *CrmContactFileUploadAndDownloadsApi) UpdateCrmContactFileUploadAndDownloads(c *gin.Context) {
	var crmContactFileUploadAndDownloads crm.CrmContactFileUploadAndDownloads
	err := c.ShouldBindJSON(&crmContactFileUploadAndDownloads)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmContactFileUploadAndDownloadsService.UpdateCrmContactFileUploadAndDownloads(crmContactFileUploadAndDownloads); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmContactFileUploadAndDownloads 用id查询crmContactFileUploadAndDownloads表
// @Tags CrmContactFileUploadAndDownloads
// @Summary 用id查询crmContactFileUploadAndDownloads表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContactFileUploadAndDownloads true "用id查询crmContactFileUploadAndDownloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactFileUploadAndDownloads/findCrmContactFileUploadAndDownloads [get]
func (crmContactFileUploadAndDownloadsApi *CrmContactFileUploadAndDownloadsApi) FindCrmContactFileUploadAndDownloads(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactFileUploadAndDownloads, err := crmContactFileUploadAndDownloadsService.GetCrmContactFileUploadAndDownloads(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContactFileUploadAndDownloads": recrmContactFileUploadAndDownloads}, c)
	}
}

// GetCrmContactFileUploadAndDownloadsList 分页获取crmContactFileUploadAndDownloads表列表
// @Tags CrmContactFileUploadAndDownloads
// @Summary 分页获取crmContactFileUploadAndDownloads表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContactFileUploadAndDownloadsSearch true "分页获取crmContactFileUploadAndDownloads表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactFileUploadAndDownloads/getCrmContactFileUploadAndDownloadsList [get]
func (crmContactFileUploadAndDownloadsApi *CrmContactFileUploadAndDownloadsApi) GetCrmContactFileUploadAndDownloadsList(c *gin.Context) {
	var pageInfo crmReq.CrmContactFileUploadAndDownloadsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmContactFileUploadAndDownloadsService.GetCrmContactFileUploadAndDownloadsInfoList(pageInfo); err != nil {
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

// GetCrmContactFileUploadAndDownloadsPublic 不需要鉴权的crmContactFileUploadAndDownloads表接口
// @Tags CrmContactFileUploadAndDownloads
// @Summary 不需要鉴权的crmContactFileUploadAndDownloads表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContactFileUploadAndDownloadsSearch true "分页获取crmContactFileUploadAndDownloads表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactFileUploadAndDownloads/getCrmContactFileUploadAndDownloadsList [get]
func (crmContactFileUploadAndDownloadsApi *CrmContactFileUploadAndDownloadsApi) GetCrmContactFileUploadAndDownloadsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的crmContactFileUploadAndDownloads表接口信息",
	}, "获取成功", c)
}
