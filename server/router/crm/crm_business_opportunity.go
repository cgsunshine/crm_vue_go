package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmBusinessOpportunityRouter struct {
}

// InitCrmBusinessOpportunityRouter 初始化 crmBusinessOpportunity表 路由信息
func (s *CrmBusinessOpportunityRouter) InitCrmBusinessOpportunityRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmBusinessOpportunityRouter := Router.Group("crmBusinessOpportunity").Use(middleware.OperationRecord())
	crmBusinessOpportunityRouterWithoutRecord := Router.Group("crmBusinessOpportunity")
	crmBusinessOpportunityRouterWithoutAuth := PublicRouter.Group("crmBusinessOpportunity")

	var crmBusinessOpportunityApi = v1.ApiGroupApp.CrmApiGroup.CrmBusinessOpportunityApi
	{
		crmBusinessOpportunityRouter.POST("createCrmBusinessOpportunity", crmBusinessOpportunityApi.CreateCrmBusinessOpportunity)             // 新建crmBusinessOpportunity表
		crmBusinessOpportunityRouter.DELETE("deleteCrmBusinessOpportunity", crmBusinessOpportunityApi.DeleteCrmBusinessOpportunity)           // 删除crmBusinessOpportunity表
		crmBusinessOpportunityRouter.DELETE("deleteCrmBusinessOpportunityByIds", crmBusinessOpportunityApi.DeleteCrmBusinessOpportunityByIds) // 批量删除crmBusinessOpportunity表
		crmBusinessOpportunityRouter.PUT("updateCrmBusinessOpportunity", crmBusinessOpportunityApi.UpdateCrmBusinessOpportunity)              // 更新crmBusinessOpportunity表
	}
	{
		crmBusinessOpportunityRouterWithoutRecord.GET("findCrmBusinessOpportunity", crmBusinessOpportunityApi.FindCrmBusinessOpportunity)       // 根据ID获取crmBusinessOpportunity表
		crmBusinessOpportunityRouterWithoutRecord.GET("getCrmBusinessOpportunityList", crmBusinessOpportunityApi.GetCrmBusinessOpportunityList) // 获取crmBusinessOpportunity表列表

	}
	{
		crmBusinessOpportunityRouterWithoutAuth.GET("getCrmBusinessOpportunityPublic", crmBusinessOpportunityApi.GetCrmBusinessOpportunityPublic) // 获取crmBusinessOpportunity表列表
	}
}
