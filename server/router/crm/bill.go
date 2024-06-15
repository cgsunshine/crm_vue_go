package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPageBillRouter struct {
}

// InitCrmBillRouter 初始化 crmBill表 路由信息
func (s *CrmBillRouter) InitCrmPageBillRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmBillRouterWithoutRecord := Router.Group("crmBill")

	crmBillRouterWithoutAuth := PublicRouter.Group("crmBill")

	var crmBillApi = v1.ApiGroupApp.CrmApiGroup.CrmBillApi
	{
		crmBillRouterWithoutRecord.GET("findPageCrmBill", crmBillApi.FindPageCrmBill) // 根据ID获取crmBill表
	}
	{
		crmBillRouterWithoutAuth.GET("getCrmPageBillList", crmBillApi.GetCrmPageBillList) // 获取crmBill表列表
	}
}
