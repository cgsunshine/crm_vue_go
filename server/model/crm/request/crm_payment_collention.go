package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CrmPaymentCollentionSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	StartAuditingTime     *time.Time `json:"startAuditingTime" form:"startAuditingTime"`
	EndAuditingTime       *time.Time `json:"endAuditingTime" form:"endAuditingTime"`
	BillId                *int       `json:"billId" form:"billId" `
	Currency              string     `json:"currency" form:"currency" `
	CustomerId            *int       `json:"customerId" form:"customerId" `
	UserId                *int       `json:"userId" form:"userId" `
	PaymentCollentionName string     `json:"paymentCollentionName" form:"paymentCollentionName"`                                  //回款名称
	ReviewStatus          string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"` //审核状态
	request.PageInfo
}
