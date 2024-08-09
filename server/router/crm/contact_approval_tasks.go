package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPageContactApprovalTasksRouter struct {
}

// InitCrmContactApprovalTasksRouter 初始化 合同审批 路由信息
func (s *CrmPageContactApprovalTasksRouter) InitCrmPageContactApprovalTasksRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmContactApprovalTasksRouter := Router.Group("crmContactApprovalTasks").Use(middleware.OperationRecord())
	crmContactApprovalTasksRouterWithoutRecord := Router.Group("crmContactApprovalTasks")
	crmContactApprovalTasksRouterWithoutAuth := PublicRouter.Group("crmContactApprovalTasks")

	var crmContactApprovalTasksApi = v1.ApiGroupApp.CrmApiGroup.CrmContactApprovalTasksApi
	{
		crmContactApprovalTasksRouter.POST("createCrmMultipleContactApprovalTasks", crmContactApprovalTasksApi.CreateCrmMultipleContactApprovalTasks) // 新建合同审批

		crmContactApprovalTasksRouter.PUT("updateCrmMultipleContactApprovalTasks", crmContactApprovalTasksApi.UpdateCrmMultipleContactApprovalTasks) // 更新合同审批
	}

	{
		crmContactApprovalTasksRouterWithoutRecord.GET("findCrmPageContactApprovalTasks", crmContactApprovalTasksApi.FindCrmPageContactApprovalTasks) // 根据ID获取合同审批
	}

	{
		crmContactApprovalTasksRouterWithoutAuth.GET("getCrmPageContactApprovalTasksList", crmContactApprovalTasksApi.GetCrmPageContactApprovalTasksList) // 获取合同审批列表
	}

}
