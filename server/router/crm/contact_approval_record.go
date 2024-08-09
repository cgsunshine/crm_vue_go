package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CrmPageContactApprovalRecordRouter struct {
}

// InitCrmContactApprovalRecordRouter 初始化 合同审批记录 路由信息
func (s *CrmPageContactApprovalRecordRouter) InitCrmPageContactApprovalRecordRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {

	crmContactApprovalRecordRouterWithoutAuth := PublicRouter.Group("crmContactApprovalRecord")

	var crmContactApprovalRecordApi = v1.ApiGroupApp.CrmApiGroup.CrmContactApprovalRecordApi

	{
		crmContactApprovalRecordRouterWithoutAuth.GET("getCrmPageContactApprovalRecordList", crmContactApprovalRecordApi.GetCrmPageContactApprovalRecordList) // 获取合同审批记录列表
	}
}
