package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmStatementAccountFileUploadAndDownloadsService struct {
}

// CreateCrmStatementAccountFileUploadAndDownloads 创建对账单上传文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountFileUploadAndDownloadsService *CrmStatementAccountFileUploadAndDownloadsService) CreateCrmStatementAccountFileUploadAndDownloads(crmStatementAccountFileUploadAndDownloads *crm.CrmStatementAccountFileUploadAndDownloads) (err error) {
	err = global.GVA_DB.Create(crmStatementAccountFileUploadAndDownloads).Error
	return err
}

// DeleteCrmStatementAccountFileUploadAndDownloads 删除对账单上传文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountFileUploadAndDownloadsService *CrmStatementAccountFileUploadAndDownloadsService)DeleteCrmStatementAccountFileUploadAndDownloads(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmStatementAccountFileUploadAndDownloads{},"id = ?",ID).Error
	return err
}

// DeleteCrmStatementAccountFileUploadAndDownloadsByIds 批量删除对账单上传文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountFileUploadAndDownloadsService *CrmStatementAccountFileUploadAndDownloadsService)DeleteCrmStatementAccountFileUploadAndDownloadsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmStatementAccountFileUploadAndDownloads{},"id in ?",IDs).Error
	return err
}

// UpdateCrmStatementAccountFileUploadAndDownloads 更新对账单上传文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountFileUploadAndDownloadsService *CrmStatementAccountFileUploadAndDownloadsService)UpdateCrmStatementAccountFileUploadAndDownloads(crmStatementAccountFileUploadAndDownloads crm.CrmStatementAccountFileUploadAndDownloads) (err error) {
	err = global.GVA_DB.Save(&crmStatementAccountFileUploadAndDownloads).Error
	return err
}

// GetCrmStatementAccountFileUploadAndDownloads 根据ID获取对账单上传文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountFileUploadAndDownloadsService *CrmStatementAccountFileUploadAndDownloadsService)GetCrmStatementAccountFileUploadAndDownloads(ID string) (crmStatementAccountFileUploadAndDownloads crm.CrmStatementAccountFileUploadAndDownloads, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmStatementAccountFileUploadAndDownloads).Error
	return
}

// GetCrmStatementAccountFileUploadAndDownloadsInfoList 分页获取对账单上传文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountFileUploadAndDownloadsService *CrmStatementAccountFileUploadAndDownloadsService)GetCrmStatementAccountFileUploadAndDownloadsInfoList(info crmReq.CrmStatementAccountFileUploadAndDownloadsSearch) (list []crm.CrmStatementAccountFileUploadAndDownloads, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmStatementAccountFileUploadAndDownloads{})
    var crmStatementAccountFileUploadAndDownloadss []crm.CrmStatementAccountFileUploadAndDownloads
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.StatementAccountId != nil {
        db = db.Where("statement_account_id = ?",info.StatementAccountId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmStatementAccountFileUploadAndDownloadss).Error
	return  crmStatementAccountFileUploadAndDownloadss, total, err
}
