package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPageBusinessOpportunityRouter struct {
}

// InitCrmBusinessOpportunityRouter 初始化 crmBusinessOpportunity表 路由信息
func (s *CrmPageBusinessOpportunityRouter) InitCrmPageBusinessOpportunityRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmBusinessOpportunityRouter := Router.Group("crmBusinessOpportunity").Use(middleware.OperationRecord())
	crmBusinessOpportunityRouterWithoutRecord := Router.Group("crmBusinessOpportunity")
	crmBusinessOpportunityRouterWithoutAuth := PublicRouter.Group("crmBusinessOpportunity")

	var crmBusinessOpportunityApi = v1.ApiGroupApp.CrmApiGroup.CrmBusinessOpportunityApi

	{
		crmBusinessOpportunityRouter.POST("createCrmPageBusinessOpportunity", crmBusinessOpportunityApi.CreateCrmPageBusinessOpportunity) // 新建商机管理
	}
	{
		crmBusinessOpportunityRouterWithoutRecord.GET("findCrmPageBusinessOpportunity", crmBusinessOpportunityApi.FindCrmPageBusinessOpportunity) // 根据ID获取商机管理
	}

	{
		crmBusinessOpportunityRouterWithoutAuth.GET("getCrmPageBusinessOpportunityList", crmBusinessOpportunityApi.GetCrmPageBusinessOpportunityList) // 获取crmBusinessOpportunity表列表
	}
}
