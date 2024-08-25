package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPageTicketsRouter struct {
}

// InitCrmTicketsRouter 初始化 工单 路由信息
func (s *CrmPageTicketsRouter) InitCrmPageTicketsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	//crmTicketsRouter := Router.Group("crmTickets").Use(middleware.OperationRecord())
	crmTicketsRouterWithoutAuth := PublicRouter.Group("crmTickets")
	crmTicketsRouterWithoutRecord := Router.Group("crmTickets")

	var crmTicketsApi = v1.ApiGroupApp.CrmApiGroup.CrmTicketsApi
	{
		crmTicketsRouterWithoutRecord.GET("findCrmPageTickets", crmTicketsApi.FindCrmPageTickets) // 根据ID获取工单
	}

	{
		crmTicketsRouterWithoutAuth.GET("getCrmSubmitterTicketsList", crmTicketsApi.GetCrmSubmitterTicketsList) // 获取发起人工单列表
		crmTicketsRouterWithoutAuth.GET("getCrmAssigneeTicketsList", crmTicketsApi.GetCrmAssigneeTicketsList)   // 获取处理人工单列表
		crmTicketsRouterWithoutAuth.POST("updCrmTicketsCompleted", crmTicketsApi.UpdCrmTicketsCompleted)        // 工单完结接口
	}
}
