// 自动生成模板CrmBusinessOpportunityProduct
package crm

// crmBusinessOpportunityProduct表 结构体  CrmBusinessOpportunityProduct
type CrmBusinessOpportunityProduct struct {
	BusinessOpportunityId *int   `json:"businessOpportunityId" form:"businessOpportunityId" gorm:"column:business_opportunity_id;comment:;"` //businessOpportunityId字段
	ProductId             *int   `json:"productId" form:"productId" gorm:"column:product_id;comment:;"`                                      //productId字段
	Quantity              *int   `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;size:10;"`                                  //产品数量
	Specifications        string `json:"specifications" form:"specifications" gorm:"column:specifications;comment:;size:10;"`                //产品规格
	BusinessOpportunity   CrmBusinessOpportunity
	Product               CrmProduct
}

// TableName crmBusinessOpportunityProduct表 CrmBusinessOpportunityProduct自定义表名 crm_business_opportunity_product
func (CrmBusinessOpportunityProduct) TableName() string {
	return "crm_business_opportunity_product"
}
