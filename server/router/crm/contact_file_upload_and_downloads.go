package crm

import (
	"github.com/gin-gonic/gin"
)

type CrmPageContactFileUploadAndDownloadsRouter struct {
}

// InitCrmContactFileUploadAndDownloadsRouter 初始化 crmContactFileUploadAndDownloads表 路由信息
func (s *CrmPageContactFileUploadAndDownloadsRouter) InitCrmPageContactFileUploadAndDownloadsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	//crmContactFileUploadAndDownloadsRouterWithoutRecord := Router.Group("crmContactFileUploadAndDownloads")
	//
	//var crmContactFileUploadAndDownloadsApi = v1.ApiGroupApp.CrmApiGroup.CrmContactFileUploadAndDownloadsApi
	//
	//{
	//	crmContactFileUploadAndDownloadsRouterWithoutRecord.POST("uploadContactFile", crmContactFileUploadAndDownloadsApi.UploadContactFile)               // 上传合同文件
	//	crmContactFileUploadAndDownloadsRouterWithoutRecord.POST("updateCrmContactFileSort", crmContactFileUploadAndDownloadsApi.UpdateCrmContactFileSort) // 修改文件排序
	//}

}
