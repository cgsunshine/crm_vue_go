package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CrmBillPaymentApi struct {
}

var crmBillPaymentService = service.ServiceGroupApp.CrmServiceGroup.CrmBillPaymentService

// CreateCrmBillPayment 创建应付账单
// @Tags CrmBillPayment
// @Summary 创建应付账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBillPayment true "创建应付账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBillPayment/createCrmBillPayment [post]
func (crmBillPaymentApi *CrmBillPaymentApi) CreateCrmBillPayment(c *gin.Context) {
	var crmBillPayment crm.CrmBillPayment
	err := c.ShouldBindJSON(&crmBillPayment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBillPaymentService.CreateCrmBillPayment(&crmBillPayment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmBillPayment 删除应付账单
// @Tags CrmBillPayment
// @Summary 删除应付账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBillPayment true "删除应付账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBillPayment/deleteCrmBillPayment [delete]
func (crmBillPaymentApi *CrmBillPaymentApi) DeleteCrmBillPayment(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmBillPaymentService.DeleteCrmBillPayment(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmBillPaymentByIds 批量删除应付账单
// @Tags CrmBillPayment
// @Summary 批量删除应付账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmBillPayment/deleteCrmBillPaymentByIds [delete]
func (crmBillPaymentApi *CrmBillPaymentApi) DeleteCrmBillPaymentByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmBillPaymentService.DeleteCrmBillPaymentByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmBillPayment 更新应付账单
// @Tags CrmBillPayment
// @Summary 更新应付账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBillPayment true "更新应付账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBillPayment/updateCrmBillPayment [put]
func (crmBillPaymentApi *CrmBillPaymentApi) UpdateCrmBillPayment(c *gin.Context) {
	var crmBillPayment crm.CrmBillPayment
	err := c.ShouldBindJSON(&crmBillPayment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBillPaymentService.UpdateCrmBillPayment(crmBillPayment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmBillPayment 用id查询应付账单
// @Tags CrmBillPayment
// @Summary 用id查询应付账单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBillPayment true "用id查询应付账单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBillPayment/findCrmBillPayment [get]
func (crmBillPaymentApi *CrmBillPaymentApi) FindCrmBillPayment(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBillPayment, err := crmBillPaymentService.GetCrmBillPayment(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmBillPayment": recrmBillPayment}, c)
	}
}

// GetCrmBillPaymentList 分页获取应付账单列表
// @Tags CrmBillPayment
// @Summary 分页获取应付账单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBillPaymentSearch true "分页获取应付账单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBillPayment/getCrmBillPaymentList [get]
func (crmBillPaymentApi *CrmBillPaymentApi) GetCrmBillPaymentList(c *gin.Context) {
	var pageInfo crmReq.CrmBillPaymentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)
	if list, total, err := crmBillPaymentService.GetCrmBillPaymentInfoList(pageInfo); err != nil {
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

// GetCrmBillPaymentPublic 不需要鉴权的应付账单接口
// @Tags CrmBillPayment
// @Summary 不需要鉴权的应付账单接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBillPaymentSearch true "分页获取应付账单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBillPayment/getCrmBillPaymentList [get]
func (crmBillPaymentApi *CrmBillPaymentApi) GetCrmBillPaymentPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的应付账单接口信息",
	}, "获取成功", c)
}
