package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmCustomerGroupService struct {
}

// CreateCrmCustomerGroup 创建crmCustomerGroup表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomerGroupService *CrmCustomerGroupService) CreateCrmCustomerGroup(crmCustomerGroup *crm.CrmCustomerGroup) (err error) {
	err = global.GVA_DB.Create(crmCustomerGroup).Error
	return err
}

// DeleteCrmCustomerGroup 删除crmCustomerGroup表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomerGroupService *CrmCustomerGroupService)DeleteCrmCustomerGroup(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmCustomerGroup{},"id = ?",ID).Error
	return err
}

// DeleteCrmCustomerGroupByIds 批量删除crmCustomerGroup表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomerGroupService *CrmCustomerGroupService)DeleteCrmCustomerGroupByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmCustomerGroup{},"id in ?",IDs).Error
	return err
}

// UpdateCrmCustomerGroup 更新crmCustomerGroup表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomerGroupService *CrmCustomerGroupService)UpdateCrmCustomerGroup(crmCustomerGroup crm.CrmCustomerGroup) (err error) {
	err = global.GVA_DB.Save(&crmCustomerGroup).Error
	return err
}

// GetCrmCustomerGroup 根据ID获取crmCustomerGroup表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomerGroupService *CrmCustomerGroupService)GetCrmCustomerGroup(ID string) (crmCustomerGroup crm.CrmCustomerGroup, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmCustomerGroup).Error
	return
}

// GetCrmCustomerGroupInfoList 分页获取crmCustomerGroup表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCustomerGroupService *CrmCustomerGroupService)GetCrmCustomerGroupInfoList(info crmReq.CrmCustomerGroupSearch) (list []crm.CrmCustomerGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmCustomerGroup{})
    var crmCustomerGroups []crm.CrmCustomerGroup
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
	
	err = db.Find(&crmCustomerGroups).Error
	return  crmCustomerGroups, total, err
}
