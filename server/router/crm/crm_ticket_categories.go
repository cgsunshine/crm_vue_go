package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CrmTicketCategoriesRouter struct {
}

// InitCrmTicketCategoriesRouter 初始化 工单类别 路由信息
func (s *CrmTicketCategoriesRouter) InitCrmTicketCategoriesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	crmTicketCategoriesRouter := Router.Group("crmTicketCategories").Use(middleware.OperationRecord())
	crmTicketCategoriesRouterWithoutRecord := Router.Group("crmTicketCategories")
	crmTicketCategoriesRouterWithoutAuth := PublicRouter.Group("crmTicketCategories")

	var crmTicketCategoriesApi = v1.ApiGroupApp.CrmApiGroup.CrmTicketCategoriesApi
	{
		crmTicketCategoriesRouter.POST("createCrmTicketCategories", crmTicketCategoriesApi.CreateCrmTicketCategories)   // 新建工单类别
		crmTicketCategoriesRouter.DELETE("deleteCrmTicketCategories", crmTicketCategoriesApi.DeleteCrmTicketCategories) // 删除工单类别
		crmTicketCategoriesRouter.DELETE("deleteCrmTicketCategoriesByIds", crmTicketCategoriesApi.DeleteCrmTicketCategoriesByIds) // 批量删除工单类别
		crmTicketCategoriesRouter.PUT("updateCrmTicketCategories", crmTicketCategoriesApi.UpdateCrmTicketCategories)    // 更新工单类别
	}
	{
		crmTicketCategoriesRouterWithoutRecord.GET("findCrmTicketCategories", crmTicketCategoriesApi.FindCrmTicketCategories)        // 根据ID获取工单类别
		crmTicketCategoriesRouterWithoutRecord.GET("getCrmTicketCategoriesList", crmTicketCategoriesApi.GetCrmTicketCategoriesList)  // 获取工单类别列表
	}
	{
	    crmTicketCategoriesRouterWithoutAuth.GET("getCrmTicketCategoriesPublic", crmTicketCategoriesApi.GetCrmTicketCategoriesPublic)  // 获取工单类别列表
	}
}
