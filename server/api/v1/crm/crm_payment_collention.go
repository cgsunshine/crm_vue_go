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
	"time"
)

type CrmPaymentCollentionApi struct {
}

var crmPaymentCollentionService = service.ServiceGroupApp.CrmServiceGroup.CrmPaymentCollentionService

// CreateCrmPaymentCollention 创建crmPaymentCollention表
// @Tags CrmPaymentCollention
// @Summary 创建crmPaymentCollention表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPaymentCollention true "创建crmPaymentCollention表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPaymentCollention/createCrmPaymentCollention [post]
func (crmPaymentCollentionApi *CrmPaymentCollentionApi) CreateCrmPaymentCollention(c *gin.Context) {
	var crmPaymentCollention crm.CrmPaymentCollention
	err := c.ShouldBindJSON(&crmPaymentCollention)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	t := time.Now()

	crmPaymentCollention.UserId = comm.GetHeaderUserId(c)
	crmPaymentCollention.PaymentTime = &t

	crmPaymentCollention.PaymentCollentionName = comm.GetBusinessNumber(comm.PaymentCollentionNumberPrefix, crmPaymentCollentionService.GetCrmPaymentCollentionTodayCount())
	if err := crmPaymentCollentionService.CreateCrmPaymentCollention(&crmPaymentCollention); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		err = crmBillService.UpdAssOrderID(crmPaymentCollention.OrderId, map[string]interface{}{
			"payment_collention_id": crmPaymentCollention.ID,
		})
		if err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmPaymentCollention 删除crmPaymentCollention表
// @Tags CrmPaymentCollention
// @Summary 删除crmPaymentCollention表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPaymentCollention true "删除crmPaymentCollention表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPaymentCollention/deleteCrmPaymentCollention [delete]
func (crmPaymentCollentionApi *CrmPaymentCollentionApi) DeleteCrmPaymentCollention(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmPaymentCollentionService.DeleteCrmPaymentCollention(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {

		//关联删除审批
		err := crmApprovalTasksService.DelCrmAssociatedIdApprovalTasks(ID, comm.PaymentCollentionApprovalType)
		if err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			response.FailWithMessage("删除失败", c)
		}

	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCrmPaymentCollentionByIds 批量删除crmPaymentCollention表
// @Tags CrmPaymentCollention
// @Summary 批量删除crmPaymentCollention表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmPaymentCollention/deleteCrmPaymentCollentionByIds [delete]
func (crmPaymentCollentionApi *CrmPaymentCollentionApi) DeleteCrmPaymentCollentionByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmPaymentCollentionService.DeleteCrmPaymentCollentionByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmPaymentCollention 更新crmPaymentCollention表
// @Tags CrmPaymentCollention
// @Summary 更新crmPaymentCollention表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPaymentCollention true "更新crmPaymentCollention表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmPaymentCollention/updateCrmPaymentCollention [put]
func (crmPaymentCollentionApi *CrmPaymentCollentionApi) UpdateCrmPaymentCollention(c *gin.Context) {
	var crmPaymentCollention crm.CrmPaymentCollention
	err := c.ShouldBindJSON(&crmPaymentCollention)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmPaymentCollentionService.UpdateCrmPaymentCollention(crmPaymentCollention); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmPaymentCollention 用id查询crmPaymentCollention表
// @Tags CrmPaymentCollention
// @Summary 用id查询crmPaymentCollention表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmPaymentCollention true "用id查询crmPaymentCollention表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPaymentCollention/findCrmPaymentCollention [get]
func (crmPaymentCollentionApi *CrmPaymentCollentionApi) FindCrmPaymentCollention(c *gin.Context) {
	ID := c.Query("ID")
	if recrmPaymentCollention, err := crmPaymentCollentionService.GetCrmPaymentCollention(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		recrmPaymentCollention.Proof, _, err = fileUploadAndDownloadService.GetFileRecordInfoIdsString(recrmPaymentCollention.Proof)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
		response.OkWithData(gin.H{"recrmPaymentCollention": recrmPaymentCollention}, c)
	}
}

// GetCrmPaymentCollentionList 分页获取crmPaymentCollention表列表
// @Tags CrmPaymentCollention
// @Summary 分页获取crmPaymentCollention表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPaymentCollentionSearch true "分页获取crmPaymentCollention表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPaymentCollention/getCrmPaymentCollentionList [get]
func (crmPaymentCollentionApi *CrmPaymentCollentionApi) GetCrmPaymentCollentionList(c *gin.Context) {
	var pageInfo crmReq.CrmPaymentCollentionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)
	if list, total, err := crmPaymentCollentionService.GetCrmPaymentCollentionInfoList(pageInfo); err != nil {
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

// GetCrmPaymentCollentionPublic 不需要鉴权的crmPaymentCollention表接口
// @Tags CrmPaymentCollention
// @Summary 不需要鉴权的crmPaymentCollention表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPaymentCollentionSearch true "分页获取crmPaymentCollention表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPaymentCollention/getCrmPaymentCollentionList [get]
func (crmPaymentCollentionApi *CrmPaymentCollentionApi) GetCrmPaymentCollentionPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的crmPaymentCollention表接口信息",
	}, "获取成功", c)
}
