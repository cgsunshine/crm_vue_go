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

type CrmApprovalRecordApi struct {
}

var crmApprovalRecordService = service.ServiceGroupApp.CrmServiceGroup.CrmApprovalRecordService


// CreateCrmApprovalRecord 创建审批记录
// @Tags CrmApprovalRecord
// @Summary 创建审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalRecord true "创建审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalRecord/createCrmApprovalRecord [post]
func (crmApprovalRecordApi *CrmApprovalRecordApi) CreateCrmApprovalRecord(c *gin.Context) {
	var crmApprovalRecord crm.CrmApprovalRecord
	err := c.ShouldBindJSON(&crmApprovalRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmApprovalRecordService.CreateCrmApprovalRecord(&crmApprovalRecord); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmApprovalRecord 删除审批记录
// @Tags CrmApprovalRecord
// @Summary 删除审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalRecord true "删除审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalRecord/deleteCrmApprovalRecord [delete]
func (crmApprovalRecordApi *CrmApprovalRecordApi) DeleteCrmApprovalRecord(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmApprovalRecordService.DeleteCrmApprovalRecord(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmApprovalRecordByIds 批量删除审批记录
// @Tags CrmApprovalRecord
// @Summary 批量删除审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmApprovalRecord/deleteCrmApprovalRecordByIds [delete]
func (crmApprovalRecordApi *CrmApprovalRecordApi) DeleteCrmApprovalRecordByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmApprovalRecordService.DeleteCrmApprovalRecordByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmApprovalRecord 更新审批记录
// @Tags CrmApprovalRecord
// @Summary 更新审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalRecord true "更新审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmApprovalRecord/updateCrmApprovalRecord [put]
func (crmApprovalRecordApi *CrmApprovalRecordApi) UpdateCrmApprovalRecord(c *gin.Context) {
	var crmApprovalRecord crm.CrmApprovalRecord
	err := c.ShouldBindJSON(&crmApprovalRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmApprovalRecordService.UpdateCrmApprovalRecord(crmApprovalRecord); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmApprovalRecord 用id查询审批记录
// @Tags CrmApprovalRecord
// @Summary 用id查询审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmApprovalRecord true "用id查询审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmApprovalRecord/findCrmApprovalRecord [get]
func (crmApprovalRecordApi *CrmApprovalRecordApi) FindCrmApprovalRecord(c *gin.Context) {
	ID := c.Query("ID")
	if recrmApprovalRecord, err := crmApprovalRecordService.GetCrmApprovalRecord(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmApprovalRecord": recrmApprovalRecord}, c)
	}
}

// GetCrmApprovalRecordList 分页获取审批记录列表
// @Tags CrmApprovalRecord
// @Summary 分页获取审批记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalRecordSearch true "分页获取审批记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalRecord/getCrmApprovalRecordList [get]
func (crmApprovalRecordApi *CrmApprovalRecordApi) GetCrmApprovalRecordList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmApprovalRecordService.GetCrmApprovalRecordInfoList(pageInfo); err != nil {
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

// GetCrmApprovalRecordPublic 不需要鉴权的审批记录接口
// @Tags CrmApprovalRecord
// @Summary 不需要鉴权的审批记录接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalRecordSearch true "分页获取审批记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalRecord/getCrmApprovalRecordList [get]
func (crmApprovalRecordApi *CrmApprovalRecordApi) GetCrmApprovalRecordPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的审批记录接口信息",
    }, "获取成功", c)
}
