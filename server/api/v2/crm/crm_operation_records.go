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

type CrmOperationRecordsApi struct {
}

var crmOperationRecordsService = service.ServiceGroupApp.CrmServiceGroup.CrmOperationRecordsService


// CreateCrmOperationRecords 创建操作记录
// @Tags CrmOperationRecords
// @Summary 创建操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmOperationRecords true "创建操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmOperationRecords/createCrmOperationRecords [post]
func (crmOperationRecordsApi *CrmOperationRecordsApi) CreateCrmOperationRecords(c *gin.Context) {
	var crmOperationRecords crm.CrmOperationRecords
	err := c.ShouldBindJSON(&crmOperationRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmOperationRecordsService.CreateCrmOperationRecords(&crmOperationRecords); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmOperationRecords 删除操作记录
// @Tags CrmOperationRecords
// @Summary 删除操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmOperationRecords true "删除操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmOperationRecords/deleteCrmOperationRecords [delete]
func (crmOperationRecordsApi *CrmOperationRecordsApi) DeleteCrmOperationRecords(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmOperationRecordsService.DeleteCrmOperationRecords(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmOperationRecordsByIds 批量删除操作记录
// @Tags CrmOperationRecords
// @Summary 批量删除操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmOperationRecords/deleteCrmOperationRecordsByIds [delete]
func (crmOperationRecordsApi *CrmOperationRecordsApi) DeleteCrmOperationRecordsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmOperationRecordsService.DeleteCrmOperationRecordsByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmOperationRecords 更新操作记录
// @Tags CrmOperationRecords
// @Summary 更新操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmOperationRecords true "更新操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmOperationRecords/updateCrmOperationRecords [put]
func (crmOperationRecordsApi *CrmOperationRecordsApi) UpdateCrmOperationRecords(c *gin.Context) {
	var crmOperationRecords crm.CrmOperationRecords
	err := c.ShouldBindJSON(&crmOperationRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmOperationRecordsService.UpdateCrmOperationRecords(crmOperationRecords); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmOperationRecords 用id查询操作记录
// @Tags CrmOperationRecords
// @Summary 用id查询操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmOperationRecords true "用id查询操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmOperationRecords/findCrmOperationRecords [get]
func (crmOperationRecordsApi *CrmOperationRecordsApi) FindCrmOperationRecords(c *gin.Context) {
	ID := c.Query("ID")
	if recrmOperationRecords, err := crmOperationRecordsService.GetCrmOperationRecords(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmOperationRecords": recrmOperationRecords}, c)
	}
}

// GetCrmOperationRecordsList 分页获取操作记录列表
// @Tags CrmOperationRecords
// @Summary 分页获取操作记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmOperationRecordsSearch true "分页获取操作记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmOperationRecords/getCrmOperationRecordsList [get]
func (crmOperationRecordsApi *CrmOperationRecordsApi) GetCrmOperationRecordsList(c *gin.Context) {
	var pageInfo crmReq.CrmOperationRecordsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmOperationRecordsService.GetCrmOperationRecordsInfoList(pageInfo); err != nil {
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

// GetCrmOperationRecordsPublic 不需要鉴权的操作记录接口
// @Tags CrmOperationRecords
// @Summary 不需要鉴权的操作记录接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmOperationRecordsSearch true "分页获取操作记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmOperationRecords/getCrmOperationRecordsList [get]
func (crmOperationRecordsApi *CrmOperationRecordsApi) GetCrmOperationRecordsPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的操作记录接口信息",
    }, "获取成功", c)
}
