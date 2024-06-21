package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmConfigService struct {
}

// CreateCrmConfig 创建系统配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmConfigService *CrmConfigService) CreateCrmConfig(crmConfig *crm.CrmConfig) (err error) {
	err = global.GVA_DB.Create(crmConfig).Error
	return err
}

// DeleteCrmConfig 删除系统配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmConfigService *CrmConfigService)DeleteCrmConfig(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmConfig{},"id = ?",ID).Error
	return err
}

// DeleteCrmConfigByIds 批量删除系统配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmConfigService *CrmConfigService)DeleteCrmConfigByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmConfig{},"id in ?",IDs).Error
	return err
}

// UpdateCrmConfig 更新系统配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmConfigService *CrmConfigService)UpdateCrmConfig(crmConfig crm.CrmConfig) (err error) {
	err = global.GVA_DB.Save(&crmConfig).Error
	return err
}

// GetCrmConfig 根据ID获取系统配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmConfigService *CrmConfigService)GetCrmConfig(ID string) (crmConfig crm.CrmConfig, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmConfig).Error
	return
}

// GetCrmConfigInfoList 分页获取系统配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmConfigService *CrmConfigService)GetCrmConfigInfoList(info crmReq.CrmConfigSearch) (list []crm.CrmConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmConfig{})
    var crmConfigs []crm.CrmConfig
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
	
	err = db.Find(&crmConfigs).Error
	return  crmConfigs, total, err
}
