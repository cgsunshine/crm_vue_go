package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
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
		example.ExaFileUploadAndDownload{}, crm.CrmTest{}, crm.CrmContractType{}, crm.CrmLoginLog{}, crm.CrmOperationRecords{}, crm.CrmProcurementContract{}, crm.CrmProductGroup{}, crm.CrmProductType{}, crm.CrmUser{}, crm.CrmApprovalNode{}, crm.CrmApprovalProcess{}, crm.CrmApprovalRecord{}, crm.CrmApprovalTasks{}, crm.CrmTicketComments{}, crm.CrmTicketResponseTemplates{}, crm.CrmTickets{}, crm.CrmContactFileUploadAndDownloads{}, crm.CrmCustomerGroup{}, crm.CrmCustomers{}, crm.CrmProduct{}, crm.CrmBusinessOpportunity{}, crm.CrmBill{}, crm.CrmContract{}, crm.CrmOrder{}, crm.CrmPaymentCollention{}, crm.CrmPayment{}, crm.CrmPurchaseOrder{}, crm.CrmStatementAccount{}, crm.CrmSupplier{}, crm.CrmCommissionRebate{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
