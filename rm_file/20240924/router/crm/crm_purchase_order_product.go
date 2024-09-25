package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPurchaseOrderProductRouter struct {
}

// InitCrmPurchaseOrderProductRouter 初始化 crmPurchaseOrderProduct表 路由信息
func (s *CrmPurchaseOrderProductRouter) InitCrmPurchaseOrderProductRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmPurchaseOrderProductRouter := Router.Group("crmPurchaseOrderProduct").Use(middleware.OperationRecord())
	crmPurchaseOrderProductRouterWithoutRecord := Router.Group("crmPurchaseOrderProduct")
	crmPurchaseOrderProductRouterWithoutAuth := PublicRouter.Group("crmPurchaseOrderProduct")

	var crmPurchaseOrderProductApi = v1.ApiGroupApp.CrmApiGroup.CrmPurchaseOrderProductApi
	{
		crmPurchaseOrderProductRouter.POST("createCrmPurchaseOrderProduct", crmPurchaseOrderProductApi.CreateCrmPurchaseOrderProduct)   // 新建crmPurchaseOrderProduct表
		crmPurchaseOrderProductRouter.DELETE("deleteCrmPurchaseOrderProduct", crmPurchaseOrderProductApi.DeleteCrmPurchaseOrderProduct) // 删除crmPurchaseOrderProduct表
		crmPurchaseOrderProductRouter.DELETE("deleteCrmPurchaseOrderProductByIds", crmPurchaseOrderProductApi.DeleteCrmPurchaseOrderProductByIds) // 批量删除crmPurchaseOrderProduct表
		crmPurchaseOrderProductRouter.PUT("updateCrmPurchaseOrderProduct", crmPurchaseOrderProductApi.UpdateCrmPurchaseOrderProduct)    // 更新crmPurchaseOrderProduct表
	}
	{
		crmPurchaseOrderProductRouterWithoutRecord.GET("findCrmPurchaseOrderProduct", crmPurchaseOrderProductApi.FindCrmPurchaseOrderProduct)        // 根据ID获取crmPurchaseOrderProduct表
		crmPurchaseOrderProductRouterWithoutRecord.GET("getCrmPurchaseOrderProductList", crmPurchaseOrderProductApi.GetCrmPurchaseOrderProductList)  // 获取crmPurchaseOrderProduct表列表
	}
	{
	    crmPurchaseOrderProductRouterWithoutAuth.GET("getCrmPurchaseOrderProductPublic", crmPurchaseOrderProductApi.GetCrmPurchaseOrderProductPublic)  // 获取crmPurchaseOrderProduct表列表
	}
}
