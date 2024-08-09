package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmOperationRecordsRouter struct {
}

// InitCrmOperationRecordsRouter 初始化 操作记录 路由信息
func (s *CrmOperationRecordsRouter) InitCrmOperationRecordsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmOperationRecordsRouter := Router.Group("crmOperationRecords").Use(middleware.OperationRecord())
	crmOperationRecordsRouterWithoutRecord := Router.Group("crmOperationRecords")
	crmOperationRecordsRouterWithoutAuth := PublicRouter.Group("crmOperationRecords")

	var crmOperationRecordsApi = v1.ApiGroupApp.CrmApiGroup.CrmOperationRecordsApi
	{
		crmOperationRecordsRouter.POST("createCrmOperationRecords", crmOperationRecordsApi.CreateCrmOperationRecords)   // 新建操作记录
		crmOperationRecordsRouter.DELETE("deleteCrmOperationRecords", crmOperationRecordsApi.DeleteCrmOperationRecords) // 删除操作记录
		crmOperationRecordsRouter.DELETE("deleteCrmOperationRecordsByIds", crmOperationRecordsApi.DeleteCrmOperationRecordsByIds) // 批量删除操作记录
		crmOperationRecordsRouter.PUT("updateCrmOperationRecords", crmOperationRecordsApi.UpdateCrmOperationRecords)    // 更新操作记录
	}
	{
		crmOperationRecordsRouterWithoutRecord.GET("findCrmOperationRecords", crmOperationRecordsApi.FindCrmOperationRecords)        // 根据ID获取操作记录
		crmOperationRecordsRouterWithoutRecord.GET("getCrmOperationRecordsList", crmOperationRecordsApi.GetCrmOperationRecordsList)  // 获取操作记录列表
	}
	{
	    crmOperationRecordsRouterWithoutAuth.GET("getCrmOperationRecordsPublic", crmOperationRecordsApi.GetCrmOperationRecordsPublic)  // 获取操作记录列表
	}
}
