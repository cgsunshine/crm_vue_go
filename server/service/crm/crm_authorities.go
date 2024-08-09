package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmAuthoritiesService struct {
}

// CreateCrmAuthorities 创建数据权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmAuthoritiesService *CrmAuthoritiesService) CreateCrmAuthorities(crmAuthorities *crm.CrmAuthorities) (err error) {
	err = global.GVA_DB.Create(crmAuthorities).Error
	return err
}

// DeleteCrmAuthorities 删除数据权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmAuthoritiesService *CrmAuthoritiesService) DeleteCrmAuthorities(ID string) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&crm.CrmAuthorities{}, "id = ?", ID).Error
	return err
}

// DeleteCrmAuthoritiesByIds 批量删除数据权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmAuthoritiesService *CrmAuthoritiesService) DeleteCrmAuthoritiesByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmAuthorities{}, "id in ?", IDs).Error
	return err
}

// UpdateCrmAuthorities 更新数据权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmAuthoritiesService *CrmAuthoritiesService) UpdateCrmAuthorities(crmAuthorities crm.CrmAuthorities) (err error) {
	err = global.GVA_DB.Save(&crmAuthorities).Error
	return err
}

// GetCrmAuthorities 根据ID获取数据权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmAuthoritiesService *CrmAuthoritiesService) GetCrmAuthorities(ID string) (crmAuthorities crm.CrmAuthorities, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmAuthorities).Error
	return
}

// GetCrmAuthoritiesInfoList 分页获取数据权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmAuthoritiesService *CrmAuthoritiesService) GetCrmAuthoritiesInfoList(info crmReq.CrmAuthoritiesSearch) (list []crm.CrmAuthorities, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmAuthorities{})
	var crmAuthoritiess []crm.CrmAuthorities
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.AuthorityId != nil {
		db = db.Where("authority_id = ?", info.AuthorityId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&crmAuthoritiess).Error
	return crmAuthoritiess, total, err
}
