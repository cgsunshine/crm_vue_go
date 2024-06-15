// 自动生成模板CrmBusinessOpportunityFile
package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 商机文件 结构体  CrmBusinessOpportunityFile
type CrmBusinessOpportunityFile struct {
 global.GVA_MODEL 
      BusinessOpportunityId  *int `json:"businessOpportunityId" form:"businessOpportunityId" gorm:"column:business_opportunity_id;comment:商机id;size:19;"`  //商机id 
      Key  string `json:"key" form:"key" gorm:"column:key;comment:编号;size:191;"`  //编号 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:文件名;size:191;"`  //文件名 
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:19;"`  //排序 
      Tag  string `json:"tag" form:"tag" gorm:"column:tag;comment:文件标签;size:191;"`  //文件标签 
      Url  string `json:"url" form:"url" gorm:"column:url;comment:文件地址;size:191;"`  //文件地址 
}


// TableName 商机文件 CrmBusinessOpportunityFile自定义表名 crm_business_opportunity_file
func (CrmBusinessOpportunityFile) TableName() string {
  return "crm_business_opportunity_file"
}

