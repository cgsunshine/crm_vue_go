package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPageContractRouter struct {
}

// InitCrmContractRouter 初始化 crmContract表 路由信息
func (s *CrmPageContractRouter) InitCrmPageContractRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmContractRouterWithoutRecord := Router.Group("crmContract")

	var crmContractApi = v1.ApiGroupApp.CrmApiGroup.CrmContractApi

	{
		crmContractRouterWithoutRecord.GET("getCrmPageContractList", crmContractApi.GetCrmPageContractList) // 获取crmContract表列表
	}

}
