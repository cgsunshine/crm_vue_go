package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmContractRouter struct {
}

// InitCrmContractRouter 初始化 合同管理 路由信息
func (s *CrmContractRouter) InitCrmContractRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmContractRouter := Router.Group("crmContract").Use(middleware.OperationRecord())
	crmContractRouterWithoutRecord := Router.Group("crmContract")
	crmContractRouterWithoutAuth := PublicRouter.Group("crmContract")

	var crmContractApi = v1.ApiGroupApp.CrmApiGroup.CrmContractApi
	{
		crmContractRouter.POST("createCrmContract", crmContractApi.CreateCrmContract)   // 新建合同管理
		crmContractRouter.DELETE("deleteCrmContract", crmContractApi.DeleteCrmContract) // 删除合同管理
		crmContractRouter.DELETE("deleteCrmContractByIds", crmContractApi.DeleteCrmContractByIds) // 批量删除合同管理
		crmContractRouter.PUT("updateCrmContract", crmContractApi.UpdateCrmContract)    // 更新合同管理
	}
	{
		crmContractRouterWithoutRecord.GET("findCrmContract", crmContractApi.FindCrmContract)        // 根据ID获取合同管理
		crmContractRouterWithoutRecord.GET("getCrmContractList", crmContractApi.GetCrmContractList)  // 获取合同管理列表
	}
	{
	    crmContractRouterWithoutAuth.GET("getCrmContractPublic", crmContractApi.GetCrmContractPublic)  // 获取合同管理列表
	}
}
