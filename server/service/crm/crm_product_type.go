package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmProductTypeService struct {
}

// CreateCrmProductType 创建产品类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductTypeService *CrmProductTypeService) CreateCrmProductType(crmProductType *crm.CrmProductType) (err error) {
	err = global.GVA_DB.Create(crmProductType).Error
	return err
}

// DeleteCrmProductType 删除产品类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductTypeService *CrmProductTypeService)DeleteCrmProductType(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmProductType{},"id = ?",ID).Error
	return err
}

// DeleteCrmProductTypeByIds 批量删除产品类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductTypeService *CrmProductTypeService)DeleteCrmProductTypeByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmProductType{},"id in ?",IDs).Error
	return err
}

// UpdateCrmProductType 更新产品类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductTypeService *CrmProductTypeService)UpdateCrmProductType(crmProductType crm.CrmProductType) (err error) {
	err = global.GVA_DB.Save(&crmProductType).Error
	return err
}

// GetCrmProductType 根据ID获取产品类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductTypeService *CrmProductTypeService)GetCrmProductType(ID string) (crmProductType crm.CrmProductType, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmProductType).Error
	return
}

// GetCrmProductTypeInfoList 分页获取产品类型记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductTypeService *CrmProductTypeService)GetCrmProductTypeInfoList(info crmReq.CrmProductTypeSearch) (list []crm.CrmProductType, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmProductType{})
    var crmProductTypes []crm.CrmProductType
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
	
	err = db.Find(&crmProductTypes).Error
	return  crmProductTypes, total, err
}
