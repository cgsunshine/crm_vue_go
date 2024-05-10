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

type CrmProductGroupApi struct {
}

var crmProductGroupService = service.ServiceGroupApp.CrmServiceGroup.CrmProductGroupService


// CreateCrmProductGroup 创建产品分组
// @Tags CrmProductGroup
// @Summary 创建产品分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmProductGroup true "创建产品分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmProductGroup/createCrmProductGroup [post]
func (crmProductGroupApi *CrmProductGroupApi) CreateCrmProductGroup(c *gin.Context) {
	var crmProductGroup crm.CrmProductGroup
	err := c.ShouldBindJSON(&crmProductGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmProductGroupService.CreateCrmProductGroup(&crmProductGroup); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCrmProductGroup 删除产品分组
// @Tags CrmProductGroup
// @Summary 删除产品分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmProductGroup true "删除产品分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmProductGroup/deleteCrmProductGroup [delete]
func (crmProductGroupApi *CrmProductGroupApi) DeleteCrmProductGroup(c *gin.Context) {
	ID := c.Query("ID")
	if err := crmProductGroupService.DeleteCrmProductGroup(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCrmProductGroupByIds 批量删除产品分组
// @Tags CrmProductGroup
// @Summary 批量删除产品分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /crmProductGroup/deleteCrmProductGroupByIds [delete]
func (crmProductGroupApi *CrmProductGroupApi) DeleteCrmProductGroupByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := crmProductGroupService.DeleteCrmProductGroupByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCrmProductGroup 更新产品分组
// @Tags CrmProductGroup
// @Summary 更新产品分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmProductGroup true "更新产品分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmProductGroup/updateCrmProductGroup [put]
func (crmProductGroupApi *CrmProductGroupApi) UpdateCrmProductGroup(c *gin.Context) {
	var crmProductGroup crm.CrmProductGroup
	err := c.ShouldBindJSON(&crmProductGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := crmProductGroupService.UpdateCrmProductGroup(crmProductGroup); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCrmProductGroup 用id查询产品分组
// @Tags CrmProductGroup
// @Summary 用id查询产品分组
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmProductGroup true "用id查询产品分组"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmProductGroup/findCrmProductGroup [get]
func (crmProductGroupApi *CrmProductGroupApi) FindCrmProductGroup(c *gin.Context) {
	ID := c.Query("ID")
	if recrmProductGroup, err := crmProductGroupService.GetCrmProductGroup(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmProductGroup": recrmProductGroup}, c)
	}
}

// GetCrmProductGroupList 分页获取产品分组列表
// @Tags CrmProductGroup
// @Summary 分页获取产品分组列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmProductGroupSearch true "分页获取产品分组列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProductGroup/getCrmProductGroupList [get]
func (crmProductGroupApi *CrmProductGroupApi) GetCrmProductGroupList(c *gin.Context) {
	var pageInfo crmReq.CrmProductGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := crmProductGroupService.GetCrmProductGroupInfoList(pageInfo); err != nil {
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

// GetCrmProductGroupPublic 不需要鉴权的产品分组接口
// @Tags CrmProductGroup
// @Summary 不需要鉴权的产品分组接口
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmProductGroupSearch true "分页获取产品分组列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProductGroup/getCrmProductGroupList [get]
func (crmProductGroupApi *CrmProductGroupApi) GetCrmProductGroupPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的产品分组接口信息",
    }, "获取成功", c)
}
