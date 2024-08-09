package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmBusinessOpportunityApprovalRecordRouter struct {
}

// InitCrmBusinessOpportunityApprovalRecordRouter 初始化 商机审批记录 路由信息
func (s *CrmBusinessOpportunityApprovalRecordRouter) InitCrmBusinessOpportunityApprovalRecordRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmBusinessOpportunityApprovalRecordRouter := Router.Group("crmBusinessOpportunityApprovalRecord").Use(middleware.OperationRecord())
	crmBusinessOpportunityApprovalRecordRouterWithoutRecord := Router.Group("crmBusinessOpportunityApprovalRecord")
	crmBusinessOpportunityApprovalRecordRouterWithoutAuth := PublicRouter.Group("crmBusinessOpportunityApprovalRecord")

	var crmBusinessOpportunityApprovalRecordApi = v1.ApiGroupApp.CrmApiGroup.CrmBusinessOpportunityApprovalRecordApi
	{
		crmBusinessOpportunityApprovalRecordRouter.POST("createCrmBusinessOpportunityApprovalRecord", crmBusinessOpportunityApprovalRecordApi.CreateCrmBusinessOpportunityApprovalRecord)   // 新建商机审批记录
		crmBusinessOpportunityApprovalRecordRouter.DELETE("deleteCrmBusinessOpportunityApprovalRecord", crmBusinessOpportunityApprovalRecordApi.DeleteCrmBusinessOpportunityApprovalRecord) // 删除商机审批记录
		crmBusinessOpportunityApprovalRecordRouter.DELETE("deleteCrmBusinessOpportunityApprovalRecordByIds", crmBusinessOpportunityApprovalRecordApi.DeleteCrmBusinessOpportunityApprovalRecordByIds) // 批量删除商机审批记录
		crmBusinessOpportunityApprovalRecordRouter.PUT("updateCrmBusinessOpportunityApprovalRecord", crmBusinessOpportunityApprovalRecordApi.UpdateCrmBusinessOpportunityApprovalRecord)    // 更新商机审批记录
	}
	{
		crmBusinessOpportunityApprovalRecordRouterWithoutRecord.GET("findCrmBusinessOpportunityApprovalRecord", crmBusinessOpportunityApprovalRecordApi.FindCrmBusinessOpportunityApprovalRecord)        // 根据ID获取商机审批记录
		crmBusinessOpportunityApprovalRecordRouterWithoutRecord.GET("getCrmBusinessOpportunityApprovalRecordList", crmBusinessOpportunityApprovalRecordApi.GetCrmBusinessOpportunityApprovalRecordList)  // 获取商机审批记录列表
	}
	{
	    crmBusinessOpportunityApprovalRecordRouterWithoutAuth.GET("getCrmBusinessOpportunityApprovalRecordPublic", crmBusinessOpportunityApprovalRecordApi.GetCrmBusinessOpportunityApprovalRecordPublic)  // 获取商机审批记录列表
	}
}
