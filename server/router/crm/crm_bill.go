package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmBillRouter struct {
}

// InitCrmBillRouter 初始化 crmBill表 路由信息
func (s *CrmBillRouter) InitCrmBillRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmBillRouter := Router.Group("crmBill").Use(middleware.OperationRecord())
	crmBillRouterWithoutRecord := Router.Group("crmBill")
	crmBillRouterWithoutAuth := PublicRouter.Group("crmBill")

	var crmBillApi = v1.ApiGroupApp.CrmApiGroup.CrmBillApi
	{
		crmBillRouter.POST("createCrmBill", crmBillApi.CreateCrmBill)   // 新建crmBill表
		crmBillRouter.DELETE("deleteCrmBill", crmBillApi.DeleteCrmBill) // 删除crmBill表
		crmBillRouter.DELETE("deleteCrmBillByIds", crmBillApi.DeleteCrmBillByIds) // 批量删除crmBill表
		crmBillRouter.PUT("updateCrmBill", crmBillApi.UpdateCrmBill)    // 更新crmBill表
	}
	{
		crmBillRouterWithoutRecord.GET("findCrmBill", crmBillApi.FindCrmBill)        // 根据ID获取crmBill表
		crmBillRouterWithoutRecord.GET("getCrmBillList", crmBillApi.GetCrmBillList)  // 获取crmBill表列表
	}
	{
	    crmBillRouterWithoutAuth.GET("getCrmBillPublic", crmBillApi.GetCrmBillPublic)  // 获取crmBill表列表
	}
}
