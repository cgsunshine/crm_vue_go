package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPageApprovalNodeRouter struct {
}

// InitCrmPageApprovalNodeRouter 初始化 节点审批 路由信息
func (s *CrmPageApprovalNodeRouter) InitCrmPageApprovalNodeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	crmApprovalNodeRouter := Router.Group("crmApprovalNode").Use(middleware.OperationRecord())
	crmApprovalNodeRouterWithoutRecord := Router.Group("crmApprovalNode")
	//crmApprovalNodeRouterWithoutAuth := PublicRouter.Group("crmApprovalNode")

	var crmApprovalNodeApi = v1.ApiGroupApp.CrmApiGroup.CrmApprovalNodeApi
	{
		crmApprovalNodeRouter.POST("createCrmPageApprovalNode", crmApprovalNodeApi.CreateCrmPageApprovalNode) // 新建多个节点审批
		//crmApprovalNodeRouter.POST("createCrmPageOneApprovalNode", crmApprovalNodeApi.CreateCrmPageOneApprovalNode) // 新建一个节点审批
	}
	{
		crmApprovalNodeRouterWithoutRecord.GET("getCrmPageApprovalNodeList", crmApprovalNodeApi.GetCrmPageApprovalNodeList) // 根据ID获取节点审批
	}
}
