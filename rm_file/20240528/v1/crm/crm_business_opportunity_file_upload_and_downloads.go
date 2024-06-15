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

type CrmBusinessOpportunityFileUploadAndDownloadsApi struct {
}

var crmBusinessOpportunityFileUploadAndDownloadsService = service.ServiceGroupApp.CrmServiceGroup.CrmBusinessOpportunityFileUploadAndDownloadsService


// CreateCrmBusinessOpportunityFileUploadAndDownloads 创建上传商机文件
// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 创建上传商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityFileUploadAndDownloads true "创建上传商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/createCrmBusinessOpportunityFileUploadAndDownloads [post]
func (crmBusinessOpportunityFileUploadAndDownloadsApi *CrmBusinessOpportunityFileUploadAndDownloadsApi) CreateCrmBusinessOpportunityFileUploadAndDownloads(c *gin.Context) {
	var crmBusinessOpportunityFileUploadAndDownloads crm.CrmBusinessOpportunityFileUploadAndDownloads
	err := c.ShouldBindJSON(&crmBusinessOpportunityFileUploadAndDownloads)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBusinessOpportunityFileUploadAndDownloadsService.CreateCrmBusinessOpportunityFileUploadAndDownloads(&crmBusinessOpportunityFileUploadAndDownloads); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmBusinessOpportunityFileUploadAndDownloads 删除上传商机文件
// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 删除上传商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityFileUploadAndDownloads true "删除上传商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/deleteCrmBusinessOpportunityFileUploadAndDownloads [delete]
func (crmBusinessOpportunityFileUploadAndDownloadsApi *CrmBusinessOpportunityFileUploadAndDownloadsApi) DeleteCrmBusinessOpportunityFileUploadAndDownloads(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmBusinessOpportunityFileUploadAndDownloadsService.DeleteCrmBusinessOpportunityFileUploadAndDownloads(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmBusinessOpportunityFileUploadAndDownloadsByIds 批量删除上传商机文件
// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 批量删除上传商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/deleteCrmBusinessOpportunityFileUploadAndDownloadsByIds [delete]
func (crmBusinessOpportunityFileUploadAndDownloadsApi *CrmBusinessOpportunityFileUploadAndDownloadsApi) DeleteCrmBusinessOpportunityFileUploadAndDownloadsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmBusinessOpportunityFileUploadAndDownloadsService.DeleteCrmBusinessOpportunityFileUploadAndDownloadsByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmBusinessOpportunityFileUploadAndDownloads 更新上传商机文件
// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 更新上传商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityFileUploadAndDownloads true "更新上传商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/updateCrmBusinessOpportunityFileUploadAndDownloads [put]
func (crmBusinessOpportunityFileUploadAndDownloadsApi *CrmBusinessOpportunityFileUploadAndDownloadsApi) UpdateCrmBusinessOpportunityFileUploadAndDownloads(c *gin.Context) {
	var crmBusinessOpportunityFileUploadAndDownloads crm.CrmBusinessOpportunityFileUploadAndDownloads
	err := c.ShouldBindJSON(&crmBusinessOpportunityFileUploadAndDownloads)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBusinessOpportunityFileUploadAndDownloadsService.UpdateCrmBusinessOpportunityFileUploadAndDownloads(crmBusinessOpportunityFileUploadAndDownloads); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmBusinessOpportunityFileUploadAndDownloads 用id查询上传商机文件
// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 用id查询上传商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBusinessOpportunityFileUploadAndDownloads true "用id查询上传商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/findCrmBusinessOpportunityFileUploadAndDownloads [get]
func (crmBusinessOpportunityFileUploadAndDownloadsApi *CrmBusinessOpportunityFileUploadAndDownloadsApi) FindCrmBusinessOpportunityFileUploadAndDownloads(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBusinessOpportunityFileUploadAndDownloads, err := crmBusinessOpportunityFileUploadAndDownloadsService.GetCrmBusinessOpportunityFileUploadAndDownloads(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmBusinessOpportunityFileUploadAndDownloads": recrmBusinessOpportunityFileUploadAndDownloads}, c)
	}
}

// GetCrmBusinessOpportunityFileUploadAndDownloadsList 分页获取上传商机文件列表
// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 分页获取上传商机文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunityFileUploadAndDownloadsSearch true "分页获取上传商机文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/getCrmBusinessOpportunityFileUploadAndDownloadsList [get]
func (crmBusinessOpportunityFileUploadAndDownloadsApi *CrmBusinessOpportunityFileUploadAndDownloadsApi) GetCrmBusinessOpportunityFileUploadAndDownloadsList(c *gin.Context) {
	var pageInfo crmReq.CrmBusinessOpportunityFileUploadAndDownloadsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmBusinessOpportunityFileUploadAndDownloadsService.GetCrmBusinessOpportunityFileUploadAndDownloadsInfoList(pageInfo); err != nil {
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

// GetCrmBusinessOpportunityFileUploadAndDownloadsPublic 不需要鉴权的上传商机文件接口
// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 不需要鉴权的上传商机文件接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunityFileUploadAndDownloadsSearch true "分页获取上传商机文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/getCrmBusinessOpportunityFileUploadAndDownloadsList [get]
func (crmBusinessOpportunityFileUploadAndDownloadsApi *CrmBusinessOpportunityFileUploadAndDownloadsApi) GetCrmBusinessOpportunityFileUploadAndDownloadsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的上传商机文件接口信息",
    }, "获取成功", c)
}
