package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmBusinessOpportunityProductRouter struct {
}

// InitCrmBusinessOpportunityProductRouter 初始化 crmBusinessOpportunityProduct表 路由信息
func (s *CrmBusinessOpportunityProductRouter) InitCrmBusinessOpportunityProductRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmBusinessOpportunityProductRouter := Router.Group("crmBusinessOpportunityProduct").Use(middleware.OperationRecord())
	crmBusinessOpportunityProductRouterWithoutRecord := Router.Group("crmBusinessOpportunityProduct")
	crmBusinessOpportunityProductRouterWithoutAuth := PublicRouter.Group("crmBusinessOpportunityProduct")

	var crmBusinessOpportunityProductApi = v1.ApiGroupApp.CrmApiGroup.CrmBusinessOpportunityProductApi
	{
		crmBusinessOpportunityProductRouter.POST("createCrmBusinessOpportunityProduct", crmBusinessOpportunityProductApi.CreateCrmBusinessOpportunityProduct)   // 新建crmBusinessOpportunityProduct表
		crmBusinessOpportunityProductRouter.DELETE("deleteCrmBusinessOpportunityProduct", crmBusinessOpportunityProductApi.DeleteCrmBusinessOpportunityProduct) // 删除crmBusinessOpportunityProduct表
		crmBusinessOpportunityProductRouter.DELETE("deleteCrmBusinessOpportunityProductByIds", crmBusinessOpportunityProductApi.DeleteCrmBusinessOpportunityProductByIds) // 批量删除crmBusinessOpportunityProduct表
		crmBusinessOpportunityProductRouter.PUT("updateCrmBusinessOpportunityProduct", crmBusinessOpportunityProductApi.UpdateCrmBusinessOpportunityProduct)    // 更新crmBusinessOpportunityProduct表
	}
	{
		crmBusinessOpportunityProductRouterWithoutRecord.GET("findCrmBusinessOpportunityProduct", crmBusinessOpportunityProductApi.FindCrmBusinessOpportunityProduct)        // 根据ID获取crmBusinessOpportunityProduct表
		crmBusinessOpportunityProductRouterWithoutRecord.GET("getCrmBusinessOpportunityProductList", crmBusinessOpportunityProductApi.GetCrmBusinessOpportunityProductList)  // 获取crmBusinessOpportunityProduct表列表
	}
	{
	    crmBusinessOpportunityProductRouterWithoutAuth.GET("getCrmBusinessOpportunityProductPublic", crmBusinessOpportunityProductApi.GetCrmBusinessOpportunityProductPublic)  // 获取crmBusinessOpportunityProduct表列表
	}
}
