package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmProductRouter struct {
}

// InitCrmProductRouter 初始化 crmProduct表 路由信息
func (s *CrmProductRouter) InitCrmProductRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmProductRouter := Router.Group("crmProduct").Use(middleware.OperationRecord())
	crmProductRouterWithoutRecord := Router.Group("crmProduct")
	crmProductRouterWithoutAuth := PublicRouter.Group("crmProduct")

	var crmProductApi = v1.ApiGroupApp.CrmApiGroup.CrmProductApi
	{
		crmProductRouter.POST("createCrmProduct", crmProductApi.CreateCrmProduct)             // 新建crmProduct表
		crmProductRouter.DELETE("deleteCrmProduct", crmProductApi.DeleteCrmProduct)           // 删除crmProduct表
		crmProductRouter.DELETE("deleteCrmProductByIds", crmProductApi.DeleteCrmProductByIds) // 批量删除crmProduct表
		crmProductRouter.PUT("updateCrmProduct", crmProductApi.UpdateCrmProduct)              // 更新crmProduct表
	}
	{
		crmProductRouterWithoutRecord.GET("findCrmProduct", crmProductApi.FindCrmProduct) // 根据ID获取crmProduct表
		//crmProductRouterWithoutRecord.GET("getCrmProductList", crmProductApi.GetCrmProductList) // 获取crmProduct表列表
	}
	{
		crmProductRouterWithoutAuth.GET("getCrmProductPublic", crmProductApi.GetCrmProductPublic) // 获取crmProduct表列表
		crmProductRouterWithoutAuth.GET("getCrmProductList", crmProductApi.GetCrmProductList)     // 获取crmProduct表列表
	}
}
