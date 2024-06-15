package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmStatementAccountFileService struct {
}

// CreateCrmStatementAccountFile 创建对账单文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountFileService *CrmStatementAccountFileService) CreateCrmStatementAccountFile(crmStatementAccountFile *crm.CrmStatementAccountFile) (err error) {
	err = global.GVA_DB.Create(crmStatementAccountFile).Error
	return err
}

// DeleteCrmStatementAccountFile 删除对账单文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountFileService *CrmStatementAccountFileService)DeleteCrmStatementAccountFile(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmStatementAccountFile{},"id = ?",ID).Error
	return err
}

// DeleteCrmStatementAccountFileByIds 批量删除对账单文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountFileService *CrmStatementAccountFileService)DeleteCrmStatementAccountFileByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmStatementAccountFile{},"id in ?",IDs).Error
	return err
}

// UpdateCrmStatementAccountFile 更新对账单文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountFileService *CrmStatementAccountFileService)UpdateCrmStatementAccountFile(crmStatementAccountFile crm.CrmStatementAccountFile) (err error) {
	err = global.GVA_DB.Save(&crmStatementAccountFile).Error
	return err
}

// GetCrmStatementAccountFile 根据ID获取对账单文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountFileService *CrmStatementAccountFileService)GetCrmStatementAccountFile(ID string) (crmStatementAccountFile crm.CrmStatementAccountFile, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmStatementAccountFile).Error
	return
}

// GetCrmStatementAccountFileInfoList 分页获取对账单文件记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountFileService *CrmStatementAccountFileService)GetCrmStatementAccountFileInfoList(info crmReq.CrmStatementAccountFileSearch) (list []crm.CrmStatementAccountFile, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmStatementAccountFile{})
    var crmStatementAccountFiles []crm.CrmStatementAccountFile
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
	
	err = db.Find(&crmStatementAccountFiles).Error
	return  crmStatementAccountFiles, total, err
}
