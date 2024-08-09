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

type CrmUserApi struct {
}

var crmUserService = service.ServiceGroupApp.CrmServiceGroup.CrmUserService


// CreateCrmUser 创建crmUser表
// @Tags CrmUser
// @Summary 创建crmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmUser true "创建crmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmUser/createCrmUser [post]
func (crmUserApi *CrmUserApi) CreateCrmUser(c *gin.Context) {
	var crmUser crm.CrmUser
	err := c.ShouldBindJSON(&crmUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmUserService.CreateCrmUser(&crmUser); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmUser 删除crmUser表
// @Tags CrmUser
// @Summary 删除crmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmUser true "删除crmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmUser/deleteCrmUser [delete]
func (crmUserApi *CrmUserApi) DeleteCrmUser(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmUserService.DeleteCrmUser(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmUserByIds 批量删除crmUser表
// @Tags CrmUser
// @Summary 批量删除crmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmUser/deleteCrmUserByIds [delete]
func (crmUserApi *CrmUserApi) DeleteCrmUserByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmUserService.DeleteCrmUserByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmUser 更新crmUser表
// @Tags CrmUser
// @Summary 更新crmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmUser true "更新crmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmUser/updateCrmUser [put]
func (crmUserApi *CrmUserApi) UpdateCrmUser(c *gin.Context) {
	var crmUser crm.CrmUser
	err := c.ShouldBindJSON(&crmUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmUserService.UpdateCrmUser(crmUser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmUser 用id查询crmUser表
// @Tags CrmUser
// @Summary 用id查询crmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmUser true "用id查询crmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmUser/findCrmUser [get]
func (crmUserApi *CrmUserApi) FindCrmUser(c *gin.Context) {
	ID := c.Query("ID")
	if recrmUser, err := crmUserService.GetCrmUser(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmUser": recrmUser}, c)
	}
}

// GetCrmUserList 分页获取crmUser表列表
// @Tags CrmUser
// @Summary 分页获取crmUser表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmUserSearch true "分页获取crmUser表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmUser/getCrmUserList [get]
func (crmUserApi *CrmUserApi) GetCrmUserList(c *gin.Context) {
	var pageInfo crmReq.CrmUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmUserService.GetCrmUserInfoList(pageInfo); err != nil {
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

// GetCrmUserPublic 不需要鉴权的crmUser表接口
// @Tags CrmUser
// @Summary 不需要鉴权的crmUser表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmUserSearch true "分页获取crmUser表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmUser/getCrmUserList [get]
func (crmUserApi *CrmUserApi) GetCrmUserPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的crmUser表接口信息",
    }, "获取成功", c)
}
