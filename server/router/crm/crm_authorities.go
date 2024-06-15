package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmAuthoritiesRouter struct {
}

// InitCrmAuthoritiesRouter 初始化 数据权限 路由信息
func (s *CrmAuthoritiesRouter) InitCrmAuthoritiesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmAuthoritiesRouter := Router.Group("crmAuthorities").Use(middleware.OperationRecord())
	crmAuthoritiesRouterWithoutRecord := Router.Group("crmAuthorities")
	crmAuthoritiesRouterWithoutAuth := PublicRouter.Group("crmAuthorities")

	var crmAuthoritiesApi = v1.ApiGroupApp.CrmApiGroup.CrmAuthoritiesApi
	{
		crmAuthoritiesRouter.POST("createCrmAuthorities", crmAuthoritiesApi.CreateCrmAuthorities)   // 新建数据权限
		crmAuthoritiesRouter.DELETE("deleteCrmAuthorities", crmAuthoritiesApi.DeleteCrmAuthorities) // 删除数据权限
		crmAuthoritiesRouter.DELETE("deleteCrmAuthoritiesByIds", crmAuthoritiesApi.DeleteCrmAuthoritiesByIds) // 批量删除数据权限
		crmAuthoritiesRouter.PUT("updateCrmAuthorities", crmAuthoritiesApi.UpdateCrmAuthorities)    // 更新数据权限
	}
	{
		crmAuthoritiesRouterWithoutRecord.GET("findCrmAuthorities", crmAuthoritiesApi.FindCrmAuthorities)        // 根据ID获取数据权限
		crmAuthoritiesRouterWithoutRecord.GET("getCrmAuthoritiesList", crmAuthoritiesApi.GetCrmAuthoritiesList)  // 获取数据权限列表
	}
	{
	    crmAuthoritiesRouterWithoutAuth.GET("getCrmAuthoritiesPublic", crmAuthoritiesApi.GetCrmAuthoritiesPublic)  // 获取数据权限列表
	}
}
