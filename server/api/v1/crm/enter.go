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
	CrmApprovalNodeApi
	CrmApprovalProcessApi
	CrmApprovalRecordApi
	CrmApprovalTasksApi
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
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
