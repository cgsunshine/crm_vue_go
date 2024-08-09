package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

// InitCrmAuthoritiesRouter 初始化 数据权限 路由信息
func (s *CrmAuthoritiesRouter) InitCustomerCrmAuthoritiesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmAuthoritiesRouter := Router.Group("crmAuthorities").Use(middleware.OperationRecord())

	var crmAuthoritiesApi = v1.ApiGroupApp.CrmApiGroup.CrmAuthoritiesApi
	{
		crmAuthoritiesRouter.PUT("updateCrmStatusAuthorities", crmAuthoritiesApi.UpdateCrmStatusAuthorities) // 更新数据权限
	}
}
