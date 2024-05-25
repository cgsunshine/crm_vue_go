package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPageCommissionRebateRouter struct {
}

// InitCrmCommissionRebateRouter 初始化 crmCommissionRebate表 路由信息
func (s *CrmPageCommissionRebateRouter) InitCrmPageCommissionRebateRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmCommissionRebateRouterWithoutRecord := Router.Group("crmCommissionRebate")

	var crmCommissionRebateApi = v1.ApiGroupApp.CrmApiGroup.CrmCommissionRebateApi

	{
		crmCommissionRebateRouterWithoutRecord.GET("getCrmPageCommissionRebateList", crmCommissionRebateApi.GetCrmPageCommissionRebateList) // 获取crmCommissionRebate表列表
	}

}
