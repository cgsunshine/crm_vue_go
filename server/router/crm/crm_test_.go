package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmTestRouter struct {
}

// InitCrmTestRouter 初始化 crmTest表 路由信息
func (s *CrmTestRouter) InitCrmTestRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmTestRouter := Router.Group("crmTest").Use(middleware.OperationRecord())
	crmTestRouterWithoutRecord := Router.Group("crmTest")
	crmTestRouterWithoutAuth := PublicRouter.Group("crmTest")

	var crmTestApi = v1.ApiGroupApp.CrmApiGroup.CrmTestApi
	{
		crmTestRouter.POST("createCrmTest", crmTestApi.CreateCrmTest)   // 新建crmTest表
		crmTestRouter.DELETE("deleteCrmTest", crmTestApi.DeleteCrmTest) // 删除crmTest表
		crmTestRouter.DELETE("deleteCrmTestByIds", crmTestApi.DeleteCrmTestByIds) // 批量删除crmTest表
		crmTestRouter.PUT("updateCrmTest", crmTestApi.UpdateCrmTest)    // 更新crmTest表
	}
	{
		crmTestRouterWithoutRecord.GET("findCrmTest", crmTestApi.FindCrmTest)        // 根据ID获取crmTest表
		crmTestRouterWithoutRecord.GET("getCrmTestList", crmTestApi.GetCrmTestList)  // 获取crmTest表列表
	}
	{
	    crmTestRouterWithoutAuth.GET("getCrmTestPublic", crmTestApi.GetCrmTestPublic)  // 获取crmTest表列表
	}
}
