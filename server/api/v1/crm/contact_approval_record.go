package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetCrmContactApprovalRecordList 分页获取合同审批记录列表
// @Tags CrmContactApprovalRecord
// @Summary 分页获取合同审批记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContactApprovalRecordSearch true "分页获取合同审批记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactApprovalRecord/getCrmContactApprovalRecordList [get]
func (crmContactApprovalRecordApi *CrmContactApprovalRecordApi) GetCrmPageContactApprovalRecordList(c *gin.Context) {
	var pageInfo crmReq.CrmContactApprovalRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	ApproverId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	pageInfo.ApproverId = &ApproverId

	if list, total, err := crmContactApprovalRecordService.GetCrmContactApprovalRecordInfoList(pageInfo); err != nil {
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
