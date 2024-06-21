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

type CrmBusinessOpportunityApprovalTasksApi struct {
}

var crmBusinessOpportunityApprovalTasksService = service.ServiceGroupApp.CrmServiceGroup.CrmBusinessOpportunityApprovalTasksService


// CreateCrmBusinessOpportunityApprovalTasks 创建商机审批
// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 创建商机审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityApprovalTasks true "创建商机审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/createCrmBusinessOpportunityApprovalTasks [post]
func (crmBusinessOpportunityApprovalTasksApi *CrmBusinessOpportunityApprovalTasksApi) CreateCrmBusinessOpportunityApprovalTasks(c *gin.Context) {
	var crmBusinessOpportunityApprovalTasks crm.CrmBusinessOpportunityApprovalTasks
	err := c.ShouldBindJSON(&crmBusinessOpportunityApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBusinessOpportunityApprovalTasksService.CreateCrmBusinessOpportunityApprovalTasks(&crmBusinessOpportunityApprovalTasks); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmBusinessOpportunityApprovalTasks 删除商机审批
// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 删除商机审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityApprovalTasks true "删除商机审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/deleteCrmBusinessOpportunityApprovalTasks [delete]
func (crmBusinessOpportunityApprovalTasksApi *CrmBusinessOpportunityApprovalTasksApi) DeleteCrmBusinessOpportunityApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmBusinessOpportunityApprovalTasksService.DeleteCrmBusinessOpportunityApprovalTasks(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmBusinessOpportunityApprovalTasksByIds 批量删除商机审批
// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 批量删除商机审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/deleteCrmBusinessOpportunityApprovalTasksByIds [delete]
func (crmBusinessOpportunityApprovalTasksApi *CrmBusinessOpportunityApprovalTasksApi) DeleteCrmBusinessOpportunityApprovalTasksByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmBusinessOpportunityApprovalTasksService.DeleteCrmBusinessOpportunityApprovalTasksByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmBusinessOpportunityApprovalTasks 更新商机审批
// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 更新商机审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityApprovalTasks true "更新商机审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/updateCrmBusinessOpportunityApprovalTasks [put]
func (crmBusinessOpportunityApprovalTasksApi *CrmBusinessOpportunityApprovalTasksApi) UpdateCrmBusinessOpportunityApprovalTasks(c *gin.Context) {
	var crmBusinessOpportunityApprovalTasks crm.CrmBusinessOpportunityApprovalTasks
	err := c.ShouldBindJSON(&crmBusinessOpportunityApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBusinessOpportunityApprovalTasksService.UpdateCrmBusinessOpportunityApprovalTasks(crmBusinessOpportunityApprovalTasks); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmBusinessOpportunityApprovalTasks 用id查询商机审批
// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 用id查询商机审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBusinessOpportunityApprovalTasks true "用id查询商机审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/findCrmBusinessOpportunityApprovalTasks [get]
func (crmBusinessOpportunityApprovalTasksApi *CrmBusinessOpportunityApprovalTasksApi) FindCrmBusinessOpportunityApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBusinessOpportunityApprovalTasks, err := crmBusinessOpportunityApprovalTasksService.GetCrmBusinessOpportunityApprovalTasks(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmBusinessOpportunityApprovalTasks": recrmBusinessOpportunityApprovalTasks}, c)
	}
}

// GetCrmBusinessOpportunityApprovalTasksList 分页获取商机审批列表
// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 分页获取商机审批列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunityApprovalTasksSearch true "分页获取商机审批列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/getCrmBusinessOpportunityApprovalTasksList [get]
func (crmBusinessOpportunityApprovalTasksApi *CrmBusinessOpportunityApprovalTasksApi) GetCrmBusinessOpportunityApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmBusinessOpportunityApprovalTasksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmBusinessOpportunityApprovalTasksService.GetCrmBusinessOpportunityApprovalTasksInfoList(pageInfo); err != nil {
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

// GetCrmBusinessOpportunityApprovalTasksPublic 不需要鉴权的商机审批接口
// @Tags CrmBusinessOpportunityApprovalTasks
// @Summary 不需要鉴权的商机审批接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunityApprovalTasksSearch true "分页获取商机审批列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityApprovalTasks/getCrmBusinessOpportunityApprovalTasksList [get]
func (crmBusinessOpportunityApprovalTasksApi *CrmBusinessOpportunityApprovalTasksApi) GetCrmBusinessOpportunityApprovalTasksPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的商机审批接口信息",
    }, "获取成功", c)
}
