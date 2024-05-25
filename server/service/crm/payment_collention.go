package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmPaymentCollentionInfoList 分页获取crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) GetCrmPagePaymentCollentionInfoList(info crmReq.CrmPaymentCollentionSearch) (list []crm.CrmPagePaymentCollention, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmPaymentCollention{})
	var crmPaymentCollentions []crm.CrmPagePaymentCollention
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.BillId != nil {
		db = db.Where("bill_id = ?", info.BillId)
	}
	if info.CustomerId != nil {
		db = db.Where("customer_id = ?", info.CustomerId)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
	}
	if info.Proof != "" {
		db = db.Where("proof = ?", info.Proof)
	}
	if info.AuditingStatus != "" {
		db = db.Where("auditing_status = ?", info.AuditingStatus)
	}
	if info.ApprovedById != "" {
		db = db.Where("approved_by_id = ?", info.ApprovedById)
	}
	if info.StartAuditingTime != nil && info.EndAuditingTime != nil {
		db = db.Where("auditing_time BETWEEN ? AND ? ", info.StartAuditingTime, info.EndAuditingTime)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_payment_collention.*,crm_customers.customer_name,sys_users.username,crm_bill.amount").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_payment_collention.user_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_payment_collention.customer_id").
		Joins("LEFT JOIN crm_bill ON crm_bill.id = crm_payment_collention.bill_id").
		Find(&crmPaymentCollentions).Error
	return crmPaymentCollentions, total, err
}
