// 自动生成模板CrmBusinessOpportunityFileUploadAndDownloads
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 上传商机文件 结构体  CrmBusinessOpportunityFileUploadAndDownloads
type CrmBusinessOpportunityFileUploadAndDownloads struct {
 global.GVA_MODEL 
      BusinessOpportunityId  *int `json:"businessOpportunityId" form:"businessOpportunityId" gorm:"column:business_opportunity_id;comment:商机id;"`  //商机id 
      Key  string `json:"key" form:"key" gorm:"column:key;comment:编号;size:191;"`  //编号 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:文件名;size:191;"`  //文件名 
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`  //排序 
      Tag  string `json:"tag" form:"tag" gorm:"column:tag;comment:文件标签;size:191;"`  //文件标签 
      Url  string `json:"url" form:"url" gorm:"column:url;comment:文件地址;size:191;"`  //文件地址 
}


// TableName 上传商机文件 CrmBusinessOpportunityFileUploadAndDownloads自定义表名 crm_business_opportunity_file_upload_and_downloads
func (CrmBusinessOpportunityFileUploadAndDownloads) TableName() string {
  return "crm_business_opportunity_file_upload_and_downloads"
}

