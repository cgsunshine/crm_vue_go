// 自动生成模板CrmPaymentCollention
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// crmPaymentCollention表 结构体  CrmPaymentCollention
type CrmPaymentCollention struct {
 global.GVA_MODEL 
      BillId  *int `json:"billId" form:"billId" gorm:"column:bill_id;comment:账单ID;size:19;"`  //账单ID 
      CustomerId  *int `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;size:19;"`  //客户ID 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;size:19;"`  //管理ID 销售代表 
      Description  string `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`  //备注 
      Currency  string `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:10;"`  //币种 
      Proof  string `json:"proof" form:"proof" gorm:"column:proof;comment:凭证;size:191;"`  //凭证 
      AuditingStatus  string `json:"auditingStatus" form:"auditingStatus" gorm:"column:auditing_status;comment:审核状态;size:191;"`  //审核状态 
      ApprovedById  string `json:"approvedById" form:"approvedById" gorm:"column:approved_by_id;comment:审批人;size:191;"`  //审批人 
      AuditingTime  *time.Time `json:"auditingTime" form:"auditingTime" gorm:"column:auditing_time;comment:审核时间;"`  //审核时间 
}


// TableName crmPaymentCollention表 CrmPaymentCollention自定义表名 crm_payment_collention
func (CrmPaymentCollention) TableName() string {
  return "crm_payment_collention"
}

