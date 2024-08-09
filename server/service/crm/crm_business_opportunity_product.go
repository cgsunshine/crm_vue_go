package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

type CrmBusinessOpportunityProductService struct {
}

// CreateCrmBusinessOpportunityProduct 创建crmBusinessOpportunityProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityProductService *CrmBusinessOpportunityProductService) CreateCrmBusinessOpportunityProduct(crmBusinessOpportunityProduct *crm.CrmBusinessOpportunityProduct) (err error) {
	err = global.GVA_DB.Create(crmBusinessOpportunityProduct).Error
	return err
}

// DeleteCrmBusinessOpportunityProduct 删除crmBusinessOpportunityProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityProductService *CrmBusinessOpportunityProductService)DeleteCrmBusinessOpportunityProduct(ID string) (err error) {
	err = global.GVA_DB.Delete(&crm.CrmBusinessOpportunityProduct{},"id = ?",ID).Error
	return err
}

// DeleteCrmBusinessOpportunityProductByIds 批量删除crmBusinessOpportunityProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityProductService *CrmBusinessOpportunityProductService)DeleteCrmBusinessOpportunityProductByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]crm.CrmBusinessOpportunityProduct{},"id in ?",IDs).Error
	return err
}

// UpdateCrmBusinessOpportunityProduct 更新crmBusinessOpportunityProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityProductService *CrmBusinessOpportunityProductService)UpdateCrmBusinessOpportunityProduct(crmBusinessOpportunityProduct crm.CrmBusinessOpportunityProduct) (err error) {
	err = global.GVA_DB.Save(&crmBusinessOpportunityProduct).Error
	return err
}

// GetCrmBusinessOpportunityProduct 根据ID获取crmBusinessOpportunityProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityProductService *CrmBusinessOpportunityProductService)GetCrmBusinessOpportunityProduct(ID string) (crmBusinessOpportunityProduct crm.CrmBusinessOpportunityProduct, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&crmBusinessOpportunityProduct).Error
	return
}

// GetCrmBusinessOpportunityProductInfoList 分页获取crmBusinessOpportunityProduct表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityProductService *CrmBusinessOpportunityProductService)GetCrmBusinessOpportunityProductInfoList(info crmReq.CrmBusinessOpportunityProductSearch) (list []crm.CrmBusinessOpportunityProduct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&crm.CrmBusinessOpportunityProduct{})
    var crmBusinessOpportunityProducts []crm.CrmBusinessOpportunityProduct
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
	
	err = db.Find(&crmBusinessOpportunityProducts).Error
	return  crmBusinessOpportunityProducts, total, err
}
