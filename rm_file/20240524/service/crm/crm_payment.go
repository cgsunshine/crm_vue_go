package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmPaymentService struct {
}

// CreateCrmPayment 创建付款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService) CreateCrmPayment(crmPayment *crm.CrmPayment) (err error) {
	err = global.GVA_DB.Create(crmPayment).Error
	return err
}

// DeleteCrmPayment 删除付款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService)DeleteCrmPayment(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmPayment{},"id = ?",ID).Error
	return err
}

// DeleteCrmPaymentByIds 批量删除付款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService)DeleteCrmPaymentByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmPayment{},"id in ?",IDs).Error
	return err
}

// UpdateCrmPayment 更新付款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService)UpdateCrmPayment(crmPayment crm.CrmPayment) (err error) {
	err = global.GVA_DB.Save(&crmPayment).Error
	return err
}

// GetCrmPayment 根据ID获取付款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService)GetCrmPayment(ID string) (crmPayment crm.CrmPayment, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmPayment).Error
	return
}

// GetCrmPaymentInfoList 分页获取付款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentService *CrmPaymentService)GetCrmPaymentInfoList(info crmReq.CrmPaymentSearch) (list []crm.CrmPayment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmPayment{})
    var crmPayments []crm.CrmPayment
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmPayments).Error
	return  crmPayments, total, err
}
