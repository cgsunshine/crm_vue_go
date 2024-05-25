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
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.OrderId != nil {
		db = db.Where("order_id = ?", info.OrderId)
	}
	if info.PaymentAmount != nil {
		db = db.Where("payment_amount = ?", info.PaymentAmount)
	}
	if info.StartPaymentTime != nil && info.EndPaymentTime != nil {
		db = db.Where("payment_time BETWEEN ? AND ? ", info.StartPaymentTime, info.EndPaymentTime)
	}
	if info.PaymentVoucher != "" {
		db = db.Where("payment_voucher = ?", info.PaymentVoucher)
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

	err = db.Select("crm_payment.*,sys_users.username").
		Joins("LEFT JOIN sys_users ON crm_payment.user_id = sys_users.id").Find(&crmPayments).Error
	return crmPayments, total, err
}
