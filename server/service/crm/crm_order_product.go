package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmOrderProductService struct {
}

// CreateCrmOrderProduct 创建crmOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderProductService *CrmOrderProductService) CreateCrmOrderProduct(crmOrderProduct *crm.CrmOrderProduct) (err error) {
	err = global.GVA_DB.Create(crmOrderProduct).Error
	return err
}

func (crmOrderProductService *CrmOrderProductService) CreateCrmOrderProductInc(crmOrderProduct interface{}) (err error) {
	err = global.GVA_DB.Create(crmOrderProduct).Error
	return err
}

// CreateCrmOrderProduct 创建crmOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderProductService *CrmOrderProductService) CreateCrmOrderProducts(crmOrderProduct []*crm.CrmOrderProduct) (err error) {
	err = global.GVA_DB.Create(crmOrderProduct).Error
	return err
}

// DeleteCrmOrderProduct 删除crmOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderProductService *CrmOrderProductService) DeleteCrmOrderProduct(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmOrderProduct{}, "id = ?", ID).Error
	return err
}

// DeleteUnscopedCrmOrderProduct 删除crmOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderProductService *CrmOrderProductService) DeleteUnscopedCrmOrderProduct(crmOrderProduct *crm.CrmOrderProduct) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&crmOrderProduct).Error
	return err
}

// DeleteCrmOrderProductByIds 批量删除crmOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderProductService *CrmOrderProductService) DeleteCrmOrderProductByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmOrderProduct{}, "id in ?", IDs).Error
	return err
}

// DeleteCrmOrderProductBy 通过oid和pid删除，订单id和产品id
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderProductService *CrmOrderProductService) DeleteCrmOrderProductBy(oid *int) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&crm.CrmOrderProduct{}, "order_id = ?", oid).Error
	return err
}

// UpdateCrmOrderProduct 更新crmOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderProductService *CrmOrderProductService) UpdateCrmOrderProduct(crmOrderProduct crm.CrmOrderProduct) (err error) {
	err = global.GVA_DB.Save(&crmOrderProduct).Error
	return err
}

// GetCrmOrderProduct 根据ID获取crmOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderProductService *CrmOrderProductService) GetCrmOrderProduct(ID string) (crmOrderProduct crm.CrmOrderProduct, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmOrderProduct).Error
	return
}

// GetCrmOrderProductInfoList 分页获取crmOrderProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderProductService *CrmOrderProductService) GetCrmOrderProductInfoList(info crmReq.CrmOrderProductSearch) (list []crm.CrmOrderProduct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmOrderProduct{})
	var crmOrderProducts []crm.CrmOrderProduct
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.OrderId != nil {
		db = db.Where("order_id = ?", info.OrderId)
	}
	if info.ProductId != nil {
		db = db.Where("product_id = ?", info.ProductId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&crmOrderProducts).Error
	return crmOrderProducts, total, err
}
