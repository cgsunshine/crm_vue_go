package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetCrmSubmitterTicketsList 分页获取工单列表-提交人
// @Tags CrmTickets
// @Summary 分页获取工单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmTicketsSearch true "分页获取工单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTickets/getCrmTicketsList [get]
func (crmTicketsApi *CrmTicketsApi) GetCrmSubmitterTicketsList(c *gin.Context) {
	var pageInfo crmReq.CrmTicketsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.SubmitterId = GetSearchUserId(pageInfo.SubmitterId, c)

	if list, total, err := crmTicketsService.GetCrmPageTicketsInfoList(pageInfo); err != nil {
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

// GetCrmAssigneeTicketsList 分页获取工单列表-指派人
// @Tags CrmTickets
// @Summary 分页获取工单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmTicketsSearch true "分页获取工单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTickets/getCrmAssigneeTicketsList [get]
func (crmTicketsApi *CrmTicketsApi) GetCrmAssigneeTicketsList(c *gin.Context) {
	var pageInfo crmReq.CrmTicketsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	AssigneeId, _ := strconv.Atoi(c.GetHeader("X-User-Id"))
	pageInfo.AssigneeId = &AssigneeId
	if list, total, err := crmTicketsService.GetCrmPageTicketsInfoList(pageInfo); err != nil {
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

// FindCrmPageTickets 用id查询工单
// @Tags CrmTickets
// @Summary 用id查询工单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmTickets true "用id查询工单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmTickets/findCrmTickets [get]
func (crmTicketsApi *CrmTicketsApi) FindCrmPageTickets(c *gin.Context) {
	ID := c.Query("ID")
	if recrmTickets, err := crmTicketsService.GetCrmPageTickets(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmTickets": recrmTickets}, c)
	}
}
