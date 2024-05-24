package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPageCustomersRouter struct {
}

// InitCrmCustomersRouter 初始化 crmCustomers表 路由信息
func (s *CrmPageCustomersRouter) InitCrmPageCustomersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmCustomersRouterWithoutRecord := Router.Group("crmCustomers")

	var crmCustomersApi = v1.ApiGroupApp.CrmApiGroup.CrmCustomersApi

	{
		crmCustomersRouterWithoutRecord.GET("getPageCrmCustomersList", crmCustomersApi.GetPageCrmCustomersList) // 获取crmCustomers表列表
	}
}
