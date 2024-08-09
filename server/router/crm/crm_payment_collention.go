package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPaymentCollentionRouter struct {
}

// InitCrmPaymentCollentionRouter 初始化 crmPaymentCollention表 路由信息
func (s *CrmPaymentCollentionRouter) InitCrmPaymentCollentionRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmPaymentCollentionRouter := Router.Group("crmPaymentCollention").Use(middleware.OperationRecord())
	crmPaymentCollentionRouterWithoutRecord := Router.Group("crmPaymentCollention")
	crmPaymentCollentionRouterWithoutAuth := PublicRouter.Group("crmPaymentCollention")

	var crmPaymentCollentionApi = v1.ApiGroupApp.CrmApiGroup.CrmPaymentCollentionApi
	{
		crmPaymentCollentionRouter.POST("createCrmPaymentCollention", crmPaymentCollentionApi.CreateCrmPaymentCollention)             // 新建crmPaymentCollention表
		crmPaymentCollentionRouter.DELETE("deleteCrmPaymentCollention", crmPaymentCollentionApi.DeleteCrmPaymentCollention)           // 删除crmPaymentCollention表
		crmPaymentCollentionRouter.DELETE("deleteCrmPaymentCollentionByIds", crmPaymentCollentionApi.DeleteCrmPaymentCollentionByIds) // 批量删除crmPaymentCollention表
		crmPaymentCollentionRouter.PUT("updateCrmPaymentCollention", crmPaymentCollentionApi.UpdateCrmPaymentCollention)              // 更新crmPaymentCollention表
	}
	{
		crmPaymentCollentionRouterWithoutRecord.GET("findCrmPaymentCollention", crmPaymentCollentionApi.FindCrmPaymentCollention) // 根据ID获取crmPaymentCollention表
		//crmPaymentCollentionRouterWithoutRecord.GET("getCrmPaymentCollentionList", crmPaymentCollentionApi.GetCrmPaymentCollentionList)  // 获取crmPaymentCollention表列表
	}
	{
		crmPaymentCollentionRouterWithoutAuth.GET("getCrmPaymentCollentionPublic", crmPaymentCollentionApi.GetCrmPaymentCollentionPublic) // 获取crmPaymentCollention表列表
		crmPaymentCollentionRouterWithoutAuth.GET("getCrmPaymentCollentionList", crmPaymentCollentionApi.GetCrmPaymentCollentionList)     // 获取crmPaymentCollention表列表
	}
}
