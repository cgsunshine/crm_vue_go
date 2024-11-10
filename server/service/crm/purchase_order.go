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
		db = db.Where(crmPurchaseOrderService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Amount != nil {
		db = db.Where(crmPurchaseOrderService.SplicingQueryConditions("amount = ?"), info.Amount)
	}
	if info.ContractId != nil {
		db = db.Where(crmPurchaseOrderService.SplicingQueryConditions("contract_id = ?"), info.ContractId)
	}
	if info.StartCreationTime != nil && info.EndCreationTime != nil {
		db = db.Where(crmPurchaseOrderService.SplicingQueryConditions("creation_time BETWEEN ? AND ? "), info.StartCreationTime, info.EndCreationTime)
	}
	if info.StartExpirationTime != nil && info.EndExpirationTime != nil {
		db = db.Where(crmPurchaseOrderService.SplicingQueryConditions("expiration_time BETWEEN ? AND ? "), info.StartExpirationTime, info.EndExpirationTime)
	}
	if info.ProductId != nil {
		db = db.Where(crmPurchaseOrderService.SplicingQueryConditions("product_id = ?"), info.ProductId)
	}
	if info.Quantity != nil {
		db = db.Where(crmPurchaseOrderService.SplicingQueryConditions("quantity = ?"), info.Quantity)
	}
	if info.UserId != nil {
		db = db.Where(crmPurchaseOrderService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}
	if info.PurchaseOrderNumber != "" {
		db = db.Where(crmPurchaseOrderService.SplicingQueryConditions("purchase_order_number = ?"), info.PurchaseOrderNumber)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_purchase_order.*,sys_users.username,crm_product.product_name,crm_currency.currency_name,crm_procurement_contract.expiration_time").
		Joins("LEFT JOIN sys_users ON crm_purchase_order.user_id = sys_users.id").
		Joins("LEFT JOIN crm_product ON crm_purchase_order.product_id = crm_product.id").
		Joins("LEFT JOIN crm_currency ON crm_currency.id = crm_purchase_order.currency_id").
		Joins("LEFT JOIN crm_procurement_contract ON crm_procurement_contract.id = crm_purchase_order.contract_id").
		Order("crm_purchase_order.created_at DESC").
		Preload("PurchaseOrderProduct.Product").
		Find(&crmPurchaseOrders).Error
	return crmPurchaseOrders, total, err
}

// GetCrmPurchaseOrder 根据ID获取订购单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderService *CrmPurchaseOrderService) GetCrmPagePurchaseOrder(ID string) (crmPurchaseOrder crm.CrmPagePurchaseOrder, err error) {
	err = global.GVA_DB.Model(&crm.CrmPurchaseOrder{}).Where("crm_purchase_order.id = ?", ID).
		Select("crm_purchase_order.*,sys_users.username,crm_product.product_name,crm_currency.currency_name,crm_procurement_contract.expiration_time").
		Joins("LEFT JOIN sys_users ON crm_purchase_order.user_id = sys_users.id").
		Joins("LEFT JOIN crm_product ON crm_purchase_order.product_id = crm_product.id").
		Joins("LEFT JOIN crm_currency ON crm_currency.id = crm_purchase_order.currency_id").
		Joins("LEFT JOIN crm_procurement_contract ON crm_procurement_contract.id = crm_purchase_order.contract_id").
		Preload("PurchaseOrderProduct.Product").
		First(&crmPurchaseOrder).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmPurchaseOrderService *CrmPurchaseOrderService) SplicingQueryConditions(condition string) string {
	return "crm_purchase_order." + condition
}
