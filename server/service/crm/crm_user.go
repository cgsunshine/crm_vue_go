package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmUserService struct {
}

// CreateCrmUser 创建crmUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmUserService *CrmUserService) CreateCrmUser(crmUser *crm.CrmUser) (err error) {
	err = global.GVA_DB.Create(crmUser).Error
	return err
}

// DeleteCrmUser 删除crmUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmUserService *CrmUserService)DeleteCrmUser(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmUser{},"id = ?",ID).Error
	return err
}

// DeleteCrmUserByIds 批量删除crmUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmUserService *CrmUserService)DeleteCrmUserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmUser{},"id in ?",IDs).Error
	return err
}

// UpdateCrmUser 更新crmUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmUserService *CrmUserService)UpdateCrmUser(crmUser crm.CrmUser) (err error) {
	err = global.GVA_DB.Save(&crmUser).Error
	return err
}

// GetCrmUser 根据ID获取crmUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmUserService *CrmUserService)GetCrmUser(ID string) (crmUser crm.CrmUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmUser).Error
	return
}

// GetCrmUserInfoList 分页获取crmUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmUserService *CrmUserService)GetCrmUserInfoList(info crmReq.CrmUserSearch) (list []crm.CrmUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmUser{})
    var crmUsers []crm.CrmUser
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
	
	err = db.Find(&crmUsers).Error
	return  crmUsers, total, err
}
