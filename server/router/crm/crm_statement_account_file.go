package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmStatementAccountFileRouter struct {
}

// InitCrmStatementAccountFileRouter 初始化 对账单文件 路由信息
func (s *CrmStatementAccountFileRouter) InitCrmStatementAccountFileRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmStatementAccountFileRouter := Router.Group("crmStatementAccountFile").Use(middleware.OperationRecord())
	crmStatementAccountFileRouterWithoutRecord := Router.Group("crmStatementAccountFile")
	crmStatementAccountFileRouterWithoutAuth := PublicRouter.Group("crmStatementAccountFile")

	var crmStatementAccountFileApi = v1.ApiGroupApp.CrmApiGroup.CrmStatementAccountFileApi
	{
		crmStatementAccountFileRouter.POST("createCrmStatementAccountFile", crmStatementAccountFileApi.CreateCrmStatementAccountFile)   // 新建对账单文件
		crmStatementAccountFileRouter.DELETE("deleteCrmStatementAccountFile", crmStatementAccountFileApi.DeleteCrmStatementAccountFile) // 删除对账单文件
		crmStatementAccountFileRouter.DELETE("deleteCrmStatementAccountFileByIds", crmStatementAccountFileApi.DeleteCrmStatementAccountFileByIds) // 批量删除对账单文件
		crmStatementAccountFileRouter.PUT("updateCrmStatementAccountFile", crmStatementAccountFileApi.UpdateCrmStatementAccountFile)    // 更新对账单文件
	}
	{
		crmStatementAccountFileRouterWithoutRecord.GET("findCrmStatementAccountFile", crmStatementAccountFileApi.FindCrmStatementAccountFile)        // 根据ID获取对账单文件
		crmStatementAccountFileRouterWithoutRecord.GET("getCrmStatementAccountFileList", crmStatementAccountFileApi.GetCrmStatementAccountFileList)  // 获取对账单文件列表
	}
	{
	    crmStatementAccountFileRouterWithoutAuth.GET("getCrmStatementAccountFilePublic", crmStatementAccountFileApi.GetCrmStatementAccountFilePublic)  // 获取对账单文件列表
	}
}
