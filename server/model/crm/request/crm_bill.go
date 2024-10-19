package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmBillSearch struct {
	baseSearchReq
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	OrderId             *int       `json:"orderId" form:"orderId" `
	PaymentCollentionId *int       `json:"paymentCollentionId" form:"paymentCollentionId"` //回款单ID
	PaymentId           *int       `json:"PaymentId" form:"PaymentId"`                     //付款ID
	UserId              *int       `json:"userId" form:"userId" `                          //管理ID 销售代表
	PaymentStatus       string     `json:"paymentStatus" form:"paymentStatus"`             //付款状态
	CustomerId          *int       `json:"customerId" form:"customerId" `                  //客户ID
	StartExpirationTime *time.Time `json:"startExpirationTime" form:"startExpirationTime"` //最迟付款时间
	EndExpirationTime   *time.Time `json:"endExpirationTime" form:"endExpirationTime"`     //最迟付款时间
	BillNumber          string     `json:"billNumber" form:"billNumber" `                  //账单编号
	OrderNumber         string     `json:"orderNumber" form:"orderNumber"`                 //订单编号
	request.PageInfo
}
