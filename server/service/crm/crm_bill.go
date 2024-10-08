package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"time"
)

type CrmBillService struct {
}

// CreateCrmBill 创建crmBill表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillService *CrmBillService) CreateCrmBill(crmBill *crm.CrmBill) (err error) {
	err = global.GVA_DB.Create(crmBill).Error
	return err
}

// DeleteCrmBill 删除crmBill表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillService *CrmBillService) DeleteCrmBill(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmBill{}, "id = ?", ID).Error
	return err
}

// DeleteCrmBillByIds 批量删除crmBill表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillService *CrmBillService) DeleteCrmBillByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmBill{}, "id in ?", IDs).Error
	return err
}

// UpdateCrmBill 更新crmBill表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillService *CrmBillService) UpdateCrmBill(crmBill crm.CrmBill) (err error) {
	err = global.GVA_DB.Save(&crmBill).Error
	return err
}

// GetCrmBillTodayCount 根据ID获取crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillService *CrmBillService) GetCrmBillTodayCount() (count int64) {
	global.GVA_DB.Model(&crm.CrmBill{}).Where("created_at >= ? ", time.Now().Format("2006-01-02")).Count(&count)
	return
}

// GetCrmBill 根据ID获取crmBill表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillService *CrmBillService) GetCrmBill(ID string) (crmBill crm.CrmBill, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmBill).Error
	return
}

// GetCrmBillInfoList 分页获取crmBill表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBillService *CrmBillService) GetCrmBillInfoList(info crmReq.CrmBillSearch) (list []crm.CrmBill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmBill{})
	var crmBills []crm.CrmBill
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&crmBills).Error
	return crmBills, total, err
}
