package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmContractRouter struct {
}

// InitCrmContractRouter 初始化 crmContract表 路由信息
func (s *CrmContractRouter) InitCrmContractRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmContractRouter := Router.Group("crmContract").Use(middleware.OperationRecord())
	crmContractRouterWithoutRecord := Router.Group("crmContract")
	crmContractRouterWithoutAuth := PublicRouter.Group("crmContract")

	var crmContractApi = v1.ApiGroupApp.CrmApiGroup.CrmContractApi
	{
		crmContractRouter.POST("createCrmContract", crmContractApi.CreateCrmContract)   // 新建crmContract表
		crmContractRouter.DELETE("deleteCrmContract", crmContractApi.DeleteCrmContract) // 删除crmContract表
		crmContractRouter.DELETE("deleteCrmContractByIds", crmContractApi.DeleteCrmContractByIds) // 批量删除crmContract表
		crmContractRouter.PUT("updateCrmContract", crmContractApi.UpdateCrmContract)    // 更新crmContract表
	}
	{
		crmContractRouterWithoutRecord.GET("findCrmContract", crmContractApi.FindCrmContract)        // 根据ID获取crmContract表
		crmContractRouterWithoutRecord.GET("getCrmContractList", crmContractApi.GetCrmContractList)  // 获取crmContract表列表
	}
	{
	    crmContractRouterWithoutAuth.GET("getCrmContractPublic", crmContractApi.GetCrmContractPublic)  // 获取crmContract表列表
	}
}
