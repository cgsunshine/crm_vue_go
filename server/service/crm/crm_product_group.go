package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmProductGroupService struct {
}

// CreateCrmProductGroup 创建产品分组记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductGroupService *CrmProductGroupService) CreateCrmProductGroup(crmProductGroup *crm.CrmProductGroup) (err error) {
	err = global.GVA_DB.Create(crmProductGroup).Error
	return err
}

// DeleteCrmProductGroup 删除产品分组记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductGroupService *CrmProductGroupService)DeleteCrmProductGroup(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmProductGroup{},"id = ?",ID).Error
	return err
}

// DeleteCrmProductGroupByIds 批量删除产品分组记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductGroupService *CrmProductGroupService)DeleteCrmProductGroupByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmProductGroup{},"id in ?",IDs).Error
	return err
}

// UpdateCrmProductGroup 更新产品分组记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductGroupService *CrmProductGroupService)UpdateCrmProductGroup(crmProductGroup crm.CrmProductGroup) (err error) {
	err = global.GVA_DB.Save(&crmProductGroup).Error
	return err
}

// GetCrmProductGroup 根据ID获取产品分组记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductGroupService *CrmProductGroupService)GetCrmProductGroup(ID string) (crmProductGroup crm.CrmProductGroup, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmProductGroup).Error
	return
}

// GetCrmProductGroupInfoList 分页获取产品分组记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmProductGroupService *CrmProductGroupService)GetCrmProductGroupInfoList(info crmReq.CrmProductGroupSearch) (list []crm.CrmProductGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmProductGroup{})
    var crmProductGroups []crm.CrmProductGroup
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
	
	err = db.Find(&crmProductGroups).Error
	return  crmProductGroups, total, err
}
