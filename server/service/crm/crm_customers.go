package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmCustomersService struct {
}

// CreateCrmCustomers 创建crmCustomers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomersService *CrmCustomersService) CreateCrmCustomers(crmCustomers *crm.CrmCustomers) (err error) {
	err = global.GVA_DB.Create(crmCustomers).Error
	return err
}

// DeleteCrmCustomers 删除crmCustomers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomersService *CrmCustomersService)DeleteCrmCustomers(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmCustomers{},"id = ?",ID).Error
	return err
}

// DeleteCrmCustomersByIds 批量删除crmCustomers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomersService *CrmCustomersService)DeleteCrmCustomersByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmCustomers{},"id in ?",IDs).Error
	return err
}

// UpdateCrmCustomers 更新crmCustomers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomersService *CrmCustomersService)UpdateCrmCustomers(crmCustomers crm.CrmCustomers) (err error) {
	err = global.GVA_DB.Save(&crmCustomers).Error
	return err
}

// GetCrmCustomers 根据ID获取crmCustomers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomersService *CrmCustomersService)GetCrmCustomers(ID string) (crmCustomers crm.CrmCustomers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmCustomers).Error
	return
}

// GetCrmCustomersInfoList 分页获取crmCustomers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomersService *CrmCustomersService)GetCrmCustomersInfoList(info crmReq.CrmCustomersSearch) (list []crm.CrmCustomers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmCustomers{})
    var crmCustomerss []crm.CrmCustomers
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.CustomerName != "" {
        db = db.Where("customer_name = ?",info.CustomerName)
    }
    if info.CustomerPhoneData != "" {
        db = db.Where("customer_phone_data = ?",info.CustomerPhoneData)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?",info.UserId)
    }
    if info.CustomerCompany != "" {
        db = db.Where("customer_company LIKE ?","%"+ info.CustomerCompany+"%")
    }
    if info.CustomerAddress != "" {
        db = db.Where("customer_address LIKE ?","%"+ info.CustomerAddress+"%")
    }
    if info.CustomerStatus != "" {
        db = db.Where("customer_status = ?",info.CustomerStatus)
    }
    if info.CustomerGroup != "" {
        db = db.Where("customer_group = ?",info.CustomerGroup)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmCustomerss).Error
	return  crmCustomerss, total, err
}
