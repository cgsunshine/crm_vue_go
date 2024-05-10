package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmSupplierRouter struct {
}

// InitCrmSupplierRouter 初始化 供应商管理 路由信息
func (s *CrmSupplierRouter) InitCrmSupplierRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmSupplierRouter := Router.Group("crmSupplier").Use(middleware.OperationRecord())
	crmSupplierRouterWithoutRecord := Router.Group("crmSupplier")
	crmSupplierRouterWithoutAuth := PublicRouter.Group("crmSupplier")

	var crmSupplierApi = v1.ApiGroupApp.CrmApiGroup.CrmSupplierApi
	{
		crmSupplierRouter.POST("createCrmSupplier", crmSupplierApi.CreateCrmSupplier)   // 新建供应商管理
		crmSupplierRouter.DELETE("deleteCrmSupplier", crmSupplierApi.DeleteCrmSupplier) // 删除供应商管理
		crmSupplierRouter.DELETE("deleteCrmSupplierByIds", crmSupplierApi.DeleteCrmSupplierByIds) // 批量删除供应商管理
		crmSupplierRouter.PUT("updateCrmSupplier", crmSupplierApi.UpdateCrmSupplier)    // 更新供应商管理
	}
	{
		crmSupplierRouterWithoutRecord.GET("findCrmSupplier", crmSupplierApi.FindCrmSupplier)        // 根据ID获取供应商管理
		crmSupplierRouterWithoutRecord.GET("getCrmSupplierList", crmSupplierApi.GetCrmSupplierList)  // 获取供应商管理列表
	}
	{
	    crmSupplierRouterWithoutAuth.GET("getCrmSupplierPublic", crmSupplierApi.GetCrmSupplierPublic)  // 获取供应商管理列表
	}
}
