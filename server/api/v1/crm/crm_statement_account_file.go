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

type CrmStatementAccountFileApi struct {
}

var crmStatementAccountFileService = service.ServiceGroupApp.CrmServiceGroup.CrmStatementAccountFileService


// CreateCrmStatementAccountFile 创建对账单文件
// @Tags CrmStatementAccountFile
// @Summary 创建对账单文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmStatementAccountFile true "创建对账单文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmStatementAccountFile/createCrmStatementAccountFile [post]
func (crmStatementAccountFileApi *CrmStatementAccountFileApi) CreateCrmStatementAccountFile(c *gin.Context) {
	var crmStatementAccountFile crm.CrmStatementAccountFile
	err := c.ShouldBindJSON(&crmStatementAccountFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmStatementAccountFileService.CreateCrmStatementAccountFile(&crmStatementAccountFile); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmStatementAccountFile 删除对账单文件
// @Tags CrmStatementAccountFile
// @Summary 删除对账单文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmStatementAccountFile true "删除对账单文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmStatementAccountFile/deleteCrmStatementAccountFile [delete]
func (crmStatementAccountFileApi *CrmStatementAccountFileApi) DeleteCrmStatementAccountFile(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmStatementAccountFileService.DeleteCrmStatementAccountFile(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmStatementAccountFileByIds 批量删除对账单文件
// @Tags CrmStatementAccountFile
// @Summary 批量删除对账单文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmStatementAccountFile/deleteCrmStatementAccountFileByIds [delete]
func (crmStatementAccountFileApi *CrmStatementAccountFileApi) DeleteCrmStatementAccountFileByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmStatementAccountFileService.DeleteCrmStatementAccountFileByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmStatementAccountFile 更新对账单文件
// @Tags CrmStatementAccountFile
// @Summary 更新对账单文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmStatementAccountFile true "更新对账单文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmStatementAccountFile/updateCrmStatementAccountFile [put]
func (crmStatementAccountFileApi *CrmStatementAccountFileApi) UpdateCrmStatementAccountFile(c *gin.Context) {
	var crmStatementAccountFile crm.CrmStatementAccountFile
	err := c.ShouldBindJSON(&crmStatementAccountFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmStatementAccountFileService.UpdateCrmStatementAccountFile(crmStatementAccountFile); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmStatementAccountFile 用id查询对账单文件
// @Tags CrmStatementAccountFile
// @Summary 用id查询对账单文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmStatementAccountFile true "用id查询对账单文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmStatementAccountFile/findCrmStatementAccountFile [get]
func (crmStatementAccountFileApi *CrmStatementAccountFileApi) FindCrmStatementAccountFile(c *gin.Context) {
	ID := c.Query("ID")
	if recrmStatementAccountFile, err := crmStatementAccountFileService.GetCrmStatementAccountFile(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmStatementAccountFile": recrmStatementAccountFile}, c)
	}
}

// GetCrmStatementAccountFileList 分页获取对账单文件列表
// @Tags CrmStatementAccountFile
// @Summary 分页获取对账单文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmStatementAccountFileSearch true "分页获取对账单文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmStatementAccountFile/getCrmStatementAccountFileList [get]
func (crmStatementAccountFileApi *CrmStatementAccountFileApi) GetCrmStatementAccountFileList(c *gin.Context) {
	var pageInfo crmReq.CrmStatementAccountFileSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmStatementAccountFileService.GetCrmStatementAccountFileInfoList(pageInfo); err != nil {
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

// GetCrmStatementAccountFilePublic 不需要鉴权的对账单文件接口
// @Tags CrmStatementAccountFile
// @Summary 不需要鉴权的对账单文件接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmStatementAccountFileSearch true "分页获取对账单文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmStatementAccountFile/getCrmStatementAccountFileList [get]
func (crmStatementAccountFileApi *CrmStatementAccountFileApi) GetCrmStatementAccountFilePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的对账单文件接口信息",
    }, "获取成功", c)
}
