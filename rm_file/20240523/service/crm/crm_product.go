package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmProductService struct {
}

// CreateCrmProduct 创建产品管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService) CreateCrmProduct(crmProduct *crm.CrmProduct) (err error) {
	err = global.GVA_DB.Create(crmProduct).Error
	return err
}

// DeleteCrmProduct 删除产品管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService)DeleteCrmProduct(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmProduct{},"id = ?",ID).Error
	return err
}

// DeleteCrmProductByIds 批量删除产品管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService)DeleteCrmProductByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmProduct{},"id in ?",IDs).Error
	return err
}

// UpdateCrmProduct 更新产品管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService)UpdateCrmProduct(crmProduct crm.CrmProduct) (err error) {
	err = global.GVA_DB.Save(&crmProduct).Error
	return err
}

// GetCrmProduct 根据ID获取产品管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService)GetCrmProduct(ID string) (crmProduct crm.CrmProduct, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmProduct).Error
	return
}

// GetCrmProductInfoList 分页获取产品管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService)GetCrmProductInfoList(info crmReq.CrmProductSearch) (list []crm.CrmProduct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmProduct{})
    var crmProducts []crm.CrmProduct
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
	
	err = db.Find(&crmProducts).Error
	return  crmProducts, total, err
}
