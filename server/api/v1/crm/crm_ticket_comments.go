package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/crm"
    crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CrmTicketCommentsApi struct {
}

var crmTicketCommentsService = service.ServiceGroupApp.CrmServiceGroup.CrmTicketCommentsService


// CreateCrmTicketComments 创建共单回复
// @Tags CrmTicketComments
// @Summary 创建共单回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTicketComments true "创建共单回复"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmTicketComments/createCrmTicketComments [post]
func (crmTicketCommentsApi *CrmTicketCommentsApi) CreateCrmTicketComments(c *gin.Context) {
	var crmTicketComments crm.CrmTicketComments
	err := c.ShouldBindJSON(&crmTicketComments)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmTicketCommentsService.CreateCrmTicketComments(&crmTicketComments); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmTicketComments 删除共单回复
// @Tags CrmTicketComments
// @Summary 删除共单回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTicketComments true "删除共单回复"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTicketComments/deleteCrmTicketComments [delete]
func (crmTicketCommentsApi *CrmTicketCommentsApi) DeleteCrmTicketComments(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmTicketCommentsService.DeleteCrmTicketComments(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmTicketCommentsByIds 批量删除共单回复
// @Tags CrmTicketComments
// @Summary 批量删除共单回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmTicketComments/deleteCrmTicketCommentsByIds [delete]
func (crmTicketCommentsApi *CrmTicketCommentsApi) DeleteCrmTicketCommentsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmTicketCommentsService.DeleteCrmTicketCommentsByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmTicketComments 更新共单回复
// @Tags CrmTicketComments
// @Summary 更新共单回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmTicketComments true "更新共单回复"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmTicketComments/updateCrmTicketComments [put]
func (crmTicketCommentsApi *CrmTicketCommentsApi) UpdateCrmTicketComments(c *gin.Context) {
	var crmTicketComments crm.CrmTicketComments
	err := c.ShouldBindJSON(&crmTicketComments)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmTicketCommentsService.UpdateCrmTicketComments(crmTicketComments); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmTicketComments 用id查询共单回复
// @Tags CrmTicketComments
// @Summary 用id查询共单回复
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmTicketComments true "用id查询共单回复"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmTicketComments/findCrmTicketComments [get]
func (crmTicketCommentsApi *CrmTicketCommentsApi) FindCrmTicketComments(c *gin.Context) {
	ID := c.Query("ID")
	if recrmTicketComments, err := crmTicketCommentsService.GetCrmTicketComments(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmTicketComments": recrmTicketComments}, c)
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
func (crmTicketCommentsApi *CrmTicketCommentsApi) GetCrmTicketCommentsList(c *gin.Context) {
	var pageInfo crmReq.CrmTicketCommentsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmTicketCommentsService.GetCrmTicketCommentsInfoList(pageInfo); err != nil {
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

// GetCrmTicketCommentsPublic 不需要鉴权的共单回复接口
// @Tags CrmTicketComments
// @Summary 不需要鉴权的共单回复接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmTicketCommentsSearch true "分页获取共单回复列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTicketComments/getCrmTicketCommentsList [get]
func (crmTicketCommentsApi *CrmTicketCommentsApi) GetCrmTicketCommentsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的共单回复接口信息",
    }, "获取成功", c)
}
