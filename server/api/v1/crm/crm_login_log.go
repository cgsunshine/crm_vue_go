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

type CrmLoginLogApi struct {
}

var crmLoginLogService = service.ServiceGroupApp.CrmServiceGroup.CrmLoginLogService


// CreateCrmLoginLog 创建登录日志
// @Tags CrmLoginLog
// @Summary 创建登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmLoginLog true "创建登录日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmLoginLog/createCrmLoginLog [post]
func (crmLoginLogApi *CrmLoginLogApi) CreateCrmLoginLog(c *gin.Context) {
	var crmLoginLog crm.CrmLoginLog
	err := c.ShouldBindJSON(&crmLoginLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmLoginLogService.CreateCrmLoginLog(&crmLoginLog); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmLoginLog 删除登录日志
// @Tags CrmLoginLog
// @Summary 删除登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmLoginLog true "删除登录日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmLoginLog/deleteCrmLoginLog [delete]
func (crmLoginLogApi *CrmLoginLogApi) DeleteCrmLoginLog(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmLoginLogService.DeleteCrmLoginLog(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmLoginLogByIds 批量删除登录日志
// @Tags CrmLoginLog
// @Summary 批量删除登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmLoginLog/deleteCrmLoginLogByIds [delete]
func (crmLoginLogApi *CrmLoginLogApi) DeleteCrmLoginLogByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmLoginLogService.DeleteCrmLoginLogByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmLoginLog 更新登录日志
// @Tags CrmLoginLog
// @Summary 更新登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmLoginLog true "更新登录日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmLoginLog/updateCrmLoginLog [put]
func (crmLoginLogApi *CrmLoginLogApi) UpdateCrmLoginLog(c *gin.Context) {
	var crmLoginLog crm.CrmLoginLog
	err := c.ShouldBindJSON(&crmLoginLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmLoginLogService.UpdateCrmLoginLog(crmLoginLog); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmLoginLog 用id查询登录日志
// @Tags CrmLoginLog
// @Summary 用id查询登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmLoginLog true "用id查询登录日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmLoginLog/findCrmLoginLog [get]
func (crmLoginLogApi *CrmLoginLogApi) FindCrmLoginLog(c *gin.Context) {
	ID := c.Query("ID")
	if recrmLoginLog, err := crmLoginLogService.GetCrmLoginLog(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmLoginLog": recrmLoginLog}, c)
	}
}

// GetCrmLoginLogList 分页获取登录日志列表
// @Tags CrmLoginLog
// @Summary 分页获取登录日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmLoginLogSearch true "分页获取登录日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmLoginLog/getCrmLoginLogList [get]
func (crmLoginLogApi *CrmLoginLogApi) GetCrmLoginLogList(c *gin.Context) {
	var pageInfo crmReq.CrmLoginLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmLoginLogService.GetCrmLoginLogInfoList(pageInfo); err != nil {
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

// GetCrmLoginLogPublic 不需要鉴权的登录日志接口
// @Tags CrmLoginLog
// @Summary 不需要鉴权的登录日志接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmLoginLogSearch true "分页获取登录日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmLoginLog/getCrmLoginLogList [get]
func (crmLoginLogApi *CrmLoginLogApi) GetCrmLoginLogPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的登录日志接口信息",
    }, "获取成功", c)
}
