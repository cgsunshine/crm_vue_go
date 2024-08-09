package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmBusinessOpportunityFileUploadAndDownloadsService struct {
}

// CreateCrmBusinessOpportunityFileUploadAndDownloads 创建上传商机文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityFileUploadAndDownloadsService *CrmBusinessOpportunityFileUploadAndDownloadsService) CreateCrmBusinessOpportunityFileUploadAndDownloads(crmBusinessOpportunityFileUploadAndDownloads *crm.CrmBusinessOpportunityFileUploadAndDownloads) (err error) {
	err = global.GVA_DB.Create(crmBusinessOpportunityFileUploadAndDownloads).Error
	return err
}

// DeleteCrmBusinessOpportunityFileUploadAndDownloads 删除上传商机文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityFileUploadAndDownloadsService *CrmBusinessOpportunityFileUploadAndDownloadsService)DeleteCrmBusinessOpportunityFileUploadAndDownloads(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmBusinessOpportunityFileUploadAndDownloads{},"id = ?",ID).Error
	return err
}

// DeleteCrmBusinessOpportunityFileUploadAndDownloadsByIds 批量删除上传商机文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityFileUploadAndDownloadsService *CrmBusinessOpportunityFileUploadAndDownloadsService)DeleteCrmBusinessOpportunityFileUploadAndDownloadsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmBusinessOpportunityFileUploadAndDownloads{},"id in ?",IDs).Error
	return err
}

// UpdateCrmBusinessOpportunityFileUploadAndDownloads 更新上传商机文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityFileUploadAndDownloadsService *CrmBusinessOpportunityFileUploadAndDownloadsService)UpdateCrmBusinessOpportunityFileUploadAndDownloads(crmBusinessOpportunityFileUploadAndDownloads crm.CrmBusinessOpportunityFileUploadAndDownloads) (err error) {
	err = global.GVA_DB.Save(&crmBusinessOpportunityFileUploadAndDownloads).Error
	return err
}

// GetCrmBusinessOpportunityFileUploadAndDownloads 根据ID获取上传商机文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityFileUploadAndDownloadsService *CrmBusinessOpportunityFileUploadAndDownloadsService)GetCrmBusinessOpportunityFileUploadAndDownloads(ID string) (crmBusinessOpportunityFileUploadAndDownloads crm.CrmBusinessOpportunityFileUploadAndDownloads, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmBusinessOpportunityFileUploadAndDownloads).Error
	return
}

// GetCrmBusinessOpportunityFileUploadAndDownloadsInfoList 分页获取上传商机文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityFileUploadAndDownloadsService *CrmBusinessOpportunityFileUploadAndDownloadsService)GetCrmBusinessOpportunityFileUploadAndDownloadsInfoList(info crmReq.CrmBusinessOpportunityFileUploadAndDownloadsSearch) (list []crm.CrmBusinessOpportunityFileUploadAndDownloads, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmBusinessOpportunityFileUploadAndDownloads{})
    var crmBusinessOpportunityFileUploadAndDownloadss []crm.CrmBusinessOpportunityFileUploadAndDownloads
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
	
	err = db.Find(&crmBusinessOpportunityFileUploadAndDownloadss).Error
	return  crmBusinessOpportunityFileUploadAndDownloadss, total, err
}
