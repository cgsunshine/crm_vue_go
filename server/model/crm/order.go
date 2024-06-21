// 自动生成模板CrmOrder
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// crmOrder表 结构体  CrmOrder
type CrmPageOrder struct {
	global.GVA_MODEL
	Currency                 string   `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:10;"`                                                               //币种
	CustomerId               *int     `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;"`                                                              //客户ID
	Description              string   `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`                                                     //备注
	DiscountRate             *float64 `json:"discountRate" form:"discountRate" gorm:"column:discount_rate;comment:折扣率;"`                                                         //折扣率
	Price                    *float64 `json:"price" form:"price" gorm:"column:price;comment:产品原价;"`                                                                              //产品原价
	ProductId                string   `json:"productId" form:"productId" gorm:"column:product_id;comment:产品id;size:191;"`                                                        //产品id
	SalesPrice               *float64 `json:"salesPrice" form:"salesPrice" gorm:"column:sales_price;comment:产品销售价格;"`                                                            //产品销售价格
	UserId                   *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;"`                                                                     //管理ID 销售代表
	SupplementaryInformation string   `json:"supplementaryInformation" form:"supplementaryInformation" gorm:"supplementary_information:description;comment:补充信息，售后专用;size:191;"` //补充信息，售后专用
	OrderName                string   `json:"orderName" form:"orderName" gorm:"column:order_name;comment:订单名称;size:10;" binding:"required"`                                      //订单名称
	ReviewStatus             string   `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`                                               //审核状态
	//以下是联表查询字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	//crm_customers 表
	CustomerName string `json:"customerName" form:"customerName" gorm:"column:customer_name;comment:客户名;size:191;"` //客户名
	//crm_product 表
	ProductName string `json:"productName" form:"productName" gorm:"column:product_name;comment:产品名称;size:191;"` //产品名称
	//crm_contract 表
	ContractName string `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合同名称;size:191;" binding:"required"` //合同名称
}
