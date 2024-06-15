package crm

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetCrmPageProcurementContractList 分页获取订购合同列表
// @Tags CrmProcurementContract
// @Summary 分页获取订购合同列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmProcurementContractSearch true "分页获取订购合同列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmProcurementContract/getCrmProcurementContractList [get]
func (crmProcurementContractApi *CrmProcurementContractApi) GetCrmPageProcurementContractList(c *gin.Context) {
	var pageInfo crmReq.CrmProcurementContractSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID, _ := strconv.Atoi(c.GetHeader("X-User-Id"))

	pageInfo.UserId = userService.FindUserDataStatusById(&userID)

	if list, total, err := crmProcurementContractService.GetCrmProcurementContractInfoList(pageInfo); err != nil {
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

// FindCrmPageProcurementContract 用id查询订购合同
// @Tags CrmProcurementContract
// @Summary 用id查询订购合同
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmProcurementContract true "用id查询订购合同"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmProcurementContract/findCrmProcurementContract [get]
func (crmProcurementContractApi *CrmProcurementContractApi) FindCrmPageProcurementContract(c *gin.Context) {
	ID := c.Query("ID")
	if recrmProcurementContract, err := crmProcurementContractService.GetCrmProcurementContract(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmProcurementContract": recrmProcurementContract}, c)
	}
}
