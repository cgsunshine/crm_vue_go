package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmApprovalRecordRouter struct {
}

// InitCrmApprovalRecordRouter 初始化 crmApprovalRecord表 路由信息
func (s *CrmApprovalRecordRouter) InitCrmApprovalRecordRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmApprovalRecordRouter := Router.Group("crmApprovalRecord").Use(middleware.OperationRecord())
	crmApprovalRecordRouterWithoutRecord := Router.Group("crmApprovalRecord")
	crmApprovalRecordRouterWithoutAuth := PublicRouter.Group("crmApprovalRecord")

	var crmApprovalRecordApi = v1.ApiGroupApp.CrmApiGroup.CrmApprovalRecordApi
	{
		crmApprovalRecordRouter.POST("createCrmApprovalRecord", crmApprovalRecordApi.CreateCrmApprovalRecord)   // 新建crmApprovalRecord表
		crmApprovalRecordRouter.DELETE("deleteCrmApprovalRecord", crmApprovalRecordApi.DeleteCrmApprovalRecord) // 删除crmApprovalRecord表
		crmApprovalRecordRouter.DELETE("deleteCrmApprovalRecordByIds", crmApprovalRecordApi.DeleteCrmApprovalRecordByIds) // 批量删除crmApprovalRecord表
		crmApprovalRecordRouter.PUT("updateCrmApprovalRecord", crmApprovalRecordApi.UpdateCrmApprovalRecord)    // 更新crmApprovalRecord表
	}
	{
		crmApprovalRecordRouterWithoutRecord.GET("findCrmApprovalRecord", crmApprovalRecordApi.FindCrmApprovalRecord)        // 根据ID获取crmApprovalRecord表
		crmApprovalRecordRouterWithoutRecord.GET("getCrmApprovalRecordList", crmApprovalRecordApi.GetCrmApprovalRecordList)  // 获取crmApprovalRecord表列表
	}
	{
	    crmApprovalRecordRouterWithoutAuth.GET("getCrmApprovalRecordPublic", crmApprovalRecordApi.GetCrmApprovalRecordPublic)  // 获取crmApprovalRecord表列表
	}
}
