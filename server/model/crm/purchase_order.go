// 自动生成模板CrmPurchaseOrder
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmPurchaseOrder表 结构体  CrmPurchaseOrder
type CrmPagePurchaseOrder struct {
	global.GVA_MODEL
	ContractId        *int       `json:"contractId" form:"contractId" gorm:"column:contract_id;comment:合同ID;"`                                 //合同ID
	ProductId         *int       `json:"productId" form:"productId" gorm:"column:product_id;comment:产品ID;"`                                    //产品ID
	Quantity          *int       `json:"quantity" form:"quantity" gorm:"column:quantity;comment:数量;"`                                          //数量
	Amount            *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`                                                //金额
	UserId            *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:负责人ID;"`                                            //负责人ID
	CreationTime      *time.Time `json:"creationTime" form:"creationTime" gorm:"column:creation_time;comment:创建时间;"`                           //创建时间
	ExpirationTime    *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:到期时间;"`                     //到期时间
	PurchaseOrderName string     `json:"purchaseOrderName" form:"purchaseOrderName" gorm:"column:purchase_order_name;comment:订购单名称;size:191;"` //订购单名称
	CurrencyId        *int       `json:"currencyId" form:"currencyId" gorm:"column:currency_id;comment:币种id;size:11;"`                         //币种ID
	//以下是联表查询字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	//crm_product 表
	ProductName string `json:"productName" form:"productName" gorm:"column:product_name;comment:产品名称;size:191;"` //产品名称
	//crm_currency 表 后面替换去掉注释
	CurrencyName string `json:"currencyName" form:"currencyName" gorm:"column:currency_name;comment:币种;size:10;"` //币种名称
}
