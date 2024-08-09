package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmContactApprovalTasksRouter struct {
}

// InitCrmContactApprovalTasksRouter 初始化 合同审批 路由信息
func (s *CrmContactApprovalTasksRouter) InitCrmContactApprovalTasksRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmContactApprovalTasksRouter := Router.Group("crmContactApprovalTasks").Use(middleware.OperationRecord())
	crmContactApprovalTasksRouterWithoutRecord := Router.Group("crmContactApprovalTasks")
	crmContactApprovalTasksRouterWithoutAuth := PublicRouter.Group("crmContactApprovalTasks")

	var crmContactApprovalTasksApi = v1.ApiGroupApp.CrmApiGroup.CrmContactApprovalTasksApi
	{
		crmContactApprovalTasksRouter.POST("createCrmContactApprovalTasks", crmContactApprovalTasksApi.CreateCrmContactApprovalTasks)   // 新建合同审批
		crmContactApprovalTasksRouter.DELETE("deleteCrmContactApprovalTasks", crmContactApprovalTasksApi.DeleteCrmContactApprovalTasks) // 删除合同审批
		crmContactApprovalTasksRouter.DELETE("deleteCrmContactApprovalTasksByIds", crmContactApprovalTasksApi.DeleteCrmContactApprovalTasksByIds) // 批量删除合同审批
		crmContactApprovalTasksRouter.PUT("updateCrmContactApprovalTasks", crmContactApprovalTasksApi.UpdateCrmContactApprovalTasks)    // 更新合同审批
	}
	{
		crmContactApprovalTasksRouterWithoutRecord.GET("findCrmContactApprovalTasks", crmContactApprovalTasksApi.FindCrmContactApprovalTasks)        // 根据ID获取合同审批
		crmContactApprovalTasksRouterWithoutRecord.GET("getCrmContactApprovalTasksList", crmContactApprovalTasksApi.GetCrmContactApprovalTasksList)  // 获取合同审批列表
	}
	{
	    crmContactApprovalTasksRouterWithoutAuth.GET("getCrmContactApprovalTasksPublic", crmContactApprovalTasksApi.GetCrmContactApprovalTasksPublic)  // 获取合同审批列表
	}
}
