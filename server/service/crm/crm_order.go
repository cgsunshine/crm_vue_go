package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"time"
)

type CrmOrderService struct {
}

// CreateCrmOrder 创建crmOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) CreateCrmOrder(crmOrder *crm.CrmOrder) (err error) {
	err = global.GVA_DB.Create(crmOrder).Error
	return err
}

// DeleteCrmOrder 删除crmOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) DeleteCrmOrder(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmOrder{}, "id = ?", ID).Error
	return err
}

// GetCrmBillTodayCount 根据ID获取crmPaymentCollention表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) GetCrmBillTodayCount() (count int64) {
	global.GVA_DB.Model(&crm.CrmOrder{}).Where("created_at >= ? ", time.Now().Format("2006-01-02")).Count(&count)
	return
}

// DeleteCrmOrderByIds 批量删除crmOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) DeleteCrmOrderByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmOrder{}, "id in ?", IDs).Error
	return err
}

// UpdateCrmOrder 更新crmOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) UpdateCrmOrder(crmOrder crm.CrmOrder) (err error) {
	err = global.GVA_DB.Save(&crmOrder).Error
	return err
}

// GetCrmOrder 根据ID获取crmOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) GetCrmOrder(ID string) (crmOrder crm.CrmOrder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmOrder).Error
	return
}

// GetCrmOrderInfoList 分页获取crmOrder表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) GetCrmOrderInfoList(info crmReq.CrmOrderSearch) (list []crm.CrmOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmOrder{})
	var crmOrders []crm.CrmOrder
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

	err = db.Find(&crmOrders).Error
	return crmOrders, total, err
}

func (crmOrderService *CrmOrderService) OrderCycleTimeAdd(userId *int, startDate, endDate *time.Time) (total int64, err error) {
	db := global.GVA_DB.Model(&crm.CrmOrder{})
	SearchCondition(db, userId, startDate, endDate)
	err = db.Count(&total).Error
	return
}

// ApprovalTasksCount 统计有效的，并且状态是待审批的数量
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) ApprovalTasksCount(userId *int, approvalStatus string, startDate, endDate *time.Time) (total int64, err error) {
	db := global.GVA_DB.Model(&crm.CrmOrder{})
	SearchCondition(db, userId, startDate, endDate)
	err = db.Where("review_status = ? ", approvalStatus).Debug().Count(&total).Error
	return
}
