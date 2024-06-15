// 自动生成模板CrmTicketResponseTemplates
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// crmTicketResponseTemplates表 结构体  CrmTicketResponseTemplates
type CrmTicketResponseTemplates struct {
 global.GVA_MODEL 
      TemplateName  string `json:"templateName" form:"templateName" gorm:"column:template_name;comment:回复模板名称;size:255;"`  //回复模板名称 
      Content  string `json:"content" form:"content" gorm:"column:content;comment:回复内容;"`  //回复内容 
      CategoryId  *int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:适用工单类别ID;size:10;"`  //适用工单类别ID 
      CreatedBy  *int `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者ID;size:10;"`  //创建者ID 
}


// TableName crmTicketResponseTemplates表 CrmTicketResponseTemplates自定义表名 crm_ticket_response_templates
func (CrmTicketResponseTemplates) TableName() string {
  return "crm_ticket_response_templates"
}

