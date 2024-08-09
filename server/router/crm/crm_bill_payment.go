package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmBillPaymentRouter struct {
}

// InitCrmBillPaymentRouter 初始化 应付账单 路由信息
func (s *CrmBillPaymentRouter) InitCrmBillPaymentRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmBillPaymentRouter := Router.Group("crmBillPayment").Use(middleware.OperationRecord())
	crmBillPaymentRouterWithoutRecord := Router.Group("crmBillPayment")
	crmBillPaymentRouterWithoutAuth := PublicRouter.Group("crmBillPayment")

	var crmBillPaymentApi = v1.ApiGroupApp.CrmApiGroup.CrmBillPaymentApi
	{
		crmBillPaymentRouter.POST("createCrmBillPayment", crmBillPaymentApi.CreateCrmBillPayment)   // 新建应付账单
		crmBillPaymentRouter.DELETE("deleteCrmBillPayment", crmBillPaymentApi.DeleteCrmBillPayment) // 删除应付账单
		crmBillPaymentRouter.DELETE("deleteCrmBillPaymentByIds", crmBillPaymentApi.DeleteCrmBillPaymentByIds) // 批量删除应付账单
		crmBillPaymentRouter.PUT("updateCrmBillPayment", crmBillPaymentApi.UpdateCrmBillPayment)    // 更新应付账单
	}
	{
		crmBillPaymentRouterWithoutRecord.GET("findCrmBillPayment", crmBillPaymentApi.FindCrmBillPayment)        // 根据ID获取应付账单
		crmBillPaymentRouterWithoutRecord.GET("getCrmBillPaymentList", crmBillPaymentApi.GetCrmBillPaymentList)  // 获取应付账单列表
	}
	{
	    crmBillPaymentRouterWithoutAuth.GET("getCrmBillPaymentPublic", crmBillPaymentApi.GetCrmBillPaymentPublic)  // 获取应付账单列表
	}
}
