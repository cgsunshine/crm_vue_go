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

type CrmApprovalTasksApi struct {
}

var crmApprovalTasksService = service.ServiceGroupApp.CrmServiceGroup.CrmApprovalTasksService


// CreateCrmApprovalTasks 创建审批任务
// @Tags CrmApprovalTasks
// @Summary 创建审批任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalTasks true "创建审批任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalTasks/createCrmApprovalTasks [post]
func (crmApprovalTasksApi *CrmApprovalTasksApi) CreateCrmApprovalTasks(c *gin.Context) {
	var crmApprovalTasks crm.CrmApprovalTasks
	err := c.ShouldBindJSON(&crmApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmApprovalTasksService.CreateCrmApprovalTasks(&crmApprovalTasks); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmApprovalTasks 删除审批任务
// @Tags CrmApprovalTasks
// @Summary 删除审批任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalTasks true "删除审批任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalTasks/deleteCrmApprovalTasks [delete]
func (crmApprovalTasksApi *CrmApprovalTasksApi) DeleteCrmApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmApprovalTasksService.DeleteCrmApprovalTasks(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmApprovalTasksByIds 批量删除审批任务
// @Tags CrmApprovalTasks
// @Summary 批量删除审批任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmApprovalTasks/deleteCrmApprovalTasksByIds [delete]
func (crmApprovalTasksApi *CrmApprovalTasksApi) DeleteCrmApprovalTasksByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmApprovalTasksService.DeleteCrmApprovalTasksByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmApprovalTasks 更新审批任务
// @Tags CrmApprovalTasks
// @Summary 更新审批任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalTasks true "更新审批任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmApprovalTasks/updateCrmApprovalTasks [put]
func (crmApprovalTasksApi *CrmApprovalTasksApi) UpdateCrmApprovalTasks(c *gin.Context) {
	var crmApprovalTasks crm.CrmApprovalTasks
	err := c.ShouldBindJSON(&crmApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmApprovalTasksService.UpdateCrmApprovalTasks(crmApprovalTasks); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmApprovalTasks 用id查询审批任务
// @Tags CrmApprovalTasks
// @Summary 用id查询审批任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmApprovalTasks true "用id查询审批任务"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmApprovalTasks/findCrmApprovalTasks [get]
func (crmApprovalTasksApi *CrmApprovalTasksApi) FindCrmApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmApprovalTasks, err := crmApprovalTasksService.GetCrmApprovalTasks(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmApprovalTasks": recrmApprovalTasks}, c)
	}
}

// GetCrmApprovalTasksList 分页获取审批任务列表
// @Tags CrmApprovalTasks
// @Summary 分页获取审批任务列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksSearch true "分页获取审批任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasks/getCrmApprovalTasksList [get]
func (crmApprovalTasksApi *CrmApprovalTasksApi) GetCrmApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalTasksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmApprovalTasksService.GetCrmApprovalTasksInfoList(pageInfo); err != nil {
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

// GetCrmApprovalTasksPublic 不需要鉴权的审批任务接口
// @Tags CrmApprovalTasks
// @Summary 不需要鉴权的审批任务接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalTasksSearch true "分页获取审批任务列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalTasks/getCrmApprovalTasksList [get]
func (crmApprovalTasksApi *CrmApprovalTasksApi) GetCrmApprovalTasksPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的审批任务接口信息",
    }, "获取成功", c)
}
