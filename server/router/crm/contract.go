package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPageContractRouter struct {
}

// InitCrmContractRouter 初始化 crmContract表 路由信息
func (s *CrmPageContractRouter) InitCrmPageContractRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmContractRouter := Router.Group("crmContract").Use(middleware.OperationRecord())
	crmContractRouterWithoutRecord := Router.Group("crmContract")
	crmContractRouterWithoutAuth := PublicRouter.Group("crmContract")

	var crmContractApi = v1.ApiGroupApp.CrmApiGroup.CrmContractApi

	{
		crmContractRouter.POST("createPageCrmContract", crmContractApi.CreatePageCrmContract) // 新建crmContract表
	}

	{
		crmContractRouterWithoutRecord.GET("findCrmPageContract", crmContractApi.FindCrmPageContract) // 根据ID获取crmContract表
	}

	{
		crmContractRouterWithoutAuth.GET("getCrmPageContractList", crmContractApi.GetCrmPageContractList) // 获取crmContract表列表
	}

}
