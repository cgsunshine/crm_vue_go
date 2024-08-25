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
		db = db.Where(crmOrderService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Currency != "" {
		db = db.Where(crmOrderService.SplicingQueryConditions("currency = ?"), info.Currency)
	}
	if info.CustomerId != nil {
		db = db.Where(crmOrderService.SplicingQueryConditions("customer_id = ?"), info.CustomerId)
	}
	if info.DiscountRate != nil {
		db = db.Where(crmOrderService.SplicingQueryConditions("discount_rate = ?"), info.DiscountRate)
	}
	if info.Price != nil {
		db = db.Where(crmOrderService.SplicingQueryConditions("price = ?"), info.Price)
	}
	if info.ProductId != "" {
		db = db.Where(crmOrderService.SplicingQueryConditions("product_id = ?"), info.ProductId)
	}

	if info.SalesPrice != nil {
		db = db.Where(crmOrderService.SplicingQueryConditions("sales_price = ?"), info.SalesPrice)
	}

	if info.UserId != nil {
		db = db.Where(crmOrderService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}

	if info.OrderName != "" {
		db = db.Where(crmOrderService.SplicingQueryConditions("order_name LIKE ?"), "%"+info.OrderName+"%")
	}

	if info.ReviewStatus != "" {
		db = db.Where(crmOrderService.SplicingQueryConditions("review_status = ?"), info.ReviewStatus)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_order.*,crm_product.product_name,crm_customers.customer_name,sys_users.username,crm_business_opportunity.business_opportunity_name").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_order.user_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_order.customer_id").
		Joins("LEFT JOIN crm_product ON crm_product.id = crm_order.product_id").
		Joins("LEFT JOIN crm_business_opportunity ON crm_business_opportunity.id = crm_order.business_opportunity_id").
		//Preload("Products.OrderProducts").
		Preload("OrderProducts.Product").
		Order("crm_order.created_at DESC").
		Find(&crmOrders).Error
	return crmOrders, total, err
}

// GetCrmOrder 根据ID获取crmOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) GetCrmPageOrder(ID string) (crmOrder crm.CrmPageOrder, err error) {
	err = global.GVA_DB.Model(&crm.CrmOrder{}).Where("crm_order.id = ?", ID).
		Select("crm_order.*,crm_product.product_name,crm_customers.customer_name,sys_users.username,crm_business_opportunity.business_opportunity_name").
		Joins("LEFT JOIN sys_users ON sys_users.id = crm_order.user_id").
		Joins("LEFT JOIN crm_customers ON crm_customers.id = crm_order.customer_id").
		Joins("LEFT JOIN crm_product ON crm_product.id = crm_order.product_id").
		Joins("LEFT JOIN crm_business_opportunity ON crm_business_opportunity.id = crm_order.business_opportunity_id").
		//Preload("Products").
		Preload("OrderProducts.Product").
		Order("crm_order.created_at DESC").
		First(&crmOrder).Error
	return
}

// GetCrmOrder 根据ID获取crmOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) GetCrmOrderId(ID *int) (crmOrder crm.CrmOrder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmOrder).Error
	return
}

// UpdApprovalStatus 修改审批状态
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) UpdApprovalStatus(ID *int, data map[string]interface{}) (err error) {
	db := global.GVA_DB.Model(&crm.CrmOrder{})
	err = db.Where("id = ?", ID).Updates(data).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmOrderService *CrmOrderService) SplicingQueryConditions(condition string) string {
	return "crm_order." + condition
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetOrderProducts
//@description: 设置一个订单的产品
//@param: id uint, authorityIds []string
//@return: err error

func (crmOrderService *CrmOrderService) SetOrderProducts(id int, productIds []int) (err error) {
	db := global.GVA_DB.Model(&crm.CrmOrderProduct{})
	err = db.Delete(&[]crm.CrmOrderProduct{}, "order_id = ?", id).Error
	if err != nil {
		return
	}

	var orderProduct []crm.CrmOrderProduct
	for _, v := range productIds {
		orderProduct = append(orderProduct, crm.CrmOrderProduct{
			OrderId: &id, ProductId: &v,
		})
	}
	err = db.Create(&orderProduct).Error
	if err != nil {
		return
	}
	return
}
