package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmUploadContactFileRouter struct {
}

// InitCrmContactFileRouter 初始化 合同文件 路由信息
func (s *CrmUploadContactFileRouter) InitCrmUploadContactFileRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmContactFileRouterWithoutRecord := Router.Group("crmContactFile")

	var crmContactFileApi = v1.ApiGroupApp.CrmApiGroup.CrmContactFileApi

	{
		crmContactFileRouterWithoutRecord.POST("uploadFile", crmContactFileApi.UploadFile)    // 获取合同文件列表
		crmContactFileRouterWithoutRecord.GET("downloadFile", crmContactFileApi.DownloadFile) // 获取合同文件列表
	}

}
