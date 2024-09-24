// 自动生成模板CrmDeposits
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 押金管理 结构体  CrmDeposits
type CrmDeposits struct {
	global.GVA_MODEL
	Amount          *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:押金金额;size:10;"`                                      //押金金额
	Currency        string     `json:"currency" form:"currency" gorm:"column:currency;comment:货币类型;"`                                        //货币类型
	CustomerId      *int       `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户id;size:10;"`                         //客户id
	Description     string     `json:"description" form:"description" gorm:"column:description;comment:描述;"`                                 //描述
	OrderId         *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id;size:10;"`                                  //订单id
	PaymentDate     *time.Time `json:"paymentDate" form:"paymentDate" gorm:"column:payment_date;comment:支付日期;"`                              //支付日期
	RefundDate      *time.Time `json:"refundDate" form:"refundDate" gorm:"column:refund_date;comment:退款日期;"`                                 //退款日期
	DepositsStatus  string     `json:"depositsStatus" form:"depositsStatus" gorm:"column:deposits_status;comment:押金状态 已支付 已退款 待处理;size:50;"` //押金状态 已支付 已退款 待处理
	ContractId      *int       `json:"contractId" form:"contractId" gorm:"column:contract_id;comment:合同ID;"`                                 //合同ID
	ReviewStatus    string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`                  //审核状态
	DepositsName    string     `json:"depositsName" form:"depositsName" gorm:"column:deposits_name;comment:押金名称;size:191;"`                  //押金名称
	RefundStatus    string     `json:"refundStatus" form:"refundStatus" gorm:"column:refund_status;comment:退款状态;size:191;"`                  //退款状态
	DepositsVoucher string     `json:"depositsVoucher" form:"depositsVoucher" gorm:"column:deposits_voucher;comment:押金凭证;size:191;"`         //押金凭证
	RefundVoucher   string     `json:"refundVoucher" form:"refundVoucher" gorm:"column:refund_voucher;comment:退款凭证;size:191;"`               //退款凭证
	UserId          *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;size:19;"`                                //管理ID 销售代表
}

// TableName 押金管理 CrmDeposits自定义表名 crm_deposits
func (CrmDeposits) TableName() string {
	return "crm_deposits"
}
