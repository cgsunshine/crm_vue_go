package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPagePaymentCollentionRouter struct {
}

// InitCrmPaymentCollentionRouter 初始化 crmPaymentCollention表 路由信息
func (s *CrmPagePaymentCollentionRouter) InitCrmPagePaymentCollentionRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmPaymentCollentionRouter := Router.Group("crmPaymentCollention").Use(middleware.OperationRecord())
	crmPaymentCollentionRouterWithoutRecord := Router.Group("crmPaymentCollention")
	crmPaymentCollentionRouterWithoutAuth := PublicRouter.Group("crmPaymentCollention")

	var crmPaymentCollentionApi = v1.ApiGroupApp.CrmApiGroup.CrmPaymentCollentionApi

	{
		crmPaymentCollentionRouter.POST("createCrmPagePaymentCollention", crmPaymentCollentionApi.CreateCrmPagePaymentCollention) // 新建crmPaymentCollention表
	}
	{
		crmPaymentCollentionRouterWithoutRecord.GET("findCrmPagePaymentCollention", crmPaymentCollentionApi.FindCrmPagePaymentCollention) // 根据ID获取crmPaymentCollention表
	}

	{
		crmPaymentCollentionRouterWithoutAuth.GET("getCrmPagePaymentCollentionInfoList", crmPaymentCollentionApi.GetCrmPagePaymentCollentionInfoList) // 获取crmPaymentCollention表列表
	}

}
