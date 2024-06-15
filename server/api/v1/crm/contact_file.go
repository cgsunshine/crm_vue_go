package crm

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/crm"
	crmRes "github.com/flipped-aurora/gin-vue-admin/server/model/crm/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// UpdateCrmContactFileSort
// @Tags CrmContactFileUploadAndDownloads
// @Summary 更新crmContactFileUploadAndDownloads表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body crm.CrmContactFileUploadAndDownloads true "更新crmContactFileUploadAndDownloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactFileUploadAndDownloads/updateCrmContactFileUploadAndDownloads [put]
func (CrmContactFileApi *CrmContactFileApi) UpdateCrmContactFileSort(c *gin.Context) {
	sorts, ok := c.GetPostFormArray("sorts")
	if ok {
		response.FailWithMessage("更新失败", c)
		return
	}
	for i, sort := range sorts {
		if err := crmContactFileService.Sort(i, sort); err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			response.FailWithMessage("更新失败", c)
			return
		}
	}
	response.OkWithMessage("更新成功", c)
}

// UploadFile
// @Tags      CrmContactFileUploadAndDownloads
// @Summary   合同文件上传
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                                                           true  "上传文件示例"
// @Success   200   {object}  response.Response{data=exampleRes.ExaFileResponse,msg=string}  "上传文件示例,返回包括文件详情"
// @Router    /fileUploadAndDownload/upload [post]
func (CrmContactFileApi *CrmContactFileApi) UploadFile(c *gin.Context) {
	var file crm.CrmContactFile
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	contact_id_str := c.PostForm("contact_id")
	contact_id, _ := strconv.Atoi(contact_id_str)

	sort_str := c.PostForm("sort")
	sort, _ := strconv.Atoi(sort_str)

	file, err = crmContactFileService.UploadFile(header, noSave, &contact_id, &sort) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(crmRes.CrmContactFileResponse{File: file}, "上传成功", c)
}

// DownloadFile
// @Tags      CrmContactFileUploadAndDownloads
// @Summary   合同文件下载
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                                                           true  "上传文件示例"
// @Success   200   {object}  response.Response{data=exampleRes.ExaFileResponse,msg=string}  "上传文件示例,返回包括文件详情"
// @Router    /fileUploadAndDownload/upload [post]
func (CrmContactFileApi *CrmContactFileApi) DownloadFile(c *gin.Context) {

	id, ok := c.GetQuery("id")
	if !ok {
		response.FailWithMessage("参数错误", c)
		return
	}
	file, file_name, err := crmContactFileService.DownloadFile(id) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file_name)) // 对下载的文件重命名
	c.Header("success", "true")
	c.Data(http.StatusOK, "application/octet-stream", file)

}
