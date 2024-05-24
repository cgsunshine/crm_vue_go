package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetPageCrmCustomersInfoList 分页获取客户管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomersService *CrmCustomersService) GetPageCrmCustomersInfoList(info crmReq.CrmCustomersSearch) (list []crm.CrmPageCustomers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmCustomers{})
	var crmCustomerss []crm.CrmPageCustomers
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

	err = db.Select("crm_customers.*,sys_users.username,crm_customer_group.group_name").
		Joins("LEFT JOIN sys_users ON crm_customers.user_id = sys_users.id").
		Joins("LEFT JOIN crm_customer_group ON crm_customers.customer_group_id = crm_customer_group.id").
		Find(&crmCustomerss).Error
	return crmCustomerss, total, err
}
