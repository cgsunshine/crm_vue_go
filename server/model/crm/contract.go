// 自动生成模板CrmContract
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// crmContract表 结构体  CrmContract
type CrmPageContract struct {
	global.GVA_MODEL
	ApplicationTime *time.Time `json:"applicationTime" form:"applicationTime" gorm:"column:application_time;comment:申请时间;"`       //申请时间
	ContractFile    string     `json:"contractFile" form:"contractFile" gorm:"column:contract_file;comment:合同文件路径;size:191;"`     //合同文件路径
	ContractName    string     `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合同名称;size:191;"`       //合同名称
	ContractStatus  string     `json:"contractStatus" form:"contractStatus" gorm:"column:contract_status;comment:合同状态;size:191;"` //合同状态
	ContractTypeId  *int       `json:"contractTypeId" form:"contractTypeId" gorm:"column:contract_type_id;comment:合同类型;"`         //合同类型
	CustomerId      *int       `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户ID;"`                      //客户ID
	Description     string     `json:"description" form:"description" gorm:"column:description;comment:备注;size:191;"`             //备注
	ExpirationTime  *time.Time `json:"expirationTime" form:"expirationTime" gorm:"column:expiration_time;comment:到期时间;"`          //到期时间
	OrderId         *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单ID;"`                               //订单ID
	ReviewResult    string     `json:"reviewResult" form:"reviewResult" gorm:"column:review_result;comment:审核结果;size:191;"`       //审核结果
	ReviewStatus    string     `json:"reviewStatus" form:"reviewStatus" gorm:"column:review_status;comment:审核状态;size:191;"`       //审核状态
	UserId          *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:管理ID 销售代表;"`                             //管理ID 销售代表
	ContractNumber  string     `json:"contractNumber" form:"contractNumber" gorm:"column:contract_number;comment:合同编号;size:191;"` //合同编号
	//以下是联表查询字段
	//sys_user 表
	Username string `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	//crm_customers 表
	CustomerName    string `json:"customerName" form:"customerName" gorm:"column:customer_name;comment:客户名;size:191;"`           //客户名
	CustomerCompany string `json:"customerCompany" form:"customerCompany" gorm:"column:customer_company;comment:客户公司;size:191;"` //客户公司
	CustomerAddress string `json:"customerAddress" form:"customerAddress" gorm:"column:customer_address;comment:客户地址;size:191;"` //客户地址
	//crm_order 表
	OrderName        string     `json:"orderName" form:"orderName" gorm:"column:order_name;comment:订单名称;size:10;" binding:"required"` //订单名称
	BillingStartTime *time.Time `json:"billingStartTime" form:"billingStartTime" gorm:"column:billing_start_time;comment:到期时间;"`      //计费开始时间
	BillingEndTime   *time.Time `json:"billingEndTime" form:"billingEndTime" gorm:"column:billing_end_time;comment:到期时间;"`            //计费结束时间
	ProductId        string     `json:"productId" form:"productId" gorm:"column:product_id;comment:产品id;size:191;"`                   //产品id
	DiscountRate     *float64   `json:"discountRate" form:"discountRate" gorm:"column:discount_rate;comment:折扣率;"`                    //折扣率
	OrderNumber      string     `json:"orderNumber" form:"orderNumber" gorm:"column:order_number;comment:订单编号;size:191;"`             //订单编号
	//crm_contract_type 表
	ContractTypeName string `json:"contractTypeName" form:"contractTypeName" gorm:"column:contract_type_name;comment:合同类型名称;size:191;"` //合同类型名称
	//crm_order 表
	Amount   *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`               //金额
	Currency string   `json:"currency" form:"currency" gorm:"column:currency;comment:币种;size:11;"` //币种
	ContractExpansion
}

// 合同补充字段
type ContractExpansion struct {
	//勾选项目开始
	//NewCustomer                       string `json:"newCustomer" form:"newCustomer" gorm:"column:new_customer;comment:新客户勾选项目;size:191;"`                                                                      //新客户
	//ExistingCustomerNewAccountNo      string `json:"existingCustomerNewAccountNo" form:"existingCustomerNewAccountNo" gorm:"column:existing_customer_new_account_no;comment:现有客户新账号;size:191;"`                //现有客户新账号
	//ExistingCustomerExistingAccountNo string `json:"existingCustomerExistingAccountNo" form:"existingCustomerExistingAccountNo" gorm:"column:existing_customer_existing_account_no;comment:现有客户老账号;size:191;"` //现有客户老账号
	//勾选项目结束
	BusinessRegistrationNo           string `json:"businessRegistrationNo" form:"businessRegistrationNo" gorm:"column:business_registration_no;comment:公司注册号;size:191;"`                                   //公司注册号
	InstallationAddress              string `json:"installationAddress" form:"installationAddress" gorm:"column:installation_address;comment:客户公司地址1;size:191;"`                                           //客户公司地址1
	InstallationAddress2             string `json:"installationAddress2" form:"installationAddress2" gorm:"column:installation_address2;comment:客户公司地址2;size:191;"`                                        //客户公司地址2
	InstallationAddress3             string `json:"installationAddress3" form:"installationAddress3" gorm:"column:installation_address3;comment:客户公司地址3;size:191;"`                                        //客户公司地址3
	CorrespondenceAddress            string `json:"correspondenceAddress" form:"correspondenceAddress" gorm:"column:correspondence_address;comment:通讯地址;size:191;"`                                        //通讯地址
	BillingAddress                   string `json:"billingAddress" form:"billingAddress" gorm:"column:billing_address;comment:账单寄送地址;size:191;"`                                                           //账单寄送地址
	ContactPerson                    string `json:"contactPerson" form:"contactPerson" gorm:"column:contact_person;comment:联系人;size:191;"`                                                                 //联系人
	Position                         string `json:"position" form:"position" gorm:"column:position;comment:职务;size:191;"`                                                                                  //职务
	CorrespondenceEmailAddress       string `json:"correspondenceEmailAddress" form:"correspondenceEmailAddress" gorm:"column:correspondence_email_address;comment:邮箱;size:191;"`                          //联系人邮箱
	PositionTelNo                    string `json:"positionTelNo" form:"positionTelNo" gorm:"column:position_tel_no;comment:电话;size:191;"`                                                                 //联系人电话
	PositionFaxNo                    string `json:"positionFaxNo" form:"positionFaxNo" gorm:"column:position_fax_no;comment:传真号;size:191;"`                                                                //联系人传真号
	MobileNo                         string `json:"mobileNo" form:"mobileNo" gorm:"column:mobile_no;comment:移动电话号码;size:191;"`                                                                             //移动电话号码
	EmergencyContactNo               string `json:"emergencyContactNo" form:"emergencyContactNo" gorm:"column:emergency_contactNo;comment:紧急联系电话;size:191;"`                                               //紧急联系电话
	TechnicalContactEmailAddress     string `json:"technicalContactEmailAddress" form:"technicalContactEmailAddress" gorm:"column:technical_contact_email_address;comment:技术联系人邮箱;size:191;"`              //技术联系人邮箱
	TechnicalContactTelNo            string `json:"technicalContactTelNo" form:"technicalContactTelNo" gorm:"column:technical_contact_tel_no;comment:电话;size:191;"`                                        //技术电话
	TechnicalContactFaxNo            string `json:"technicalContactFaxNo" form:"technicalContactFaxNo" gorm:"column:technical_contact_fax_no;comment:传真号;size:191;"`                                       //技术传真号
	DirectoryListing                 string `json:"directoryListing" form:"directoryListing" gorm:"column:directory_listing;comment:目录列表;size:191;"`                                                       //目录列表YN 限制只能填写Y或者N
	EBillEmailAddressReceivingEBills string `json:"eBillEmailAddressReceivingEBills" form:"eBillEmailAddressReceivingEBills" gorm:"column:eBill_email_address_receiving_e_bills;comment:账单邮件地址;size:191;"` //账单邮件地址
	PaymentMethod                    string `json:"paymentMethod" form:"paymentMethod" gorm:"column:payment_method;comment:支付方式;size:191;"`                                                                //支付方式
}
