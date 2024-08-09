package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmUploadStatementAccountFileRouter struct {
}

// InitCrmStatementAccountFileRouter 初始化 对账单文件 路由信息
func (s *CrmUploadStatementAccountFileRouter) InitCrmUploadStatementAccountFileRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {

	crmStatementAccountFileRouterWithoutRecord := Router.Group("crmStatementAccountFile")

	var crmStatementAccountFileApi = v1.ApiGroupApp.CrmApiGroup.CrmStatementAccountFileApi

	{
		crmStatementAccountFileRouterWithoutRecord.POST("uploadFile", crmStatementAccountFileApi.UploadFile)    // 获取对账单文件列表
		crmStatementAccountFileRouterWithoutRecord.GET("downloadFile", crmStatementAccountFileApi.DownloadFile) // 获取对账单文件列表
	}

}
