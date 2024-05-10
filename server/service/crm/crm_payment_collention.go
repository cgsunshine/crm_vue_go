package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmPaymentCollentionService struct {
}

// CreateCrmPaymentCollention 创建回款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService) CreateCrmPaymentCollention(crmPaymentCollention *crm.CrmPaymentCollention) (err error) {
	err = global.GVA_DB.Create(crmPaymentCollention).Error
	return err
}

// DeleteCrmPaymentCollention 删除回款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService)DeleteCrmPaymentCollention(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmPaymentCollention{},"id = ?",ID).Error
	return err
}

// DeleteCrmPaymentCollentionByIds 批量删除回款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService)DeleteCrmPaymentCollentionByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmPaymentCollention{},"id in ?",IDs).Error
	return err
}

// UpdateCrmPaymentCollention 更新回款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService)UpdateCrmPaymentCollention(crmPaymentCollention crm.CrmPaymentCollention) (err error) {
	err = global.GVA_DB.Save(&crmPaymentCollention).Error
	return err
}

// GetCrmPaymentCollention 根据ID获取回款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService)GetCrmPaymentCollention(ID string) (crmPaymentCollention crm.CrmPaymentCollention, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmPaymentCollention).Error
	return
}

// GetCrmPaymentCollentionInfoList 分页获取回款管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPaymentCollentionService *CrmPaymentCollentionService)GetCrmPaymentCollentionInfoList(info crmReq.CrmPaymentCollentionSearch) (list []crm.CrmPaymentCollention, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmPaymentCollention{})
    var crmPaymentCollentions []crm.CrmPaymentCollention
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
	
	err = db.Find(&crmPaymentCollentions).Error
	return  crmPaymentCollentions, total, err
}
