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

type CrmPaymentCollentionApprovalTasksApi struct {
}

var crmPaymentCollentionApprovalTasksService = service.ServiceGroupApp.CrmServiceGroup.CrmPaymentCollentionApprovalTasksService


// CreateCrmPaymentCollentionApprovalTasks 创建crmPaymentCollentionApprovalTasks表
// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 创建crmPaymentCollentionApprovalTasks表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPaymentCollentionApprovalTasks true "创建crmPaymentCollentionApprovalTasks表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPaymentCollentionApprovalTasks/createCrmPaymentCollentionApprovalTasks [post]
func (crmPaymentCollentionApprovalTasksApi *CrmPaymentCollentionApprovalTasksApi) CreateCrmPaymentCollentionApprovalTasks(c *gin.Context) {
	var crmPaymentCollentionApprovalTasks crm.CrmPaymentCollentionApprovalTasks
	err := c.ShouldBindJSON(&crmPaymentCollentionApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmPaymentCollentionApprovalTasksService.CreateCrmPaymentCollentionApprovalTasks(&crmPaymentCollentionApprovalTasks); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmPaymentCollentionApprovalTasks 删除crmPaymentCollentionApprovalTasks表
// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 删除crmPaymentCollentionApprovalTasks表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPaymentCollentionApprovalTasks true "删除crmPaymentCollentionApprovalTasks表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPaymentCollentionApprovalTasks/deleteCrmPaymentCollentionApprovalTasks [delete]
func (crmPaymentCollentionApprovalTasksApi *CrmPaymentCollentionApprovalTasksApi) DeleteCrmPaymentCollentionApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmPaymentCollentionApprovalTasksService.DeleteCrmPaymentCollentionApprovalTasks(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmPaymentCollentionApprovalTasksByIds 批量删除crmPaymentCollentionApprovalTasks表
// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 批量删除crmPaymentCollentionApprovalTasks表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmPaymentCollentionApprovalTasks/deleteCrmPaymentCollentionApprovalTasksByIds [delete]
func (crmPaymentCollentionApprovalTasksApi *CrmPaymentCollentionApprovalTasksApi) DeleteCrmPaymentCollentionApprovalTasksByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmPaymentCollentionApprovalTasksService.DeleteCrmPaymentCollentionApprovalTasksByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmPaymentCollentionApprovalTasks 更新crmPaymentCollentionApprovalTasks表
// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 更新crmPaymentCollentionApprovalTasks表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPaymentCollentionApprovalTasks true "更新crmPaymentCollentionApprovalTasks表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmPaymentCollentionApprovalTasks/updateCrmPaymentCollentionApprovalTasks [put]
func (crmPaymentCollentionApprovalTasksApi *CrmPaymentCollentionApprovalTasksApi) UpdateCrmPaymentCollentionApprovalTasks(c *gin.Context) {
	var crmPaymentCollentionApprovalTasks crm.CrmPaymentCollentionApprovalTasks
	err := c.ShouldBindJSON(&crmPaymentCollentionApprovalTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmPaymentCollentionApprovalTasksService.UpdateCrmPaymentCollentionApprovalTasks(crmPaymentCollentionApprovalTasks); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmPaymentCollentionApprovalTasks 用id查询crmPaymentCollentionApprovalTasks表
// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 用id查询crmPaymentCollentionApprovalTasks表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmPaymentCollentionApprovalTasks true "用id查询crmPaymentCollentionApprovalTasks表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPaymentCollentionApprovalTasks/findCrmPaymentCollentionApprovalTasks [get]
func (crmPaymentCollentionApprovalTasksApi *CrmPaymentCollentionApprovalTasksApi) FindCrmPaymentCollentionApprovalTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmPaymentCollentionApprovalTasks, err := crmPaymentCollentionApprovalTasksService.GetCrmPaymentCollentionApprovalTasks(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmPaymentCollentionApprovalTasks": recrmPaymentCollentionApprovalTasks}, c)
	}
}

// GetCrmPaymentCollentionApprovalTasksList 分页获取crmPaymentCollentionApprovalTasks表列表
// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 分页获取crmPaymentCollentionApprovalTasks表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPaymentCollentionApprovalTasksSearch true "分页获取crmPaymentCollentionApprovalTasks表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPaymentCollentionApprovalTasks/getCrmPaymentCollentionApprovalTasksList [get]
func (crmPaymentCollentionApprovalTasksApi *CrmPaymentCollentionApprovalTasksApi) GetCrmPaymentCollentionApprovalTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmPaymentCollentionApprovalTasksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmPaymentCollentionApprovalTasksService.GetCrmPaymentCollentionApprovalTasksInfoList(pageInfo); err != nil {
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

// GetCrmPaymentCollentionApprovalTasksPublic 不需要鉴权的crmPaymentCollentionApprovalTasks表接口
// @Tags CrmPaymentCollentionApprovalTasks
// @Summary 不需要鉴权的crmPaymentCollentionApprovalTasks表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPaymentCollentionApprovalTasksSearch true "分页获取crmPaymentCollentionApprovalTasks表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPaymentCollentionApprovalTasks/getCrmPaymentCollentionApprovalTasksList [get]
func (crmPaymentCollentionApprovalTasksApi *CrmPaymentCollentionApprovalTasksApi) GetCrmPaymentCollentionApprovalTasksPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的crmPaymentCollentionApprovalTasks表接口信息",
    }, "获取成功", c)
}
