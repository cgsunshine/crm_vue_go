package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmProcurementContractRouter struct {
}

// InitCrmProcurementContractRouter 初始化 订购合同 路由信息
func (s *CrmProcurementContractRouter) InitCrmProcurementContractRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmProcurementContractRouter := Router.Group("crmProcurementContract").Use(middleware.OperationRecord())
	crmProcurementContractRouterWithoutRecord := Router.Group("crmProcurementContract")
	crmProcurementContractRouterWithoutAuth := PublicRouter.Group("crmProcurementContract")

	var crmProcurementContractApi = v1.ApiGroupApp.CrmApiGroup.CrmProcurementContractApi
	{
		crmProcurementContractRouter.POST("createCrmProcurementContract", crmProcurementContractApi.CreateCrmProcurementContract)             // 新建订购合同
		crmProcurementContractRouter.DELETE("deleteCrmProcurementContract", crmProcurementContractApi.DeleteCrmProcurementContract)           // 删除订购合同
		crmProcurementContractRouter.DELETE("deleteCrmProcurementContractByIds", crmProcurementContractApi.DeleteCrmProcurementContractByIds) // 批量删除订购合同
		crmProcurementContractRouter.PUT("updateCrmProcurementContract", crmProcurementContractApi.UpdateCrmProcurementContract)              // 更新订购合同
	}
	{
		crmProcurementContractRouterWithoutRecord.GET("findCrmProcurementContract", crmProcurementContractApi.FindCrmProcurementContract) // 根据ID获取订购合同
		//crmProcurementContractRouterWithoutRecord.GET("getCrmProcurementContractList", crmProcurementContractApi.GetCrmProcurementContractList)  // 获取订购合同列表
	}
	{
		crmProcurementContractRouterWithoutAuth.GET("getCrmProcurementContractPublic", crmProcurementContractApi.GetCrmProcurementContractPublic) // 获取订购合同列表
		crmProcurementContractRouterWithoutAuth.GET("getCrmProcurementContractList", crmProcurementContractApi.GetCrmProcurementContractList)     // 获取订购合同列表
	}
}
