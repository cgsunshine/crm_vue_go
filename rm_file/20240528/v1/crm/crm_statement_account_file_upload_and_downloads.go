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

type CrmStatementAccountFileUploadAndDownloadsApi struct {
}

var crmStatementAccountFileUploadAndDownloadsService = service.ServiceGroupApp.CrmServiceGroup.CrmStatementAccountFileUploadAndDownloadsService


// CreateCrmStatementAccountFileUploadAndDownloads 创建对账单上传文件
// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 创建对账单上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmStatementAccountFileUploadAndDownloads true "创建对账单上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/createCrmStatementAccountFileUploadAndDownloads [post]
func (crmStatementAccountFileUploadAndDownloadsApi *CrmStatementAccountFileUploadAndDownloadsApi) CreateCrmStatementAccountFileUploadAndDownloads(c *gin.Context) {
	var crmStatementAccountFileUploadAndDownloads crm.CrmStatementAccountFileUploadAndDownloads
	err := c.ShouldBindJSON(&crmStatementAccountFileUploadAndDownloads)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmStatementAccountFileUploadAndDownloadsService.CreateCrmStatementAccountFileUploadAndDownloads(&crmStatementAccountFileUploadAndDownloads); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmStatementAccountFileUploadAndDownloads 删除对账单上传文件
// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 删除对账单上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmStatementAccountFileUploadAndDownloads true "删除对账单上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/deleteCrmStatementAccountFileUploadAndDownloads [delete]
func (crmStatementAccountFileUploadAndDownloadsApi *CrmStatementAccountFileUploadAndDownloadsApi) DeleteCrmStatementAccountFileUploadAndDownloads(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmStatementAccountFileUploadAndDownloadsService.DeleteCrmStatementAccountFileUploadAndDownloads(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmStatementAccountFileUploadAndDownloadsByIds 批量删除对账单上传文件
// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 批量删除对账单上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/deleteCrmStatementAccountFileUploadAndDownloadsByIds [delete]
func (crmStatementAccountFileUploadAndDownloadsApi *CrmStatementAccountFileUploadAndDownloadsApi) DeleteCrmStatementAccountFileUploadAndDownloadsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmStatementAccountFileUploadAndDownloadsService.DeleteCrmStatementAccountFileUploadAndDownloadsByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmStatementAccountFileUploadAndDownloads 更新对账单上传文件
// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 更新对账单上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmStatementAccountFileUploadAndDownloads true "更新对账单上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/updateCrmStatementAccountFileUploadAndDownloads [put]
func (crmStatementAccountFileUploadAndDownloadsApi *CrmStatementAccountFileUploadAndDownloadsApi) UpdateCrmStatementAccountFileUploadAndDownloads(c *gin.Context) {
	var crmStatementAccountFileUploadAndDownloads crm.CrmStatementAccountFileUploadAndDownloads
	err := c.ShouldBindJSON(&crmStatementAccountFileUploadAndDownloads)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmStatementAccountFileUploadAndDownloadsService.UpdateCrmStatementAccountFileUploadAndDownloads(crmStatementAccountFileUploadAndDownloads); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmStatementAccountFileUploadAndDownloads 用id查询对账单上传文件
// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 用id查询对账单上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmStatementAccountFileUploadAndDownloads true "用id查询对账单上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/findCrmStatementAccountFileUploadAndDownloads [get]
func (crmStatementAccountFileUploadAndDownloadsApi *CrmStatementAccountFileUploadAndDownloadsApi) FindCrmStatementAccountFileUploadAndDownloads(c *gin.Context) {
	ID := c.Query("ID")
	if recrmStatementAccountFileUploadAndDownloads, err := crmStatementAccountFileUploadAndDownloadsService.GetCrmStatementAccountFileUploadAndDownloads(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmStatementAccountFileUploadAndDownloads": recrmStatementAccountFileUploadAndDownloads}, c)
	}
}

// GetCrmStatementAccountFileUploadAndDownloadsList 分页获取对账单上传文件列表
// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 分页获取对账单上传文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmStatementAccountFileUploadAndDownloadsSearch true "分页获取对账单上传文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/getCrmStatementAccountFileUploadAndDownloadsList [get]
func (crmStatementAccountFileUploadAndDownloadsApi *CrmStatementAccountFileUploadAndDownloadsApi) GetCrmStatementAccountFileUploadAndDownloadsList(c *gin.Context) {
	var pageInfo crmReq.CrmStatementAccountFileUploadAndDownloadsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmStatementAccountFileUploadAndDownloadsService.GetCrmStatementAccountFileUploadAndDownloadsInfoList(pageInfo); err != nil {
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

// GetCrmStatementAccountFileUploadAndDownloadsPublic 不需要鉴权的对账单上传文件接口
// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 不需要鉴权的对账单上传文件接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmStatementAccountFileUploadAndDownloadsSearch true "分页获取对账单上传文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/getCrmStatementAccountFileUploadAndDownloadsList [get]
func (crmStatementAccountFileUploadAndDownloadsApi *CrmStatementAccountFileUploadAndDownloadsApi) GetCrmStatementAccountFileUploadAndDownloadsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的对账单上传文件接口信息",
    }, "获取成功", c)
}
