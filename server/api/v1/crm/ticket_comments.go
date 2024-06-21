package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateCrmTicketComments 创建共单回复
// @Tags CrmTicketComments
// @Summary 创建共单回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTicketComments true "创建共单回复"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmTicketComments/createCrmTicketComments [post]
func (crmTicketCommentsApi *CrmTicketCommentsApi) CreateCrmPageTicketComments(c *gin.Context) {
	var crmTicketComments crm.CrmTicketComments
	err := c.ShouldBindJSON(&crmTicketComments)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmTicketComments.UserId = comm.GetHeaderUserId(c)

	if err := crmTicketCommentsService.CreateCrmTicketComments(&crmTicketComments); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// GetCrmTicketCommentsList 分页获取共单回复列表
// @Tags CrmTicketComments
// @Summary 分页获取共单回复列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmTicketCommentsSearch true "分页获取共单回复列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTicketComments/getCrmTicketCommentsList [get]
func (crmTicketCommentsApi *CrmTicketCommentsApi) GetCrmPageTicketCommentsList(c *gin.Context) {
	var pageInfo crmReq.CrmTicketCommentsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmTicketCommentsService.GetCrmPageTicketCommentsInfoList(pageInfo); err != nil {
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

// FindCrmPageTicketComments 用id查询共单回复
// @Tags CrmTicketComments
// @Summary 用id查询共单回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmTicketComments true "用id查询共单回复"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmTicketComments/findCrmTicketComments [get]
func (crmTicketCommentsApi *CrmTicketCommentsApi) FindCrmPageTicketComments(c *gin.Context) {
	ID := c.Query("ID")
	if recrmTicketComments, err := crmTicketCommentsService.GetCrmPageTicketComments(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmTicketComments": recrmTicketComments}, c)
	}
}
