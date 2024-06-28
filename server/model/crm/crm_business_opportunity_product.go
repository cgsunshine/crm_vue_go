// 自动生成模板CrmBusinessOpportunityProduct
package crm

// crmBusinessOpportunityProduct表 结构体  CrmBusinessOpportunityProduct
type CrmBusinessOpportunityProduct struct {
	BusinessOpportunityId uint `json:"businessOpportunityId" form:"businessOpportunityId" gorm:"column:business_opportunity_id;comment:;"` //businessOpportunityId字段
	ProductId             uint `json:"productId" form:"productId" gorm:"column:product_id;comment:;"`                                      //productId字段
}

// TableName crmBusinessOpportunityProduct表 CrmBusinessOpportunityProduct自定义表名 crm_business_opportunity_product
func (CrmBusinessOpportunityProduct) TableName() string {
	return "crm_business_opportunity_product"
}
