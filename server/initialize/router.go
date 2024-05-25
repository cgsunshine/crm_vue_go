package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/flipped-aurora/gin-vue-admin/server/routerv2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	InstallPlugin(Router)
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.GVA_CONFIG.Local.StorePath)})

	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{

		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitInitRouter(PublicGroup)
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitSystemRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitAutoCodeRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
		systemRouter.InitSysDictionaryRouter(PrivateGroup)
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)
		systemRouter.InitSysExportTemplateRouter(PrivateGroup)
		exampleRouter.InitCustomerRouter(PrivateGroup)
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)
	}

	{
		crmRouter := router.RouterGroupApp.Crm
		crmRouter.InitCrmTestRouter(PrivateGroup, PublicGroup)

		crmRouter.InitCrmContractTypeRouter(PrivateGroup, PublicGroup)

		crmRouter.InitCrmLoginLogRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmOperationRecordsRouter(PrivateGroup, PublicGroup)

		crmRouter.InitCrmProcurementContractRouter(PrivateGroup, PublicGroup)

		crmRouter.InitCrmProductGroupRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmProductTypeRouter(PrivateGroup, PublicGroup)

		crmRouter.InitCrmUserRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmApprovalNodeRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmApprovalProcessRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmApprovalRecordRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmApprovalTasksRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmTicketCommentsRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmTicketResponseTemplatesRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmTicketsRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmContactFileUploadAndDownloadsRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmCustomerGroupRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmCustomersRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmProductRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmBusinessOpportunityRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmBillRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmContractRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmOrderRouter(PrivateGroup, PublicGroup)

		crmRouter.InitCrmPageCustomersRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmPageProductRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmPageBusinessOpportunityRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmPaymentRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmPurchaseOrderRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmStatementAccountRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmSupplierRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmCommissionRebateRouter(PrivateGroup, PublicGroup)

		crmRouter.InitCrmPageCommissionRebateRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmPagePaymentRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmPagePurchaseOrderRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmPageStatementAccountRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmPageSupplierRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmPageBillRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmPageContractRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmPagOrderRouter(PrivateGroup, PublicGroup)

		crmRouter.InitCrmPagePaymentCollentionRouter(PrivateGroup, PublicGroup)
		crmRouter.InitCrmPaymentCollentionRouter(PrivateGroup, PublicGroup)
	}

	{
		systemRouterV2 := routerv2.RouterGroupApp.System
		PrivateGroupV2 := Router.Group("v2")
		systemRouterV2.InitBaseRouter(PrivateGroupV2)

		PrivateGroupV2.Use(middleware.JWTAuth())
		systemRouterV2.InitMenuRouter(PrivateGroupV2)
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
