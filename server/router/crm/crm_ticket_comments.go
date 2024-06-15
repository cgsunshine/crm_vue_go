package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmTicketCommentsRouter struct {
}

// InitCrmTicketCommentsRouter 初始化 共单回复 路由信息
func (s *CrmTicketCommentsRouter) InitCrmTicketCommentsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmTicketCommentsRouter := Router.Group("crmTicketComments").Use(middleware.OperationRecord())
	crmTicketCommentsRouterWithoutRecord := Router.Group("crmTicketComments")
	crmTicketCommentsRouterWithoutAuth := PublicRouter.Group("crmTicketComments")

	var crmTicketCommentsApi = v1.ApiGroupApp.CrmApiGroup.CrmTicketCommentsApi
	{
		crmTicketCommentsRouter.POST("createCrmTicketComments", crmTicketCommentsApi.CreateCrmTicketComments)   // 新建共单回复
		crmTicketCommentsRouter.DELETE("deleteCrmTicketComments", crmTicketCommentsApi.DeleteCrmTicketComments) // 删除共单回复
		crmTicketCommentsRouter.DELETE("deleteCrmTicketCommentsByIds", crmTicketCommentsApi.DeleteCrmTicketCommentsByIds) // 批量删除共单回复
		crmTicketCommentsRouter.PUT("updateCrmTicketComments", crmTicketCommentsApi.UpdateCrmTicketComments)    // 更新共单回复
	}
	{
		crmTicketCommentsRouterWithoutRecord.GET("findCrmTicketComments", crmTicketCommentsApi.FindCrmTicketComments)        // 根据ID获取共单回复
		crmTicketCommentsRouterWithoutRecord.GET("getCrmTicketCommentsList", crmTicketCommentsApi.GetCrmTicketCommentsList)  // 获取共单回复列表
	}
	{
	    crmTicketCommentsRouterWithoutAuth.GET("getCrmTicketCommentsPublic", crmTicketCommentsApi.GetCrmTicketCommentsPublic)  // 获取共单回复列表
	}
}
