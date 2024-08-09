package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPageCustomersRouter struct {
}

// InitCrmCustomersRouter 初始化 crmCustomers表 路由信息
func (s *CrmPageCustomersRouter) InitCrmPageCustomersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmCustomersRouter := Router.Group("crmCustomers").Use(middleware.OperationRecord())
	crmCustomersRouterWithoutRecord := Router.Group("crmCustomers")
	crmCustomersRouterWithoutAuth := PublicRouter.Group("crmCustomers")

	var crmCustomersApi = v1.ApiGroupApp.CrmApiGroup.CrmCustomersApi

	{
		crmCustomersRouter.POST("createCrmPageCustomers", crmCustomersApi.CreateCrmPageCustomers) // 新建crmCustomers表
	}

	{

		crmCustomersRouterWithoutRecord.GET("findCrmPageCustomers", crmCustomersApi.FindCrmPageCustomers) // 根据ID获取crmCustomers表
	}

	{
		crmCustomersRouterWithoutAuth.GET("getCrmPageCustomersList", crmCustomersApi.GetCrmPageCustomersList)       // 获取crmCustomers表列表
		crmCustomersRouterWithoutAuth.GET("getCrmPageCustomersAllList", crmCustomersApi.GetCrmPageCustomersAllList) // 获取crmCustomers表列表
	}
}
