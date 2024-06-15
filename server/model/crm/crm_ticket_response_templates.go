// 自动生成模板CrmTicketResponseTemplates
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 快捷回复模板 结构体  CrmTicketResponseTemplates
type CrmTicketResponseTemplates struct {
 global.GVA_MODEL 
      CategoryId  *int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:适用工单类别ID;"`  //适用工单类别ID 
      Content  string `json:"content" form:"content" gorm:"column:content;comment:回复内容;size:191;"`  //回复内容 
      CreatedBy  *int `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者ID;"`  //创建者ID 
      TemplateName  string `json:"templateName" form:"templateName" gorm:"column:template_name;comment:回复模板名称;size:255;"`  //回复模板名称 
}


// TableName 快捷回复模板 CrmTicketResponseTemplates自定义表名 crm_ticket_response_templates
func (CrmTicketResponseTemplates) TableName() string {
  return "crm_ticket_response_templates"
}

