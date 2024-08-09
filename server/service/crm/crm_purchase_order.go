package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmPurchaseOrderService struct {
}

// CreateCrmPurchaseOrder 创建订购单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderService *CrmPurchaseOrderService) CreateCrmPurchaseOrder(crmPurchaseOrder *crm.CrmPurchaseOrder) (err error) {
	err = global.GVA_DB.Create(crmPurchaseOrder).Error
	return err
}

// DeleteCrmPurchaseOrder 删除订购单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderService *CrmPurchaseOrderService) DeleteCrmPurchaseOrder(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmPurchaseOrder{}, "id = ?", ID).Error
	return err
}

// DeleteCrmPurchaseOrderByIds 批量删除订购单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderService *CrmPurchaseOrderService) DeleteCrmPurchaseOrderByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmPurchaseOrder{}, "id in ?", IDs).Error
	return err
}

// UpdateCrmPurchaseOrder 更新订购单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderService *CrmPurchaseOrderService) UpdateCrmPurchaseOrder(crmPurchaseOrder crm.CrmPurchaseOrder) (err error) {
	err = global.GVA_DB.Save(&crmPurchaseOrder).Error
	return err
}

// GetCrmPurchaseOrder 根据ID获取订购单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderService *CrmPurchaseOrderService) GetCrmPurchaseOrder(ID string) (crmPurchaseOrder crm.CrmPurchaseOrder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmPurchaseOrder).Error
	return
}

// GetCrmPurchaseOrderInfoList 分页获取订购单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmPurchaseOrderService *CrmPurchaseOrderService) GetCrmPurchaseOrderInfoList(info crmReq.CrmPurchaseOrderSearch) (list []crm.CrmPurchaseOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmPurchaseOrder{})
	var crmPurchaseOrders []crm.CrmPurchaseOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Amount != nil {
		db = db.Where("amount = ?", info.Amount)
	}
	if info.ContractId != nil {
		db = db.Where("contract_id = ?", info.ContractId)
	}
	if info.StartCreationTime != nil && info.EndCreationTime != nil {
		db = db.Where("creation_time BETWEEN ? AND ? ", info.StartCreationTime, info.EndCreationTime)
	}
	if info.StartExpirationTime != nil && info.EndExpirationTime != nil {
		db = db.Where("expiration_time BETWEEN ? AND ? ", info.StartExpirationTime, info.EndExpirationTime)
	}
	if info.ProductId != nil {
		db = db.Where("product_id = ?", info.ProductId)
	}
	if info.Quantity != nil {
		db = db.Where("quantity = ?", info.Quantity)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}

	if info.PurchaseOrderName != "" {
		db = db.Where("payee LIKE ?", "%"+info.PurchaseOrderName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&crmPurchaseOrders).Error
	return crmPurchaseOrders, total, err
}
