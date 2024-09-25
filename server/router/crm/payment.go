package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPagePaymentRouter struct {
}

// InitCrmPaymentRouter 初始化 crmPayment表 路由信息
func (s *CrmPagePaymentRouter) InitCrmPagePaymentRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmPaymentRouterWithoutRecord := Router.Group("crmPayment")
	crmPaymentRouterWithoutAuth := PublicRouter.Group("crmPayment")

	var crmPaymentApi = v1.ApiGroupApp.CrmApiGroup.CrmPaymentApi

	{
		crmPaymentRouterWithoutRecord.GET("findCrmPagePayment", crmPaymentApi.FindCrmPagePayment) // 根据ID获取crmPayment表
	}

	{
		crmPaymentRouterWithoutAuth.GET("getCrmPagePaymentList", crmPaymentApi.GetCrmPagePaymentList) // 获取crmPayment表列表
		crmPaymentRouterWithoutAuth.POST("paymentCompleted", crmPaymentApi.PaymentCompleted)          // 获取crmPayment表列表
	}
}
