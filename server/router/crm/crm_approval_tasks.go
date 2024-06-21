package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmApprovalTasksRouter struct {
}

// InitCrmApprovalTasksRouter 初始化 审批任务 路由信息
func (s *CrmApprovalTasksRouter) InitCrmApprovalTasksRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmApprovalTasksRouter := Router.Group("crmApprovalTasks").Use(middleware.OperationRecord())
	crmApprovalTasksRouterWithoutRecord := Router.Group("crmApprovalTasks")
	crmApprovalTasksRouterWithoutAuth := PublicRouter.Group("crmApprovalTasks")

	var crmApprovalTasksApi = v1.ApiGroupApp.CrmApiGroup.CrmApprovalTasksApi
	{
		crmApprovalTasksRouter.POST("createCrmApprovalTasks", crmApprovalTasksApi.CreateCrmApprovalTasks)   // 新建审批任务
		crmApprovalTasksRouter.DELETE("deleteCrmApprovalTasks", crmApprovalTasksApi.DeleteCrmApprovalTasks) // 删除审批任务
		crmApprovalTasksRouter.DELETE("deleteCrmApprovalTasksByIds", crmApprovalTasksApi.DeleteCrmApprovalTasksByIds) // 批量删除审批任务
		crmApprovalTasksRouter.PUT("updateCrmApprovalTasks", crmApprovalTasksApi.UpdateCrmApprovalTasks)    // 更新审批任务
	}
	{
		crmApprovalTasksRouterWithoutRecord.GET("findCrmApprovalTasks", crmApprovalTasksApi.FindCrmApprovalTasks)        // 根据ID获取审批任务
		crmApprovalTasksRouterWithoutRecord.GET("getCrmApprovalTasksList", crmApprovalTasksApi.GetCrmApprovalTasksList)  // 获取审批任务列表
	}
	{
	    crmApprovalTasksRouterWithoutAuth.GET("getCrmApprovalTasksPublic", crmApprovalTasksApi.GetCrmApprovalTasksPublic)  // 获取审批任务列表
	}
}
