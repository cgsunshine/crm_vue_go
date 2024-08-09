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

type CrmCustomerGroupApi struct {
}

var crmCustomerGroupService = service.ServiceGroupApp.CrmServiceGroup.CrmCustomerGroupService


// CreateCrmCustomerGroup 创建crmCustomerGroup表
// @Tags CrmCustomerGroup
// @Summary 创建crmCustomerGroup表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCustomerGroup true "创建crmCustomerGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmCustomerGroup/createCrmCustomerGroup [post]
func (crmCustomerGroupApi *CrmCustomerGroupApi) CreateCrmCustomerGroup(c *gin.Context) {
	var crmCustomerGroup crm.CrmCustomerGroup
	err := c.ShouldBindJSON(&crmCustomerGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmCustomerGroupService.CreateCrmCustomerGroup(&crmCustomerGroup); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmCustomerGroup 删除crmCustomerGroup表
// @Tags CrmCustomerGroup
// @Summary 删除crmCustomerGroup表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCustomerGroup true "删除crmCustomerGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCustomerGroup/deleteCrmCustomerGroup [delete]
func (crmCustomerGroupApi *CrmCustomerGroupApi) DeleteCrmCustomerGroup(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmCustomerGroupService.DeleteCrmCustomerGroup(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmCustomerGroupByIds 批量删除crmCustomerGroup表
// @Tags CrmCustomerGroup
// @Summary 批量删除crmCustomerGroup表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmCustomerGroup/deleteCrmCustomerGroupByIds [delete]
func (crmCustomerGroupApi *CrmCustomerGroupApi) DeleteCrmCustomerGroupByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmCustomerGroupService.DeleteCrmCustomerGroupByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmCustomerGroup 更新crmCustomerGroup表
// @Tags CrmCustomerGroup
// @Summary 更新crmCustomerGroup表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmCustomerGroup true "更新crmCustomerGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmCustomerGroup/updateCrmCustomerGroup [put]
func (crmCustomerGroupApi *CrmCustomerGroupApi) UpdateCrmCustomerGroup(c *gin.Context) {
	var crmCustomerGroup crm.CrmCustomerGroup
	err := c.ShouldBindJSON(&crmCustomerGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmCustomerGroupService.UpdateCrmCustomerGroup(crmCustomerGroup); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmCustomerGroup 用id查询crmCustomerGroup表
// @Tags CrmCustomerGroup
// @Summary 用id查询crmCustomerGroup表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmCustomerGroup true "用id查询crmCustomerGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmCustomerGroup/findCrmCustomerGroup [get]
func (crmCustomerGroupApi *CrmCustomerGroupApi) FindCrmCustomerGroup(c *gin.Context) {
	ID := c.Query("ID")
	if recrmCustomerGroup, err := crmCustomerGroupService.GetCrmCustomerGroup(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmCustomerGroup": recrmCustomerGroup}, c)
	}
}

// GetCrmCustomerGroupList 分页获取crmCustomerGroup表列表
// @Tags CrmCustomerGroup
// @Summary 分页获取crmCustomerGroup表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmCustomerGroupSearch true "分页获取crmCustomerGroup表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCustomerGroup/getCrmCustomerGroupList [get]
func (crmCustomerGroupApi *CrmCustomerGroupApi) GetCrmCustomerGroupList(c *gin.Context) {
	var pageInfo crmReq.CrmCustomerGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmCustomerGroupService.GetCrmCustomerGroupInfoList(pageInfo); err != nil {
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

// GetCrmCustomerGroupPublic 不需要鉴权的crmCustomerGroup表接口
// @Tags CrmCustomerGroup
// @Summary 不需要鉴权的crmCustomerGroup表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmCustomerGroupSearch true "分页获取crmCustomerGroup表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCustomerGroup/getCrmCustomerGroupList [get]
func (crmCustomerGroupApi *CrmCustomerGroupApi) GetCrmCustomerGroupPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的crmCustomerGroup表接口信息",
    }, "获取成功", c)
}
