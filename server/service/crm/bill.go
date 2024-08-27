package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmBill 根据ID获取crmBill表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillService *CrmBillService) GetCrmPageBill(ID string) (crmBill crm.CrmPageBill, err error) {
	err = global.GVA_DB.Model(&crm.CrmBill{}).
		Where("crm_bill.id = ?", ID).
		Select("crm_bill.*,crm_customers.customer_name,sys_users.username").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_bill.user_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_bill.customer_id").
		First(&crmBill).Error
	return
}

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
		db = db.Where(crmBillService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.StartExpirationTime != nil && info.EndExpirationTime != nil {
		db = db.Where(crmBillService.SplicingQueryConditions("expiration_time BETWEEN ? AND ?"), info.StartExpirationTime, info.EndExpirationTime)
	}
	if info.OrderId != nil {
		db = db.Where(crmBillService.SplicingQueryConditions("order_id = ?"), info.OrderId)
	}
	if info.PaymentCollentionId != nil {
		db = db.Where(crmBillService.SplicingQueryConditions("payment_collention_id = ?"), info.PaymentCollentionId)
	}
	if info.PaymentId != nil {
		db = db.Where(crmBillService.SplicingQueryConditions("payment_id = ?"), info.PaymentId)
	}
	if info.UserId != nil {
		db = db.Where(crmBillService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}

	if info.PaymentStatus != "" {
		db = db.Where(crmBillService.SplicingQueryConditions("payment_status = ?"), info.PaymentStatus)
	}
	if info.CustomerId != nil {
		db = db.Where(crmBillService.SplicingQueryConditions("customer_id = ?"), info.CustomerId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_bill.*," +
		"crm_order.order_name,crm_order.customer_id," +
		"crm_customers.customer_name," +
		"sys_users.username," +
		"crm_payment_collention.payment_collention_name,crm_payment_collention.payment_time").
		Joins("LEFT JOIN crm_order ON crm_order.id = crm_bill.order_id").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_order.user_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_bill.customer_id").
		Joins("LEFT JOIN crm_payment_collention ON crm_payment_collention.bill_id = crm_bill.id").
		Order("crm_bill.created_at DESC").
		Find(&crmBills).Error
	return crmBills, total, err
}

// UpdApprovalStatus 修改关联id
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillService *CrmBillService) UpdAssOrderID(ID *int, data map[string]interface{}) (err error) {
	db := global.GVA_DB.Model(&crm.CrmBill{})
	err = db.Where("order_id = ?", ID).Updates(data).Error
	return
}

// UpdApprovalStatus 修改审批状态
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillService *CrmBillService) UpdApprovalStatus(ID *int, data map[string]interface{}) (err error) {
	db := global.GVA_DB.Model(&crm.CrmBill{})
	err = db.Where("id = ?", ID).Updates(data).Error
	return
}

// GetCrmBill 根据ID获取crmBill表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillService *CrmBillService) GetCrmPageIdBill(ID string) (crmBill crm.CrmPageBill, err error) {
	db := global.GVA_DB.Model(&crm.CrmBill{})
	err = db.Where("crm_bill.id = ?", ID).
		Select("crm_bill.*," +
			"crm_order.order_name,crm_order.customer_id," +
			"crm_customers.customer_name," +
			"sys_users.username," +
			"crm_payment_collention.payment_collention_name,crm_payment_collention.payment_time").
		Joins("LEFT JOIN crm_order ON crm_order.id = crm_bill.order_id").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_order.user_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_bill.customer_id").
		Joins("LEFT JOIN crm_payment_collention ON crm_payment_collention.id = crm_bill.payment_collention_id").
		Joins("LEFT JOIN crm_payment ON crm_payment.id = crm_bill.payment_id").
		First(&crmBill).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmBillService *CrmBillService) SplicingQueryConditions(condition string) string {
	return "crm_bill." + condition
}
