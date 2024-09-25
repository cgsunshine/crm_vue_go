package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmPurchaseOrderProductService struct {
}

// CreateCrmPurchaseOrderProduct 创建crmPurchaseOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderProductService *CrmPurchaseOrderProductService) CreateCrmPurchaseOrderProduct(crmPurchaseOrderProduct *crm.CrmPurchaseOrderProduct) (err error) {
	err = global.GVA_DB.Create(crmPurchaseOrderProduct).Error
	return err
}

// CreateCrmPurchaseOrderProduct 创建crmPurchaseOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderProductService *CrmPurchaseOrderProductService) CreateCrmPurchaseOrderProducts(crmPurchaseOrderProduct []*crm.CrmPurchaseOrderProduct) (err error) {
	err = global.GVA_DB.Create(crmPurchaseOrderProduct).Error
	return err
}

// DeleteCrmPurchaseOrderProduct 删除crmPurchaseOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderProductService *CrmPurchaseOrderProductService) DeleteCrmPurchaseOrderProduct(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmPurchaseOrderProduct{}, "id = ?", ID).Error
	return err
}

// DeleteCrmPurchaseOrderProductByIds 批量删除crmPurchaseOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderProductService *CrmPurchaseOrderProductService) DeleteCrmPurchaseOrderProductByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmPurchaseOrderProduct{}, "id in ?", IDs).Error
	return err
}

// UpdateCrmPurchaseOrderProduct 更新crmPurchaseOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderProductService *CrmPurchaseOrderProductService) UpdateCrmPurchaseOrderProduct(crmPurchaseOrderProduct crm.CrmPurchaseOrderProduct) (err error) {
	err = global.GVA_DB.Save(&crmPurchaseOrderProduct).Error
	return err
}

// GetCrmPurchaseOrderProduct 根据ID获取crmPurchaseOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderProductService *CrmPurchaseOrderProductService) GetCrmPurchaseOrderProduct(ID string) (crmPurchaseOrderProduct crm.CrmPurchaseOrderProduct, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmPurchaseOrderProduct).Error
	return
}

// GetCrmPurchaseOrderProductInfoList 分页获取crmPurchaseOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderProductService *CrmPurchaseOrderProductService) GetCrmPurchaseOrderProductInfoList(info crmReq.CrmPurchaseOrderProductSearch) (list []crm.CrmPurchaseOrderProduct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmPurchaseOrderProduct{})
	var crmPurchaseOrderProducts []crm.CrmPurchaseOrderProduct
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

	err = db.Find(&crmPurchaseOrderProducts).Error
	return crmPurchaseOrderProducts, total, err
}
