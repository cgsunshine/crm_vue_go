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

type CrmApprovalNodeApi struct {
}

var crmApprovalNodeService = service.ServiceGroupApp.CrmServiceGroup.CrmApprovalNodeService


// CreateCrmApprovalNode 创建crmApprovalNode表
// @Tags CrmApprovalNode
// @Summary 创建crmApprovalNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalNode true "创建crmApprovalNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalNode/createCrmApprovalNode [post]
func (crmApprovalNodeApi *CrmApprovalNodeApi) CreateCrmApprovalNode(c *gin.Context) {
	var crmApprovalNode crm.CrmApprovalNode
	err := c.ShouldBindJSON(&crmApprovalNode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmApprovalNodeService.CreateCrmApprovalNode(&crmApprovalNode); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmApprovalNode 删除crmApprovalNode表
// @Tags CrmApprovalNode
// @Summary 删除crmApprovalNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalNode true "删除crmApprovalNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalNode/deleteCrmApprovalNode [delete]
func (crmApprovalNodeApi *CrmApprovalNodeApi) DeleteCrmApprovalNode(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmApprovalNodeService.DeleteCrmApprovalNode(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmApprovalNodeByIds 批量删除crmApprovalNode表
// @Tags CrmApprovalNode
// @Summary 批量删除crmApprovalNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmApprovalNode/deleteCrmApprovalNodeByIds [delete]
func (crmApprovalNodeApi *CrmApprovalNodeApi) DeleteCrmApprovalNodeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmApprovalNodeService.DeleteCrmApprovalNodeByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmApprovalNode 更新crmApprovalNode表
// @Tags CrmApprovalNode
// @Summary 更新crmApprovalNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmApprovalNode true "更新crmApprovalNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmApprovalNode/updateCrmApprovalNode [put]
func (crmApprovalNodeApi *CrmApprovalNodeApi) UpdateCrmApprovalNode(c *gin.Context) {
	var crmApprovalNode crm.CrmApprovalNode
	err := c.ShouldBindJSON(&crmApprovalNode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmApprovalNodeService.UpdateCrmApprovalNode(crmApprovalNode); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmApprovalNode 用id查询crmApprovalNode表
// @Tags CrmApprovalNode
// @Summary 用id查询crmApprovalNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmApprovalNode true "用id查询crmApprovalNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmApprovalNode/findCrmApprovalNode [get]
func (crmApprovalNodeApi *CrmApprovalNodeApi) FindCrmApprovalNode(c *gin.Context) {
	ID := c.Query("ID")
	if recrmApprovalNode, err := crmApprovalNodeService.GetCrmApprovalNode(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmApprovalNode": recrmApprovalNode}, c)
	}
}

// GetCrmApprovalNodeList 分页获取crmApprovalNode表列表
// @Tags CrmApprovalNode
// @Summary 分页获取crmApprovalNode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalNodeSearch true "分页获取crmApprovalNode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalNode/getCrmApprovalNodeList [get]
func (crmApprovalNodeApi *CrmApprovalNodeApi) GetCrmApprovalNodeList(c *gin.Context) {
	var pageInfo crmReq.CrmApprovalNodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmApprovalNodeService.GetCrmApprovalNodeInfoList(pageInfo); err != nil {
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

// GetCrmApprovalNodePublic 不需要鉴权的crmApprovalNode表接口
// @Tags CrmApprovalNode
// @Summary 不需要鉴权的crmApprovalNode表接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmApprovalNodeSearch true "分页获取crmApprovalNode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalNode/getCrmApprovalNodeList [get]
func (crmApprovalNodeApi *CrmApprovalNodeApi) GetCrmApprovalNodePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的crmApprovalNode表接口信息",
    }, "获取成功", c)
}
