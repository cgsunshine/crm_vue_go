package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmCurrencyRouter struct {
}

// InitCrmCurrencyRouter 初始化 币种管理 路由信息
func (s *CrmCurrencyRouter) InitCrmCurrencyRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmCurrencyRouter := Router.Group("crmCurrency").Use(middleware.OperationRecord())
	crmCurrencyRouterWithoutRecord := Router.Group("crmCurrency")
	crmCurrencyRouterWithoutAuth := PublicRouter.Group("crmCurrency")

	var crmCurrencyApi = v1.ApiGroupApp.CrmApiGroup.CrmCurrencyApi
	{
		crmCurrencyRouter.POST("createCrmCurrency", crmCurrencyApi.CreateCrmCurrency)   // 新建币种管理
		crmCurrencyRouter.DELETE("deleteCrmCurrency", crmCurrencyApi.DeleteCrmCurrency) // 删除币种管理
		crmCurrencyRouter.DELETE("deleteCrmCurrencyByIds", crmCurrencyApi.DeleteCrmCurrencyByIds) // 批量删除币种管理
		crmCurrencyRouter.PUT("updateCrmCurrency", crmCurrencyApi.UpdateCrmCurrency)    // 更新币种管理
	}
	{
		crmCurrencyRouterWithoutRecord.GET("findCrmCurrency", crmCurrencyApi.FindCrmCurrency)        // 根据ID获取币种管理
		crmCurrencyRouterWithoutRecord.GET("getCrmCurrencyList", crmCurrencyApi.GetCrmCurrencyList)  // 获取币种管理列表
	}
	{
	    crmCurrencyRouterWithoutAuth.GET("getCrmCurrencyPublic", crmCurrencyApi.GetCrmCurrencyPublic)  // 获取币种管理列表
	}
}
