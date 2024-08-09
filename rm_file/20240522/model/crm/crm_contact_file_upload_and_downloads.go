// 自动生成模板CrmContactFileUploadAndDownloads
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// crmContactFileUploadAndDownloads表 结构体  CrmContactFileUploadAndDownloads
type CrmContactFileUploadAndDownloads struct {
 global.GVA_MODEL 
      ContactId  *int `json:"contactId" form:"contactId" gorm:"column:contact_id;comment:合同id;size:10;"`  //合同id 
      Key  string `json:"key" form:"key" gorm:"column:key;comment:编号;size:191;"`  //编号 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:文件名;size:191;"`  //文件名 
      Tag  string `json:"tag" form:"tag" gorm:"column:tag;comment:文件标签;size:191;"`  //文件标签 
      Url  string `json:"url" form:"url" gorm:"column:url;comment:文件地址;size:191;"`  //文件地址 
}


// TableName crmContactFileUploadAndDownloads表 CrmContactFileUploadAndDownloads自定义表名 crm_contact_file_upload_and_downloads
func (CrmContactFileUploadAndDownloads) TableName() string {
  return "crm_contact_file_upload_and_downloads"
}

