package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmConfigRouter struct {
}

// InitCrmConfigRouter 初始化 系统配置 路由信息
func (s *CrmConfigRouter) InitCrmConfigRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmConfigRouter := Router.Group("crmConfig").Use(middleware.OperationRecord())
	crmConfigRouterWithoutRecord := Router.Group("crmConfig")
	crmConfigRouterWithoutAuth := PublicRouter.Group("crmConfig")

	var crmConfigApi = v1.ApiGroupApp.CrmApiGroup.CrmConfigApi
	{
		crmConfigRouter.POST("createCrmConfig", crmConfigApi.CreateCrmConfig)   // 新建系统配置
		crmConfigRouter.DELETE("deleteCrmConfig", crmConfigApi.DeleteCrmConfig) // 删除系统配置
		crmConfigRouter.DELETE("deleteCrmConfigByIds", crmConfigApi.DeleteCrmConfigByIds) // 批量删除系统配置
		crmConfigRouter.PUT("updateCrmConfig", crmConfigApi.UpdateCrmConfig)    // 更新系统配置
	}
	{
		crmConfigRouterWithoutRecord.GET("findCrmConfig", crmConfigApi.FindCrmConfig)        // 根据ID获取系统配置
		crmConfigRouterWithoutRecord.GET("getCrmConfigList", crmConfigApi.GetCrmConfigList)  // 获取系统配置列表
	}
	{
	    crmConfigRouterWithoutAuth.GET("getCrmConfigPublic", crmConfigApi.GetCrmConfigPublic)  // 获取系统配置列表
	}
}
