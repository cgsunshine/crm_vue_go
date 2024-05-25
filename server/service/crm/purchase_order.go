package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmPurchaseOrderInfoList 分页获取crmPurchaseOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderService *CrmPurchaseOrderService) GetCrmPagePurchaseOrderInfoList(info crmReq.CrmPurchaseOrderSearch) (list []crm.CrmPagePurchaseOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmPurchaseOrder{})
	var crmPurchaseOrders []crm.CrmPagePurchaseOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ContractId != nil {
		db = db.Where("contract_id = ?", info.ContractId)
	}
	if info.ProductId != nil {
		db = db.Where("product_id = ?", info.ProductId)
	}
	if info.Quantity != nil {
		db = db.Where("quantity = ?", info.Quantity)
	}
	if info.Amount != nil {
		db = db.Where("amount = ?", info.Amount)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.StartCreationTime != nil && info.EndCreationTime != nil {
		db = db.Where("creation_time BETWEEN ? AND ? ", info.StartCreationTime, info.EndCreationTime)
	}
	if info.StartExpirationTime != nil && info.EndExpirationTime != nil {
		db = db.Where("expiration_time BETWEEN ? AND ? ", info.StartExpirationTime, info.EndExpirationTime)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_purchase_order.*,sys_users.username,crm_product.product_name").
		Joins("LEFT JOIN sys_users ON crm_purchase_order.user_id = sys_users.id").
		Joins("LEFT JOIN crm_product ON crm_purchase_order.product_id = crm_product.id").
		Find(&crmPurchaseOrders).Error
	return crmPurchaseOrders, total, err
}
