package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmLoginLogRouter struct {
}

// InitCrmLoginLogRouter 初始化 登录日志 路由信息
func (s *CrmLoginLogRouter) InitCrmLoginLogRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmLoginLogRouter := Router.Group("crmLoginLog").Use(middleware.OperationRecord())
	crmLoginLogRouterWithoutRecord := Router.Group("crmLoginLog")
	crmLoginLogRouterWithoutAuth := PublicRouter.Group("crmLoginLog")

	var crmLoginLogApi = v1.ApiGroupApp.CrmApiGroup.CrmLoginLogApi
	{
		crmLoginLogRouter.POST("createCrmLoginLog", crmLoginLogApi.CreateCrmLoginLog)   // 新建登录日志
		crmLoginLogRouter.DELETE("deleteCrmLoginLog", crmLoginLogApi.DeleteCrmLoginLog) // 删除登录日志
		crmLoginLogRouter.DELETE("deleteCrmLoginLogByIds", crmLoginLogApi.DeleteCrmLoginLogByIds) // 批量删除登录日志
		crmLoginLogRouter.PUT("updateCrmLoginLog", crmLoginLogApi.UpdateCrmLoginLog)    // 更新登录日志
	}
	{
		crmLoginLogRouterWithoutRecord.GET("findCrmLoginLog", crmLoginLogApi.FindCrmLoginLog)        // 根据ID获取登录日志
		crmLoginLogRouterWithoutRecord.GET("getCrmLoginLogList", crmLoginLogApi.GetCrmLoginLogList)  // 获取登录日志列表
	}
	{
	    crmLoginLogRouterWithoutAuth.GET("getCrmLoginLogPublic", crmLoginLogApi.GetCrmLoginLogPublic)  // 获取登录日志列表
	}
}
