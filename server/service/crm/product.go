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
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
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
		Find(&crmProducts).Error
	return crmProducts, total, err
}
