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

type CrmPaymentCollentionApprovalRecordApi struct {
}

var crmPaymentCollentionApprovalRecordService = service.ServiceGroupApp.CrmServiceGroup.CrmPaymentCollentionApprovalRecordService


// CreateCrmPaymentCollentionApprovalRecord 创建回款审批记录
// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 创建回款审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPaymentCollentionApprovalRecord true "创建回款审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPaymentCollentionApprovalRecord/createCrmPaymentCollentionApprovalRecord [post]
func (crmPaymentCollentionApprovalRecordApi *CrmPaymentCollentionApprovalRecordApi) CreateCrmPaymentCollentionApprovalRecord(c *gin.Context) {
	var crmPaymentCollentionApprovalRecord crm.CrmPaymentCollentionApprovalRecord
	err := c.ShouldBindJSON(&crmPaymentCollentionApprovalRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmPaymentCollentionApprovalRecordService.CreateCrmPaymentCollentionApprovalRecord(&crmPaymentCollentionApprovalRecord); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmPaymentCollentionApprovalRecord 删除回款审批记录
// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 删除回款审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPaymentCollentionApprovalRecord true "删除回款审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPaymentCollentionApprovalRecord/deleteCrmPaymentCollentionApprovalRecord [delete]
func (crmPaymentCollentionApprovalRecordApi *CrmPaymentCollentionApprovalRecordApi) DeleteCrmPaymentCollentionApprovalRecord(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmPaymentCollentionApprovalRecordService.DeleteCrmPaymentCollentionApprovalRecord(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmPaymentCollentionApprovalRecordByIds 批量删除回款审批记录
// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 批量删除回款审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmPaymentCollentionApprovalRecord/deleteCrmPaymentCollentionApprovalRecordByIds [delete]
func (crmPaymentCollentionApprovalRecordApi *CrmPaymentCollentionApprovalRecordApi) DeleteCrmPaymentCollentionApprovalRecordByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmPaymentCollentionApprovalRecordService.DeleteCrmPaymentCollentionApprovalRecordByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmPaymentCollentionApprovalRecord 更新回款审批记录
// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 更新回款审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmPaymentCollentionApprovalRecord true "更新回款审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmPaymentCollentionApprovalRecord/updateCrmPaymentCollentionApprovalRecord [put]
func (crmPaymentCollentionApprovalRecordApi *CrmPaymentCollentionApprovalRecordApi) UpdateCrmPaymentCollentionApprovalRecord(c *gin.Context) {
	var crmPaymentCollentionApprovalRecord crm.CrmPaymentCollentionApprovalRecord
	err := c.ShouldBindJSON(&crmPaymentCollentionApprovalRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmPaymentCollentionApprovalRecordService.UpdateCrmPaymentCollentionApprovalRecord(crmPaymentCollentionApprovalRecord); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmPaymentCollentionApprovalRecord 用id查询回款审批记录
// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 用id查询回款审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmPaymentCollentionApprovalRecord true "用id查询回款审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPaymentCollentionApprovalRecord/findCrmPaymentCollentionApprovalRecord [get]
func (crmPaymentCollentionApprovalRecordApi *CrmPaymentCollentionApprovalRecordApi) FindCrmPaymentCollentionApprovalRecord(c *gin.Context) {
	ID := c.Query("ID")
	if recrmPaymentCollentionApprovalRecord, err := crmPaymentCollentionApprovalRecordService.GetCrmPaymentCollentionApprovalRecord(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmPaymentCollentionApprovalRecord": recrmPaymentCollentionApprovalRecord}, c)
	}
}

// GetCrmPaymentCollentionApprovalRecordList 分页获取回款审批记录列表
// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 分页获取回款审批记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPaymentCollentionApprovalRecordSearch true "分页获取回款审批记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPaymentCollentionApprovalRecord/getCrmPaymentCollentionApprovalRecordList [get]
func (crmPaymentCollentionApprovalRecordApi *CrmPaymentCollentionApprovalRecordApi) GetCrmPaymentCollentionApprovalRecordList(c *gin.Context) {
	var pageInfo crmReq.CrmPaymentCollentionApprovalRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmPaymentCollentionApprovalRecordService.GetCrmPaymentCollentionApprovalRecordInfoList(pageInfo); err != nil {
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

// GetCrmPaymentCollentionApprovalRecordPublic 不需要鉴权的回款审批记录接口
// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 不需要鉴权的回款审批记录接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmPaymentCollentionApprovalRecordSearch true "分页获取回款审批记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPaymentCollentionApprovalRecord/getCrmPaymentCollentionApprovalRecordList [get]
func (crmPaymentCollentionApprovalRecordApi *CrmPaymentCollentionApprovalRecordApi) GetCrmPaymentCollentionApprovalRecordPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的回款审批记录接口信息",
    }, "获取成功", c)
}
