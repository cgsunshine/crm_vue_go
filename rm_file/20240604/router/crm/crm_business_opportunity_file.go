package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmBusinessOpportunityFileRouter struct {
}

// InitCrmBusinessOpportunityFileRouter 初始化 商机文件 路由信息
func (s *CrmBusinessOpportunityFileRouter) InitCrmBusinessOpportunityFileRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmBusinessOpportunityFileRouter := Router.Group("crmBusinessOpportunityFile").Use(middleware.OperationRecord())
	crmBusinessOpportunityFileRouterWithoutRecord := Router.Group("crmBusinessOpportunityFile")
	crmBusinessOpportunityFileRouterWithoutAuth := PublicRouter.Group("crmBusinessOpportunityFile")

	var crmBusinessOpportunityFileApi = v1.ApiGroupApp.CrmApiGroup.CrmBusinessOpportunityFileApi
	{
		crmBusinessOpportunityFileRouter.POST("createCrmBusinessOpportunityFile", crmBusinessOpportunityFileApi.CreateCrmBusinessOpportunityFile)   // 新建商机文件
		crmBusinessOpportunityFileRouter.DELETE("deleteCrmBusinessOpportunityFile", crmBusinessOpportunityFileApi.DeleteCrmBusinessOpportunityFile) // 删除商机文件
		crmBusinessOpportunityFileRouter.DELETE("deleteCrmBusinessOpportunityFileByIds", crmBusinessOpportunityFileApi.DeleteCrmBusinessOpportunityFileByIds) // 批量删除商机文件
		crmBusinessOpportunityFileRouter.PUT("updateCrmBusinessOpportunityFile", crmBusinessOpportunityFileApi.UpdateCrmBusinessOpportunityFile)    // 更新商机文件
	}
	{
		crmBusinessOpportunityFileRouterWithoutRecord.GET("findCrmBusinessOpportunityFile", crmBusinessOpportunityFileApi.FindCrmBusinessOpportunityFile)        // 根据ID获取商机文件
		crmBusinessOpportunityFileRouterWithoutRecord.GET("getCrmBusinessOpportunityFileList", crmBusinessOpportunityFileApi.GetCrmBusinessOpportunityFileList)  // 获取商机文件列表
	}
	{
	    crmBusinessOpportunityFileRouterWithoutAuth.GET("getCrmBusinessOpportunityFilePublic", crmBusinessOpportunityFileApi.GetCrmBusinessOpportunityFilePublic)  // 获取商机文件列表
	}
}
