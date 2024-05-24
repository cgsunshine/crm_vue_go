package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPagePaymentCollentionRouter struct {
}

// InitCrmPaymentCollentionRouter 初始化 crmPaymentCollention表 路由信息
func (s *CrmPagePaymentCollentionRouter) InitCrmPagePaymentCollentionRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmPaymentCollentionRouterWithoutRecord := Router.Group("crmPaymentCollention")

	var crmPaymentCollentionApi = v1.ApiGroupApp.CrmApiGroup.CrmPaymentCollentionApi

	{
		crmPaymentCollentionRouterWithoutRecord.GET("getCrmPagePaymentCollentionInfoList", crmPaymentCollentionApi.GetCrmPagePaymentCollentionInfoList) // 获取crmPaymentCollention表列表
	}

}
