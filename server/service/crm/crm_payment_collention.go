package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"time"
)

type CrmPaymentCollentionService struct {
}

// CreateCrmPaymentCollention 创建crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) CreateCrmPaymentCollention(crmPaymentCollention *crm.CrmPaymentCollention) (err error) {
	err = global.GVA_DB.Create(crmPaymentCollention).Error
	return err
}

// DeleteCrmPaymentCollention 删除crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) DeleteCrmPaymentCollention(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmPaymentCollention{}, "id = ?", ID).Error
	return err
}

// DeleteCrmPaymentCollentionByIds 批量删除crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) DeleteCrmPaymentCollentionByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmPaymentCollention{}, "id in ?", IDs).Error
	return err
}

// UpdateCrmPaymentCollention 更新crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) UpdateCrmPaymentCollention(crmPaymentCollention crm.CrmPaymentCollention) (err error) {
	err = global.GVA_DB.Save(&crmPaymentCollention).Error
	return err
}

// GetCrmPaymentCollention 根据ID获取crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) GetCrmPaymentCollention(ID string) (crmPaymentCollention crm.CrmPaymentCollention, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmPaymentCollention).Error
	return
}

// GetCrmPaymentCollention 根据ID获取crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) GetCrmPaymentCollentionTodayCount() (count int64) {
	global.GVA_DB.Model(&crm.CrmPaymentCollention{}).Where("created_at >= ? ", time.Now().Format("2006-01-02")).Count(&count)
	return
}

// GetCrmPaymentCollentionInfoList 分页获取crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) GetCrmPaymentCollentionInfoList(info crmReq.CrmPaymentCollentionSearch) (list []crm.CrmPaymentCollention, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmPaymentCollention{})
	var crmPaymentCollentions []crm.CrmPaymentCollention
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

	err = db.Find(&crmPaymentCollentions).Error
	return crmPaymentCollentions, total, err
}

// ApprovalTasksCount 统计有效的，并且状态是待审批的数量
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) ApprovalTasksCount(userId *int, approvalStatus string, startDate, endDate *time.Time) (total int64, err error) {
	db := global.GVA_DB.Model(&crm.CrmPaymentCollention{})
	SearchCondition(db, userId, startDate, endDate)
	err = db.Where("review_status = ? ", approvalStatus).Debug().Count(&total).Error
	return
}
