package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmLoginLogService struct {
}

// CreateCrmLoginLog 创建登录日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmLoginLogService *CrmLoginLogService) CreateCrmLoginLog(crmLoginLog *crm.CrmLoginLog) (err error) {
	err = global.GVA_DB.Create(crmLoginLog).Error
	return err
}

// DeleteCrmLoginLog 删除登录日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmLoginLogService *CrmLoginLogService)DeleteCrmLoginLog(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmLoginLog{},"id = ?",ID).Error
	return err
}

// DeleteCrmLoginLogByIds 批量删除登录日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmLoginLogService *CrmLoginLogService)DeleteCrmLoginLogByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmLoginLog{},"id in ?",IDs).Error
	return err
}

// UpdateCrmLoginLog 更新登录日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmLoginLogService *CrmLoginLogService)UpdateCrmLoginLog(crmLoginLog crm.CrmLoginLog) (err error) {
	err = global.GVA_DB.Save(&crmLoginLog).Error
	return err
}

// GetCrmLoginLog 根据ID获取登录日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmLoginLogService *CrmLoginLogService)GetCrmLoginLog(ID string) (crmLoginLog crm.CrmLoginLog, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmLoginLog).Error
	return
}

// GetCrmLoginLogInfoList 分页获取登录日志记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmLoginLogService *CrmLoginLogService)GetCrmLoginLogInfoList(info crmReq.CrmLoginLogSearch) (list []crm.CrmLoginLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmLoginLog{})
    var crmLoginLogs []crm.CrmLoginLog
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
	
	err = db.Find(&crmLoginLogs).Error
	return  crmLoginLogs, total, err
}
