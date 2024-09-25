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

type CrmRefundTasksApi struct {
}

var crmRefundTasksService = service.ServiceGroupApp.CrmServiceGroup.CrmRefundTasksService

// CreateCrmRefundTasks 创建退款管理
// @Tags CrmRefundTasks
// @Summary 创建退款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmRefundTasks true "创建退款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmRefundTasks/createCrmRefundTasks [post]
func (crmRefundTasksApi *CrmRefundTasksApi) CreateCrmRefundTasks(c *gin.Context) {
	var crmRefundTasks crm.CrmRefundTasks
	err := c.ShouldBindJSON(&crmRefundTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmRefundTasksService.CreateCrmRefundTasks(&crmRefundTasks); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// ProcessingCrmRefundTasks 处理审批
// @Tags CrmRefundTasks
// @Summary 处理审批
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmRefundTasks true "处理审批"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmRefundTasks/updateCrmRefundTasks [put]
func (crmRefundTasksApi *CrmRefundTasksApi) ProcessingCrmRefundTasks(c *gin.Context) {
	var crmRefundTasks crm.CrmRefundTasks
	err := c.ShouldBindJSON(&crmRefundTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmRefundTasksInfo, err := crmRefundTasksService.GetCrmRefundTasksInfo(crmRefundTasks.ID)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	if crmRefundTasksInfo.RefundStatus != comm.Deposits_Processing_Status_Unprocessed {
		global.GVA_LOG.Error("已处理，无需重复操作!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	//if crmRefundTasks.RefundStatus == comm.Deposits_Processing_Status_Processed {
	//
	//}

	//退款更新关联表
	err = crmDepositsService.UpdDepositsInfo(crmRefundTasksInfo.AssociatedId, map[string]interface{}{
		"deposits_status": comm.Deposits_Status_Refund,
		"refund_date":     time.Now(),
		"refund_status":   comm.Deposits_Status_Refund,
		"refund_voucher":  crmRefundTasks.RefundVoucher,
	})
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	id := int(crmRefundTasksInfo.ID)
	//退款更新关联表
	err = crmRefundTasksService.UpdDepositsInfo(&id, map[string]interface{}{
		"refund_status":  comm.Deposits_Processing_Status_Processed,
		"refund_voucher": crmRefundTasks.RefundVoucher,
	})
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	response.OkWithMessage("更新成功", c)

}

// DeleteCrmRefundTasks 删除退款管理
// @Tags CrmRefundTasks
// @Summary 删除退款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmRefundTasks true "删除退款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmRefundTasks/deleteCrmRefundTasks [delete]
func (crmRefundTasksApi *CrmRefundTasksApi) DeleteCrmRefundTasks(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmRefundTasksService.DeleteCrmRefundTasks(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmRefundTasksByIds 批量删除退款管理
// @Tags CrmRefundTasks
// @Summary 批量删除退款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmRefundTasks/deleteCrmRefundTasksByIds [delete]
func (crmRefundTasksApi *CrmRefundTasksApi) DeleteCrmRefundTasksByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmRefundTasksService.DeleteCrmRefundTasksByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmRefundTasks 更新退款管理
// @Tags CrmRefundTasks
// @Summary 更新退款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmRefundTasks true "更新退款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmRefundTasks/updateCrmRefundTasks [put]
func (crmRefundTasksApi *CrmRefundTasksApi) UpdateCrmRefundTasks(c *gin.Context) {
	var crmRefundTasks crm.CrmRefundTasks
	err := c.ShouldBindJSON(&crmRefundTasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmRefundTasksService.UpdateCrmRefundTasks(crmRefundTasks); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmRefundTasks 用id查询退款管理
// @Tags CrmRefundTasks
// @Summary 用id查询退款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmRefundTasks true "用id查询退款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmRefundTasks/findCrmRefundTasks [get]
func (crmRefundTasksApi *CrmRefundTasksApi) FindCrmRefundTasks(c *gin.Context) {
	ID := c.Query("ID")
	if recrmRefundTasks, err := crmRefundTasksService.GetCrmRefundTasks(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		recrmRefundTasks.RefundVoucher, _, err = fileUploadAndDownloadService.GetFileRecordInfoIdsString(recrmRefundTasks.RefundVoucher)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
		response.OkWithData(gin.H{"recrmRefundTasks": recrmRefundTasks}, c)
	}
}

// GetCrmRefundTasksList 分页获取退款管理列表
// @Tags CrmRefundTasks
// @Summary 分页获取退款管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmRefundTasksSearch true "分页获取退款管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmRefundTasks/getCrmRefundTasksList [get]
func (crmRefundTasksApi *CrmRefundTasksApi) GetCrmRefundTasksList(c *gin.Context) {
	var pageInfo crmReq.CrmRefundTasksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmRefundTasksService.GetCrmRefundTasksInfoList(pageInfo); err != nil {
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

// GetCrmRefundTasksPublic 不需要鉴权的退款管理接口
// @Tags CrmRefundTasks
// @Summary 不需要鉴权的退款管理接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmRefundTasksSearch true "分页获取退款管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmRefundTasks/getCrmRefundTasksList [get]
func (crmRefundTasksApi *CrmRefundTasksApi) GetCrmRefundTasksPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的退款管理接口信息",
	}, "获取成功", c)
}
