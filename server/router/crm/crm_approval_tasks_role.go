package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmApprovalTasksRoleRouter struct {
}

// InitCrmApprovalTasksRoleRouter 初始化 角色审批表 路由信息
func (s *CrmApprovalTasksRoleRouter) InitCrmApprovalTasksRoleRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmApprovalTasksRoleRouter := Router.Group("crmApprovalTasksRole").Use(middleware.OperationRecord())
	crmApprovalTasksRoleRouterWithoutRecord := Router.Group("crmApprovalTasksRole")
	crmApprovalTasksRoleRouterWithoutAuth := PublicRouter.Group("crmApprovalTasksRole")

	var crmApprovalTasksRoleApi = v1.ApiGroupApp.CrmApiGroup.CrmApprovalTasksRoleApi
	{
		crmApprovalTasksRoleRouter.POST("createCrmApprovalTasksRole", crmApprovalTasksRoleApi.CreateCrmApprovalTasksRole)   // 新建角色审批表
		crmApprovalTasksRoleRouter.DELETE("deleteCrmApprovalTasksRole", crmApprovalTasksRoleApi.DeleteCrmApprovalTasksRole) // 删除角色审批表
		crmApprovalTasksRoleRouter.DELETE("deleteCrmApprovalTasksRoleByIds", crmApprovalTasksRoleApi.DeleteCrmApprovalTasksRoleByIds) // 批量删除角色审批表
		crmApprovalTasksRoleRouter.PUT("updateCrmApprovalTasksRole", crmApprovalTasksRoleApi.UpdateCrmApprovalTasksRole)    // 更新角色审批表
	}
	{
		crmApprovalTasksRoleRouterWithoutRecord.GET("findCrmApprovalTasksRole", crmApprovalTasksRoleApi.FindCrmApprovalTasksRole)        // 根据ID获取角色审批表
		crmApprovalTasksRoleRouterWithoutRecord.GET("getCrmApprovalTasksRoleList", crmApprovalTasksRoleApi.GetCrmApprovalTasksRoleList)  // 获取角色审批表列表
	}
	{
	    crmApprovalTasksRoleRouterWithoutAuth.GET("getCrmApprovalTasksRolePublic", crmApprovalTasksRoleApi.GetCrmApprovalTasksRolePublic)  // 获取角色审批表列表
	}
}
