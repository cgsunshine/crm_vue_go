package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmContactApprovalRecordRouter struct {
}

// InitCrmContactApprovalRecordRouter 初始化 合同审批记录 路由信息
func (s *CrmContactApprovalRecordRouter) InitCrmContactApprovalRecordRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmContactApprovalRecordRouter := Router.Group("crmContactApprovalRecord").Use(middleware.OperationRecord())
	crmContactApprovalRecordRouterWithoutRecord := Router.Group("crmContactApprovalRecord")
	crmContactApprovalRecordRouterWithoutAuth := PublicRouter.Group("crmContactApprovalRecord")

	var crmContactApprovalRecordApi = v1.ApiGroupApp.CrmApiGroup.CrmContactApprovalRecordApi
	{
		crmContactApprovalRecordRouter.POST("createCrmContactApprovalRecord", crmContactApprovalRecordApi.CreateCrmContactApprovalRecord)   // 新建合同审批记录
		crmContactApprovalRecordRouter.DELETE("deleteCrmContactApprovalRecord", crmContactApprovalRecordApi.DeleteCrmContactApprovalRecord) // 删除合同审批记录
		crmContactApprovalRecordRouter.DELETE("deleteCrmContactApprovalRecordByIds", crmContactApprovalRecordApi.DeleteCrmContactApprovalRecordByIds) // 批量删除合同审批记录
		crmContactApprovalRecordRouter.PUT("updateCrmContactApprovalRecord", crmContactApprovalRecordApi.UpdateCrmContactApprovalRecord)    // 更新合同审批记录
	}
	{
		crmContactApprovalRecordRouterWithoutRecord.GET("findCrmContactApprovalRecord", crmContactApprovalRecordApi.FindCrmContactApprovalRecord)        // 根据ID获取合同审批记录
		crmContactApprovalRecordRouterWithoutRecord.GET("getCrmContactApprovalRecordList", crmContactApprovalRecordApi.GetCrmContactApprovalRecordList)  // 获取合同审批记录列表
	}
	{
	    crmContactApprovalRecordRouterWithoutAuth.GET("getCrmContactApprovalRecordPublic", crmContactApprovalRecordApi.GetCrmContactApprovalRecordPublic)  // 获取合同审批记录列表
	}
}
