package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmBusinessOpportunityRouter struct {
}

// InitCrmBusinessOpportunityRouter 初始化 商机管理 路由信息
func (s *CrmBusinessOpportunityRouter) InitCrmBusinessOpportunityRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmBusinessOpportunityRouter := Router.Group("crmBusinessOpportunity").Use(middleware.OperationRecord())
	crmBusinessOpportunityRouterWithoutRecord := Router.Group("crmBusinessOpportunity")
	crmBusinessOpportunityRouterWithoutAuth := PublicRouter.Group("crmBusinessOpportunity")

	var crmBusinessOpportunityApi = v1.ApiGroupApp.CrmApiGroup.CrmBusinessOpportunityApi
	{
		crmBusinessOpportunityRouter.POST("createCrmBusinessOpportunity", crmBusinessOpportunityApi.CreateCrmBusinessOpportunity)   // 新建商机管理
		crmBusinessOpportunityRouter.DELETE("deleteCrmBusinessOpportunity", crmBusinessOpportunityApi.DeleteCrmBusinessOpportunity) // 删除商机管理
		crmBusinessOpportunityRouter.DELETE("deleteCrmBusinessOpportunityByIds", crmBusinessOpportunityApi.DeleteCrmBusinessOpportunityByIds) // 批量删除商机管理
		crmBusinessOpportunityRouter.PUT("updateCrmBusinessOpportunity", crmBusinessOpportunityApi.UpdateCrmBusinessOpportunity)    // 更新商机管理
	}
	{
		crmBusinessOpportunityRouterWithoutRecord.GET("findCrmBusinessOpportunity", crmBusinessOpportunityApi.FindCrmBusinessOpportunity)        // 根据ID获取商机管理
		crmBusinessOpportunityRouterWithoutRecord.GET("getCrmBusinessOpportunityList", crmBusinessOpportunityApi.GetCrmBusinessOpportunityList)  // 获取商机管理列表
	}
	{
	    crmBusinessOpportunityRouterWithoutAuth.GET("getCrmBusinessOpportunityPublic", crmBusinessOpportunityApi.GetCrmBusinessOpportunityPublic)  // 获取商机管理列表
	}
}
