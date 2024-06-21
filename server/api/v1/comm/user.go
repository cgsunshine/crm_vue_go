package comm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func GetHeaderUserId(c *gin.Context) *int {
	userID, err := strconv.Atoi(c.GetHeader("X-User-Id"))
	if err != nil {
		global.GVA_LOG.Error("获取userId失败!", zap.Error(err))
		return nil
	}
	return &userID
}
