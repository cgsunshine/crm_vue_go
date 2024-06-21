package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPaymentCollentionApprovalTasksRouter struct {
}

// InitCrmPaymentCollentionApprovalTasksRouter 初始化 crmPaymentCollentionApprovalTasks表 路由信息
func (s *CrmPaymentCollentionApprovalTasksRouter) InitCrmPaymentCollentionApprovalTasksRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmPaymentCollentionApprovalTasksRouter := Router.Group("crmPaymentCollentionApprovalTasks").Use(middleware.OperationRecord())
	crmPaymentCollentionApprovalTasksRouterWithoutRecord := Router.Group("crmPaymentCollentionApprovalTasks")
	crmPaymentCollentionApprovalTasksRouterWithoutAuth := PublicRouter.Group("crmPaymentCollentionApprovalTasks")

	var crmPaymentCollentionApprovalTasksApi = v1.ApiGroupApp.CrmApiGroup.CrmPaymentCollentionApprovalTasksApi
	{
		crmPaymentCollentionApprovalTasksRouter.POST("createCrmPaymentCollentionApprovalTasks", crmPaymentCollentionApprovalTasksApi.CreateCrmPaymentCollentionApprovalTasks)   // 新建crmPaymentCollentionApprovalTasks表
		crmPaymentCollentionApprovalTasksRouter.DELETE("deleteCrmPaymentCollentionApprovalTasks", crmPaymentCollentionApprovalTasksApi.DeleteCrmPaymentCollentionApprovalTasks) // 删除crmPaymentCollentionApprovalTasks表
		crmPaymentCollentionApprovalTasksRouter.DELETE("deleteCrmPaymentCollentionApprovalTasksByIds", crmPaymentCollentionApprovalTasksApi.DeleteCrmPaymentCollentionApprovalTasksByIds) // 批量删除crmPaymentCollentionApprovalTasks表
		crmPaymentCollentionApprovalTasksRouter.PUT("updateCrmPaymentCollentionApprovalTasks", crmPaymentCollentionApprovalTasksApi.UpdateCrmPaymentCollentionApprovalTasks)    // 更新crmPaymentCollentionApprovalTasks表
	}
	{
		crmPaymentCollentionApprovalTasksRouterWithoutRecord.GET("findCrmPaymentCollentionApprovalTasks", crmPaymentCollentionApprovalTasksApi.FindCrmPaymentCollentionApprovalTasks)        // 根据ID获取crmPaymentCollentionApprovalTasks表
		crmPaymentCollentionApprovalTasksRouterWithoutRecord.GET("getCrmPaymentCollentionApprovalTasksList", crmPaymentCollentionApprovalTasksApi.GetCrmPaymentCollentionApprovalTasksList)  // 获取crmPaymentCollentionApprovalTasks表列表
	}
	{
	    crmPaymentCollentionApprovalTasksRouterWithoutAuth.GET("getCrmPaymentCollentionApprovalTasksPublic", crmPaymentCollentionApprovalTasksApi.GetCrmPaymentCollentionApprovalTasksPublic)  // 获取crmPaymentCollentionApprovalTasks表列表
	}
}
