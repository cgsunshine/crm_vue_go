package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmPaymentInfoList 分页获取crmPayment表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) GetCrmPagePaymentInfoList(info crmReq.CrmPaymentSearch) (list []crm.CrmPagePayment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmPayment{})
	var crmPayments []crm.CrmPagePayment
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmPaymentService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StatementAccountId != nil {
		db = db.Where(crmPaymentService.SplicingQueryConditions("statement_account_id = ?"), info.StatementAccountId)
	}
	if info.PaymentAmount != nil {
		db = db.Where(crmPaymentService.SplicingQueryConditions("payment_amount = ?"), info.PaymentAmount)
	}
	if info.StartPaymentTime != nil && info.EndPaymentTime != nil {
		db = db.Where(crmPaymentService.SplicingQueryConditions("payment_time BETWEEN ? AND ? "), info.StartPaymentTime, info.EndPaymentTime)
	}
	if info.PaymentVoucher != "" {
		db = db.Where(crmPaymentService.SplicingQueryConditions("payment_voucher = ?"), info.PaymentVoucher)
	}
	if info.UserId != nil {
		db = db.Where(crmPaymentService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}

	if info.PaymentNumber != "" {
		db = db.Where(crmPaymentService.SplicingQueryConditions("payment_number = ?"), info.PaymentNumber)
	}

	//func UniversalSearchBillPaymentNumber(db *gorm.DB, value string) {
	//	if value != "" {
	//		db = db.Where("crm_bill_payment.bill_payment_number = ?", value)
	//	}
	//}

	//if info.BillPaymentNumber != "" {
	//	db = db.Where("crm_bill_payment.bill_payment_number = ?", info.BillPaymentNumber)
	//}

	UniversalSearchBillPaymentNumber(db, info.BillPaymentNumber)

	db.Select("crm_payment.*,sys_users.username,crm_bill_payment.bill_payment_number").
		Joins("LEFT JOIN sys_users ON crm_payment.user_id = sys_users.id").
		Joins("LEFT JOIN crm_bill_payment ON crm_bill_payment.id = crm_payment.bill_payment_id").
		Order("crm_payment.created_at DESC")

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&crmPayments).Error
	return crmPayments, total, err
}

// GetCrmPayment 根据ID获取crmPayment表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) GetCrmPagePayment(ID string) (crmPayment crm.CrmPagePayment, err error) {
	err = global.GVA_DB.Model(&crm.CrmPayment{}).Where("crm_payment.id = ?", ID).
		Select("crm_payment.*,sys_users.username,crm_bill_payment.bill_payment_number").
		Joins("LEFT JOIN sys_users ON crm_payment.user_id = sys_users.id").
		Joins("LEFT JOIN crm_bill_payment ON crm_bill_payment.id = crm_payment.bill_payment_id").
		First(&crmPayment).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmPaymentService *CrmPaymentService) SplicingQueryConditions(condition string) string {
	return "crm_payment." + condition
}

// UpdApprovalStatus 修改审批状态
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) UpdApprovalStatus(ID *int, data map[string]interface{}) (err error) {
	db := global.GVA_DB.Model(&crm.CrmPayment{})
	err = db.Where("id = ?", ID).Updates(data).Error
	return
}
