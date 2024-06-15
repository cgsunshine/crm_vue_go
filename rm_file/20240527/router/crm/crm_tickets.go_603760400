package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmTicketsRouter struct {
}

// InitCrmTicketsRouter 初始化 工单 路由信息
func (s *CrmTicketsRouter) InitCrmTicketsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmTicketsRouter := Router.Group("crmTickets").Use(middleware.OperationRecord())
	crmTicketsRouterWithoutRecord := Router.Group("crmTickets")
	crmTicketsRouterWithoutAuth := PublicRouter.Group("crmTickets")

	var crmTicketsApi = v1.ApiGroupApp.CrmApiGroup.CrmTicketsApi
	{
		crmTicketsRouter.POST("createCrmTickets", crmTicketsApi.CreateCrmTickets)   // 新建工单
		crmTicketsRouter.DELETE("deleteCrmTickets", crmTicketsApi.DeleteCrmTickets) // 删除工单
		crmTicketsRouter.DELETE("deleteCrmTicketsByIds", crmTicketsApi.DeleteCrmTicketsByIds) // 批量删除工单
		crmTicketsRouter.PUT("updateCrmTickets", crmTicketsApi.UpdateCrmTickets)    // 更新工单
	}
	{
		crmTicketsRouterWithoutRecord.GET("findCrmTickets", crmTicketsApi.FindCrmTickets)        // 根据ID获取工单
		crmTicketsRouterWithoutRecord.GET("getCrmTicketsList", crmTicketsApi.GetCrmTicketsList)  // 获取工单列表
	}
	{
	    crmTicketsRouterWithoutAuth.GET("getCrmTicketsPublic", crmTicketsApi.GetCrmTicketsPublic)  // 获取工单列表
	}
}
