package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmBusinessOpportunitySearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Amount                   *float64   `json:"amount" form:"amount" `
	BusinessOpportunityName  string     `json:"businessOpportunityName" form:"businessOpportunityName" `
	CustomerId               *int       `json:"customerId" form:"customerId" `
	StartInputTime           *time.Time `json:"startInputTime" form:"startInputTime"`
	EndInputTime             *time.Time `json:"endInputTime" form:"endInputTime"`
	StartOfferValidityPeriod *time.Time `json:"startOfferValidityPeriod" form:"startOfferValidityPeriod"`
	EndOfferValidityPeriod   *time.Time `json:"endOfferValidityPeriod" form:"endOfferValidityPeriod"`
	ProductId                *int       `json:"productId" form:"productId" `
	Status                   string     `json:"status" form:"status" `
	UserId                   *int       `json:"userId" form:"userId" `
	ReviewStatus             string     `json:"reviewStatus" form:"reviewStatus"` //审核状态
	request.PageInfo
}

// 商机管理 结构体  CrmBusinessOpportunity
type CrmReqBusinessOpportunity struct {
	global.GVA_MODEL
	Amount                  *float64           `json:"amount" form:"amount" gorm:"column:amount;comment:商机金额;size:22;" binding:"required"`                                                       //商机金额
	BusinessOpportunityName string             `json:"businessOpportunityName" form:"businessOpportunityName" gorm:"column:business_opportunity_name;comment:商机名称;size:191;" binding:"required"` //商机名称
	CustomerId              *int               `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;size:19;" binding:"required"`                                          //客户ID
	Description             string             `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`                                                            //备注
	InputTime               *time.Time         `json:"inputTime" form:"inputTime" gorm:"column:input_time;comment:商机录入时间;"`                                                                      //商机录入时间
	OfferValidityPeriod     *time.Time         `json:"offerValidityPeriod" form:"offerValidityPeriod" gorm:"column:offer_validity_period;comment:报价有效期;"`                                        //报价有效期
	ProductIds              []uint             `json:"productIds" form:"productIds" gorm:"column:product_id;comment:产品id;size:19;"`                                                              //产品id
	Status                  string             `json:"status" form:"status" gorm:"column:status;comment:商机状态;size:191;"`                                                                         //商机状态
	UserId                  *int               `json:"userId" form:"userId" gorm:"column:user_id;comment:员工id;size:19;"`                                                                         //员工id
	ReviewStatus            string             `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`                                                      //审核状态
	Currency                string             `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:11;"`
	OrderProduct            []*CrmOrderProduct `json:"orderProduct" form:"orderProduct"` //产品规格信息
}

type CrmOrderProduct struct {
	OrderId        *int   `json:"orderId" form:"orderId" gorm:"column:order_id;comment:;size:10;"`                     //orderId字段
	ProductId      *int   `json:"productId" form:"productId" gorm:"column:product_id;comment:;size:10;"`               //productId字段
	Quantity       *int   `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;size:10;"`                   //产品数量
	Specifications string `json:"specifications" form:"specifications" gorm:"column:specifications;comment:;size:10;"` //产品规格
}
