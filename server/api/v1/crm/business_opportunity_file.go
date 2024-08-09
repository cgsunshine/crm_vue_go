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

// UploadFile
// @Tags      CrmStatementAccountFileApi
// @Summary   商机文件上传
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                                                           true  "上传文件示例"
// @Success   200   {object}  response.Response{data=crmRes.CrmBusinessOpportunityFileResponse,msg=string}  "上传文件示例,返回包括文件详情"
// @Router    /fileUploadAndDownload/upload [post]
func (b *CrmBusinessOpportunityFileApi) UploadFile(c *gin.Context) {
	var file crm.CrmBusinessOpportunityFile
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	business_opportunity_id_str := c.PostForm("business_opportunity_id")
	business_opportunity_id, _ := strconv.Atoi(business_opportunity_id_str)

	sort_str := c.PostForm("sort")
	sort, _ := strconv.Atoi(sort_str)

	file, err = crmBusinessOpportunityFileService.UploadFile(header, noSave, &business_opportunity_id, &sort) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(crmRes.CrmBusinessOpportunityFileResponse{File: file}, "上传成功", c)
}

// DownloadFile
// @Tags      CrmContactFileUploadAndDownloads
// @Summary   合同文件下载
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   octet-stream
// @Param     file  formData  file  true  "上传文件示例"
// @Success   200   "file"  "上传文件示例,返回包括文件详情"
// @Router    /fileUploadAndDownload/upload [post]
func (CrmContactFileApi *CrmBusinessOpportunityFileApi) DownloadFile(c *gin.Context) {

	id, ok := c.GetQuery("id")
	if !ok {
		response.FailWithMessage("参数错误", c)
		return
	}
	file, file_name, err := crmBusinessOpportunityFileService.DownloadFile(id) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file_name)) // 对下载的文件重命名
	c.Header("success", "true")
	c.Data(http.StatusOK, "application/octet-stream", file)
}
