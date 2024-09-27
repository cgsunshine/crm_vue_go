package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过id获取用户信息 并且查询权限
//@param: uuid string
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserDataStatusById(id *int) *int {
	var u system.SysUser
	if err := global.GVA_DB.Where("id = ?", id).First(&u).Error; err != nil {
		return id
	}
	var a system.SysAuthority
	if err := global.GVA_DB.Where("authority_id = ?", u.AuthorityId).First(&a).Error; err != nil {
		return id
	}
	if a.Status == "1" {
		return id
	}
	return nil
}

func (userService *UserService) FindUserDataStatusByIdRole(id *int) *int {
	var u system.SysUser
	AuthorityId := 0
	if err := global.GVA_DB.Where("id = ?", id).First(&u).Error; err != nil {
		AuthorityId = int(u.AuthorityId)
		return &AuthorityId
	}
	var a system.SysAuthority
	if err := global.GVA_DB.Where("authority_id = ?", u.AuthorityId).First(&a).Error; err != nil {
		AuthorityId = int(u.AuthorityId)
		return &AuthorityId
	}
	if a.Status == "1" {
		AuthorityId = int(u.AuthorityId)
		return &AuthorityId
	}
	return nil
}
