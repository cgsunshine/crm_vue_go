package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmContactFileService struct {
}

// CreateCrmContactFile 创建crmContactFile表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactFileService *CrmContactFileService) CreateCrmContactFile(crmContactFile *crm.CrmContactFile) (err error) {
	err = global.GVA_DB.Create(crmContactFile).Error
	return err
}

// DeleteCrmContactFile 删除crmContactFile表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactFileService *CrmContactFileService)DeleteCrmContactFile(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmContactFile{},"id = ?",ID).Error
	return err
}

// DeleteCrmContactFileByIds 批量删除crmContactFile表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactFileService *CrmContactFileService)DeleteCrmContactFileByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmContactFile{},"id in ?",IDs).Error
	return err
}

// UpdateCrmContactFile 更新crmContactFile表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactFileService *CrmContactFileService)UpdateCrmContactFile(crmContactFile crm.CrmContactFile) (err error) {
	err = global.GVA_DB.Save(&crmContactFile).Error
	return err
}

// GetCrmContactFile 根据ID获取crmContactFile表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactFileService *CrmContactFileService)GetCrmContactFile(ID string) (crmContactFile crm.CrmContactFile, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmContactFile).Error
	return
}

// GetCrmContactFileInfoList 分页获取crmContactFile表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmContactFileService *CrmContactFileService)GetCrmContactFileInfoList(info crmReq.CrmContactFileSearch) (list []crm.CrmContactFile, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmContactFile{})
    var crmContactFiles []crm.CrmContactFile
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ContactId != nil {
        db = db.Where("contact_id = ?",info.ContactId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmContactFiles).Error
	return  crmContactFiles, total, err
}
