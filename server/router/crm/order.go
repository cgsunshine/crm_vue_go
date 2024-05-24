package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPageOrderRouter struct {
}

// InitCrmOrderRouter 初始化 crmOrder表 路由信息
func (s *CrmPageOrderRouter) InitCrmPagOrderRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmOrderRouterWithoutRecord := Router.Group("crmOrder")

	var crmOrderApi = v1.ApiGroupApp.CrmApiGroup.CrmOrderApi

	{
		crmOrderRouterWithoutRecord.GET("getCrmPageOrderList", crmOrderApi.GetCrmPageOrderList) // 获取crmOrder表列表
	}
}
