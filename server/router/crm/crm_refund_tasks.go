package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmRefundTasksRouter struct {
}

// InitCrmRefundTasksRouter 初始化 退款管理 路由信息
func (s *CrmRefundTasksRouter) InitCrmRefundTasksRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmRefundTasksRouter := Router.Group("crmRefundTasks").Use(middleware.OperationRecord())
	//crmRefundTasksRouterWithoutRecord := Router.Group("crmRefundTasks")
	crmRefundTasksRouterWithoutAuth := PublicRouter.Group("crmRefundTasks")

	var crmRefundTasksApi = v1.ApiGroupApp.CrmApiGroup.CrmRefundTasksApi
	{
		crmRefundTasksRouter.POST("createCrmRefundTasks", crmRefundTasksApi.CreateCrmRefundTasks)             // 新建退款管理
		crmRefundTasksRouter.DELETE("deleteCrmRefundTasks", crmRefundTasksApi.DeleteCrmRefundTasks)           // 删除退款管理
		crmRefundTasksRouter.DELETE("deleteCrmRefundTasksByIds", crmRefundTasksApi.DeleteCrmRefundTasksByIds) // 批量删除退款管理
		crmRefundTasksRouter.PUT("updateCrmRefundTasks", crmRefundTasksApi.UpdateCrmRefundTasks)              // 更新退款管理
	}
	{
		//crmRefundTasksRouterWithoutRecord.GET("findCrmRefundTasks", crmRefundTasksApi.FindCrmRefundTasks)        // 根据ID获取退款管理
		//crmRefundTasksRouterWithoutRecord.GET("getCrmRefundTasksList", crmRefundTasksApi.GetCrmRefundTasksList)  // 获取退款管理列表
	}
	{
		crmRefundTasksRouterWithoutAuth.GET("getCrmRefundTasksPublic", crmRefundTasksApi.GetCrmRefundTasksPublic)   // 获取退款管理列表
		crmRefundTasksRouterWithoutAuth.GET("findCrmRefundTasks", crmRefundTasksApi.FindCrmRefundTasks)             // 根据ID获取退款管理
		crmRefundTasksRouterWithoutAuth.GET("getCrmRefundTasksList", crmRefundTasksApi.GetCrmRefundTasksList)       // 获取退款管理列表
		crmRefundTasksRouterWithoutAuth.PUT("processingCrmRefundTasks", crmRefundTasksApi.ProcessingCrmRefundTasks) // 获取退款管理列表
	}
}
