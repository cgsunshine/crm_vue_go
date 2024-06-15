package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetCrmSupplierInfoList 分页获取crmSupplier表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService) GetCrmPageSupplierInfoList(info crmReq.CrmSupplierSearch) (list []crm.CrmPageSupplier, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmSupplier{})
	var crmSuppliers []crm.CrmPageSupplier
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmSupplierService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CommpanyName != "" {
		db = db.Where(crmSupplierService.SplicingQueryConditions("commpany_name LIKE ?"), "%"+info.CommpanyName+"%")
	}
	if info.ContactPerson != "" {
		db = db.Where(crmSupplierService.SplicingQueryConditions("contact_person LIKE ?"), "%"+info.ContactPerson+"%")
	}
	if info.Email != "" {
		db = db.Where(crmSupplierService.SplicingQueryConditions("email LIKE ?"), "%"+info.Email+"%")
	}
	if info.Telephone != "" {
		db = db.Where(crmSupplierService.SplicingQueryConditions("telephone = ?"), info.Telephone)
	}
	if info.StartNoteAddTime != nil && info.EndNoteAddTime != nil {
		db = db.Where(crmSupplierService.SplicingQueryConditions("note_add_time BETWEEN ? AND ? "), info.StartNoteAddTime, info.EndNoteAddTime)
	}
	if info.UserId != nil {
		db = db.Where(crmSupplierService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_supplier.*,sys_users.username").
		Joins("LEFT JOIN sys_users ON crm_supplier.user_id = sys_users.id").
		Find(&crmSuppliers).Error
	return crmSuppliers, total, err
}

// GetCrmSupplier 根据ID获取crmSupplier表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService) GetCrmPageSupplier(ID string) (crmSupplier crm.CrmPageSupplier, err error) {
	err = global.GVA_DB.Model(&crm.CrmSupplier{}).Where("crm_supplier.id = ?", ID).
		Select("crm_supplier.*,sys_users.username").
		Joins("LEFT JOIN sys_users ON crm_supplier.user_id = sys_users.id").
		First(&crmSupplier).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmSupplierService *CrmSupplierService) SplicingQueryConditions(condition string) string {
	return "crm_supplier." + condition
}
