package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type CrmLoginResponse struct {
	Data    CrmLSysUserResponse `json:"data"`
	Success bool                `json:"success"`
}

type CrmLSysUserResponse struct {
	AccessToken  string   `json:"accessToken"`
	Avatar       string   `json:"avatar"`
	Expires      string   `json:"expires"`
	Nickname     string   `json:"nickname"`
	RefreshToken string   `json:"refreshToken"`
	Roles        []string `json:"roles"`
	Username     string   `json:"username"`
}
