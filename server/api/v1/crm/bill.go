package crm

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	crmReq "github.com/flipped-aurora/gin-vue-admin/server/model/crm/request"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

// GetCrmBillList 分页获取crmBill表列表
// @Tags CrmBill
// @Summary 分页获取crmBill表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crmReq.CrmBillSearch true "分页获取crmBill表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBill/getCrmBillList [get]
func (crmBillApi *CrmBillApi) GetCrmPageBillList(c *gin.Context) {
	var pageInfo crmReq.CrmBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = GetSearchUserId(pageInfo.UserId, c)

	if list, total, err := crmBillService.GetCrmPageBillInfoList(pageInfo); err != nil {
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

// FindPageCrmBill 用id查询crmBill表
// @Tags CrmBill
// @Summary 用id查询crmBill表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBill true "用id查询crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBill/findCrmBill [get]
func (crmBillApi *CrmBillApi) FindPageCrmBill(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBill, err := crmBillService.GetCrmPageBill(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recrmBill": recrmBill}, c)
	}
}

// DownloadPageCrmBillExcel 下载账单Excel
// @Tags CrmBill
// @Summary 下载账单Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query crm.CrmBill true "用id查询crmBill表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBill/downloadPageCrmBillExcel [get]
func (crmBillApi *CrmBillApi) DownloadPageCrmBillExcel(c *gin.Context) {
	ID := c.Query("ID")
	if recrmBill, err := crmBillService.GetCrmPageIdBill(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		// 打开一个Excel文件
		f, err := excelize.OpenFile("./public/DownloadTemplate/ASLINE_INVOICE.xlsx")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = f.SetCellValue("Sheet1", "I2", "#"+recrmBill.BillNumber)
		err = f.SetCellValue("Sheet1", "D6", "#"+recrmBill.CustomerCompany)
		err = f.SetCellValue("Sheet1", "D7", "#"+recrmBill.CustomerAddress)
		err = f.SetCellValue("Sheet1", "J6", *recrmBill.Amount)
		if recrmBill.ExpirationTime != nil {
			err = f.SetCellValue("Sheet1", "J8", recrmBill.ExpirationTime.Format("2006-01-02"))
		}
		err = f.SetCellValue("Sheet1", "B12", "账期？")
		if recrmBill.BillingStartTime != nil {
			err = f.SetCellValue("Sheet1", "D12", recrmBill.BillingStartTime.Format("2006-01-02"))
		}
		if recrmBill.BillingEndTime != nil {
			err = f.SetCellValue("Sheet1", "F12", recrmBill.BillingEndTime.Format("2006-01-02"))
		}

		err = f.SetCellValue("Sheet1", "H12", recrmBill.Username)
		err = f.SetCellValue("Sheet1", "K12", *recrmBill.DiscountRate)
		order, err := crmOrderService.GetCrmOrderId(recrmBill.OrderId)
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
			return
		}
		startTable := 15    //需要插入起始单元格
		startTableStep := 2 //插入单元格步长
		for i, product := range order.OrderProducts {
			//if i != 0 {
			//	// 在第startTable行之后插入两行，并复制第1行和第2行的数据
			//	err = f.InsertRows("Sheet1", startTable+i*startTableStep, 2)
			//}
			err = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", startTable+i*startTableStep), i+1)
			err = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", startTable+i*startTableStep), product.Product.ProductName)
			err = f.SetCellValue("Sheet1", fmt.Sprintf("F%d", startTable+i*startTableStep), product.Product.DataCenter)
			if product.Quantity != nil {
				err = f.SetCellValue("Sheet1", fmt.Sprintf("H%d", startTable+i*startTableStep), *product.Quantity)
			}
			if product.Product.DiscountPrice != nil {
				err = f.SetCellValue("Sheet1", fmt.Sprintf("J%d", startTable+i*startTableStep), *product.Product.DiscountPrice)
			}
			if product.Product.SalesPrice != nil {
				err = f.SetCellValue("Sheet1", fmt.Sprintf("K%d", startTable+i*startTableStep), *product.Product.SalesPrice)
			}
		}
		//重新设置统计函数
		//err = f.SetCellFormula("Sheet1", "A5", fmt.Sprintf("SUM(K15:K%d)", startTable+len(order.OrderProducts)*startTableStep-1))
		//导出文档，统计函数不执行 需处理
		//err = f.SetCellFormula("Sheet1", "A5", "SUM(K15:K20)")
		err = f.UpdateLinkedValue()
		fileName := fmt.Sprintf("%d%s.xlsx", time.Now().Unix(), recrmBill.BillNumber)
		// 保存到文件
		if err := f.SaveAs("./public/DownloadBill/" + fileName); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("文件已保存")

		// 打开文件
		file, err := os.ReadFile("./public/DownloadBill/" + fileName)
		if err != nil {
			fmt.Println(err)
			return
		}

		//获取到文件地址然后下载
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", recrmBill.BillNumber+".xlsx")) // 对下载的文件重命名
		c.Header("success", "true")
		c.Data(http.StatusOK, "application/octet-stream", file)
		//response.OkWithData(gin.H{"recrmBill": recrmBill}, c)
	}

}
