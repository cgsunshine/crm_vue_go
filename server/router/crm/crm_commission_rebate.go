package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmCommissionRebateRouter struct {
}

// InitCrmCommissionRebateRouter 初始化 crmCommissionRebate表 路由信息
func (s *CrmCommissionRebateRouter) InitCrmCommissionRebateRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmCommissionRebateRouter := Router.Group("crmCommissionRebate").Use(middleware.OperationRecord())
	crmCommissionRebateRouterWithoutRecord := Router.Group("crmCommissionRebate")
	crmCommissionRebateRouterWithoutAuth := PublicRouter.Group("crmCommissionRebate")

	var crmCommissionRebateApi = v1.ApiGroupApp.CrmApiGroup.CrmCommissionRebateApi
	{
		crmCommissionRebateRouter.POST("createCrmCommissionRebate", crmCommissionRebateApi.CreateCrmCommissionRebate)   // 新建crmCommissionRebate表
		crmCommissionRebateRouter.DELETE("deleteCrmCommissionRebate", crmCommissionRebateApi.DeleteCrmCommissionRebate) // 删除crmCommissionRebate表
		crmCommissionRebateRouter.DELETE("deleteCrmCommissionRebateByIds", crmCommissionRebateApi.DeleteCrmCommissionRebateByIds) // 批量删除crmCommissionRebate表
		crmCommissionRebateRouter.PUT("updateCrmCommissionRebate", crmCommissionRebateApi.UpdateCrmCommissionRebate)    // 更新crmCommissionRebate表
	}
	{
		crmCommissionRebateRouterWithoutRecord.GET("findCrmCommissionRebate", crmCommissionRebateApi.FindCrmCommissionRebate)        // 根据ID获取crmCommissionRebate表
		crmCommissionRebateRouterWithoutRecord.GET("getCrmCommissionRebateList", crmCommissionRebateApi.GetCrmCommissionRebateList)  // 获取crmCommissionRebate表列表
	}
	{
	    crmCommissionRebateRouterWithoutAuth.GET("getCrmCommissionRebatePublic", crmCommissionRebateApi.GetCrmCommissionRebatePublic)  // 获取crmCommissionRebate表列表
	}
}
