package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmApprovalProcessRouter struct {
}

// InitCrmApprovalProcessRouter 初始化 审批流程 路由信息
func (s *CrmApprovalProcessRouter) InitCrmApprovalProcessRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmApprovalProcessRouter := Router.Group("crmApprovalProcess").Use(middleware.OperationRecord())
	crmApprovalProcessRouterWithoutRecord := Router.Group("crmApprovalProcess")
	crmApprovalProcessRouterWithoutAuth := PublicRouter.Group("crmApprovalProcess")

	var crmApprovalProcessApi = v1.ApiGroupApp.CrmApiGroup.CrmApprovalProcessApi
	{
		crmApprovalProcessRouter.POST("createCrmApprovalProcess", crmApprovalProcessApi.CreateCrmApprovalProcess)   // 新建审批流程
		crmApprovalProcessRouter.DELETE("deleteCrmApprovalProcess", crmApprovalProcessApi.DeleteCrmApprovalProcess) // 删除审批流程
		crmApprovalProcessRouter.DELETE("deleteCrmApprovalProcessByIds", crmApprovalProcessApi.DeleteCrmApprovalProcessByIds) // 批量删除审批流程
		crmApprovalProcessRouter.PUT("updateCrmApprovalProcess", crmApprovalProcessApi.UpdateCrmApprovalProcess)    // 更新审批流程
	}
	{
		crmApprovalProcessRouterWithoutRecord.GET("findCrmApprovalProcess", crmApprovalProcessApi.FindCrmApprovalProcess)        // 根据ID获取审批流程
		crmApprovalProcessRouterWithoutRecord.GET("getCrmApprovalProcessList", crmApprovalProcessApi.GetCrmApprovalProcessList)  // 获取审批流程列表
	}
	{
	    crmApprovalProcessRouterWithoutAuth.GET("getCrmApprovalProcessPublic", crmApprovalProcessApi.GetCrmApprovalProcessPublic)  // 获取审批流程列表
	}
}
