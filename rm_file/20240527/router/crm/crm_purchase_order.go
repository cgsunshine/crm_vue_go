package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPurchaseOrderRouter struct {
}

// InitCrmPurchaseOrderRouter 初始化 crmPurchaseOrder表 路由信息
func (s *CrmPurchaseOrderRouter) InitCrmPurchaseOrderRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmPurchaseOrderRouter := Router.Group("crmPurchaseOrder").Use(middleware.OperationRecord())
	crmPurchaseOrderRouterWithoutRecord := Router.Group("crmPurchaseOrder")
	crmPurchaseOrderRouterWithoutAuth := PublicRouter.Group("crmPurchaseOrder")

	var crmPurchaseOrderApi = v1.ApiGroupApp.CrmApiGroup.CrmPurchaseOrderApi
	{
		crmPurchaseOrderRouter.POST("createCrmPurchaseOrder", crmPurchaseOrderApi.CreateCrmPurchaseOrder)   // 新建crmPurchaseOrder表
		crmPurchaseOrderRouter.DELETE("deleteCrmPurchaseOrder", crmPurchaseOrderApi.DeleteCrmPurchaseOrder) // 删除crmPurchaseOrder表
		crmPurchaseOrderRouter.DELETE("deleteCrmPurchaseOrderByIds", crmPurchaseOrderApi.DeleteCrmPurchaseOrderByIds) // 批量删除crmPurchaseOrder表
		crmPurchaseOrderRouter.PUT("updateCrmPurchaseOrder", crmPurchaseOrderApi.UpdateCrmPurchaseOrder)    // 更新crmPurchaseOrder表
	}
	{
		crmPurchaseOrderRouterWithoutRecord.GET("findCrmPurchaseOrder", crmPurchaseOrderApi.FindCrmPurchaseOrder)        // 根据ID获取crmPurchaseOrder表
		crmPurchaseOrderRouterWithoutRecord.GET("getCrmPurchaseOrderList", crmPurchaseOrderApi.GetCrmPurchaseOrderList)  // 获取crmPurchaseOrder表列表
	}
	{
	    crmPurchaseOrderRouterWithoutAuth.GET("getCrmPurchaseOrderPublic", crmPurchaseOrderApi.GetCrmPurchaseOrderPublic)  // 获取crmPurchaseOrder表列表
	}
}
