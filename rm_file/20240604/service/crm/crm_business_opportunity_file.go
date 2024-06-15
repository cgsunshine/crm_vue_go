package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmBusinessOpportunityFileService struct {
}

// CreateCrmBusinessOpportunityFile 创建商机文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityFileService *CrmBusinessOpportunityFileService) CreateCrmBusinessOpportunityFile(crmBusinessOpportunityFile *crm.CrmBusinessOpportunityFile) (err error) {
	err = global.GVA_DB.Create(crmBusinessOpportunityFile).Error
	return err
}

// DeleteCrmBusinessOpportunityFile 删除商机文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityFileService *CrmBusinessOpportunityFileService)DeleteCrmBusinessOpportunityFile(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmBusinessOpportunityFile{},"id = ?",ID).Error
	return err
}

// DeleteCrmBusinessOpportunityFileByIds 批量删除商机文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityFileService *CrmBusinessOpportunityFileService)DeleteCrmBusinessOpportunityFileByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmBusinessOpportunityFile{},"id in ?",IDs).Error
	return err
}

// UpdateCrmBusinessOpportunityFile 更新商机文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityFileService *CrmBusinessOpportunityFileService)UpdateCrmBusinessOpportunityFile(crmBusinessOpportunityFile crm.CrmBusinessOpportunityFile) (err error) {
	err = global.GVA_DB.Save(&crmBusinessOpportunityFile).Error
	return err
}

// GetCrmBusinessOpportunityFile 根据ID获取商机文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityFileService *CrmBusinessOpportunityFileService)GetCrmBusinessOpportunityFile(ID string) (crmBusinessOpportunityFile crm.CrmBusinessOpportunityFile, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmBusinessOpportunityFile).Error
	return
}

// GetCrmBusinessOpportunityFileInfoList 分页获取商机文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityFileService *CrmBusinessOpportunityFileService)GetCrmBusinessOpportunityFileInfoList(info crmReq.CrmBusinessOpportunityFileSearch) (list []crm.CrmBusinessOpportunityFile, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmBusinessOpportunityFile{})
    var crmBusinessOpportunityFiles []crm.CrmBusinessOpportunityFile
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.BusinessOpportunityId != nil {
        db = db.Where("business_opportunity_id = ?",info.BusinessOpportunityId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmBusinessOpportunityFiles).Error
	return  crmBusinessOpportunityFiles, total, err
}
