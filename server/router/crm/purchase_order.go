package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPagePurchaseOrderRouter struct {
}

// InitCrmPurchaseOrderRouter 初始化 crmPurchaseOrder表 路由信息
func (s *CrmPagePurchaseOrderRouter) InitCrmPagePurchaseOrderRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmPurchaseOrderRouterWithoutRecord := Router.Group("crmPurchaseOrder")
	crmPurchaseOrderRouterWithoutAuth := PublicRouter.Group("crmPurchaseOrder")

	var crmPurchaseOrderApi = v1.ApiGroupApp.CrmApiGroup.CrmPurchaseOrderApi

	{
		crmPurchaseOrderRouterWithoutRecord.GET("findCrmPagePurchaseOrder", crmPurchaseOrderApi.FindCrmPagePurchaseOrder) // 根据ID获取订购单管理
	}

	{
		crmPurchaseOrderRouterWithoutAuth.GET("getCrmPagePurchaseOrderList", crmPurchaseOrderApi.GetCrmPagePurchaseOrderList) // 获取crmPurchaseOrder表列表
	}

}
