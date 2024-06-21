package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmCommissionRebateInfoList 分页获取crmCommissionRebate表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService) GetCrmPageCommissionRebateInfoList(info crmReq.CrmCommissionRebateSearch) (list []crm.CrmPageCommissionRebate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmCommissionRebate{})
	var crmCommissionRebates []crm.CrmPageCommissionRebate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmCommissionRebateService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.OrderId != nil {
		db = db.Where(crmCommissionRebateService.SplicingQueryConditions("order_id = ?"), info.OrderId)
	}
	if info.UserId != nil {
		db = db.Where(crmCommissionRebateService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}
	if info.Payee != "" {
		db = db.Where(crmCommissionRebateService.SplicingQueryConditions("payee LIKE ?"), "%"+info.Payee+"%")
	}
	if info.PaymentMethod != "" {
		db = db.Where(crmCommissionRebateService.SplicingQueryConditions("payment_method = ?"), info.PaymentMethod)
	}
	if info.Account != "" {
		db = db.Where(crmCommissionRebateService.SplicingQueryConditions("account = ?"), info.Account)
	}
	if info.Amount != nil {
		db = db.Where(crmCommissionRebateService.SplicingQueryConditions("amount = ?"), info.Amount)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_commission_rebate.*,sys_users.username,crm_order.order_name").
		Joins("LEFT JOIN sys_users ON crm_commission_rebate.user_id = sys_users.id").
		Joins("LEFT JOIN crm_order ON crm_order.id = crm_commission_rebate.order_id").
		Find(&crmCommissionRebates).Error
	return crmCommissionRebates, total, err
}

// GetCrmCommissionRebate 根据ID获取crmCommissionRebate表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService) GetCrmPageCommissionRebate(ID string) (crmCommissionRebate crm.CrmPageCommissionRebate, err error) {
	err = global.GVA_DB.Model(&crm.CrmCommissionRebate{}).Where("crm_commission_rebate.id = ?", ID).
		Select("crm_commission_rebate.*,sys_users.username,crm_order.order_name").
		Joins("LEFT JOIN sys_users ON crm_commission_rebate.user_id = sys_users.id").
		Joins("LEFT JOIN crm_order ON crm_order.id = crm_commission_rebate.order_id").
		First(&crmCommissionRebate).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmCommissionRebateService *CrmCommissionRebateService) SplicingQueryConditions(condition string) string {
	return "crm_commission_rebate." + condition
}
