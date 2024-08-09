package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmProductInfoList 分页获取产品管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService) GetCrmPageProductInfoList(info crmReq.CrmProductSearch) (list []crm.CrmPageProduct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmProduct{})
	var crmProducts []crm.CrmPageProduct
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmProductService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ProductName != "" {
		db = db.Where(crmProductService.SplicingQueryConditions("product_name = ?"), info.ProductName)
	}
	if info.ProductGroupId != nil {
		db = db.Where(crmProductService.SplicingQueryConditions("product_group_id = ?"), info.ProductGroupId)
	}
	if info.ProductTypeId != nil {
		db = db.Where(crmProductService.SplicingQueryConditions("product_type_id = ?"), info.ProductTypeId)
	}
	if info.Inventory != nil {
		db = db.Where(crmProductService.SplicingQueryConditions("inventory <> ?"), info.Inventory)
	}
	if info.Price != nil {
		db = db.Where(crmProductService.SplicingQueryConditions("price <> ?"), info.Price)
	}
	if info.DiscountPrice != nil {
		db = db.Where(crmProductService.SplicingQueryConditions("discount_price <> ?"), info.DiscountPrice)
	}
	if info.SalesPrice != nil {
		db = db.Where(crmProductService.SplicingQueryConditions("sales_price <> ?"), info.SalesPrice)
	}
	if info.ResourceId != nil {
		db = db.Where(crmProductService.SplicingQueryConditions("resource_id = ?"), info.ResourceId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_product.*,crm_product_group.group_name,crm_product_type.type_name,sys_users.username").
		Joins("LEFT JOIN sys_users ON crm_product.user_id = sys_users.id").
		Joins("LEFT JOIN crm_product_group ON crm_product.product_group_id = crm_product_group.id").
		Joins("LEFT JOIN crm_product_type ON crm_product.product_type_id = crm_product_type.id").
		Order("crm_product.created_at DESC").
		Find(&crmProducts).Error
	return crmProducts, total, err
}

// GetCrmProduct 根据ID获取crmProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductService *CrmProductService) GetCrmPageProduct(ID string) (crmProduct crm.CrmPageProduct, err error) {
	err = global.GVA_DB.Model(&crm.CrmProduct{}).Where("crm_product.id = ?", ID).
		Select("crm_product.*,crm_product_group.group_name,crm_product_type.type_name,sys_users.username").
		Joins("LEFT JOIN sys_users ON crm_product.user_id = sys_users.id").
		Joins("LEFT JOIN crm_product_group ON crm_product.product_group_id = crm_product_group.id").
		Joins("LEFT JOIN crm_product_type ON crm_product.product_type_id = crm_product_type.id").
		First(&crmProduct).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmProductService *CrmProductService) SplicingQueryConditions(condition string) string {
	return "crm_product." + condition
}
