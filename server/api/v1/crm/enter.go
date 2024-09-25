package crm

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CrmTestApi
	CrmContractTypeApi
	CrmLoginLogApi
	CrmOperationRecordsApi
	CrmProductGroupApi
	CrmProductTypeApi
	CrmUserApi
	CrmCustomerGroupApi
	CrmCustomersApi
	CrmProductApi
	CrmBillApi
	CrmContractApi
	CrmOrderApi
	CrmPaymentApi
	CrmStatementAccountApi
	CrmSupplierApi
	CrmCommissionRebateApi
	CrmTicketCategoriesApi
	CrmTicketResponseTemplatesApi
	CrmTicketsApi
	CrmTicketCommentsApi
	CrmStatementAccountFileApi
	CrmContactFileApi
	CrmContactApprovalTasksApi
	CrmContactApprovalRecordApi
	CrmPurchaseOrderApi
	CrmBusinessOpportunityApi
	CrmBusinessOpportunityFileApi
	CrmAuthoritiesApi
	CrmPaymentCollentionApi
	CrmProcurementContractApi
	CrmApprovalProcessApi
	CrmConfigApi
	CrmApprovalNodeApi
	CrmBusinessOpportunityApprovalRecordApi
	CrmBusinessOpportunityApprovalTasksApi
	CrmPaymentCollentionApprovalRecordApi
	CrmPaymentCollentionApprovalTasksApi
	CrmApprovalRecordApi
	CrmApprovalTasksApi
	CrmCurrencyApi
	CrmOrderProductApi
	CrmBusinessOpportunityProductApi
	CrmDepositsApi
	CrmRefundTasksApi
	CrmBillPaymentApi
	AdminHome
	CrmPurchaseOrderProductApi
}

var (
	userService			= service.ServiceGroupApp.SystemServiceGroup.UserService
	customerService			= service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService	= service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
