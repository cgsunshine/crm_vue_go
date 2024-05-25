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
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PurchaseOrderId != nil {
		db = db.Where("purchase_order_id = ?", info.PurchaseOrderId)
	}
	if info.Amount != nil {
		db = db.Where("amount = ?", info.Amount)
	}
	if info.StartCreationTime != nil && info.EndCreationTime != nil {
		db = db.Where("creation_time BETWEEN ? AND ? ", info.StartCreationTime, info.EndCreationTime)
	}
	if info.StartPayableTime != nil && info.EndPayableTime != nil {
		db = db.Where("payable_time BETWEEN ? AND ? ", info.StartPayableTime, info.EndPayableTime)
	}
	if info.PaymentStatus != "" {
		db = db.Where("payment_status = ?", info.PaymentStatus)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_statement_account.*,sys_users.username").
		Joins("LEFT JOIN sys_users ON crm_statement_account.user_id = sys_users.id").
		Find(&crmStatementAccounts).Error
	return crmStatementAccounts, total, err
}
