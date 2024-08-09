package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmProductTypeRouter struct {
}

// InitCrmProductTypeRouter 初始化 产品类型 路由信息
func (s *CrmProductTypeRouter) InitCrmProductTypeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmProductTypeRouter := Router.Group("crmProductType").Use(middleware.OperationRecord())
	crmProductTypeRouterWithoutRecord := Router.Group("crmProductType")
	crmProductTypeRouterWithoutAuth := PublicRouter.Group("crmProductType")

	var crmProductTypeApi = v1.ApiGroupApp.CrmApiGroup.CrmProductTypeApi
	{
		crmProductTypeRouter.POST("createCrmProductType", crmProductTypeApi.CreateCrmProductType)             // 新建产品类型
		crmProductTypeRouter.DELETE("deleteCrmProductType", crmProductTypeApi.DeleteCrmProductType)           // 删除产品类型
		crmProductTypeRouter.DELETE("deleteCrmProductTypeByIds", crmProductTypeApi.DeleteCrmProductTypeByIds) // 批量删除产品类型
		crmProductTypeRouter.PUT("updateCrmProductType", crmProductTypeApi.UpdateCrmProductType)              // 更新产品类型
	}
	{
		crmProductTypeRouterWithoutRecord.GET("findCrmProductType", crmProductTypeApi.FindCrmProductType) // 根据ID获取产品类型
		//crmProductTypeRouterWithoutRecord.GET("getCrmProductTypeList", crmProductTypeApi.GetCrmProductTypeList)  // 获取产品类型列表
	}
	{
		crmProductTypeRouterWithoutAuth.GET("getCrmProductTypePublic", crmProductTypeApi.GetCrmProductTypePublic) // 获取产品类型列表
		crmProductTypeRouterWithoutAuth.GET("getCrmProductTypeList", crmProductTypeApi.GetCrmProductTypeList)     // 获取产品类型列表
	}
}
