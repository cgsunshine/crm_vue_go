package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmApprovalNodeRouter struct {
}

// InitCrmApprovalNodeRouter 初始化 节点审批 路由信息
func (s *CrmApprovalNodeRouter) InitCrmApprovalNodeRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmApprovalNodeRouter := Router.Group("crmApprovalNode").Use(middleware.OperationRecord())
	crmApprovalNodeRouterWithoutRecord := Router.Group("crmApprovalNode")
	crmApprovalNodeRouterWithoutAuth := PublicRouter.Group("crmApprovalNode")

	var crmApprovalNodeApi = v1.ApiGroupApp.CrmApiGroup.CrmApprovalNodeApi
	{
		crmApprovalNodeRouter.POST("createCrmApprovalNode", crmApprovalNodeApi.CreateCrmApprovalNode)   // 新建节点审批
		crmApprovalNodeRouter.DELETE("deleteCrmApprovalNode", crmApprovalNodeApi.DeleteCrmApprovalNode) // 删除节点审批
		crmApprovalNodeRouter.DELETE("deleteCrmApprovalNodeByIds", crmApprovalNodeApi.DeleteCrmApprovalNodeByIds) // 批量删除节点审批
		crmApprovalNodeRouter.PUT("updateCrmApprovalNode", crmApprovalNodeApi.UpdateCrmApprovalNode)    // 更新节点审批
	}
	{
		crmApprovalNodeRouterWithoutRecord.GET("findCrmApprovalNode", crmApprovalNodeApi.FindCrmApprovalNode)        // 根据ID获取节点审批
		crmApprovalNodeRouterWithoutRecord.GET("getCrmApprovalNodeList", crmApprovalNodeApi.GetCrmApprovalNodeList)  // 获取节点审批列表
	}
	{
	    crmApprovalNodeRouterWithoutAuth.GET("getCrmApprovalNodePublic", crmApprovalNodeApi.GetCrmApprovalNodePublic)  // 获取节点审批列表
	}
}
