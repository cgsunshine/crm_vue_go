package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmBusinessOpportunityApprovalTasksRouter struct {
}

// InitCrmBusinessOpportunityApprovalTasksRouter 初始化 商机审批 路由信息
func (s *CrmBusinessOpportunityApprovalTasksRouter) InitCrmBusinessOpportunityApprovalTasksRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmBusinessOpportunityApprovalTasksRouter := Router.Group("crmBusinessOpportunityApprovalTasks").Use(middleware.OperationRecord())
	crmBusinessOpportunityApprovalTasksRouterWithoutRecord := Router.Group("crmBusinessOpportunityApprovalTasks")
	crmBusinessOpportunityApprovalTasksRouterWithoutAuth := PublicRouter.Group("crmBusinessOpportunityApprovalTasks")

	var crmBusinessOpportunityApprovalTasksApi = v1.ApiGroupApp.CrmApiGroup.CrmBusinessOpportunityApprovalTasksApi
	{
		crmBusinessOpportunityApprovalTasksRouter.POST("createCrmBusinessOpportunityApprovalTasks", crmBusinessOpportunityApprovalTasksApi.CreateCrmBusinessOpportunityApprovalTasks)   // 新建商机审批
		crmBusinessOpportunityApprovalTasksRouter.DELETE("deleteCrmBusinessOpportunityApprovalTasks", crmBusinessOpportunityApprovalTasksApi.DeleteCrmBusinessOpportunityApprovalTasks) // 删除商机审批
		crmBusinessOpportunityApprovalTasksRouter.DELETE("deleteCrmBusinessOpportunityApprovalTasksByIds", crmBusinessOpportunityApprovalTasksApi.DeleteCrmBusinessOpportunityApprovalTasksByIds) // 批量删除商机审批
		crmBusinessOpportunityApprovalTasksRouter.PUT("updateCrmBusinessOpportunityApprovalTasks", crmBusinessOpportunityApprovalTasksApi.UpdateCrmBusinessOpportunityApprovalTasks)    // 更新商机审批
	}
	{
		crmBusinessOpportunityApprovalTasksRouterWithoutRecord.GET("findCrmBusinessOpportunityApprovalTasks", crmBusinessOpportunityApprovalTasksApi.FindCrmBusinessOpportunityApprovalTasks)        // 根据ID获取商机审批
		crmBusinessOpportunityApprovalTasksRouterWithoutRecord.GET("getCrmBusinessOpportunityApprovalTasksList", crmBusinessOpportunityApprovalTasksApi.GetCrmBusinessOpportunityApprovalTasksList)  // 获取商机审批列表
	}
	{
	    crmBusinessOpportunityApprovalTasksRouterWithoutAuth.GET("getCrmBusinessOpportunityApprovalTasksPublic", crmBusinessOpportunityApprovalTasksApi.GetCrmBusinessOpportunityApprovalTasksPublic)  // 获取商机审批列表
	}
}
