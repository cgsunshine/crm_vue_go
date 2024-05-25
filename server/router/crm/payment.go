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

	var crmPaymentApi = v1.ApiGroupApp.CrmApiGroup.CrmPaymentApi

	{
		crmPaymentRouterWithoutRecord.GET("getCrmPagePaymentList", crmPaymentApi.GetCrmPagePaymentList) // 获取crmPayment表列表
	}
}
