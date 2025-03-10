package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmStatementAccountInfoList 分页获取crmStatementAccount表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountService *CrmStatementAccountService) GetCrmPageStatementAccountInfoList(info crmReq.CrmStatementAccountSearch) (list []crm.CrmPageStatementAccount, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmStatementAccount{})
	var crmStatementAccounts []crm.CrmPageStatementAccount
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmStatementAccountService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PurchaseOrderId != nil {
		db = db.Where(crmStatementAccountService.SplicingQueryConditions("purchase_order_id = ?"), info.PurchaseOrderId)
	}
	if info.Amount != nil {
		db = db.Where(crmStatementAccountService.SplicingQueryConditions("amount = ?"), info.Amount)
	}
	if info.StartCreationTime != nil && info.EndCreationTime != nil {
		db = db.Where(crmStatementAccountService.SplicingQueryConditions("creation_time BETWEEN ? AND ? "), info.StartCreationTime, info.EndCreationTime)
	}
	if info.StartPayableTime != nil && info.EndPayableTime != nil {
		db = db.Where(crmStatementAccountService.SplicingQueryConditions("payable_time BETWEEN ? AND ? "), info.StartPayableTime, info.EndPayableTime)
	}
	if info.PaymentStatus != "" {
		db = db.Where(crmStatementAccountService.SplicingQueryConditions("payment_status = ?"), info.PaymentStatus)
	}
	if info.UserId != nil {
		db = db.Where(crmStatementAccountService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}
	if info.StatementAccountNumber != "" {
		db = db.Where(crmStatementAccountService.SplicingQueryConditions("statement_account_number = ?"), info.StatementAccountNumber)
	}

	UniversalSearchPurchaseOrderNumber(db, info.PurchaseOrderNumber)

	db.Select("crm_statement_account.*,sys_users.username,crm_purchase_order.purchase_order_name,crm_purchase_order.purchase_order_number").
		Joins("LEFT JOIN sys_users ON crm_statement_account.user_id = sys_users.id").
		Joins("LEFT JOIN crm_purchase_order ON crm_statement_account.purchase_order_id = crm_purchase_order.id").
		Order("crm_statement_account.created_at DESC")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Debug().Find(&crmStatementAccounts).Error
	return crmStatementAccounts, total, err
}

// GetCrmStatementAccount 根据ID获取crmStatementAccount表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountService *CrmStatementAccountService) GetCrmPageStatementAccount(ID string) (crmStatementAccount crm.CrmPageStatementAccount, err error) {
	err = global.GVA_DB.Model(&crm.CrmStatementAccount{}).Where("crm_statement_account.id = ?", ID).
		Select("crm_statement_account.*,sys_users.username,crm_purchase_order.purchase_order_name,crm_purchase_order.purchase_order_number").
		Joins("LEFT JOIN sys_users ON crm_statement_account.user_id = sys_users.id").
		Joins("LEFT JOIN crm_purchase_order ON crm_statement_account.purchase_order_id = crm_purchase_order.id").
		First(&crmStatementAccount).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmStatementAccountService *CrmStatementAccountService) SplicingQueryConditions(condition string) string {
	return "crm_statement_account." + condition
}
