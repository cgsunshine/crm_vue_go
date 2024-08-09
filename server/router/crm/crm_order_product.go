package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmOrderProductRouter struct {
}

// InitCrmOrderProductRouter 初始化 crmOrderProduct表 路由信息
func (s *CrmOrderProductRouter) InitCrmOrderProductRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmOrderProductRouter := Router.Group("crmOrderProduct").Use(middleware.OperationRecord())
	crmOrderProductRouterWithoutRecord := Router.Group("crmOrderProduct")
	crmOrderProductRouterWithoutAuth := PublicRouter.Group("crmOrderProduct")

	var crmOrderProductApi = v1.ApiGroupApp.CrmApiGroup.CrmOrderProductApi
	{
		crmOrderProductRouter.POST("createCrmOrderProduct", crmOrderProductApi.CreateCrmOrderProduct)   // 新建crmOrderProduct表
		crmOrderProductRouter.DELETE("deleteCrmOrderProduct", crmOrderProductApi.DeleteCrmOrderProduct) // 删除crmOrderProduct表
		crmOrderProductRouter.DELETE("deleteCrmOrderProductByIds", crmOrderProductApi.DeleteCrmOrderProductByIds) // 批量删除crmOrderProduct表
		crmOrderProductRouter.PUT("updateCrmOrderProduct", crmOrderProductApi.UpdateCrmOrderProduct)    // 更新crmOrderProduct表
	}
	{
		crmOrderProductRouterWithoutRecord.GET("findCrmOrderProduct", crmOrderProductApi.FindCrmOrderProduct)        // 根据ID获取crmOrderProduct表
		crmOrderProductRouterWithoutRecord.GET("getCrmOrderProductList", crmOrderProductApi.GetCrmOrderProductList)  // 获取crmOrderProduct表列表
	}
	{
	    crmOrderProductRouterWithoutAuth.GET("getCrmOrderProductPublic", crmOrderProductApi.GetCrmOrderProductPublic)  // 获取crmOrderProduct表列表
	}
}
