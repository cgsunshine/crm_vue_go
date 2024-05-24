package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPaymentRouter struct {
}

// InitCrmPaymentRouter 初始化 crmPayment表 路由信息
func (s *CrmPaymentRouter) InitCrmPaymentRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmPaymentRouter := Router.Group("crmPayment").Use(middleware.OperationRecord())
	crmPaymentRouterWithoutRecord := Router.Group("crmPayment")
	crmPaymentRouterWithoutAuth := PublicRouter.Group("crmPayment")

	var crmPaymentApi = v1.ApiGroupApp.CrmApiGroup.CrmPaymentApi
	{
		crmPaymentRouter.POST("createCrmPayment", crmPaymentApi.CreateCrmPayment)   // 新建crmPayment表
		crmPaymentRouter.DELETE("deleteCrmPayment", crmPaymentApi.DeleteCrmPayment) // 删除crmPayment表
		crmPaymentRouter.DELETE("deleteCrmPaymentByIds", crmPaymentApi.DeleteCrmPaymentByIds) // 批量删除crmPayment表
		crmPaymentRouter.PUT("updateCrmPayment", crmPaymentApi.UpdateCrmPayment)    // 更新crmPayment表
	}
	{
		crmPaymentRouterWithoutRecord.GET("findCrmPayment", crmPaymentApi.FindCrmPayment)        // 根据ID获取crmPayment表
		crmPaymentRouterWithoutRecord.GET("getCrmPaymentList", crmPaymentApi.GetCrmPaymentList)  // 获取crmPayment表列表
	}
	{
	    crmPaymentRouterWithoutAuth.GET("getCrmPaymentPublic", crmPaymentApi.GetCrmPaymentPublic)  // 获取crmPayment表列表
	}
}
