package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmTicketResponseTemplatesRouter struct {
}

// InitCrmTicketResponseTemplatesRouter 初始化 快捷回复模板 路由信息
func (s *CrmTicketResponseTemplatesRouter) InitCrmTicketResponseTemplatesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmTicketResponseTemplatesRouter := Router.Group("crmTicketResponseTemplates").Use(middleware.OperationRecord())
	crmTicketResponseTemplatesRouterWithoutRecord := Router.Group("crmTicketResponseTemplates")
	crmTicketResponseTemplatesRouterWithoutAuth := PublicRouter.Group("crmTicketResponseTemplates")

	var crmTicketResponseTemplatesApi = v1.ApiGroupApp.CrmApiGroup.CrmTicketResponseTemplatesApi
	{
		crmTicketResponseTemplatesRouter.POST("createCrmTicketResponseTemplates", crmTicketResponseTemplatesApi.CreateCrmTicketResponseTemplates)             // 新建快捷回复模板
		crmTicketResponseTemplatesRouter.DELETE("deleteCrmTicketResponseTemplates", crmTicketResponseTemplatesApi.DeleteCrmTicketResponseTemplates)           // 删除快捷回复模板
		crmTicketResponseTemplatesRouter.DELETE("deleteCrmTicketResponseTemplatesByIds", crmTicketResponseTemplatesApi.DeleteCrmTicketResponseTemplatesByIds) // 批量删除快捷回复模板
		crmTicketResponseTemplatesRouter.PUT("updateCrmTicketResponseTemplates", crmTicketResponseTemplatesApi.UpdateCrmTicketResponseTemplates)              // 更新快捷回复模板
	}
	{
		crmTicketResponseTemplatesRouterWithoutRecord.GET("findCrmTicketResponseTemplates", crmTicketResponseTemplatesApi.FindCrmTicketResponseTemplates) // 根据ID获取快捷回复模板
		//crmTicketResponseTemplatesRouterWithoutRecord.GET("getCrmTicketResponseTemplatesList", crmTicketResponseTemplatesApi.GetCrmTicketResponseTemplatesList)  // 获取快捷回复模板列表
	}
	{
		crmTicketResponseTemplatesRouterWithoutAuth.GET("getCrmTicketResponseTemplatesPublic", crmTicketResponseTemplatesApi.GetCrmTicketResponseTemplatesPublic) // 获取快捷回复模板列表
		crmTicketResponseTemplatesRouterWithoutAuth.GET("getCrmTicketResponseTemplatesList", crmTicketResponseTemplatesApi.GetCrmTicketResponseTemplatesList)     // 获取快捷回复模板列表
	}
}
