package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmContractTypeRouter struct {
}

// InitCrmContractTypeRouter 初始化 合同类型 路由信息
func (s *CrmContractTypeRouter) InitCrmContractTypeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmContractTypeRouter := Router.Group("crmContractType").Use(middleware.OperationRecord())
	crmContractTypeRouterWithoutRecord := Router.Group("crmContractType")
	crmContractTypeRouterWithoutAuth := PublicRouter.Group("crmContractType")

	var crmContractTypeApi = v1.ApiGroupApp.CrmApiGroup.CrmContractTypeApi
	{
		crmContractTypeRouter.POST("createCrmContractType", crmContractTypeApi.CreateCrmContractType)             // 新建合同类型
		crmContractTypeRouter.DELETE("deleteCrmContractType", crmContractTypeApi.DeleteCrmContractType)           // 删除合同类型
		crmContractTypeRouter.DELETE("deleteCrmContractTypeByIds", crmContractTypeApi.DeleteCrmContractTypeByIds) // 批量删除合同类型
		crmContractTypeRouter.PUT("updateCrmContractType", crmContractTypeApi.UpdateCrmContractType)              // 更新合同类型
	}
	{
		crmContractTypeRouterWithoutRecord.GET("findCrmContractType", crmContractTypeApi.FindCrmContractType) // 根据ID获取合同类型
		//crmContractTypeRouterWithoutRecord.GET("getCrmContractTypeList", crmContractTypeApi.GetCrmContractTypeList)  // 获取合同类型列表
	}
	{
		crmContractTypeRouterWithoutAuth.GET("getCrmContractTypePublic", crmContractTypeApi.GetCrmContractTypePublic) // 获取合同类型列表
		crmContractTypeRouterWithoutAuth.GET("getCrmContractTypeList", crmContractTypeApi.GetCrmContractTypeList)     // 获取合同类型列表
	}
}
