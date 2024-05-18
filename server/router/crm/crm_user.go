package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmUserRouter struct {
}

// InitCrmUserRouter 初始化 crmUser表 路由信息
func (s *CrmUserRouter) InitCrmUserRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmUserRouter := Router.Group("crmUser").Use(middleware.OperationRecord())
	crmUserRouterWithoutRecord := Router.Group("crmUser")
	crmUserRouterWithoutAuth := PublicRouter.Group("crmUser")

	var crmUserApi = v1.ApiGroupApp.CrmApiGroup.CrmUserApi
	{
		crmUserRouter.POST("createCrmUser", crmUserApi.CreateCrmUser)   // 新建crmUser表
		crmUserRouter.DELETE("deleteCrmUser", crmUserApi.DeleteCrmUser) // 删除crmUser表
		crmUserRouter.DELETE("deleteCrmUserByIds", crmUserApi.DeleteCrmUserByIds) // 批量删除crmUser表
		crmUserRouter.PUT("updateCrmUser", crmUserApi.UpdateCrmUser)    // 更新crmUser表
	}
	{
		crmUserRouterWithoutRecord.GET("findCrmUser", crmUserApi.FindCrmUser)        // 根据ID获取crmUser表
		crmUserRouterWithoutRecord.GET("getCrmUserList", crmUserApi.GetCrmUserList)  // 获取crmUser表列表
	}
	{
	    crmUserRouterWithoutAuth.GET("getCrmUserPublic", crmUserApi.GetCrmUserPublic)  // 获取crmUser表列表
	}
}
