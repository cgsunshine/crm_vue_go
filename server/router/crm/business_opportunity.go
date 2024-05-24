package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPageBusinessOpportunityRouter struct {
}

// InitCrmBusinessOpportunityRouter 初始化 crmBusinessOpportunity表 路由信息
func (s *CrmPageBusinessOpportunityRouter) InitCrmPageBusinessOpportunityRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmBusinessOpportunityRouterWithoutRecord := Router.Group("crmBusinessOpportunity")

	var crmBusinessOpportunityApi = v1.ApiGroupApp.CrmApiGroup.CrmBusinessOpportunityApi

	{
		crmBusinessOpportunityRouterWithoutRecord.GET("getCrmPageBusinessOpportunityList", crmBusinessOpportunityApi.GetCrmPageBusinessOpportunityList) // 获取crmBusinessOpportunity表列表
	}
}
