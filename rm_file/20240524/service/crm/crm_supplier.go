package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmSupplierService struct {
}

// CreateCrmSupplier 创建供应商管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService) CreateCrmSupplier(crmSupplier *crm.CrmSupplier) (err error) {
	err = global.GVA_DB.Create(crmSupplier).Error
	return err
}

// DeleteCrmSupplier 删除供应商管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService)DeleteCrmSupplier(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmSupplier{},"id = ?",ID).Error
	return err
}

// DeleteCrmSupplierByIds 批量删除供应商管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService)DeleteCrmSupplierByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmSupplier{},"id in ?",IDs).Error
	return err
}

// UpdateCrmSupplier 更新供应商管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService)UpdateCrmSupplier(crmSupplier crm.CrmSupplier) (err error) {
	err = global.GVA_DB.Save(&crmSupplier).Error
	return err
}

// GetCrmSupplier 根据ID获取供应商管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService)GetCrmSupplier(ID string) (crmSupplier crm.CrmSupplier, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmSupplier).Error
	return
}

// GetCrmSupplierInfoList 分页获取供应商管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService)GetCrmSupplierInfoList(info crmReq.CrmSupplierSearch) (list []crm.CrmSupplier, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmSupplier{})
    var crmSuppliers []crm.CrmSupplier
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
	
	err = db.Find(&crmSuppliers).Error
	return  crmSuppliers, total, err
}
