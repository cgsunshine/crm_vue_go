package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"time"
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
		db = db.Where(crmPaymentCollentionService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.BillId != nil {
		db = db.Where(crmPaymentCollentionService.SplicingQueryConditions("bill_id = ?"), info.BillId)
	}
	if info.CustomerId != nil {
		db = db.Where(crmPaymentCollentionService.SplicingQueryConditions("customer_id = ?"), info.CustomerId)
	}
	if info.UserId != nil {
		db = db.Where(crmPaymentCollentionService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}
	if info.Currency != "" {
		db = db.Where(crmPaymentCollentionService.SplicingQueryConditions("currency = ?"), info.Currency)
	}

	if info.StartAuditingTime != nil && info.EndAuditingTime != nil {
		db = db.Where(crmPaymentCollentionService.SplicingQueryConditions("auditing_time BETWEEN ? AND ? "), info.StartAuditingTime, info.EndAuditingTime)
	}

	if info.PaymentCollentionName != "" {
		db = db.Where(crmPaymentCollentionService.SplicingQueryConditions("payment_collention_name LIKE ?"), "%"+info.PaymentCollentionName+"%")
	}

	if info.ReviewStatus != "" {
		db = db.Where(crmPaymentCollentionService.SplicingQueryConditions("review_status = ?"), info.ReviewStatus)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.
		Select("crm_payment_collention.*,crm_customers.customer_name,sys_users.username,crm_business_opportunity.business_opportunity_name,crm_order.price,crm_order.order_name,crm_order.customer_id,crm_bill.bill_name").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_payment_collention.user_id").
		Joins("LEFT JOIN crm_business_opportunity ON crm_business_opportunity.id = crm_payment_collention.business_opportunity_id").
		Joins("LEFT JOIN crm_order ON crm_order.id = crm_payment_collention.order_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_order.customer_id").
		Joins("LEFT JOIN crm_bill ON crm_bill.id = crm_payment_collention.bill_id").
		Order("crm_payment_collention.created_at DESC").
		Find(&crmPaymentCollentions).Error
	return crmPaymentCollentions, total, err
}

// GetCrmPaymentCollention 根据ID获取crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) GetCrmPagePaymentCollention(ID string) (crmPaymentCollention crm.CrmPagePaymentCollention, err error) {
	err = global.GVA_DB.Model(&crm.CrmPaymentCollention{}).Where("crm_payment_collention.id = ?", ID).
		Select("crm_payment_collention.*,crm_customers.customer_name,sys_users.username,crm_business_opportunity.business_opportunity_name,crm_order.price,crm_order.order_name,crm_order.customer_id,crm_bill.bill_name").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_payment_collention.user_id").
		Joins("LEFT JOIN crm_business_opportunity ON crm_business_opportunity.id = crm_payment_collention.business_opportunity_id").
		Joins("LEFT JOIN crm_order ON crm_order.id = crm_payment_collention.order_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_order.customer_id").
		Joins("LEFT JOIN crm_bill ON crm_bill.id = crm_payment_collention.bill_id").
		First(&crmPaymentCollention).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmPaymentCollentionService *CrmPaymentCollentionService) SplicingQueryConditions(condition string) string {
	return "crm_payment_collention." + condition
}

// UpdApprovalStatus 修改审批状态
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) UpdApprovalStatus(ID *int, data map[string]interface{}) (err error) {
	db := global.GVA_DB.Model(&crm.CrmPaymentCollention{})
	err = db.Where("id = ?", ID).Updates(data).Error
	return
}

// GetCrmPaymentCollention 根据ID获取crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) GetCrmPaymentIdCollention(ID int) (crmPaymentCollention crm.CrmPaymentCollention, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmPaymentCollention).Error
	return
}

// 周期内回款金额总计
func (crmPaymentCollentionService *CrmPaymentCollentionService) PaymentCollentionAmountCycleTime(userId *int, startDate, endDate *time.Time) (total int64, err error) {
	db := global.GVA_DB.Model(&crm.CrmContract{})
	SearchCondition(db, userId, startDate, endDate)
	db.Select("SUM(amount)")
	db.Where("review_status = ?", comm.Approval_Status_Pass)
	err = db.Count(&total).Error
	return
}
