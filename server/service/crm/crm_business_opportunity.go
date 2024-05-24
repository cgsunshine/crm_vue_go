package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmBusinessOpportunityService struct {
}

// CreateCrmBusinessOpportunity 创建crmBusinessOpportunity表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityService *CrmBusinessOpportunityService) CreateCrmBusinessOpportunity(crmBusinessOpportunity *crm.CrmBusinessOpportunity) (err error) {
	err = global.GVA_DB.Create(crmBusinessOpportunity).Error
	return err
}

// DeleteCrmBusinessOpportunity 删除crmBusinessOpportunity表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityService *CrmBusinessOpportunityService)DeleteCrmBusinessOpportunity(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmBusinessOpportunity{},"id = ?",ID).Error
	return err
}

// DeleteCrmBusinessOpportunityByIds 批量删除crmBusinessOpportunity表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityService *CrmBusinessOpportunityService)DeleteCrmBusinessOpportunityByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmBusinessOpportunity{},"id in ?",IDs).Error
	return err
}

// UpdateCrmBusinessOpportunity 更新crmBusinessOpportunity表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityService *CrmBusinessOpportunityService)UpdateCrmBusinessOpportunity(crmBusinessOpportunity crm.CrmBusinessOpportunity) (err error) {
	err = global.GVA_DB.Save(&crmBusinessOpportunity).Error
	return err
}

// GetCrmBusinessOpportunity 根据ID获取crmBusinessOpportunity表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityService *CrmBusinessOpportunityService)GetCrmBusinessOpportunity(ID string) (crmBusinessOpportunity crm.CrmBusinessOpportunity, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmBusinessOpportunity).Error
	return
}

// GetCrmBusinessOpportunityInfoList 分页获取crmBusinessOpportunity表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityService *CrmBusinessOpportunityService)GetCrmBusinessOpportunityInfoList(info crmReq.CrmBusinessOpportunitySearch) (list []crm.CrmBusinessOpportunity, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmBusinessOpportunity{})
    var crmBusinessOpportunitys []crm.CrmBusinessOpportunity
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.BusinessOpportunityName != "" {
        db = db.Where("business_opportunity_name LIKE ?","%"+ info.BusinessOpportunityName+"%")
    }
    if info.CustomerId != nil {
        db = db.Where("customer_id = ?",info.CustomerId)
    }
        if info.StartInputTime != nil && info.EndInputTime != nil {
            db = db.Where("input_time BETWEEN ? AND ? ",info.StartInputTime,info.EndInputTime)
        }
        if info.StartOfferValidityPeriod != nil && info.EndOfferValidityPeriod != nil {
            db = db.Where("offer_validity_period BETWEEN ? AND ? ",info.StartOfferValidityPeriod,info.EndOfferValidityPeriod)
        }
    if info.Price != nil {
        db = db.Where("price <> ?",info.Price)
    }
    if info.ProductId != nil {
        db = db.Where("product_id = ?",info.ProductId)
    }
    if info.Status != "" {
        db = db.Where("status = ?",info.Status)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?",info.UserId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&crmBusinessOpportunitys).Error
	return  crmBusinessOpportunitys, total, err
}
