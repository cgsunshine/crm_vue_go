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

type CrmBillApi struct {
}

var crmBillService = service.ServiceGroupApp.CrmServiceGroup.CrmBillService

// CreateCrmBill 创建crmBill表
// @Tags CrmBill
// @Summary 创建crmBill表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBill true "创建crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBill/createCrmBill [post]
func (crmBillApi *CrmBillApi) CreateCrmBill(c *gin.Context) {
	var crmBill crm.CrmBill
	err := c.ShouldBindJSON(&crmBill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmBill.UserId = comm.GetHeaderUserId(c)

	crmBill.PaymentStatus = comm.PaymentStatusUnpaid

	if err := crmBillService.CreateCrmBill(&crmBill); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmBill 删除crmBill表
// @Tags CrmBill
// @Summary 删除crmBill表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBill true "删除crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBill/deleteCrmBill [delete]
func (crmBillApi *CrmBillApi) DeleteCrmBill(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmBillService.DeleteCrmBill(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmBillByIds 批量删除crmBill表
// @Tags CrmBill
// @Summary 批量删除crmBill表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmBill/deleteCrmBillByIds [delete]
func (crmBillApi *CrmBillApi) DeleteCrmBillByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmBillService.DeleteCrmBillByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmBill 更新crmBill表
// @Tags CrmBill
// @Summary 更新crmBill表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBill true "更新crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBill/updateCrmBill [put]
func (crmBillApi *CrmBillApi) UpdateCrmBill(c *gin.Context) {
	var crmBill crm.CrmBill
	err := c.ShouldBindJSON(&crmBill)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBillService.UpdateCrmBill(crmBill); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmBill 用id查询crmBill表
// @Tags CrmBill
// @Summary 用id查询crmBill表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBill true "用id查询crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBill/findCrmBill [get]
func (crmBillApi *CrmBillApi) FindCrmBill(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBill, err := crmBillService.GetCrmPageIdBill(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmBill": recrmBill}, c)
	}
}

// GetCrmBillList 分页获取crmBill表列表
// @Tags CrmBill
// @Summary 分页获取crmBill表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBillSearch true "分页获取crmBill表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBill/getCrmBillList [get]
func (crmBillApi *CrmBillApi) GetCrmBillList(c *gin.Context) {
	var pageInfo crmReq.CrmBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmBillService.GetCrmBillInfoList(pageInfo); err != nil {
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

// GetCrmBillPublic 不需要鉴权的crmBill表接口
// @Tags CrmBill
// @Summary 不需要鉴权的crmBill表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBillSearch true "分页获取crmBill表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBill/getCrmBillList [get]
func (crmBillApi *CrmBillApi) GetCrmBillPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的crmBill表接口信息",
	}, "获取成功", c)
}
