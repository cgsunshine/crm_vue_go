package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmOrderRouter struct {
}

// InitCrmOrderRouter 初始化 crmOrder表 路由信息
func (s *CrmOrderRouter) InitCrmOrderRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmOrderRouter := Router.Group("crmOrder").Use(middleware.OperationRecord())
	crmOrderRouterWithoutRecord := Router.Group("crmOrder")
	crmOrderRouterWithoutAuth := PublicRouter.Group("crmOrder")

	var crmOrderApi = v1.ApiGroupApp.CrmApiGroup.CrmOrderApi
	{
		crmOrderRouter.POST("createCrmOrder", crmOrderApi.CreateCrmOrder)             // 新建crmOrder表
		crmOrderRouter.DELETE("deleteCrmOrder", crmOrderApi.DeleteCrmOrder)           // 删除crmOrder表
		crmOrderRouter.DELETE("deleteCrmOrderByIds", crmOrderApi.DeleteCrmOrderByIds) // 批量删除crmOrder表
		crmOrderRouter.PUT("updateCrmOrder", crmOrderApi.UpdateCrmOrder)              // 更新crmOrder表
	}
	{
		crmOrderRouterWithoutRecord.GET("findCrmOrder", crmOrderApi.FindCrmOrder) // 根据ID获取crmOrder表
		//crmOrderRouterWithoutRecord.GET("getCrmOrderList", crmOrderApi.GetCrmOrderList) // 获取crmOrder表列表
	}
	{
		crmOrderRouterWithoutAuth.GET("getCrmOrderPublic", crmOrderApi.GetCrmOrderPublic) // 获取crmOrder表列表
		crmOrderRouterWithoutAuth.GET("getCrmOrderList", crmOrderApi.GetCrmOrderList)     // 获取crmOrder表列表
	}
}
