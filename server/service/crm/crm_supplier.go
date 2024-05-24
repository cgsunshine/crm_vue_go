package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmSupplierService struct {
}

// CreateCrmSupplier 创建crmSupplier表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService) CreateCrmSupplier(crmSupplier *crm.CrmSupplier) (err error) {
	err = global.GVA_DB.Create(crmSupplier).Error
	return err
}

// DeleteCrmSupplier 删除crmSupplier表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService)DeleteCrmSupplier(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmSupplier{},"id = ?",ID).Error
	return err
}

// DeleteCrmSupplierByIds 批量删除crmSupplier表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService)DeleteCrmSupplierByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmSupplier{},"id in ?",IDs).Error
	return err
}

// UpdateCrmSupplier 更新crmSupplier表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService)UpdateCrmSupplier(crmSupplier crm.CrmSupplier) (err error) {
	err = global.GVA_DB.Save(&crmSupplier).Error
	return err
}

// GetCrmSupplier 根据ID获取crmSupplier表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService)GetCrmSupplier(ID string) (crmSupplier crm.CrmSupplier, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmSupplier).Error
	return
}

// GetCrmSupplierInfoList 分页获取crmSupplier表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmSupplierService *CrmSupplierService)GetCrmSupplierInfoList(info crmReq.CrmSupplierSearch) (list []crm.CrmSupplier, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmSupplier{})
    var crmSuppliers []crm.CrmSupplier
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.CommpanyName != "" {
        db = db.Where("commpany_name LIKE ?","%"+ info.CommpanyName+"%")
    }
    if info.ContactPerson != "" {
        db = db.Where("contact_person LIKE ?","%"+ info.ContactPerson+"%")
    }
    if info.Email != "" {
        db = db.Where("email LIKE ?","%"+ info.Email+"%")
    }
    if info.Telephone != "" {
        db = db.Where("telephone = ?",info.Telephone)
    }
        if info.StartNoteAddTime != nil && info.EndNoteAddTime != nil {
            db = db.Where("note_add_time BETWEEN ? AND ? ",info.StartNoteAddTime,info.EndNoteAddTime)
        }
    if info.UserId != nil {
        db = db.Where("user_id = ?",info.UserId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmSuppliers).Error
	return  crmSuppliers, total, err
}
