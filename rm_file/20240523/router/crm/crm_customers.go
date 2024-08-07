package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmCustomersRouter struct {
}

// InitCrmCustomersRouter 初始化 客户管理 路由信息
func (s *CrmCustomersRouter) InitCrmCustomersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmCustomersRouter := Router.Group("crmCustomers").Use(middleware.OperationRecord())
	crmCustomersRouterWithoutRecord := Router.Group("crmCustomers")
	crmCustomersRouterWithoutAuth := PublicRouter.Group("crmCustomers")

	var crmCustomersApi = v1.ApiGroupApp.CrmApiGroup.CrmCustomersApi
	{
		crmCustomersRouter.POST("createCrmCustomers", crmCustomersApi.CreateCrmCustomers)             // 新建客户管理
		crmCustomersRouter.DELETE("deleteCrmCustomers", crmCustomersApi.DeleteCrmCustomers)           // 删除客户管理
		crmCustomersRouter.DELETE("deleteCrmCustomersByIds", crmCustomersApi.DeleteCrmCustomersByIds) // 批量删除客户管理
		crmCustomersRouter.PUT("updateCrmCustomers", crmCustomersApi.UpdateCrmCustomers)              // 更新客户管理
	}
	{
		crmCustomersRouterWithoutRecord.GET("findCrmCustomers", crmCustomersApi.FindCrmCustomers)       // 根据ID获取客户管理
		crmCustomersRouterWithoutRecord.GET("getCrmCustomersList", crmCustomersApi.GetCrmCustomersList) // 获取客户管理列表

		crmCustomersRouterWithoutRecord.GET("getPageCrmCustomersList", crmCustomersApi.GetPageCrmCustomersList) // 获取客户管理列表 完整
	}
	{
		crmCustomersRouterWithoutAuth.GET("getCrmCustomersPublic", crmCustomersApi.GetCrmCustomersPublic) // 获取客户管理列表
	}
}
