package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
)

// UpdateCrmAuthorities 更新数据权限记录
// Author [piexlmax](https://github.com/piexlmax)
func (crmAuthoritiesService *CrmAuthoritiesService) UpdateCrmStatusAuthorities(crmAuthorities crm.CrmAuthorities) (err error) {

	err = global.GVA_DB.Table("sys_authorities").Debug().Where("authority_id = ?", crmAuthorities.AuthorityId).Update("status", crmAuthorities.Status).Error
	return err
}
