package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmOrderService struct {
}

// CreateCrmOrder 创建订单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService) CreateCrmOrder(crmOrder *crm.CrmOrder) (err error) {
	err = global.GVA_DB.Create(crmOrder).Error
	return err
}

// DeleteCrmOrder 删除订单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService)DeleteCrmOrder(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmOrder{},"id = ?",ID).Error
	return err
}

// DeleteCrmOrderByIds 批量删除订单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService)DeleteCrmOrderByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmOrder{},"id in ?",IDs).Error
	return err
}

// UpdateCrmOrder 更新订单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService)UpdateCrmOrder(crmOrder crm.CrmOrder) (err error) {
	err = global.GVA_DB.Save(&crmOrder).Error
	return err
}

// GetCrmOrder 根据ID获取订单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService)GetCrmOrder(ID string) (crmOrder crm.CrmOrder, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmOrder).Error
	return
}

// GetCrmOrderInfoList 分页获取订单管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmOrderService *CrmOrderService)GetCrmOrderInfoList(info crmReq.CrmOrderSearch) (list []crm.CrmOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmOrder{})
    var crmOrders []crm.CrmOrder
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmOrders).Error
	return  crmOrders, total, err
}
