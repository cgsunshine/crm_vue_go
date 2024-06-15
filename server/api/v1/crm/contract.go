package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetCrmContractList 分页获取crmContract表列表
// @Tags CrmContract
// @Summary 分页获取crmContract表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmContractSearch true "分页获取crmContract表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContract/getCrmContractList [get]
func (crmContractApi *CrmContractApi) GetCrmPageContractList(c *gin.Context) {
	var pageInfo crmReq.CrmContractSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	pageInfo.UserId = userService.FindUserDataStatusById(&userID)

	if list, total, err := crmContractService.GetCrmPageContractInfoList(pageInfo); err != nil {
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

// CreateCrmContract 创建crmContract表
// @Tags CrmContract
// @Summary 创建crmContract表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContract true "创建crmContract表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContract/createCrmContract [post]
func (crmContractApi *CrmContractApi) CreatePageCrmContract(c *gin.Context) {
	var crmContract crm.CrmContract
	err := c.ShouldBindJSON(&crmContract)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	crmContract.ContractStatus = "1"
	crmContract.ReviewStatus = "1"
	crmContract.ReviewResult = "1"

	if err := crmContractService.CreateCrmContract(&crmContract); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// FindPageCrmContract 用id查询crmContract表
// @Tags CrmContract
// @Summary 用id查询crmContract表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContract true "用id查询crmContract表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContract/findCrmContract [get]
func (crmContractApi *CrmContractApi) FindCrmPageContract(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContract, err := crmContractService.GetCrmPageContract(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmContract": recrmContract}, c)
	}
}
