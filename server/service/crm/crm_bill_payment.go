package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"time"
)

type CrmBillPaymentService struct {
}

// CreateCrmBillPayment 创建应付账单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillPaymentService *CrmBillPaymentService) CreateCrmBillPayment(crmBillPayment *crm.CrmBillPayment) (err error) {
	err = global.GVA_DB.Create(crmBillPayment).Error
	return err
}

// DeleteCrmBillPayment 删除应付账单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillPaymentService *CrmBillPaymentService) DeleteCrmBillPayment(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmBillPayment{}, "id = ?", ID).Error
	return err
}

// DeleteCrmBillPaymentByIds 批量删除应付账单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillPaymentService *CrmBillPaymentService) DeleteCrmBillPaymentByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmBillPayment{}, "id in ?", IDs).Error
	return err
}

// GetCrmBillPaymentTodayCount 根据ID获取crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillPaymentService *CrmBillPaymentService) GetCrmBillPaymentTodayCount() (count int64) {
	global.GVA_DB.Model(&crm.CrmBillPayment{}).Where("created_at >= ? ", time.Now().Format("2006-01-02")).Count(&count)
	return
}

// UpdateCrmBillPayment 更新应付账单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillPaymentService *CrmBillPaymentService) UpdateCrmBillPayment(crmBillPayment crm.CrmBillPayment) (err error) {
	err = global.GVA_DB.Save(&crmBillPayment).Error
	return err
}

// GetCrmBillPayment 根据ID获取应付账单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillPaymentService *CrmBillPaymentService) GetCrmBillPayment(ID string) (crmPageBillPayment crm.CrmPageBillPayment, err error) {
	err = global.GVA_DB.Model(&crm.CrmBillPayment{}).Where("crm_bill_payment.id = ?", ID).
		Select("crm_bill_payment.*,crm_statement_account.statement_account_name,crm_statement_account.statement_account_number").
		Joins("LEFT JOIN crm_statement_account ON crm_statement_account.id = crm_bill_payment.statement_account_id").
		Order("crm_bill_payment.created_at DESC").First(&crmPageBillPayment).Error
	return
}

func (crmBillPaymentService *CrmBillPaymentService) GetCrmBillPaymentID(ID *int) (crmPageBillPayment crm.CrmPageBillPayment, err error) {
	err = global.GVA_DB.Model(&crm.CrmBillPayment{}).Where("crm_bill_payment.id = ?", ID).
		Select("crm_bill_payment.*,crm_statement_account.statement_account_name,crm_statement_account.statement_account_number").
		Joins("LEFT JOIN crm_statement_account ON crm_statement_account.id = crm_bill_payment.statement_account_id").
		Order("crm_bill_payment.created_at DESC").First(&crmPageBillPayment).Error
	return
}

// GetCrmBillPaymentInfoList 分页获取应付账单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillPaymentService *CrmBillPaymentService) GetCrmBillPaymentInfoList(info crmReq.CrmBillPaymentSearch) (list []crm.CrmPageBillPayment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmBillPayment{})
	var crmBillPayments []crm.CrmPageBillPayment
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmBillPaymentService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Amount != nil {
		db = db.Where(crmBillPaymentService.SplicingQueryConditions("amount = ?"), info.Amount)
	}
	if info.BillPaymentName != "" {
		db = db.Where(crmBillPaymentService.SplicingQueryConditions("bill_payment_name LIKE ?"), "%"+info.BillPaymentName+"%")
	}
	if info.Currency != "" {
		db = db.Where(crmBillPaymentService.SplicingQueryConditions("currency = ?"), info.Currency)
	}
	if info.CustomerId != nil {
		db = db.Where(crmBillPaymentService.SplicingQueryConditions("customer_id = ?"), info.CustomerId)
	}
	if info.PaymentStatus != "" {
		db = db.Where(crmBillPaymentService.SplicingQueryConditions("payment_status = ?"), info.PaymentStatus)
	}
	if info.StatementAccountId != nil {
		db = db.Where(crmBillPaymentService.SplicingQueryConditions("statement_account_id = ?"), info.StatementAccountId)
	}
	if info.UserId != nil {
		db = db.Where(crmBillPaymentService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// var OrderStr string
	// orderMap := make(map[string]bool)
	//  	orderMap["amount"] = true
	//if orderMap[info.Sort] {
	//   OrderStr = info.Sort
	//   if info.Order == "descending" {
	//      OrderStr = OrderStr + " desc"
	//   }
	//   db = db.Order(OrderStr)
	//}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_bill_payment.*,crm_statement_account.statement_account_name,crm_statement_account.statement_account_number").
		Joins("LEFT JOIN crm_statement_account ON crm_statement_account.id = crm_bill_payment.statement_account_id").
		Order("crm_bill_payment.created_at DESC").
		Find(&crmBillPayments).Error
	return crmBillPayments, total, err
}

// UpdApprovalStatus 修改审批状态
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillPaymentService *CrmBillPaymentService) UpdApprovalStatus(ID *int, data map[string]interface{}) (err error) {
	db := global.GVA_DB.Model(&crm.CrmBillPayment{})
	err = db.Where("id = ?", ID).Updates(data).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmBillPaymentService *CrmBillPaymentService) SplicingQueryConditions(condition string) string {
	return "crm_bill_payment." + condition
}
