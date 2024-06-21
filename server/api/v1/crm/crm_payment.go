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

type CrmPaymentApi struct {
}

var crmPaymentService = service.ServiceGroupApp.CrmServiceGroup.CrmPaymentService

// CreateCrmPayment 创建crmPayment表
// @Tags CrmPayment
// @Summary 创建crmPayment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPayment true "创建crmPayment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPayment/createCrmPayment [post]
func (crmPaymentApi *CrmPaymentApi) CreateCrmPayment(c *gin.Context) {
	var crmPayment crm.CrmPayment
	err := c.ShouldBindJSON(&crmPayment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmPayment.UserId = comm.GetHeaderUserId(c)

	if err := crmPaymentService.CreateCrmPayment(&crmPayment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmPayment 删除crmPayment表
// @Tags CrmPayment
// @Summary 删除crmPayment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPayment true "删除crmPayment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPayment/deleteCrmPayment [delete]
func (crmPaymentApi *CrmPaymentApi) DeleteCrmPayment(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmPaymentService.DeleteCrmPayment(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmPaymentByIds 批量删除crmPayment表
// @Tags CrmPayment
// @Summary 批量删除crmPayment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmPayment/deleteCrmPaymentByIds [delete]
func (crmPaymentApi *CrmPaymentApi) DeleteCrmPaymentByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmPaymentService.DeleteCrmPaymentByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmPayment 更新crmPayment表
// @Tags CrmPayment
// @Summary 更新crmPayment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPayment true "更新crmPayment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmPayment/updateCrmPayment [put]
func (crmPaymentApi *CrmPaymentApi) UpdateCrmPayment(c *gin.Context) {
	var crmPayment crm.CrmPayment
	err := c.ShouldBindJSON(&crmPayment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmPaymentService.UpdateCrmPayment(crmPayment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmPayment 用id查询crmPayment表
// @Tags CrmPayment
// @Summary 用id查询crmPayment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmPayment true "用id查询crmPayment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPayment/findCrmPayment [get]
func (crmPaymentApi *CrmPaymentApi) FindCrmPayment(c *gin.Context) {
	ID := c.Query("ID")
	if recrmPayment, err := crmPaymentService.GetCrmPayment(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmPayment": recrmPayment}, c)
	}
}

// GetCrmPaymentList 分页获取crmPayment表列表
// @Tags CrmPayment
// @Summary 分页获取crmPayment表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPaymentSearch true "分页获取crmPayment表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPayment/getCrmPaymentList [get]
func (crmPaymentApi *CrmPaymentApi) GetCrmPaymentList(c *gin.Context) {
	var pageInfo crmReq.CrmPaymentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmPaymentService.GetCrmPaymentInfoList(pageInfo); err != nil {
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

// GetCrmPaymentPublic 不需要鉴权的crmPayment表接口
// @Tags CrmPayment
// @Summary 不需要鉴权的crmPayment表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPaymentSearch true "分页获取crmPayment表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPayment/getCrmPaymentList [get]
func (crmPaymentApi *CrmPaymentApi) GetCrmPaymentPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的crmPayment表接口信息",
	}, "获取成功", c)
}
