package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmCustomerGroupRouter struct {
}

// InitCrmCustomerGroupRouter 初始化 crmCustomerGroup表 路由信息
func (s *CrmCustomerGroupRouter) InitCrmCustomerGroupRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmCustomerGroupRouter := Router.Group("crmCustomerGroup").Use(middleware.OperationRecord())
	crmCustomerGroupRouterWithoutRecord := Router.Group("crmCustomerGroup")
	crmCustomerGroupRouterWithoutAuth := PublicRouter.Group("crmCustomerGroup")

	var crmCustomerGroupApi = v1.ApiGroupApp.CrmApiGroup.CrmCustomerGroupApi
	{
		crmCustomerGroupRouter.POST("createCrmCustomerGroup", crmCustomerGroupApi.CreateCrmCustomerGroup)             // 新建crmCustomerGroup表
		crmCustomerGroupRouter.DELETE("deleteCrmCustomerGroup", crmCustomerGroupApi.DeleteCrmCustomerGroup)           // 删除crmCustomerGroup表
		crmCustomerGroupRouter.DELETE("deleteCrmCustomerGroupByIds", crmCustomerGroupApi.DeleteCrmCustomerGroupByIds) // 批量删除crmCustomerGroup表
		crmCustomerGroupRouter.PUT("updateCrmCustomerGroup", crmCustomerGroupApi.UpdateCrmCustomerGroup)              // 更新crmCustomerGroup表
	}
	{
		crmCustomerGroupRouterWithoutRecord.GET("findCrmCustomerGroup", crmCustomerGroupApi.FindCrmCustomerGroup) // 根据ID获取crmCustomerGroup表
		//crmCustomerGroupRouterWithoutRecord.GET("getCrmCustomerGroupList", crmCustomerGroupApi.GetCrmCustomerGroupList)  // 获取crmCustomerGroup表列表
	}
	{
		crmCustomerGroupRouterWithoutAuth.GET("getCrmCustomerGroupPublic", crmCustomerGroupApi.GetCrmCustomerGroupPublic) // 获取crmCustomerGroup表列表
		crmCustomerGroupRouterWithoutAuth.GET("getCrmCustomerGroupList", crmCustomerGroupApi.GetCrmCustomerGroupList)     // 获取crmCustomerGroup表列表
	}
}
