package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmCommissionRebateService struct {
}

// CreateCrmCommissionRebate 创建返佣管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService) CreateCrmCommissionRebate(crmCommissionRebate *crm.CrmCommissionRebate) (err error) {
	err = global.GVA_DB.Create(crmCommissionRebate).Error
	return err
}

// DeleteCrmCommissionRebate 删除返佣管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService)DeleteCrmCommissionRebate(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmCommissionRebate{},"id = ?",ID).Error
	return err
}

// DeleteCrmCommissionRebateByIds 批量删除返佣管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService)DeleteCrmCommissionRebateByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmCommissionRebate{},"id in ?",IDs).Error
	return err
}

// UpdateCrmCommissionRebate 更新返佣管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService)UpdateCrmCommissionRebate(crmCommissionRebate crm.CrmCommissionRebate) (err error) {
	err = global.GVA_DB.Save(&crmCommissionRebate).Error
	return err
}

// GetCrmCommissionRebate 根据ID获取返佣管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService)GetCrmCommissionRebate(ID string) (crmCommissionRebate crm.CrmCommissionRebate, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmCommissionRebate).Error
	return
}

// GetCrmCommissionRebateInfoList 分页获取返佣管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService)GetCrmCommissionRebateInfoList(info crmReq.CrmCommissionRebateSearch) (list []crm.CrmCommissionRebate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmCommissionRebate{})
    var crmCommissionRebates []crm.CrmCommissionRebate
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
	
	err = db.Find(&crmCommissionRebates).Error
	return  crmCommissionRebates, total, err
}
