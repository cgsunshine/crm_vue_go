package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmStatementAccountService struct {
}

// CreateCrmStatementAccount 创建对账单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountService *CrmStatementAccountService) CreateCrmStatementAccount(crmStatementAccount *crm.CrmStatementAccount) (err error) {
	err = global.GVA_DB.Create(crmStatementAccount).Error
	return err
}

// DeleteCrmStatementAccount 删除对账单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountService *CrmStatementAccountService)DeleteCrmStatementAccount(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmStatementAccount{},"id = ?",ID).Error
	return err
}

// DeleteCrmStatementAccountByIds 批量删除对账单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountService *CrmStatementAccountService)DeleteCrmStatementAccountByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmStatementAccount{},"id in ?",IDs).Error
	return err
}

// UpdateCrmStatementAccount 更新对账单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountService *CrmStatementAccountService)UpdateCrmStatementAccount(crmStatementAccount crm.CrmStatementAccount) (err error) {
	err = global.GVA_DB.Save(&crmStatementAccount).Error
	return err
}

// GetCrmStatementAccount 根据ID获取对账单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountService *CrmStatementAccountService)GetCrmStatementAccount(ID string) (crmStatementAccount crm.CrmStatementAccount, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmStatementAccount).Error
	return
}

// GetCrmStatementAccountInfoList 分页获取对账单记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmStatementAccountService *CrmStatementAccountService)GetCrmStatementAccountInfoList(info crmReq.CrmStatementAccountSearch) (list []crm.CrmStatementAccount, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmStatementAccount{})
    var crmStatementAccounts []crm.CrmStatementAccount
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
	
	err = db.Find(&crmStatementAccounts).Error
	return  crmStatementAccounts, total, err
}
