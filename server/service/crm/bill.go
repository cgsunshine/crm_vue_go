package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmBillInfoList 分页获取crmBill表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillService *CrmBillService) GetCrmPageBillInfoList(info crmReq.CrmBillSearch) (list []crm.CrmPageBill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmBill{})
	var crmBills []crm.CrmPageBill
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Amount != nil {
		db = db.Where("amount = ?", info.Amount)
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
	}
	if info.CustomerId != nil {
		db = db.Where("customer_id = ?", info.CustomerId)
	}
	if info.StartExpirationTime != nil && info.EndExpirationTime != nil {
		db = db.Where("expiration_time BETWEEN ? AND ? ", info.StartExpirationTime, info.EndExpirationTime)
	}
	if info.OrderId != nil {
		db = db.Where("order_id = ?", info.OrderId)
	}
	if info.PaymentCollention != nil {
		db = db.Where("payment_collention = ?", info.PaymentCollention)
	}
	if info.PaymentStatus != "" {
		db = db.Where("payment_status = ?", info.PaymentStatus)
	}
	if info.StartPaymentTime != nil && info.EndPaymentTime != nil {
		db = db.Where("payment_time BETWEEN ? AND ? ", info.StartPaymentTime, info.EndPaymentTime)
	}
	if info.PaymentType != "" {
		db = db.Where("payment_type = ?", info.PaymentType)
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

	err = db.Select("crm_bill.*,crm_customers.customer_name,sys_users.username").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_bill.user_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_bill.customer_id").
		Find(&crmBills).Error
	return crmBills, total, err
}
