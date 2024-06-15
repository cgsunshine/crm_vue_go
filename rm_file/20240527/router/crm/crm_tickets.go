package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmTicketsRouter struct {
}

// InitCrmTicketsRouter 初始化 crmTickets表 路由信息
func (s *CrmTicketsRouter) InitCrmTicketsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmTicketsRouter := Router.Group("crmTickets").Use(middleware.OperationRecord())
	crmTicketsRouterWithoutRecord := Router.Group("crmTickets")
	crmTicketsRouterWithoutAuth := PublicRouter.Group("crmTickets")

	var crmTicketsApi = v1.ApiGroupApp.CrmApiGroup.CrmTicketsApi
	{
		crmTicketsRouter.POST("createCrmTickets", crmTicketsApi.CreateCrmTickets)   // 新建crmTickets表
		crmTicketsRouter.DELETE("deleteCrmTickets", crmTicketsApi.DeleteCrmTickets) // 删除crmTickets表
		crmTicketsRouter.DELETE("deleteCrmTicketsByIds", crmTicketsApi.DeleteCrmTicketsByIds) // 批量删除crmTickets表
		crmTicketsRouter.PUT("updateCrmTickets", crmTicketsApi.UpdateCrmTickets)    // 更新crmTickets表
	}
	{
		crmTicketsRouterWithoutRecord.GET("findCrmTickets", crmTicketsApi.FindCrmTickets)        // 根据ID获取crmTickets表
		crmTicketsRouterWithoutRecord.GET("getCrmTicketsList", crmTicketsApi.GetCrmTicketsList)  // 获取crmTickets表列表
	}
	{
	    crmTicketsRouterWithoutAuth.GET("getCrmTicketsPublic", crmTicketsApi.GetCrmTicketsPublic)  // 获取crmTickets表列表
	}
}
