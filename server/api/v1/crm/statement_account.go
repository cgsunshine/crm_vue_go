package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetCrmStatementAccountList 分页获取crmStatementAccount表列表
// @Tags CrmStatementAccount
// @Summary 分页获取crmStatementAccount表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmStatementAccountSearch true "分页获取crmStatementAccount表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmStatementAccount/getCrmStatementAccountList [get]
func (crmStatementAccountApi *CrmStatementAccountApi) GetCrmPageStatementAccountList(c *gin.Context) {
	var pageInfo crmReq.CrmStatementAccountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	pageInfo.UserId = userService.FindUserDataStatusById(&userID)

	if list, total, err := crmStatementAccountService.GetCrmPageStatementAccountInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// FindCrmPageStatementAccount 用id查询crmStatementAccount表
// @Tags CrmStatementAccount
// @Summary 用id查询crmStatementAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmStatementAccount true "用id查询crmStatementAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmStatementAccount/findCrmStatementAccount [get]
func (crmStatementAccountApi *CrmStatementAccountApi) FindCrmPageStatementAccount(c *gin.Context) {
	ID := c.Query("ID")
	if recrmStatementAccount, err := crmStatementAccountService.GetCrmPageStatementAccount(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmStatementAccount": recrmStatementAccount}, c)
	}
}
