package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmContactFileRouter struct {
}

// InitCrmContactFileRouter 初始化 crmContactFile表 路由信息
func (s *CrmContactFileRouter) InitCrmContactFileRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmContactFileRouter := Router.Group("crmContactFile").Use(middleware.OperationRecord())
	crmContactFileRouterWithoutRecord := Router.Group("crmContactFile")
	crmContactFileRouterWithoutAuth := PublicRouter.Group("crmContactFile")

	var crmContactFileApi = v1.ApiGroupApp.CrmApiGroup.CrmContactFileApi
	{
		crmContactFileRouter.POST("createCrmContactFile", crmContactFileApi.CreateCrmContactFile)   // 新建crmContactFile表
		crmContactFileRouter.DELETE("deleteCrmContactFile", crmContactFileApi.DeleteCrmContactFile) // 删除crmContactFile表
		crmContactFileRouter.DELETE("deleteCrmContactFileByIds", crmContactFileApi.DeleteCrmContactFileByIds) // 批量删除crmContactFile表
		crmContactFileRouter.PUT("updateCrmContactFile", crmContactFileApi.UpdateCrmContactFile)    // 更新crmContactFile表
	}
	{
		crmContactFileRouterWithoutRecord.GET("findCrmContactFile", crmContactFileApi.FindCrmContactFile)        // 根据ID获取crmContactFile表
		crmContactFileRouterWithoutRecord.GET("getCrmContactFileList", crmContactFileApi.GetCrmContactFileList)  // 获取crmContactFile表列表
	}
	{
	    crmContactFileRouterWithoutAuth.GET("getCrmContactFilePublic", crmContactFileApi.GetCrmContactFilePublic)  // 获取crmContactFile表列表
	}
}
