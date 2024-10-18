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
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
func (crmContractApi *CrmContractApi) DownloadPageCrmContactExcel(c *gin.Context) {
	ID := c.Query("ID")
	if recrmContract, err := crmContractService.GetCrmPageContract(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		// 打开一个Excel文件
		f, err := excelize.OpenFile("./public/DownloadTemplate/ASLINE_ORDER_CONFIRMATION.xlsm")
		if err != nil {
			fmt.Println(err)
			return
		}
		//客户信息 A7 A8 A9 根据具体情况勾选一个
		//err = f.SetCellValue("OC", "A7", "#"+recrmContract.BillNumber)
		err = f.SetCellValue("OC", "G10", recrmContract.CustomerCompany)

		err = f.SetCellValue("OC", "E11", recrmContract.BusinessRegistrationNo)
		err = f.SetCellValue("OC", "E12", recrmContract.CustomerAddress)
		err = f.SetCellValue("OC", "E13", recrmContract.InstallationAddress2)
		err = f.SetCellValue("OC", "E14", recrmContract.InstallationAddress3)
		err = f.SetCellValue("OC", "D16", recrmContract.BillingAddress)
		err = f.SetCellValue("OC", "E15", recrmContract.CorrespondenceAddress)
		err = f.SetCellValue("OC", "D17", recrmContract.ContactPerson)
		err = f.SetCellValue("OC", "I17", recrmContract.Position)
		err = f.SetCellValue("OC", "K12", recrmContract.CorrespondenceEmailAddress)
		err = f.SetCellValue("OC", "C18", recrmContract.PositionTelNo)
		err = f.SetCellValue("OC", "G18", recrmContract.PositionFaxNo)
		err = f.SetCellValue("OC", "L18", recrmContract.MobileNo)
		err = f.SetCellValue("OC", "U18", recrmContract.EmergencyContactNo)
		err = f.SetCellValue("OC", "F19", recrmContract.TechnicalContactEmailAddress)
		err = f.SetCellValue("OC", "G22", recrmContract.EBillEmailAddressReceivingEBills)
		err = f.SetCellValue("OC", "A24", recrmContract.PaymentMethod)

		err = f.SetCellValue("OC", "E1", "测试")

		//合同对应订单，订单包含多个产品
		//这里可能会有多个，需要处理

		order, err := crmOrderService.GetCrmOrderId(recrmContract.OrderId)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
			return
		}
		startTable := 30    //需要插入起始单元格
		startTableStep := 1 //插入单元格步长

		dv := excelize.NewDataValidation(true)
		dv.Sqref = "D30:D47"
		err = dv.SetDropList([]string{"New", "IR", "ER"})
		dv.SetError(excelize.DataValidationErrorStyleStop, "输入错误", "只能填入：New / IR /ER 中一个")
		err = f.AddDataValidation("OC", dv)
		for i, product := range order.OrderProducts {
			//if i != 0 {
			//	// 在第startTable行之后插入两行，并复制第1行和第2行的数据
			//	err = f.InsertRows("OC", startTable+i*startTableStep, 2)
			//}
			err = f.SetCellValue("OC", fmt.Sprintf("A%d", startTable+i*startTableStep), i+1)
			err = f.SetCellValue("OC", fmt.Sprintf("C%d", startTable+i*startTableStep), *product.Quantity)
			err = f.SetCellValue("OC", fmt.Sprintf("D%d", startTable+i*startTableStep), "New") //三个选项要做成固定的 New / IR /ER
			err = f.SetCellValue("OC", fmt.Sprintf("F%d", startTable+i*startTableStep), product.Product.ProductName)
			err = f.SetCellValue("OC", fmt.Sprintf("R%d", startTable+i*startTableStep), 10)  //一次收费
			err = f.SetCellValue("OC", fmt.Sprintf("U%d", startTable+i*startTableStep), 260) //月收费
		}

		err = f.SetCellValue("OC", "J50", recrmContract.CreatedAt.Format("02/01/2006")) //"服务激活日期
		////重新设置统计函数
		//err = f.SetCellFormula("OC", "A5", fmt.Sprintf("SUM(K15:K%d)", startTable+len(order.OrderProducts)*startTableStep-1))
		err = f.UpdateLinkedValue()
		fileName := fmt.Sprintf("%d%s.xlsm", time.Now().Unix(), recrmContract.ContractNumber)
		// 保存到文件
		if err := f.SaveAs("./public/DownloadContact/" + fileName); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("文件已保存")

		// 打开文件
		file, err := os.ReadFile("./public/DownloadContact/" + fileName)
		if err != nil {
			fmt.Println(err)
			return
		}

		//获取到文件地址然后下载
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", recrmContract.ContractNumber+".xlsm")) // 对下载的文件重命名
		c.Header("success", "true")
		c.Data(http.StatusOK, "application/octet-stream", file)
	}

}
