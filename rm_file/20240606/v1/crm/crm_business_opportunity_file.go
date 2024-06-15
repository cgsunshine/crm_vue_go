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

type CrmBusinessOpportunityFileApi struct {
}

var crmBusinessOpportunityFileService = service.ServiceGroupApp.CrmServiceGroup.CrmBusinessOpportunityFileService


// CreateCrmBusinessOpportunityFile 创建商机文件
// @Tags CrmBusinessOpportunityFile
// @Summary 创建商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityFile true "创建商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunityFile/createCrmBusinessOpportunityFile [post]
func (crmBusinessOpportunityFileApi *CrmBusinessOpportunityFileApi) CreateCrmBusinessOpportunityFile(c *gin.Context) {
	var crmBusinessOpportunityFile crm.CrmBusinessOpportunityFile
	err := c.ShouldBindJSON(&crmBusinessOpportunityFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBusinessOpportunityFileService.CreateCrmBusinessOpportunityFile(&crmBusinessOpportunityFile); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmBusinessOpportunityFile 删除商机文件
// @Tags CrmBusinessOpportunityFile
// @Summary 删除商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityFile true "删除商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityFile/deleteCrmBusinessOpportunityFile [delete]
func (crmBusinessOpportunityFileApi *CrmBusinessOpportunityFileApi) DeleteCrmBusinessOpportunityFile(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmBusinessOpportunityFileService.DeleteCrmBusinessOpportunityFile(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmBusinessOpportunityFileByIds 批量删除商机文件
// @Tags CrmBusinessOpportunityFile
// @Summary 批量删除商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmBusinessOpportunityFile/deleteCrmBusinessOpportunityFileByIds [delete]
func (crmBusinessOpportunityFileApi *CrmBusinessOpportunityFileApi) DeleteCrmBusinessOpportunityFileByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmBusinessOpportunityFileService.DeleteCrmBusinessOpportunityFileByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmBusinessOpportunityFile 更新商机文件
// @Tags CrmBusinessOpportunityFile
// @Summary 更新商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityFile true "更新商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBusinessOpportunityFile/updateCrmBusinessOpportunityFile [put]
func (crmBusinessOpportunityFileApi *CrmBusinessOpportunityFileApi) UpdateCrmBusinessOpportunityFile(c *gin.Context) {
	var crmBusinessOpportunityFile crm.CrmBusinessOpportunityFile
	err := c.ShouldBindJSON(&crmBusinessOpportunityFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBusinessOpportunityFileService.UpdateCrmBusinessOpportunityFile(crmBusinessOpportunityFile); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmBusinessOpportunityFile 用id查询商机文件
// @Tags CrmBusinessOpportunityFile
// @Summary 用id查询商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBusinessOpportunityFile true "用id查询商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunityFile/findCrmBusinessOpportunityFile [get]
func (crmBusinessOpportunityFileApi *CrmBusinessOpportunityFileApi) FindCrmBusinessOpportunityFile(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBusinessOpportunityFile, err := crmBusinessOpportunityFileService.GetCrmBusinessOpportunityFile(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmBusinessOpportunityFile": recrmBusinessOpportunityFile}, c)
	}
}

// GetCrmBusinessOpportunityFileList 分页获取商机文件列表
// @Tags CrmBusinessOpportunityFile
// @Summary 分页获取商机文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunityFileSearch true "分页获取商机文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityFile/getCrmBusinessOpportunityFileList [get]
func (crmBusinessOpportunityFileApi *CrmBusinessOpportunityFileApi) GetCrmBusinessOpportunityFileList(c *gin.Context) {
	var pageInfo crmReq.CrmBusinessOpportunityFileSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmBusinessOpportunityFileService.GetCrmBusinessOpportunityFileInfoList(pageInfo); err != nil {
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

// GetCrmBusinessOpportunityFilePublic 不需要鉴权的商机文件接口
// @Tags CrmBusinessOpportunityFile
// @Summary 不需要鉴权的商机文件接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunityFileSearch true "分页获取商机文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityFile/getCrmBusinessOpportunityFileList [get]
func (crmBusinessOpportunityFileApi *CrmBusinessOpportunityFileApi) GetCrmBusinessOpportunityFilePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的商机文件接口信息",
    }, "获取成功", c)
}
