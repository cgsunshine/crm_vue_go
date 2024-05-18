package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmTicketCommentsRouter struct {
}

// InitCrmTicketCommentsRouter 初始化 crmTicketComments表 路由信息
func (s *CrmTicketCommentsRouter) InitCrmTicketCommentsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmTicketCommentsRouter := Router.Group("crmTicketComments").Use(middleware.OperationRecord())
	crmTicketCommentsRouterWithoutRecord := Router.Group("crmTicketComments")
	crmTicketCommentsRouterWithoutAuth := PublicRouter.Group("crmTicketComments")

	var crmTicketCommentsApi = v1.ApiGroupApp.CrmApiGroup.CrmTicketCommentsApi
	{
		crmTicketCommentsRouter.POST("createCrmTicketComments", crmTicketCommentsApi.CreateCrmTicketComments)   // 新建crmTicketComments表
		crmTicketCommentsRouter.DELETE("deleteCrmTicketComments", crmTicketCommentsApi.DeleteCrmTicketComments) // 删除crmTicketComments表
		crmTicketCommentsRouter.DELETE("deleteCrmTicketCommentsByIds", crmTicketCommentsApi.DeleteCrmTicketCommentsByIds) // 批量删除crmTicketComments表
		crmTicketCommentsRouter.PUT("updateCrmTicketComments", crmTicketCommentsApi.UpdateCrmTicketComments)    // 更新crmTicketComments表
	}
	{
		crmTicketCommentsRouterWithoutRecord.GET("findCrmTicketComments", crmTicketCommentsApi.FindCrmTicketComments)        // 根据ID获取crmTicketComments表
		crmTicketCommentsRouterWithoutRecord.GET("getCrmTicketCommentsList", crmTicketCommentsApi.GetCrmTicketCommentsList)  // 获取crmTicketComments表列表
	}
	{
	    crmTicketCommentsRouterWithoutAuth.GET("getCrmTicketCommentsPublic", crmTicketCommentsApi.GetCrmTicketCommentsPublic)  // 获取crmTicketComments表列表
	}
}
