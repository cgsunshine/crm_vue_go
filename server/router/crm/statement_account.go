package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPageStatementAccountRouter struct {
}

// InitCrmStatementAccountRouter 初始化 crmStatementAccount表 路由信息
func (s *CrmPageStatementAccountRouter) InitCrmPageStatementAccountRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmStatementAccountRouterWithoutRecord := Router.Group("crmStatementAccount")

	var crmStatementAccountApi = v1.ApiGroupApp.CrmApiGroup.CrmStatementAccountApi

	{
		crmStatementAccountRouterWithoutRecord.GET("getCrmPageStatementAccountList", crmStatementAccountApi.GetCrmPageStatementAccountList) // 获取crmStatementAccount表列表
	}

}
