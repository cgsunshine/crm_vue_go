// 自动生成模板CrmBusinessOpportunity
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmBusinessOpportunity表 结构体  CrmBusinessOpportunity
type CrmPageBusinessOpportunity struct {
	global.GVA_MODEL
	Amount                  *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:商机金额;size:22;"`                                                       //商机金额
	BusinessOpportunityName string     `json:"businessOpportunityName" form:"businessOpportunityName" gorm:"column:business_opportunity_name;comment:商机名称;size:191;"` //商机名称
	CustomerId              *int       `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;"`                                                  //客户ID
	Description             string     `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`                                         //备注
	InputTime               *time.Time `json:"inputTime" form:"inputTime" gorm:"column:input_time;comment:商机录入时间;"`                                                   //商机录入时间
	OfferValidityPeriod     *time.Time `json:"offerValidityPeriod" form:"offerValidityPeriod" gorm:"column:offer_validity_period;comment:报价有效期;"`                     //报价有效期
	Price                   *float64   `json:"price" form:"price" gorm:"column:price;comment:商机金额;"`                                                                  //商机金额
	ProductId               *int       `json:"productId" form:"productId" gorm:"column:product_id;comment:产品id;"`                                                     //产品id
	Status                  string     `json:"status" form:"status" gorm:"column:status;comment:商机状态;size:191;"`                                                      //商机状态
	UserId                  *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:员工id;"`                                                              //员工id
	ReviewStatus            string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`                                   //审核状态
	Currency                string     `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:11;"`                                                   //币种
	//以下是联表字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	//crm_customers 表
	CustomerName string `json:"customerName" form:"customerName" gorm:"column:customer_name;comment:客户名;size:191;"` //客户名
	//crm_product 表
	ProductName string `json:"productName" form:"productName" gorm:"column:product_name;comment:产品名称;size:191;"` //产品名称

	BusinessOpportunityProduct []CrmBusinessOpportunityProduct `gorm:"foreignKey:BusinessOpportunityId"`
	//OrderProducts []CrmOrderProduct `gorm:"foreignKey:OrderId"`
}
