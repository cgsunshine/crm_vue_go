package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmOrderInfoList 分页获取crmOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) GetCrmPageOrderInfoList(info crmReq.CrmOrderSearch) (list []crm.CrmPageOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmOrder{})
	var crmOrders []crm.CrmPageOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
	}
	if info.CustomerId != nil {
		db = db.Where("customer_id = ?", info.CustomerId)
	}
	if info.DiscountRate != nil {
		db = db.Where("discount_rate = ?", info.DiscountRate)
	}
	if info.Price != nil {
		db = db.Where("price = ?", info.Price)
	}
	if info.ProductId != "" {
		db = db.Where("product_id = ?", info.ProductId)
	}
	if info.SalesPrice != nil {
		db = db.Where("sales_price = ?", info.SalesPrice)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_order.*,crm_product.product_name,crm_customers.customer_name,sys_users.username").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_order.user_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_order.customer_id").
		Joins("LEFT JOIN crm_product ON crm_product.id = crm_order.product_id").
		Find(&crmOrders).Error
	return crmOrders, total, err
}
