package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmProductGroupRouter struct {
}

// InitCrmProductGroupRouter 初始化 产品分组 路由信息
func (s *CrmProductGroupRouter) InitCrmProductGroupRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmProductGroupRouter := Router.Group("crmProductGroup").Use(middleware.OperationRecord())
	crmProductGroupRouterWithoutRecord := Router.Group("crmProductGroup")
	crmProductGroupRouterWithoutAuth := PublicRouter.Group("crmProductGroup")

	var crmProductGroupApi = v1.ApiGroupApp.CrmApiGroup.CrmProductGroupApi
	{
		crmProductGroupRouter.POST("createCrmProductGroup", crmProductGroupApi.CreateCrmProductGroup)   // 新建产品分组
		crmProductGroupRouter.DELETE("deleteCrmProductGroup", crmProductGroupApi.DeleteCrmProductGroup) // 删除产品分组
		crmProductGroupRouter.DELETE("deleteCrmProductGroupByIds", crmProductGroupApi.DeleteCrmProductGroupByIds) // 批量删除产品分组
		crmProductGroupRouter.PUT("updateCrmProductGroup", crmProductGroupApi.UpdateCrmProductGroup)    // 更新产品分组
	}
	{
		crmProductGroupRouterWithoutRecord.GET("findCrmProductGroup", crmProductGroupApi.FindCrmProductGroup)        // 根据ID获取产品分组
		crmProductGroupRouterWithoutRecord.GET("getCrmProductGroupList", crmProductGroupApi.GetCrmProductGroupList)  // 获取产品分组列表
	}
	{
	    crmProductGroupRouterWithoutAuth.GET("getCrmProductGroupPublic", crmProductGroupApi.GetCrmProductGroupPublic)  // 获取产品分组列表
	}
}
