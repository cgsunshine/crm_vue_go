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

type CrmContactApprovalRecordApi struct {
}

var crmContactApprovalRecordService = service.ServiceGroupApp.CrmServiceGroup.CrmContactApprovalRecordService


// CreateCrmContactApprovalRecord 创建合同审批记录
// @Tags CrmContactApprovalRecord
// @Summary 创建合同审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactApprovalRecord true "创建合同审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContactApprovalRecord/createCrmContactApprovalRecord [post]
func (crmContactApprovalRecordApi *CrmContactApprovalRecordApi) CreateCrmContactApprovalRecord(c *gin.Context) {
	var crmContactApprovalRecord crm.CrmContactApprovalRecord
	err := c.ShouldBindJSON(&crmContactApprovalRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmContactApprovalRecordService.CreateCrmContactApprovalRecord(&crmContactApprovalRecord); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmContactApprovalRecord 删除合同审批记录
// @Tags CrmContactApprovalRecord
// @Summary 删除合同审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactApprovalRecord true "删除合同审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContactApprovalRecord/deleteCrmContactApprovalRecord [delete]
func (crmContactApprovalRecordApi *CrmContactApprovalRecordApi) DeleteCrmContactApprovalRecord(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmContactApprovalRecordService.DeleteCrmContactApprovalRecord(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmContactApprovalRecordByIds 批量删除合同审批记录
// @Tags CrmContactApprovalRecord
// @Summary 批量删除合同审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmContactApprovalRecord/deleteCrmContactApprovalRecordByIds [delete]
func (crmContactApprovalRecordApi *CrmContactApprovalRecordApi) DeleteCrmContactApprovalRecordByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmContactApprovalRecordService.DeleteCrmContactApprovalRecordByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmContactApprovalRecord 更新合同审批记录
// @Tags CrmContactApprovalRecord
// @Summary 更新合同审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactApprovalRecord true "更新合同审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactApprovalRecord/updateCrmContactApprovalRecord [put]
func (crmContactApprovalRecordApi *CrmContactApprovalRecordApi) UpdateCrmContactApprovalRecord(c *gin.Context) {
	var crmContactApprovalRecord crm.CrmContactApprovalRecord
	err := c.ShouldBindJSON(&crmContactApprovalRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmContactApprovalRecordService.UpdateCrmContactApprovalRecord(crmContactApprovalRecord); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmContactApprovalRecord 用id查询合同审批记录
// @Tags CrmContactApprovalRecord
// @Summary 用id查询合同审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContactApprovalRecord true "用id查询合同审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactApprovalRecord/findCrmContactApprovalRecord [get]
func (crmContactApprovalRecordApi *CrmContactApprovalRecordApi) FindCrmContactApprovalRecord(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContactApprovalRecord, err := crmContactApprovalRecordService.GetCrmContactApprovalRecord(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContactApprovalRecord": recrmContactApprovalRecord}, c)
	}
}

// GetCrmContactApprovalRecordList 分页获取合同审批记录列表
// @Tags CrmContactApprovalRecord
// @Summary 分页获取合同审批记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContactApprovalRecordSearch true "分页获取合同审批记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactApprovalRecord/getCrmContactApprovalRecordList [get]
func (crmContactApprovalRecordApi *CrmContactApprovalRecordApi) GetCrmContactApprovalRecordList(c *gin.Context) {
	var pageInfo crmReq.CrmContactApprovalRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmContactApprovalRecordService.GetCrmContactApprovalRecordInfoList(pageInfo); err != nil {
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

// GetCrmContactApprovalRecordPublic 不需要鉴权的合同审批记录接口
// @Tags CrmContactApprovalRecord
// @Summary 不需要鉴权的合同审批记录接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContactApprovalRecordSearch true "分页获取合同审批记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactApprovalRecord/getCrmContactApprovalRecordList [get]
func (crmContactApprovalRecordApi *CrmContactApprovalRecordApi) GetCrmContactApprovalRecordPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的合同审批记录接口信息",
    }, "获取成功", c)
}
