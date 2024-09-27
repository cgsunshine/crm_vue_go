package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"time"
)

type CrmPaymentService struct {
}

// CreateCrmPayment 创建crmPayment表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) CreateCrmPayment(crmPayment *crm.CrmPayment) (err error) {
	err = global.GVA_DB.Create(crmPayment).Error
	return err
}

// GetCrmStatementAccountBillPaymentCount 根据ID获取crmStatementAccount表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) GetCrmStatementAccountBillPaymentCount(billPaymentID *int) (count int64, err error) {
	err = global.GVA_DB.Model(&crm.CrmPayment{}).Where("bill_payment_id = ?", billPaymentID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return
}

// DeleteCrmPayment 删除crmPayment表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) DeleteCrmPayment(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmPayment{}, "id = ?", ID).Error
	return err
}

// DeleteCrmPaymentByIds 批量删除crmPayment表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) DeleteCrmPaymentByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmPayment{}, "id in ?", IDs).Error
	return err
}

// UpdateCrmPayment 更新crmPayment表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) UpdateCrmPayment(crmPayment crm.CrmPayment) (err error) {
	err = global.GVA_DB.Save(&crmPayment).Error
	return err
}

// GetCrmPayment 根据ID获取crmPayment表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) GetCrmPayment(ID string) (crmPayment crm.CrmPayment, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmPayment).Error
	return
}

// GetCrmPaymentTodayCount 根据ID获取crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) GetCrmPaymentTodayCount() (count int64) {
	global.GVA_DB.Model(&crm.CrmPayment{}).Where("created_at >= ? ", time.Now().Format("2006-01-02")).Count(&count)
	return
}

// GetCrmPayment 根据ID获取crmPayment表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) GetCrmPaymentIdInfo(ID *int) (crmPayment crm.CrmPayment, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmPayment).Error
	return
}

// GetCrmPaymentInfoList 分页获取crmPayment表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) GetCrmPaymentInfoList(info crmReq.CrmPaymentSearch) (list []crm.CrmPayment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmPayment{})
	var crmPayments []crm.CrmPayment
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	//if info.OrderId != nil {
	//    db = db.Where("order_id = ?",info.OrderId)
	//}
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

	err = db.Find(&crmPayments).Error
	return crmPayments, total, err
}

// ApprovalTasksCount 统计有效的，并且状态是待审批的数量
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) ApprovalTasksCount(userId *int, approvalStatus string, startDate, endDate *time.Time) (total int64, err error) {
	db := global.GVA_DB.Model(&crm.CrmPayment{})
	SearchCondition(db, userId, startDate, endDate)
	err = db.Where("review_status = ? ", approvalStatus).Debug().Count(&total).Error
	return
}
