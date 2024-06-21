package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPageApprovalTasksRouter struct {
}

// InitCrmPageApprovalTasksRouter 初始化 审批任务 路由信息
func (s *CrmPageApprovalTasksRouter) InitCrmPageApprovalTasksRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmApprovalTasksRouter := Router.Group("crmApprovalTasks").Use(middleware.OperationRecord())
	crmApprovalTasksRouterWithoutRecord := Router.Group("crmApprovalTasks")

	var crmApprovalTasksApi = v1.ApiGroupApp.CrmApiGroup.CrmApprovalTasksApi
	{
		crmApprovalTasksRouter.PUT("updateCrmMultipleApprovalTasks", crmApprovalTasksApi.UpdateCrmMultipleApprovalTasks) // 更新审批任务
	}
	{
		crmApprovalTasksRouterWithoutRecord.GET("findCrmPageContactApprovalTasks", crmApprovalTasksApi.FindCrmPageContactApprovalTasks)                                       // 根据ID获取合同审批任务
		crmApprovalTasksRouterWithoutRecord.GET("findCrmPageBusinessOpportunityApprovalTasks", crmApprovalTasksApi.FindCrmPageBusinessOpportunityApprovalTasks)               // 根据ID获取商机审批任务
		crmApprovalTasksRouterWithoutRecord.GET("findCrmPagePaymentCollentionApprovalTasks", crmApprovalTasksApi.FindCrmPagePaymentCollentionApprovalTasks)                   // 根据ID获取回款审批任务
		crmApprovalTasksRouterWithoutRecord.GET("getCrmContractApprovalTasksList", crmApprovalTasksApi.GetCrmContractApprovalTasksList)                                       // 获取合同审批任务列表
		crmApprovalTasksRouterWithoutRecord.GET("getCrmBusinessOpportunityContractApprovalTasksList", crmApprovalTasksApi.GetCrmBusinessOpportunityContractApprovalTasksList) // 获取商机审批任务列表
		crmApprovalTasksRouterWithoutRecord.GET("getCrmPaymentCollentionApprovalTasksList", crmApprovalTasksApi.GetCrmPaymentCollentionApprovalTasksList)                     // 获取回款审批任务列表
		crmApprovalTasksRouterWithoutRecord.GET("getCrmOrderApprovalTasksList", crmApprovalTasksApi.GetCrmOrderApprovalTasksList)                                             // 获取回款审批任务列表
	}
}
