package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmPaymentCollentionApprovalRecordRouter struct {
}

// InitCrmPaymentCollentionApprovalRecordRouter 初始化 回款审批记录 路由信息
func (s *CrmPaymentCollentionApprovalRecordRouter) InitCrmPaymentCollentionApprovalRecordRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmPaymentCollentionApprovalRecordRouter := Router.Group("crmPaymentCollentionApprovalRecord").Use(middleware.OperationRecord())
	crmPaymentCollentionApprovalRecordRouterWithoutRecord := Router.Group("crmPaymentCollentionApprovalRecord")
	crmPaymentCollentionApprovalRecordRouterWithoutAuth := PublicRouter.Group("crmPaymentCollentionApprovalRecord")

	var crmPaymentCollentionApprovalRecordApi = v1.ApiGroupApp.CrmApiGroup.CrmPaymentCollentionApprovalRecordApi
	{
		crmPaymentCollentionApprovalRecordRouter.POST("createCrmPaymentCollentionApprovalRecord", crmPaymentCollentionApprovalRecordApi.CreateCrmPaymentCollentionApprovalRecord)   // 新建回款审批记录
		crmPaymentCollentionApprovalRecordRouter.DELETE("deleteCrmPaymentCollentionApprovalRecord", crmPaymentCollentionApprovalRecordApi.DeleteCrmPaymentCollentionApprovalRecord) // 删除回款审批记录
		crmPaymentCollentionApprovalRecordRouter.DELETE("deleteCrmPaymentCollentionApprovalRecordByIds", crmPaymentCollentionApprovalRecordApi.DeleteCrmPaymentCollentionApprovalRecordByIds) // 批量删除回款审批记录
		crmPaymentCollentionApprovalRecordRouter.PUT("updateCrmPaymentCollentionApprovalRecord", crmPaymentCollentionApprovalRecordApi.UpdateCrmPaymentCollentionApprovalRecord)    // 更新回款审批记录
	}
	{
		crmPaymentCollentionApprovalRecordRouterWithoutRecord.GET("findCrmPaymentCollentionApprovalRecord", crmPaymentCollentionApprovalRecordApi.FindCrmPaymentCollentionApprovalRecord)        // 根据ID获取回款审批记录
		crmPaymentCollentionApprovalRecordRouterWithoutRecord.GET("getCrmPaymentCollentionApprovalRecordList", crmPaymentCollentionApprovalRecordApi.GetCrmPaymentCollentionApprovalRecordList)  // 获取回款审批记录列表
	}
	{
	    crmPaymentCollentionApprovalRecordRouterWithoutAuth.GET("getCrmPaymentCollentionApprovalRecordPublic", crmPaymentCollentionApprovalRecordApi.GetCrmPaymentCollentionApprovalRecordPublic)  // 获取回款审批记录列表
	}
}
