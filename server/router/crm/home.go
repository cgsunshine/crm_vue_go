package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmAdminHomeRouter struct {
}

// CrmAdminHome 初始化 首页 路由信息
func (s *CrmAdminHomeRouter) InitCrmAdminHomeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	//crmBillPaymentRouter := Router.Group("crmBillPayment").Use(middleware.OperationRecord())
	//crmBillPaymentRouterWithoutRecord := Router.Group("crmBillPayment")
	crmBillPaymentRouterWithoutAuth := PublicRouter.Group("crmAdminHome")

	var adminHome = v1.ApiGroupApp.CrmApiGroup.AdminHome

	{
		crmBillPaymentRouterWithoutAuth.GET("homeData", adminHome.HomeData) // 首页信息
	}
}
