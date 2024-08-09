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

type CrmBusinessOpportunityApprovalRecordApi struct {
}

var crmBusinessOpportunityApprovalRecordService = service.ServiceGroupApp.CrmServiceGroup.CrmBusinessOpportunityApprovalRecordService


// CreateCrmBusinessOpportunityApprovalRecord 创建商机审批记录
// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 创建商机审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityApprovalRecord true "创建商机审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/createCrmBusinessOpportunityApprovalRecord [post]
func (crmBusinessOpportunityApprovalRecordApi *CrmBusinessOpportunityApprovalRecordApi) CreateCrmBusinessOpportunityApprovalRecord(c *gin.Context) {
	var crmBusinessOpportunityApprovalRecord crm.CrmBusinessOpportunityApprovalRecord
	err := c.ShouldBindJSON(&crmBusinessOpportunityApprovalRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBusinessOpportunityApprovalRecordService.CreateCrmBusinessOpportunityApprovalRecord(&crmBusinessOpportunityApprovalRecord); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmBusinessOpportunityApprovalRecord 删除商机审批记录
// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 删除商机审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityApprovalRecord true "删除商机审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/deleteCrmBusinessOpportunityApprovalRecord [delete]
func (crmBusinessOpportunityApprovalRecordApi *CrmBusinessOpportunityApprovalRecordApi) DeleteCrmBusinessOpportunityApprovalRecord(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmBusinessOpportunityApprovalRecordService.DeleteCrmBusinessOpportunityApprovalRecord(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmBusinessOpportunityApprovalRecordByIds 批量删除商机审批记录
// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 批量删除商机审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/deleteCrmBusinessOpportunityApprovalRecordByIds [delete]
func (crmBusinessOpportunityApprovalRecordApi *CrmBusinessOpportunityApprovalRecordApi) DeleteCrmBusinessOpportunityApprovalRecordByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmBusinessOpportunityApprovalRecordService.DeleteCrmBusinessOpportunityApprovalRecordByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmBusinessOpportunityApprovalRecord 更新商机审批记录
// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 更新商机审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmBusinessOpportunityApprovalRecord true "更新商机审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/updateCrmBusinessOpportunityApprovalRecord [put]
func (crmBusinessOpportunityApprovalRecordApi *CrmBusinessOpportunityApprovalRecordApi) UpdateCrmBusinessOpportunityApprovalRecord(c *gin.Context) {
	var crmBusinessOpportunityApprovalRecord crm.CrmBusinessOpportunityApprovalRecord
	err := c.ShouldBindJSON(&crmBusinessOpportunityApprovalRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmBusinessOpportunityApprovalRecordService.UpdateCrmBusinessOpportunityApprovalRecord(crmBusinessOpportunityApprovalRecord); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmBusinessOpportunityApprovalRecord 用id查询商机审批记录
// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 用id查询商机审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBusinessOpportunityApprovalRecord true "用id查询商机审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/findCrmBusinessOpportunityApprovalRecord [get]
func (crmBusinessOpportunityApprovalRecordApi *CrmBusinessOpportunityApprovalRecordApi) FindCrmBusinessOpportunityApprovalRecord(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBusinessOpportunityApprovalRecord, err := crmBusinessOpportunityApprovalRecordService.GetCrmBusinessOpportunityApprovalRecord(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmBusinessOpportunityApprovalRecord": recrmBusinessOpportunityApprovalRecord}, c)
	}
}

// GetCrmBusinessOpportunityApprovalRecordList 分页获取商机审批记录列表
// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 分页获取商机审批记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunityApprovalRecordSearch true "分页获取商机审批记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/getCrmBusinessOpportunityApprovalRecordList [get]
func (crmBusinessOpportunityApprovalRecordApi *CrmBusinessOpportunityApprovalRecordApi) GetCrmBusinessOpportunityApprovalRecordList(c *gin.Context) {
	var pageInfo crmReq.CrmBusinessOpportunityApprovalRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmBusinessOpportunityApprovalRecordService.GetCrmBusinessOpportunityApprovalRecordInfoList(pageInfo); err != nil {
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

// GetCrmBusinessOpportunityApprovalRecordPublic 不需要鉴权的商机审批记录接口
// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 不需要鉴权的商机审批记录接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBusinessOpportunityApprovalRecordSearch true "分页获取商机审批记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/getCrmBusinessOpportunityApprovalRecordList [get]
func (crmBusinessOpportunityApprovalRecordApi *CrmBusinessOpportunityApprovalRecordApi) GetCrmBusinessOpportunityApprovalRecordPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的商机审批记录接口信息",
    }, "获取成功", c)
}
