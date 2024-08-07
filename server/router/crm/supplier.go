package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPageSupplierRouter struct {
}

// InitCrmSupplierRouter 初始化 crmSupplier表 路由信息
func (s *CrmPageSupplierRouter) InitCrmPageSupplierRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmSupplierRouterWithoutRecord := Router.Group("crmSupplier")
	crmSupplierRouterWithoutAuth := PublicRouter.Group("crmSupplier")

	var crmSupplierApi = v1.ApiGroupApp.CrmApiGroup.CrmSupplierApi

	{
		crmSupplierRouterWithoutRecord.GET("findCrmPageSupplier", crmSupplierApi.FindCrmPageSupplier) // 根据ID获取crmSupplier表
	}

	{
		crmSupplierRouterWithoutAuth.GET("getCrmPageSupplierList", crmSupplierApi.GetCrmPageSupplierList) // 获取crmSupplier表列表
	}
}
