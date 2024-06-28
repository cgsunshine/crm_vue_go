package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},
		system.SysExportTemplate{},
		system.Condition{},
		system.JoinTemplate{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{}, crm.CrmTest{}, crm.CrmContractType{}, crm.CrmLoginLog{}, crm.CrmOperationRecords{}, crm.CrmProductGroup{}, crm.CrmProductType{}, crm.CrmUser{}, crm.CrmCustomerGroup{}, crm.CrmCustomers{}, crm.CrmProduct{}, crm.CrmBill{}, crm.CrmContract{}, crm.CrmOrder{}, crm.CrmPayment{}, crm.CrmStatementAccount{}, crm.CrmSupplier{}, crm.CrmCommissionRebate{}, crm.CrmTicketCategories{}, crm.CrmTicketResponseTemplates{}, crm.CrmTickets{}, crm.CrmTicketComments{}, crm.CrmStatementAccountFile{}, crm.CrmContactFile{}, crm.CrmContactApprovalTasks{}, crm.CrmContactApprovalRecord{}, crm.CrmPurchaseOrder{}, crm.CrmBusinessOpportunity{}, crm.CrmBusinessOpportunityFile{}, crm.CrmAuthorities{}, crm.CrmPaymentCollention{}, crm.CrmProcurementContract{}, crm.CrmApprovalProcess{}, crm.CrmConfig{}, crm.CrmApprovalNode{}, crm.CrmBusinessOpportunityApprovalRecord{}, crm.CrmBusinessOpportunityApprovalTasks{}, crm.CrmPaymentCollentionApprovalRecord{}, crm.CrmPaymentCollentionApprovalTasks{}, crm.CrmApprovalRecord{}, crm.CrmApprovalTasks{}, crm.CrmCurrency{}, crm.CrmOrderProduct{}, crm.CrmBusinessOpportunityProduct{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
