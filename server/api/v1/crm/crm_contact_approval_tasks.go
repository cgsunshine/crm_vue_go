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

type CrmContactApprovalTasksApi struct {
}

var crmContactApprovalTasksService = service.ServiceGroupApp.CrmServiceGroup.CrmContactApprovalTasksService


// CreateCrmContactApprovalTasks 创建合同审批
// @Tags CrmContactApprovalTasks
// @Summary 创建合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactApprovalTasks true "创建合同审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContactApprovalTasks/createCrmContactApprovalTasks [post]
func (crmContactApprovalTasksApi *CrmContactApprovalTasksApi) CreateCrmContactApprovalTasks(c *gin.Context) {
	var crmContactApprovalTasks crm.CrmContactApprovalTasks
	err := c.ShouldBindJSON(&crmContactApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmContactApprovalTasksService.CreateCrmContactApprovalTasks(&crmContactApprovalTasks); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmContactApprovalTasks 删除合同审批
// @Tags CrmContactApprovalTasks
// @Summary 删除合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactApprovalTasks true "删除合同审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContactApprovalTasks/deleteCrmContactApprovalTasks [delete]
func (crmContactApprovalTasksApi *CrmContactApprovalTasksApi) DeleteCrmContactApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmContactApprovalTasksService.DeleteCrmContactApprovalTasks(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmContactApprovalTasksByIds 批量删除合同审批
// @Tags CrmContactApprovalTasks
// @Summary 批量删除合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmContactApprovalTasks/deleteCrmContactApprovalTasksByIds [delete]
func (crmContactApprovalTasksApi *CrmContactApprovalTasksApi) DeleteCrmContactApprovalTasksByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmContactApprovalTasksService.DeleteCrmContactApprovalTasksByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmContactApprovalTasks 更新合同审批
// @Tags CrmContactApprovalTasks
// @Summary 更新合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactApprovalTasks true "更新合同审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactApprovalTasks/updateCrmContactApprovalTasks [put]
func (crmContactApprovalTasksApi *CrmContactApprovalTasksApi) UpdateCrmContactApprovalTasks(c *gin.Context) {
	var crmContactApprovalTasks crm.CrmContactApprovalTasks
	err := c.ShouldBindJSON(&crmContactApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmContactApprovalTasksService.UpdateCrmContactApprovalTasks(crmContactApprovalTasks); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmContactApprovalTasks 用id查询合同审批
// @Tags CrmContactApprovalTasks
// @Summary 用id查询合同审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContactApprovalTasks true "用id查询合同审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactApprovalTasks/findCrmContactApprovalTasks [get]
func (crmContactApprovalTasksApi *CrmContactApprovalTasksApi) FindCrmContactApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactApprovalTasks, err := crmContactApprovalTasksService.GetCrmContactApprovalTasks(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContactApprovalTasks": recrmContactApprovalTasks}, c)
	}
}

// GetCrmContactApprovalTasksList 分页获取合同审批列表
// @Tags CrmContactApprovalTasks
// @Summary 分页获取合同审批列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContactApprovalTasksSearch true "分页获取合同审批列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactApprovalTasks/getCrmContactApprovalTasksList [get]
func (crmContactApprovalTasksApi *CrmContactApprovalTasksApi) GetCrmContactApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmContactApprovalTasksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmContactApprovalTasksService.GetCrmContactApprovalTasksInfoList(pageInfo); err != nil {
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

// GetCrmContactApprovalTasksPublic 不需要鉴权的合同审批接口
// @Tags CrmContactApprovalTasks
// @Summary 不需要鉴权的合同审批接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContactApprovalTasksSearch true "分页获取合同审批列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactApprovalTasks/getCrmContactApprovalTasksList [get]
func (crmContactApprovalTasksApi *CrmContactApprovalTasksApi) GetCrmContactApprovalTasksPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的合同审批接口信息",
    }, "获取成功", c)
}
