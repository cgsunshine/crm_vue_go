package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmStatementAccountRouter struct {
}

// InitCrmStatementAccountRouter 初始化 crmStatementAccount表 路由信息
func (s *CrmStatementAccountRouter) InitCrmStatementAccountRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmStatementAccountRouter := Router.Group("crmStatementAccount").Use(middleware.OperationRecord())
	crmStatementAccountRouterWithoutRecord := Router.Group("crmStatementAccount")
	crmStatementAccountRouterWithoutAuth := PublicRouter.Group("crmStatementAccount")

	var crmStatementAccountApi = v1.ApiGroupApp.CrmApiGroup.CrmStatementAccountApi
	{
		crmStatementAccountRouter.POST("createCrmStatementAccount", crmStatementAccountApi.CreateCrmStatementAccount)             // 新建crmStatementAccount表
		crmStatementAccountRouter.DELETE("deleteCrmStatementAccount", crmStatementAccountApi.DeleteCrmStatementAccount)           // 删除crmStatementAccount表
		crmStatementAccountRouter.DELETE("deleteCrmStatementAccountByIds", crmStatementAccountApi.DeleteCrmStatementAccountByIds) // 批量删除crmStatementAccount表
		crmStatementAccountRouter.PUT("updateCrmStatementAccount", crmStatementAccountApi.UpdateCrmStatementAccount)              // 更新crmStatementAccount表
	}
	{
		crmStatementAccountRouterWithoutRecord.GET("findCrmStatementAccount", crmStatementAccountApi.FindCrmStatementAccount) // 根据ID获取crmStatementAccount表
		//crmStatementAccountRouterWithoutRecord.GET("getCrmStatementAccountList", crmStatementAccountApi.GetCrmStatementAccountList)  // 获取crmStatementAccount表列表
	}
	{
		crmStatementAccountRouterWithoutAuth.GET("getCrmStatementAccountPublic", crmStatementAccountApi.GetCrmStatementAccountPublic) // 获取crmStatementAccount表列表
		crmStatementAccountRouterWithoutAuth.GET("getCrmStatementAccountList", crmStatementAccountApi.GetCrmStatementAccountList)     // 获取crmStatementAccount表列表
	}
}
