package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmCommissionRebateService struct {
}

// CreateCrmCommissionRebate 创建crmCommissionRebate表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService) CreateCrmCommissionRebate(crmCommissionRebate *crm.CrmCommissionRebate) (err error) {
	err = global.GVA_DB.Create(crmCommissionRebate).Error
	return err
}

// DeleteCrmCommissionRebate 删除crmCommissionRebate表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService)DeleteCrmCommissionRebate(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmCommissionRebate{},"id = ?",ID).Error
	return err
}

// DeleteCrmCommissionRebateByIds 批量删除crmCommissionRebate表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService)DeleteCrmCommissionRebateByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmCommissionRebate{},"id in ?",IDs).Error
	return err
}

// UpdateCrmCommissionRebate 更新crmCommissionRebate表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService)UpdateCrmCommissionRebate(crmCommissionRebate crm.CrmCommissionRebate) (err error) {
	err = global.GVA_DB.Save(&crmCommissionRebate).Error
	return err
}

// GetCrmCommissionRebate 根据ID获取crmCommissionRebate表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmCommissionRebateService *CrmCommissionRebateService)GetCrmCommissionRebate(ID string) (crmCommissionRebate crm.CrmCommissionRebate, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmCommissionRebate).Error
	return
}

// GetCrmCommissionRebateInfoList 分页获取crmCommissionRebate表记录
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
    if info.OrderId != nil {
        db = db.Where("order_id = ?",info.OrderId)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?",info.UserId)
    }
    if info.Payee != "" {
        db = db.Where("payee LIKE ?","%"+ info.Payee+"%")
    }
    if info.PaymentMethod != "" {
        db = db.Where("payment_method = ?",info.PaymentMethod)
    }
    if info.Account != "" {
        db = db.Where("account = ?",info.Account)
    }
    if info.Amount != nil {
        db = db.Where("amount = ?",info.Amount)
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
