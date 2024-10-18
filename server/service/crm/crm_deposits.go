package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"time"
)

type CrmDepositsService struct {
}

// CreateCrmDeposits 创建押金管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmDepositsService *CrmDepositsService) CreateCrmDeposits(crmDeposits *crm.CrmDeposits) (err error) {
	err = global.GVA_DB.Create(crmDeposits).Error
	return err
}

// DeleteCrmDeposits 删除押金管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmDepositsService *CrmDepositsService) DeleteCrmDeposits(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmDeposits{}, "id = ?", ID).Error
	return err
}

// DeleteCrmDepositsByIds 批量删除押金管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmDepositsService *CrmDepositsService) DeleteCrmDepositsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmDeposits{}, "id in ?", IDs).Error
	return err
}

// GetCrmBillTodayCount 根据ID获取crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmDepositsService *CrmDepositsService) GetCrmBillTodayCount() (count int64) {
	global.GVA_DB.Model(&crm.CrmDeposits{}).Where("created_at >= ? ", time.Now().Format("2006-01-02")).Count(&count)
	return
}

// UpdateCrmDeposits 更新押金管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmDepositsService *CrmDepositsService) UpdateCrmDeposits(crmDeposits crm.CrmDeposits) (err error) {
	err = global.GVA_DB.Save(&crmDeposits).Error
	return err
}

// GetCrmDeposits 根据ID获取押金管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmDepositsService *CrmDepositsService) GetCrmDeposits(ID string) (crmDeposits crm.CrmPageDeposits, err error) {
	err = global.GVA_DB.Model(&crm.CrmDeposits{}).Where("crm_deposits.id = ?", ID).
		Select("crm_deposits.*,crm_customers.customer_name,crm_contract.contract_name,crm_contract.contract_number").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_deposits.customer_id").
		Joins("LEFT JOIN crm_contract ON crm_contract.id = crm_deposits.contract_id").
		First(&crmDeposits).Error
	return
}

// GetCrmDepositsInfoList 分页获取押金管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmDepositsService *CrmDepositsService) GetCrmDepositsInfoList(info crmReq.CrmDepositsSearch) (list []crm.CrmPageDeposits, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmDeposits{})
	var crmDepositss []crm.CrmPageDeposits
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmDepositsService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Amount != nil {
		db = db.Where(crmDepositsService.SplicingQueryConditions("amount = ?"), info.Amount)
	}
	if info.Currency != "" {
		db = db.Where(crmDepositsService.SplicingQueryConditions("currency = ?"), info.Currency)
	}
	if info.CustomerId != nil {
		db = db.Where(crmDepositsService.SplicingQueryConditions("customer_id = ?"), info.CustomerId)
	}
	if info.OrderId != nil {
		db = db.Where(crmDepositsService.SplicingQueryConditions("order_id = ?"), info.OrderId)
	}
	if info.StartPaymentDate != nil && info.EndPaymentDate != nil {
		db = db.Where(crmDepositsService.SplicingQueryConditions("payment_date BETWEEN ? AND ? "), info.StartPaymentDate, info.EndPaymentDate)
	}
	if info.StartRefundDate != nil && info.EndRefundDate != nil {
		db = db.Where(crmDepositsService.SplicingQueryConditions("refund_date BETWEEN ? AND ? "), info.StartRefundDate, info.EndRefundDate)
	}
	if info.UserId != nil {
		db = db.Where(crmDepositsService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}
	if info.ReviewStatus != "" {
		db = db.Where(crmDepositsService.SplicingQueryConditions("review_status = ?"), info.ReviewStatus)
	}
	if info.RefundStatus != "" {
		db = db.Where(crmDepositsService.SplicingQueryConditions("refund_status = ?"), info.RefundStatus)
	}
	if info.DepositsNumber != "" {
		db = db.Where(crmDepositsService.SplicingQueryConditions("deposits_number = ?"), info.DepositsNumber)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_deposits.*,crm_customers.customer_name,crm_contract.contract_name,crm_contract.contract_number").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_deposits.customer_id").
		Joins("LEFT JOIN crm_contract ON crm_contract.id = crm_deposits.contract_id").
		Order("crm_deposits.created_at DESC").
		Find(&crmDepositss).Error
	return crmDepositss, total, err
}

// UpdApprovalStatus 修改押金信息
// Author [piexlmax](https://github.com/piexlmax)
func (crmDepositsService *CrmDepositsService) UpdDepositsInfo(ID *int, data map[string]interface{}) (err error) {
	db := global.GVA_DB.Model(&crm.CrmDeposits{})
	err = db.Where("id = ?", ID).Updates(data).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmDepositsService *CrmDepositsService) SplicingQueryConditions(condition string) string {
	return "crm_deposits." + condition
}

// ApprovalTasksCount 统计有效的，并且状态是待审批的数量
// Author [piexlmax](https://github.com/piexlmax)
func (crmDepositsService *CrmDepositsService) ApprovalTasksCount(userId *int, approvalStatus string, startDate, endDate *time.Time) (total int64, err error) {
	db := global.GVA_DB.Model(&crm.CrmDeposits{})
	SearchCondition(db, userId, startDate, endDate)
	err = db.Where("review_status = ? ", approvalStatus).Debug().Count(&total).Error
	return
}
