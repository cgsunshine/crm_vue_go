// 自动生成模板CrmBusinessOpportunity
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// 商机管理 结构体  CrmBusinessOpportunity
type CrmBusinessOpportunity struct {
 global.GVA_MODEL 
      BusinessOpportunityName  string `json:"businessOpportunityName" form:"businessOpportunityName" gorm:"column:business_opportunity_name;comment:商机名称;size:191;"`  //商机名称 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`  //备注 
      CustomerId  *int `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;size:19;"`  //客户ID 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:员工id;size:19;"`  //员工id 
      ProductId  *int `json:"productId" form:"productId" gorm:"column:product_id;comment:产品id;size:19;"`  //产品id 
      Amount  *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:商机金额;size:22;"`  //商机金额 
      Status  string `json:"status" form:"status" gorm:"column:status;comment:商机状态;size:191;"`  //商机状态 
      InputTime  *time.Time `json:"inputTime" form:"inputTime" gorm:"column:input_time;comment:商机录入时间;"`  //商机录入时间 
      OfferValidityPeriod  *time.Time `json:"offerValidityPeriod" form:"offerValidityPeriod" gorm:"column:offer_validity_period;comment:报价有效期;"`  //报价有效期 
}


// TableName 商机管理 CrmBusinessOpportunity自定义表名 crm_business_opportunity
func (CrmBusinessOpportunity) TableName() string {
  return "crm_business_opportunity"
}

