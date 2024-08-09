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

type CrmContactFileApi struct {
}

var crmContactFileService = service.ServiceGroupApp.CrmServiceGroup.CrmContactFileService


// CreateCrmContactFile 创建crmContactFile表
// @Tags CrmContactFile
// @Summary 创建crmContactFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactFile true "创建crmContactFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContactFile/createCrmContactFile [post]
func (crmContactFileApi *CrmContactFileApi) CreateCrmContactFile(c *gin.Context) {
	var crmContactFile crm.CrmContactFile
	err := c.ShouldBindJSON(&crmContactFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmContactFileService.CreateCrmContactFile(&crmContactFile); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmContactFile 删除crmContactFile表
// @Tags CrmContactFile
// @Summary 删除crmContactFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactFile true "删除crmContactFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContactFile/deleteCrmContactFile [delete]
func (crmContactFileApi *CrmContactFileApi) DeleteCrmContactFile(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmContactFileService.DeleteCrmContactFile(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmContactFileByIds 批量删除crmContactFile表
// @Tags CrmContactFile
// @Summary 批量删除crmContactFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmContactFile/deleteCrmContactFileByIds [delete]
func (crmContactFileApi *CrmContactFileApi) DeleteCrmContactFileByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmContactFileService.DeleteCrmContactFileByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmContactFile 更新crmContactFile表
// @Tags CrmContactFile
// @Summary 更新crmContactFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactFile true "更新crmContactFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactFile/updateCrmContactFile [put]
func (crmContactFileApi *CrmContactFileApi) UpdateCrmContactFile(c *gin.Context) {
	var crmContactFile crm.CrmContactFile
	err := c.ShouldBindJSON(&crmContactFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmContactFileService.UpdateCrmContactFile(crmContactFile); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmContactFile 用id查询crmContactFile表
// @Tags CrmContactFile
// @Summary 用id查询crmContactFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContactFile true "用id查询crmContactFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactFile/findCrmContactFile [get]
func (crmContactFileApi *CrmContactFileApi) FindCrmContactFile(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactFile, err := crmContactFileService.GetCrmContactFile(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContactFile": recrmContactFile}, c)
	}
}

// GetCrmContactFileList 分页获取crmContactFile表列表
// @Tags CrmContactFile
// @Summary 分页获取crmContactFile表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContactFileSearch true "分页获取crmContactFile表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactFile/getCrmContactFileList [get]
func (crmContactFileApi *CrmContactFileApi) GetCrmContactFileList(c *gin.Context) {
	var pageInfo crmReq.CrmContactFileSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmContactFileService.GetCrmContactFileInfoList(pageInfo); err != nil {
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

// GetCrmContactFilePublic 不需要鉴权的crmContactFile表接口
// @Tags CrmContactFile
// @Summary 不需要鉴权的crmContactFile表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContactFileSearch true "分页获取crmContactFile表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactFile/getCrmContactFileList [get]
func (crmContactFileApi *CrmContactFileApi) GetCrmContactFilePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的crmContactFile表接口信息",
    }, "获取成功", c)
}
