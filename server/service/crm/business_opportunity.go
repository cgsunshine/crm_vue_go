package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
)

// GetPageCrmBusinessOpportunityInfoList 分页获取crmBusinessOpportunity表记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityService *CrmBusinessOpportunityService) GetPageCrmBusinessOpportunityInfoList(info crmReq.CrmBusinessOpportunitySearch) (list []crm.CrmPageBusinessOpportunity, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&crm.CrmBusinessOpportunity{})
	var crmBusinessOpportunitys []crm.CrmPageBusinessOpportunity
	// 如果有条件搜索 下方会自动创建搜索语句
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(crmBusinessOpportunityService.SplicingQueryConditions("created_at BETWEEN ? AND ?"), info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Amount != nil {
		db = db.Where(crmBusinessOpportunityService.SplicingQueryConditions("amount = ?"), info.Amount)
	}
	if info.BusinessOpportunityName != "" {
		db = db.Where(crmBusinessOpportunityService.SplicingQueryConditions("business_opportunity_name LIKE ?"), "%"+info.BusinessOpportunityName+"%")
	}
	if info.CustomerId != nil {
		db = db.Where(crmBusinessOpportunityService.SplicingQueryConditions("customer_id = ?"), info.CustomerId)
	}
	if info.StartInputTime != nil && info.EndInputTime != nil {
		db = db.Where(crmBusinessOpportunityService.SplicingQueryConditions("input_time BETWEEN ? AND ? "), info.StartInputTime, info.EndInputTime)
	}
	//if info.OfferValidityPeriod != nil {
	//	db = db.Where("offer_validity_period = ?", info.OfferValidityPeriod)
	//}
	if info.ProductId != nil {
		db = db.Where(crmBusinessOpportunityService.SplicingQueryConditions("product_id = ?"), info.ProductId)
	}
	if info.Status != "" {
		db = db.Where(crmBusinessOpportunityService.SplicingQueryConditions("status = ?"), info.Status)
	}
	if info.UserId != nil {
		db = db.Where(crmBusinessOpportunityService.SplicingQueryConditions("user_id = ?"), info.UserId)
	}

	if info.ReviewStatus != "" {
		db = db.Where(crmBusinessOpportunityService.SplicingQueryConditions("review_status = ?"), info.ReviewStatus)
	}
	if info.BusinessOpportunityNumber != "" {
		db = db.Where(crmBusinessOpportunityService.SplicingQueryConditions("business_opportunity_number = ?"), info.BusinessOpportunityNumber)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Select("crm_business_opportunity.*,sys_users.username,crm_product.product_name,crm_customers.customer_name").
		Joins("LEFT JOIN crm_customers ON crm_business_opportunity.customer_id = crm_customers.id").
		Joins("LEFT JOIN sys_users ON crm_customers.user_id = sys_users.id").
		Joins("LEFT JOIN crm_product ON crm_business_opportunity.product_id = crm_product.id").
		Order("crm_business_opportunity.created_at DESC").
		Preload("BusinessOpportunityProduct.Product").
		Find(&crmBusinessOpportunitys).Error
	return crmBusinessOpportunitys, total, err
}

// GetCrmBusinessOpportunity 根据ID获取商机管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityService *CrmBusinessOpportunityService) GetCrmPageBusinessOpportunity(ID string) (crmBusinessOpportunity crm.CrmPageBusinessOpportunity, err error) {
	err = global.GVA_DB.Model(&crm.CrmBusinessOpportunity{}).Where("crm_business_opportunity.id = ?", ID).
		Select("crm_business_opportunity.*,sys_users.username,crm_product.product_name,crm_customers.customer_name").
		Joins("LEFT JOIN crm_customers ON crm_business_opportunity.customer_id = crm_customers.id").
		Joins("LEFT JOIN sys_users ON crm_customers.user_id = sys_users.id").
		Joins("LEFT JOIN crm_product ON crm_business_opportunity.product_id = crm_product.id").
		Preload("BusinessOpportunityProduct.Product").
		First(&crmBusinessOpportunity).Error
	return
}

// UpdApprovalStatus 修改审批状态
// Author [piexlmax](https://github.com/piexlmax)
func (crmBusinessOpportunityService *CrmBusinessOpportunityService) UpdApprovalStatus(ID *int, data map[string]interface{}) (err error) {
	db := global.GVA_DB.Model(&crm.CrmBusinessOpportunity{})
	err = db.Where("id = ?", ID).Updates(data).Error
	return
}

// SplicingQueryConditions 拼接条件
func (crmBusinessOpportunityService *CrmBusinessOpportunityService) SplicingQueryConditions(condition string) string {
	return "crm_business_opportunity." + condition
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetOrderProducts
//@description: 设置一个商机的产品
//@param: id uint, authorityIds []string
//@return: err error

func (crmBusinessOpportunityService *CrmBusinessOpportunityService) SetBusinessOpportunityProducts(id *int, productIds []*int) (err error) {
	db := global.GVA_DB.Model(&crm.CrmBusinessOpportunityProduct{})
	err = db.Delete(&[]crm.CrmBusinessOpportunityProduct{}, "business_opportunity_id = ?", id).Error
	if err != nil {
		return
	}

	var orderProduct []crm.CrmBusinessOpportunityProduct
	for _, v := range productIds {
		orderProduct = append(orderProduct, crm.CrmBusinessOpportunityProduct{
			BusinessOpportunityId: id, ProductId: v,
		})
	}
	err = db.Create(&orderProduct).Error
	if err != nil {
		return
	}
	return
}
