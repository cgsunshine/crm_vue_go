package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmDepositsRouter struct {
}

// InitCrmDepositsRouter 初始化 押金管理 路由信息
func (s *CrmDepositsRouter) InitCrmDepositsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmDepositsRouter := Router.Group("crmDeposits").Use(middleware.OperationRecord())
	crmDepositsRouterWithoutRecord := Router.Group("crmDeposits")
	crmDepositsRouterWithoutAuth := PublicRouter.Group("crmDeposits")

	var crmDepositsApi = v1.ApiGroupApp.CrmApiGroup.CrmDepositsApi
	{
		crmDepositsRouter.POST("createCrmDeposits", crmDepositsApi.CreateCrmDeposits)             // 新建押金管理
		crmDepositsRouter.DELETE("deleteCrmDeposits", crmDepositsApi.DeleteCrmDeposits)           // 删除押金管理
		crmDepositsRouter.DELETE("deleteCrmDepositsByIds", crmDepositsApi.DeleteCrmDepositsByIds) // 批量删除押金管理
		crmDepositsRouter.PUT("updateCrmDeposits", crmDepositsApi.UpdateCrmDeposits)              // 更新押金管理
	}
	{
		//crmDepositsRouterWithoutRecord.GET("findCrmDeposits", crmDepositsApi.FindCrmDeposits)       // 根据ID获取押金管理
		crmDepositsRouterWithoutRecord.GET("getCrmDepositsList", crmDepositsApi.GetCrmDepositsList) // 获取押金管理列表
	}
	{
		crmDepositsRouterWithoutAuth.GET("getCrmDepositsPublic", crmDepositsApi.GetCrmDepositsPublic)        // 获取押金管理列表
		crmDepositsRouterWithoutAuth.POST("createCrmDepositsRefund", crmDepositsApi.CreateCrmDepositsRefund) // 创建退款押金
		crmDepositsRouterWithoutAuth.GET("findCrmDeposits", crmDepositsApi.FindCrmDeposits)                  // 根据ID获取押金管理
	}
}
