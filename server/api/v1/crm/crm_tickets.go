package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CrmTicketsApi struct {
}

var crmTicketsService = service.ServiceGroupApp.CrmServiceGroup.CrmTicketsService

// CreateCrmTickets 创建工单
// @Tags CrmTickets
// @Summary 创建工单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTickets true "创建工单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmTickets/createCrmTickets [post]
func (crmTicketsApi *CrmTicketsApi) CreateCrmTickets(c *gin.Context) {
	var crmTickets crm.CrmTickets
	err := c.ShouldBindJSON(&crmTickets)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	crmTickets.SubmitterId = comm.GetHeaderUserId(c)
	crmTickets.TicketStatus = comm.Deposits_Processing_Status_Tickets_Unprocessed
	if err := crmTicketsService.CreateCrmTickets(&crmTickets); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmTickets 删除工单
// @Tags CrmTickets
// @Summary 删除工单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTickets true "删除工单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTickets/deleteCrmTickets [delete]
func (crmTicketsApi *CrmTicketsApi) DeleteCrmTickets(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmTicketsService.DeleteCrmTickets(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmTicketsByIds 批量删除工单
// @Tags CrmTickets
// @Summary 批量删除工单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmTickets/deleteCrmTicketsByIds [delete]
func (crmTicketsApi *CrmTicketsApi) DeleteCrmTicketsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmTicketsService.DeleteCrmTicketsByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmTickets 更新工单
// @Tags CrmTickets
// @Summary 更新工单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTickets true "更新工单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmTickets/updateCrmTickets [put]
func (crmTicketsApi *CrmTicketsApi) UpdateCrmTickets(c *gin.Context) {
	var crmTickets crm.CrmTickets
	err := c.ShouldBindJSON(&crmTickets)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmTicketsService.UpdateCrmTickets(crmTickets); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmTickets 用id查询工单
// @Tags CrmTickets
// @Summary 用id查询工单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmTickets true "用id查询工单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmTickets/findCrmTickets [get]
func (crmTicketsApi *CrmTicketsApi) FindCrmTickets(c *gin.Context) {
	ID := c.Query("ID")
	if recrmTickets, err := crmTicketsService.GetCrmTickets(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmTickets": recrmTickets}, c)
	}
}

// GetCrmTicketsList 分页获取工单列表
// @Tags CrmTickets
// @Summary 分页获取工单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmTicketsSearch true "分页获取工单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTickets/getCrmTicketsList [get]
func (crmTicketsApi *CrmTicketsApi) GetCrmTicketsList(c *gin.Context) {
	var pageInfo crmReq.CrmTicketsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmTicketsService.GetCrmTicketsInfoList(pageInfo); err != nil {
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

// GetCrmTicketsPublic 不需要鉴权的工单接口
// @Tags CrmTickets
// @Summary 不需要鉴权的工单接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmTicketsSearch true "分页获取工单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTickets/getCrmTicketsList [get]
func (crmTicketsApi *CrmTicketsApi) GetCrmTicketsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的工单接口信息",
	}, "获取成功", c)
}
