package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmContactFileRouter struct {
}

// InitCrmContactFileRouter 初始化 合同文件 路由信息
func (s *CrmContactFileRouter) InitCrmContactFileRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmContactFileRouter := Router.Group("crmContactFile").Use(middleware.OperationRecord())
	crmContactFileRouterWithoutRecord := Router.Group("crmContactFile")
	crmContactFileRouterWithoutAuth := PublicRouter.Group("crmContactFile")

	var crmContactFileApi = v1.ApiGroupApp.CrmApiGroup.CrmContactFileApi
	{
		crmContactFileRouter.POST("createCrmContactFile", crmContactFileApi.CreateCrmContactFile)   // 新建合同文件
		crmContactFileRouter.DELETE("deleteCrmContactFile", crmContactFileApi.DeleteCrmContactFile) // 删除合同文件
		crmContactFileRouter.DELETE("deleteCrmContactFileByIds", crmContactFileApi.DeleteCrmContactFileByIds) // 批量删除合同文件
		crmContactFileRouter.PUT("updateCrmContactFile", crmContactFileApi.UpdateCrmContactFile)    // 更新合同文件
	}
	{
		crmContactFileRouterWithoutRecord.GET("findCrmContactFile", crmContactFileApi.FindCrmContactFile)        // 根据ID获取合同文件
		crmContactFileRouterWithoutRecord.GET("getCrmContactFileList", crmContactFileApi.GetCrmContactFileList)  // 获取合同文件列表
	}
	{
	    crmContactFileRouterWithoutAuth.GET("getCrmContactFilePublic", crmContactFileApi.GetCrmContactFilePublic)  // 获取合同文件列表
	}
}
