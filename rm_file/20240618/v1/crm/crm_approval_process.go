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

type CrmApprovalProcessApi struct {
}

var crmApprovalProcessService = service.ServiceGroupApp.CrmServiceGroup.CrmApprovalProcessService


// CreateCrmApprovalProcess 创建crmApprovalProcess表
// @Tags CrmApprovalProcess
// @Summary 创建crmApprovalProcess表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalProcess true "创建crmApprovalProcess表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalProcess/createCrmApprovalProcess [post]
func (crmApprovalProcessApi *CrmApprovalProcessApi) CreateCrmApprovalProcess(c *gin.Context) {
	var crmApprovalProcess crm.CrmApprovalProcess
	err := c.ShouldBindJSON(&crmApprovalProcess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmApprovalProcessService.CreateCrmApprovalProcess(&crmApprovalProcess); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmApprovalProcess 删除crmApprovalProcess表
// @Tags CrmApprovalProcess
// @Summary 删除crmApprovalProcess表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalProcess true "删除crmApprovalProcess表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalProcess/deleteCrmApprovalProcess [delete]
func (crmApprovalProcessApi *CrmApprovalProcessApi) DeleteCrmApprovalProcess(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmApprovalProcessService.DeleteCrmApprovalProcess(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmApprovalProcessByIds 批量删除crmApprovalProcess表
// @Tags CrmApprovalProcess
// @Summary 批量删除crmApprovalProcess表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmApprovalProcess/deleteCrmApprovalProcessByIds [delete]
func (crmApprovalProcessApi *CrmApprovalProcessApi) DeleteCrmApprovalProcessByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmApprovalProcessService.DeleteCrmApprovalProcessByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmApprovalProcess 更新crmApprovalProcess表
// @Tags CrmApprovalProcess
// @Summary 更新crmApprovalProcess表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalProcess true "更新crmApprovalProcess表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmApprovalProcess/updateCrmApprovalProcess [put]
func (crmApprovalProcessApi *CrmApprovalProcessApi) UpdateCrmApprovalProcess(c *gin.Context) {
	var crmApprovalProcess crm.CrmApprovalProcess
	err := c.ShouldBindJSON(&crmApprovalProcess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmApprovalProcessService.UpdateCrmApprovalProcess(crmApprovalProcess); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmApprovalProcess 用id查询crmApprovalProcess表
// @Tags CrmApprovalProcess
// @Summary 用id查询crmApprovalProcess表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmApprovalProcess true "用id查询crmApprovalProcess表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmApprovalProcess/findCrmApprovalProcess [get]
func (crmApprovalProcessApi *CrmApprovalProcessApi) FindCrmApprovalProcess(c *gin.Context) {
	ID := c.Query("ID")
	if recrmApprovalProcess, err := crmApprovalProcessService.GetCrmApprovalProcess(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmApprovalProcess": recrmApprovalProcess}, c)
	}
}

// GetCrmApprovalProcessList 分页获取crmApprovalProcess表列表
// @Tags CrmApprovalProcess
// @Summary 分页获取crmApprovalProcess表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalProcessSearch true "分页获取crmApprovalProcess表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalProcess/getCrmApprovalProcessList [get]
func (crmApprovalProcessApi *CrmApprovalProcessApi) GetCrmApprovalProcessList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalProcessSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmApprovalProcessService.GetCrmApprovalProcessInfoList(pageInfo); err != nil {
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

// GetCrmApprovalProcessPublic 不需要鉴权的crmApprovalProcess表接口
// @Tags CrmApprovalProcess
// @Summary 不需要鉴权的crmApprovalProcess表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalProcessSearch true "分页获取crmApprovalProcess表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalProcess/getCrmApprovalProcessList [get]
func (crmApprovalProcessApi *CrmApprovalProcessApi) GetCrmApprovalProcessPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的crmApprovalProcess表接口信息",
    }, "获取成功", c)
}
