package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmContactFileUploadAndDownloadsRouter struct {
}

// InitCrmContactFileUploadAndDownloadsRouter 初始化 crmContactFileUploadAndDownloads表 路由信息
func (s *CrmContactFileUploadAndDownloadsRouter) InitCrmContactFileUploadAndDownloadsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmContactFileUploadAndDownloadsRouter := Router.Group("crmContactFileUploadAndDownloads").Use(middleware.OperationRecord())
	crmContactFileUploadAndDownloadsRouterWithoutRecord := Router.Group("crmContactFileUploadAndDownloads")
	crmContactFileUploadAndDownloadsRouterWithoutAuth := PublicRouter.Group("crmContactFileUploadAndDownloads")

	var crmContactFileUploadAndDownloadsApi = v1.ApiGroupApp.CrmApiGroup.CrmContactFileUploadAndDownloadsApi
	{
		crmContactFileUploadAndDownloadsRouter.POST("createCrmContactFileUploadAndDownloads", crmContactFileUploadAndDownloadsApi.CreateCrmContactFileUploadAndDownloads)   // 新建crmContactFileUploadAndDownloads表
		crmContactFileUploadAndDownloadsRouter.DELETE("deleteCrmContactFileUploadAndDownloads", crmContactFileUploadAndDownloadsApi.DeleteCrmContactFileUploadAndDownloads) // 删除crmContactFileUploadAndDownloads表
		crmContactFileUploadAndDownloadsRouter.DELETE("deleteCrmContactFileUploadAndDownloadsByIds", crmContactFileUploadAndDownloadsApi.DeleteCrmContactFileUploadAndDownloadsByIds) // 批量删除crmContactFileUploadAndDownloads表
		crmContactFileUploadAndDownloadsRouter.PUT("updateCrmContactFileUploadAndDownloads", crmContactFileUploadAndDownloadsApi.UpdateCrmContactFileUploadAndDownloads)    // 更新crmContactFileUploadAndDownloads表
	}
	{
		crmContactFileUploadAndDownloadsRouterWithoutRecord.GET("findCrmContactFileUploadAndDownloads", crmContactFileUploadAndDownloadsApi.FindCrmContactFileUploadAndDownloads)        // 根据ID获取crmContactFileUploadAndDownloads表
		crmContactFileUploadAndDownloadsRouterWithoutRecord.GET("getCrmContactFileUploadAndDownloadsList", crmContactFileUploadAndDownloadsApi.GetCrmContactFileUploadAndDownloadsList)  // 获取crmContactFileUploadAndDownloads表列表
	}
	{
	    crmContactFileUploadAndDownloadsRouterWithoutAuth.GET("getCrmContactFileUploadAndDownloadsPublic", crmContactFileUploadAndDownloadsApi.GetCrmContactFileUploadAndDownloadsPublic)  // 获取crmContactFileUploadAndDownloads表列表
	}
	{

	}
}
