package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmUploadBusinessOpportunityFileRouter struct {
}

// InitCrmBusinessOpportunityFileRouter 初始化 商机文件 路由信息
func (s *CrmUploadBusinessOpportunityFileRouter) InitCrmUploadBusinessOpportunityFileRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmBusinessOpportunityFileRouterWithoutRecord := Router.Group("crmBusinessOpportunityFile")

	var crmBusinessOpportunityFileApi = v1.ApiGroupApp.CrmApiGroup.CrmBusinessOpportunityFileApi

	{
		crmBusinessOpportunityFileRouterWithoutRecord.POST("uploadFile", crmBusinessOpportunityFileApi.UploadFile)    // 上传文件
		crmBusinessOpportunityFileRouterWithoutRecord.GET("downloadFile", crmBusinessOpportunityFileApi.DownloadFile) // 上传文件
	}

}
