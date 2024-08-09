package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPageProductRouter struct {
}

// InitCrmProductRouter 初始化 crmProduct表 路由信息
func (s *CrmPageProductRouter) InitCrmPageProductRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmProductRouterWithoutRecord := Router.Group("crmProduct")
	crmProductRouterWithoutAuth := PublicRouter.Group("crmProduct")

	var crmProductApi = v1.ApiGroupApp.CrmApiGroup.CrmProductApi

	{
		crmProductRouterWithoutRecord.GET("findCrmPageProduct", crmProductApi.FindCrmPageProduct) // 根据ID获取crmProduct表
	}

	{
		crmProductRouterWithoutAuth.GET("getCrmPageProductList", crmProductApi.GetCrmPageProductList) // 获取crmProduct表列表 完整
	}

}
