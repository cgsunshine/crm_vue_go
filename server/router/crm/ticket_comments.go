package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPageTicketCommentsRouter struct {
}

// InitCrmTicketCommentsRouter 初始化 共单回复 路由信息
func (s *CrmPageTicketCommentsRouter) InitCrmPageTicketCommentsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmTicketCommentsRouter := Router.Group("crmTicketComments").Use(middleware.OperationRecord())
	crmTicketCommentsRouterWithoutRecord := Router.Group("crmTicketComments")
	crmTicketCommentsRouterWithoutAuth := PublicRouter.Group("crmTicketComments")

	var crmTicketCommentsApi = v1.ApiGroupApp.CrmApiGroup.CrmTicketCommentsApi
	{
		crmTicketCommentsRouter.POST("createCrmPageTicketComments", crmTicketCommentsApi.CreateCrmPageTicketComments) // 新建共单回复
	}

	{
		crmTicketCommentsRouterWithoutRecord.GET("findCrmPageTicketComments", crmTicketCommentsApi.FindCrmPageTicketComments) // 根据ID获取共单回复
	}

	{
		crmTicketCommentsRouterWithoutAuth.GET("getCrmPageTicketCommentsList", crmTicketCommentsApi.GetCrmPageTicketCommentsList) // 获取共单回复列表
	}
}
