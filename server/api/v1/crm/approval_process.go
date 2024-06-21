package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreateCrmPageApprovalProcess 创建审批流程
// @Tags CrmApprovalProcess
// @Summary 创建审批流程
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalProcess true "创建审批流程"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalProcess/createCrmApprovalProcess [post]
func (crmApprovalProcessApi *CrmApprovalProcessApi) CreateCrmPageApprovalProcess(c *gin.Context) {
	var crmApprovalProcess crm.CrmApprovalProcess
	err := c.ShouldBindJSON(&crmApprovalProcess)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	crmApprovalProcess.UserId = userService.FindUserDataStatusById(&userID)

	if err := crmApprovalProcessService.CreateCrmApprovalProcess(&crmApprovalProcess); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
