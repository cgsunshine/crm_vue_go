package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/gin-gonic/gin"
)

// 所有人数据权限，和条件筛选处理
func GetSearchUserId(userId *int, c *gin.Context) *int {
	if userId == nil {
		return userService.FindUserDataStatusById(comm.GetHeaderUserId(c))
	}
	return userId
}
