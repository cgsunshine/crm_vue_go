package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmStatementAccountFileUploadAndDownloadsRouter struct {
}

// InitCrmStatementAccountFileUploadAndDownloadsRouter 初始化 对账单上传文件 路由信息
func (s *CrmStatementAccountFileUploadAndDownloadsRouter) InitCrmStatementAccountFileUploadAndDownloadsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmStatementAccountFileUploadAndDownloadsRouter := Router.Group("crmStatementAccountFileUploadAndDownloads").Use(middleware.OperationRecord())
	crmStatementAccountFileUploadAndDownloadsRouterWithoutRecord := Router.Group("crmStatementAccountFileUploadAndDownloads")
	crmStatementAccountFileUploadAndDownloadsRouterWithoutAuth := PublicRouter.Group("crmStatementAccountFileUploadAndDownloads")

	var crmStatementAccountFileUploadAndDownloadsApi = v1.ApiGroupApp.CrmApiGroup.CrmStatementAccountFileUploadAndDownloadsApi
	{
		crmStatementAccountFileUploadAndDownloadsRouter.POST("createCrmStatementAccountFileUploadAndDownloads", crmStatementAccountFileUploadAndDownloadsApi.CreateCrmStatementAccountFileUploadAndDownloads)   // 新建对账单上传文件
		crmStatementAccountFileUploadAndDownloadsRouter.DELETE("deleteCrmStatementAccountFileUploadAndDownloads", crmStatementAccountFileUploadAndDownloadsApi.DeleteCrmStatementAccountFileUploadAndDownloads) // 删除对账单上传文件
		crmStatementAccountFileUploadAndDownloadsRouter.DELETE("deleteCrmStatementAccountFileUploadAndDownloadsByIds", crmStatementAccountFileUploadAndDownloadsApi.DeleteCrmStatementAccountFileUploadAndDownloadsByIds) // 批量删除对账单上传文件
		crmStatementAccountFileUploadAndDownloadsRouter.PUT("updateCrmStatementAccountFileUploadAndDownloads", crmStatementAccountFileUploadAndDownloadsApi.UpdateCrmStatementAccountFileUploadAndDownloads)    // 更新对账单上传文件
	}
	{
		crmStatementAccountFileUploadAndDownloadsRouterWithoutRecord.GET("findCrmStatementAccountFileUploadAndDownloads", crmStatementAccountFileUploadAndDownloadsApi.FindCrmStatementAccountFileUploadAndDownloads)        // 根据ID获取对账单上传文件
		crmStatementAccountFileUploadAndDownloadsRouterWithoutRecord.GET("getCrmStatementAccountFileUploadAndDownloadsList", crmStatementAccountFileUploadAndDownloadsApi.GetCrmStatementAccountFileUploadAndDownloadsList)  // 获取对账单上传文件列表
	}
	{
	    crmStatementAccountFileUploadAndDownloadsRouterWithoutAuth.GET("getCrmStatementAccountFileUploadAndDownloadsPublic", crmStatementAccountFileUploadAndDownloadsApi.GetCrmStatementAccountFileUploadAndDownloadsPublic)  // 获取对账单上传文件列表
	}
}
