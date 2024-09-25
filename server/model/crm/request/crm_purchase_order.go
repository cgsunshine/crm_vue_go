package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmPurchaseOrderSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Amount              *float64   `json:"amount" form:"amount" `
	ContractId          *int       `json:"contractId" form:"contractId" `
	StartCreationTime   *time.Time `json:"startCreationTime" form:"startCreationTime"`
	EndCreationTime     *time.Time `json:"endCreationTime" form:"endCreationTime"`
	StartExpirationTime *time.Time `json:"startExpirationTime" form:"startExpirationTime"`
	EndExpirationTime   *time.Time `json:"endExpirationTime" form:"endExpirationTime"`
	ProductId           *int       `json:"productId" form:"productId" `
	Quantity            *int       `json:"quantity" form:"quantity" `
	UserId              *int       `json:"userId" form:"userId" `
	PurchaseOrderName   string     `json:"purchaseOrderName" form:"purchaseOrderName"` //订购单名称
	request.PageInfo
}

// 订购单管理 结构体  CrmPurchaseOrder
type CrmReqPurchaseOrder struct {
	global.GVA_MODEL
	Amount            *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`                                                //金额
	ContractId        *int       `json:"contractId" form:"contractId" gorm:"column:contract_id;comment:合同ID;"`                                 //合同ID
	CreationTime      *time.Time `json:"creationTime" form:"creationTime" gorm:"column:creation_time;comment:创建时间;"`                           //创建时间
	Description       string     `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`                        //备注
	ExpirationTime    *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:到期时间;"`                     //到期时间
	ProductId         *int       `json:"productId" form:"productId" gorm:"column:product_id;comment:产品ID;"`                                    //产品ID
	Quantity          *int       `json:"quantity" form:"quantity" gorm:"column:quantity;comment:数量;"`                                          //数量
	UserId            *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:负责人ID;"`                                            //负责人ID
	PurchaseOrderName string     `json:"purchaseOrderName" form:"purchaseOrderName" gorm:"column:purchase_order_name;comment:订购单名称;size:191;"` //订购单名称
	Currency          string     `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:11;"`                                  //币种
	ReviewStatus      string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`                  //审核状态

	ProductsInfo []CrmPurchaseOrderProduct `json:"productsInfo" form:"productsInfo"`
}

// CrmPurchaseOrderProduct 结构体  CrmBusinessOpportunityProduct
type CrmPurchaseOrderProduct struct {
	PurchaseOrderId *int   `json:"purchaseOrderId" form:"purchaseOrderId" gorm:"column:purchase_order_id;comment:;"`    //purchaseOrderId字段
	ProductId       *int   `json:"productId" form:"productId" gorm:"column:product_id;comment:;"`                       //productId字段
	Quantity        *int   `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;size:10;"`                   //产品数量
	Specifications  string `json:"specifications" form:"specifications" gorm:"column:specifications;comment:;size:10;"` //产品规格
}
