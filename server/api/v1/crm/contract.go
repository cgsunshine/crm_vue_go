package crm

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/comm"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"strconv"
	"strings"
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

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)

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

	crmContract.UserId = comm.GetHeaderUserId(c)

	crmContract.ContractStatus = "1"
	crmContract.ReviewStatus = comm.Approval_Status_Pending
	crmContract.ReviewResult = "1"
	//生成编号
	crmContract.ContractNumber = comm.GetBusinessNumber(comm.ContractNumberPrefix, crmContractService.GetCrmBillTodayCount())
	//先查询合同审批对应的流程

	//在查出第一步对应的角色id
	roleInfo, err := crmConfigService.GetCrmNameToConfig(comm.ContractApproval)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	//ids, err := userService.GetRoleUsers(roleInfo.RoleIds)
	//if err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//	return
	//}

	if err := crmContractService.CreateCrmContract(&crmContract); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		contactId := int(crmContract.ID)
		//插入角色id对应的用户的审批记录
		//for _, userAuth := range ids {
		//	assigneeId := int(userAuth.SysUserId)
		//
		//	if err := crmApprovalTasksService.CreateCrmApprovalTasks(&crm.CrmApprovalTasks{
		//		AssigneeId:     &assigneeId,
		//		ApprovalStatus: comm.Approval_Status_Under,
		//		AssociatedId:   &contactId,
		//		Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
		//		StepId:         roleInfo.NodeId,
		//		ApprovalType:   utils.Pointer(comm.ContractApprovalType),
		//	}); err != nil {
		//		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		//		response.FailWithMessage("创建失败", c)
		//		return
		//	}
		//}

		items := strings.Split(roleInfo.RoleIds, ",")
		for _, item := range items {
			roleId, _ := strconv.Atoi(item)
			if err := crmApprovalTasksRoleService.CreateCrmApprovalTasksRole(&crm.CrmApprovalTasksRole{
				RoleId:         &roleId,
				ApprovalStatus: comm.Approval_Status_Under,
				AssociatedId:   &contactId,
				Valid:          utils.Pointer(comm.Contact_Approval_Tasks_valid_Effective),
				StepId:         roleInfo.NodeId,
				ApprovalType:   utils.Pointer(comm.ContractApprovalType),
			}); err != nil {
				global.GVA_LOG.Error("创建失败!", zap.Error(err))
				response.FailWithMessage("创建失败", c)
				return
			}
		}
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

// FindCrmPageFileContract 用id查询crmContract表 获取相关文件图片
// @Tags CrmContract
// @Summary 用id查询crmContract表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContract true "用id查询crmContract表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContract/findCrmContract [get]
func (crmContractApi *CrmContractApi) FindCrmPageFileContract(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContract, err := crmContractService.GetCrmPageContract(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		//查询关联文件
		list, _, err := fileUploadAndDownloadService.GetFileRecordInfoIdsList(recrmContract.ContractFile)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
		response.OkWithData(gin.H{"recrmContract": list}, c)
	}
}

// DownloadCrmPageFileContractExcel 用id查询crmContract表 下载合同Excel
// @Tags CrmContract
// @Summary 下载合同Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmContract true "用id查询crmContract表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContract/downloadCrmPageFileContractExcel [get]
func (crmContractApi *CrmContractApi) DownloadCrmPageFileContractExcel(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContract, err := crmContractService.GetCrmPageContract(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		//查询关联文件
		list, _, err := fileUploadAndDownloadService.GetFileRecordInfoIdsList(recrmContract.ContractFile)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
		response.OkWithData(gin.H{"recrmContract": list}, c)
	}
}

// DownloadPageCrmContactExcel 下载账单Contact
// @Tags CrmContract
// @Summary 下载账单Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBill true "用id查询crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBill/downloadPageCrmContact [get]
func (crmBillApi *CrmBillApi) DownloadPageCrmContactExcel(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContract, err := crmContractService.GetCrmPageContract(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		// 打开一个Excel文件
		f, err := excelize.OpenFile("ASLINE_ORDER_CONFIRMATION.xlsm")
		if err != nil {
			fmt.Println(err)
			return
		}
		//客户信息 A7 A8 A9 根据具体情况勾选一个
		//err = f.SetCellValue("Sheet1", "A7", "#"+recrmContract.BillNumber)
		err = f.SetCellValue("Sheet1", "G10", recrmContract.CustomerCompany)

		err = f.SetCellValue("Sheet1", "E11", "公司注册号？")
		err = f.SetCellValue("Sheet1", "E12", recrmContract.CustomerAddress)
		err = f.SetCellValue("Sheet1", "E13", "公司地址2")
		err = f.SetCellValue("Sheet1", "E14", "公司地址3")
		err = f.SetCellValue("Sheet1", "D16", "账单寄送地址")
		err = f.SetCellValue("Sheet1", "E15", "通讯地址")
		err = f.SetCellValue("Sheet1", "D17", "联系人")
		err = f.SetCellValue("Sheet1", "I17", "职位")
		err = f.SetCellValue("Sheet1", "K12", "通信电子邮件地址")
		err = f.SetCellValue("Sheet1", "C18", "电话")
		err = f.SetCellValue("Sheet1", "G18", "传真")
		err = f.SetCellValue("Sheet1", "L18", "移动电话")
		err = f.SetCellValue("Sheet1", "U18", "紧急联系电话")
		err = f.SetCellValue("Sheet1", "F19", "技术人员邮箱")
		err = f.SetCellValue("Sheet1", "G22", "账单邮件地址")
		err = f.SetCellValue("Sheet1", "A24", "支付方式")

		//这里可能会有多个，需要处理
		err = f.SetCellValue("Sheet1", "A30", "编号")
		err = f.SetCellValue("Sheet1", "C30", "数量")
		err = f.SetCellValue("Sheet1", "D30", "New")
		err = f.SetCellValue("Sheet1", "F30", "服务内容")
		err = f.SetCellValue("Sheet1", "R30", "一次收费")
		err = f.SetCellValue("Sheet1", "U30", "月付金额")

		err = f.SetCellValue("Sheet1", "J50", "服务激活日期")
		order, err := crmOrderService.GetCrmOrderId(recrmContract.OrderId)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
			return
		}
		startTable := 30    //需要插入起始单元格
		startTableStep := 1 //插入单元格步长
		for i, product := range order.OrderProducts {
			if i != 0 {
				// 在第startTable行之后插入两行，并复制第1行和第2行的数据
				err = f.InsertRows("Sheet1", startTable+i*startTableStep, 2)
			}
			err = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", startTable+i*startTableStep), i)
			err = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", startTable+i*startTableStep), product.Product.ProductName)
			err = f.SetCellValue("Sheet1", fmt.Sprintf("F%d", startTable+i*startTableStep), product.Product.DataCenter)
			err = f.SetCellValue("Sheet1", fmt.Sprintf("H%d", startTable+i*startTableStep), product.Quantity)
			err = f.SetCellValue("Sheet1", fmt.Sprintf("J%d", startTable+i*startTableStep), product.Product.DiscountPrice)
		}
		//重新设置统计函数
		err = f.SetCellFormula("Sheet1", "A5", fmt.Sprintf("SUM(K15:K%d)", startTable+len(order.OrderProducts)*startTableStep-1))
		// 保存到文件
		if err := f.SaveAs("example.xlsx"); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("文件已保存")
		//获取到文件地址然后下载
		//response.OkWithData(gin.H{"recrmContract": recrmContract}, c)
		//c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "file_name")) // 对下载的文件重命名
		//c.Header("success", "true")
		//c.Data(http.StatusOK, "application/octet-stream", file)
	}

}
