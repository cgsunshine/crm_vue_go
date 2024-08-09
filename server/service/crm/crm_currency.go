package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmCurrencyService struct {
}

// CreateCrmCurrency 创建币种管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCurrencyService *CrmCurrencyService) CreateCrmCurrency(crmCurrency *crm.CrmCurrency) (err error) {
	err = global.GVA_DB.Create(crmCurrency).Error
	return err
}

// DeleteCrmCurrency 删除币种管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCurrencyService *CrmCurrencyService)DeleteCrmCurrency(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmCurrency{},"id = ?",ID).Error
	return err
}

// DeleteCrmCurrencyByIds 批量删除币种管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCurrencyService *CrmCurrencyService)DeleteCrmCurrencyByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmCurrency{},"id in ?",IDs).Error
	return err
}

// UpdateCrmCurrency 更新币种管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCurrencyService *CrmCurrencyService)UpdateCrmCurrency(crmCurrency crm.CrmCurrency) (err error) {
	err = global.GVA_DB.Save(&crmCurrency).Error
	return err
}

// GetCrmCurrency 根据ID获取币种管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCurrencyService *CrmCurrencyService)GetCrmCurrency(ID string) (crmCurrency crm.CrmCurrency, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmCurrency).Error
	return
}

// GetCrmCurrencyInfoList 分页获取币种管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCurrencyService *CrmCurrencyService)GetCrmCurrencyInfoList(info crmReq.CrmCurrencySearch) (list []crm.CrmCurrency, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmCurrency{})
    var crmCurrencys []crm.CrmCurrency
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
	
	err = db.Find(&crmCurrencys).Error
	return  crmCurrencys, total, err
}
