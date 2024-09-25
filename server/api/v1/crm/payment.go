package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// GetCrmPaymentList 分页获取crmPayment表列表
// @Tags CrmPayment
// @Summary 分页获取crmPayment表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPaymentSearch true "分页获取crmPayment表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPayment/getCrmPaymentList [get]
func (crmPaymentApi *CrmPaymentApi) GetCrmPagePaymentList(c *gin.Context) {
	var pageInfo crmReq.CrmPaymentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)

	if list, total, err := crmPaymentService.GetCrmPagePaymentInfoList(pageInfo); err != nil {
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

// FindCrmPagePayment 用id查询crmPayment表
// @Tags CrmPayment
// @Summary 用id查询crmPayment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmPayment true "用id查询crmPayment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPayment/findCrmPayment [get]
func (crmPaymentApi *CrmPaymentApi) FindCrmPagePayment(c *gin.Context) {
	ID := c.Query("ID")
	if recrmPayment, err := crmPaymentService.GetCrmPagePayment(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {

		recrmPayment.PaymentVoucher, _, err = fileUploadAndDownloadService.GetFileRecordInfoIdsString(recrmPayment.PaymentVoucher)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
		response.OkWithData(gin.H{"recrmPayment": recrmPayment}, c)
	}
}

// PaymentCompleted 付款完成上传凭证修改状态
// @Tags CrmPayment
// @Summary 用id查询crmPayment表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmPayment true "用id查询crmPayment表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPayment/paymentCompleted [get]
func (crmPaymentApi *CrmPaymentApi) PaymentCompleted(c *gin.Context) {
	//ID := c.Query("ID")
	//paymentVoucher := c.Query("paymentVoucher")
	var crmPayment crm.CrmPayment
	err := c.ShouldBindJSON(&crmPayment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	id := int(crmPayment.ID)
	if recrmPayment, err := crmPaymentService.GetCrmPaymentIdInfo(&id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		if recrmPayment.ReviewStatus != comm.Approval_Status_Pass {
			response.FailWithMessage("必须审批通过才能提付款", c)
			return
		}

		//修改付款管理付款状态
		err = crmPaymentService.UpdApprovalStatus(&id, map[string]interface{}{
			"payment_status":  comm.PaymentStatusPaid,
			"payment_voucher": crmPayment.PaymentVoucher,
			"payment_time":    time.Now(),
		})
		if err != nil {
			global.GVA_LOG.Error("上传失败!", zap.Error(err))
			response.FailWithMessage("上传失败", c)
			return
		}
		//修改应付账单付款状态
		err = crmBillPaymentService.UpdApprovalStatus(recrmPayment.BillPaymentId, map[string]interface{}{
			"payment_status": comm.PaymentStatusPaid,
		})
		if err != nil {
			global.GVA_LOG.Error("上传失败!", zap.Error(err))
			response.FailWithMessage("上传失败", c)
			return
		}
		//修改对账单付款状态
		err = crmStatementAccountService.UpdApprovalStatus(recrmPayment.StatementAccountId, map[string]interface{}{
			"payment_status": comm.PaymentStatusPaid,
		})
		if err != nil {
			global.GVA_LOG.Error("上传失败!", zap.Error(err))
			response.FailWithMessage("上传失败", c)
			return
		}
		response.OkWithData(gin.H{"recrmPayment": recrmPayment}, c)
	}
}
