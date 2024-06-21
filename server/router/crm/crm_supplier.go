package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmSupplierRouter struct {
}

// InitCrmSupplierRouter 初始化 crmSupplier表 路由信息
func (s *CrmSupplierRouter) InitCrmSupplierRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmSupplierRouter := Router.Group("crmSupplier").Use(middleware.OperationRecord())
	crmSupplierRouterWithoutRecord := Router.Group("crmSupplier")
	crmSupplierRouterWithoutAuth := PublicRouter.Group("crmSupplier")

	var crmSupplierApi = v1.ApiGroupApp.CrmApiGroup.CrmSupplierApi
	{
		crmSupplierRouter.POST("createCrmSupplier", crmSupplierApi.CreateCrmSupplier)             // 新建crmSupplier表
		crmSupplierRouter.DELETE("deleteCrmSupplier", crmSupplierApi.DeleteCrmSupplier)           // 删除crmSupplier表
		crmSupplierRouter.DELETE("deleteCrmSupplierByIds", crmSupplierApi.DeleteCrmSupplierByIds) // 批量删除crmSupplier表
		crmSupplierRouter.PUT("updateCrmSupplier", crmSupplierApi.UpdateCrmSupplier)              // 更新crmSupplier表
	}
	{
		crmSupplierRouterWithoutRecord.GET("findCrmSupplier", crmSupplierApi.FindCrmSupplier) // 根据ID获取crmSupplier表
		//crmSupplierRouterWithoutRecord.GET("getCrmSupplierList", crmSupplierApi.GetCrmSupplierList) // 获取crmSupplier表列表
	}
	{
		crmSupplierRouterWithoutAuth.GET("getCrmSupplierPublic", crmSupplierApi.GetCrmSupplierPublic) // 获取crmSupplier表列表
		crmSupplierRouterWithoutAuth.GET("getCrmSupplierList", crmSupplierApi.GetCrmSupplierList)     // 获取crmSupplier表列表
	}
}
