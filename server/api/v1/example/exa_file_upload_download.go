package example

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleRes "github.com/flipped-aurora/gin-vue-admin/server/model/example/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type FileUploadAndDownloadApi struct{}

// UploadFile
// @Tags      ExaFileUploadAndDownload
// @Summary   上传文件示例
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                                                           true  "上传文件示例"
// @Success   200   {object}  response.Response{data=exampleRes.ExaFileResponse,msg=string}  "上传文件示例,返回包括文件详情"
// @Router    /fileUploadAndDownload/upload [post]
func (b *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = fileUploadAndDownloadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(exampleRes.ExaFileResponse{File: file}, "上传成功", c)
}

//// UploadFile
//// @Tags      ExaFileUploadAndDownload
//// @Summary   合同文件上传
//// @Security  ApiKeyAuth
//// @accept    multipart/form-data
//// @Produce   application/json
//// @Param     file  formData  file                                                           true  "上传文件示例"
//// @Success   200   {object}  response.Response{data=exampleRes.ExaFileResponse,msg=string}  "上传文件示例,返回包括文件详情"
//// @Router    /fileUploadAndDownload/upload [post]
//func (b *FileUploadAndDownloadApi) UploadContactFile(c *gin.Context) {
//	//var file crm.CrmContactFile
//	//noSave := c.DefaultQuery("noSave", "0")
//	//_, header, err := c.Request.FormFile("file")
//	//if err != nil {
//	//	global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
//	//	response.FailWithMessage("接收文件失败", c)
//	//	return
//	//}
//	//contact_id_str := c.PostForm("contact_id")
//	//contact_id, _ := strconv.Atoi(contact_id_str)
//	//
//	//sort_str := c.PostForm("sort")
//	//sort, _ := strconv.Atoi(sort_str)
//	//
//	////file, err = crmContactFileUploadAndDownloadsService.UploadFile(header, noSave, &contact_id, &sort) // 文件上传后拿到文件路径
//	////if err != nil {
//	////	global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
//	////	response.FailWithMessage("修改数据库链接失败", c)
//	////	return
//	////}
//	//response.OkWithDetailed(crmRes.CrmContactFileResponse{File: file}, "上传成功", c)
//}

//
//// UploadMultipartFile
//// @Tags      ExaFileUploadAndDownload
//// @Summary   上传多文件示例
//// @Security  ApiKeyAuth
//// @accept    multipart/form-data
//// @Produce   application/json
//// @Param     file  formData  file                                                           true  "上传文件示例"
//// @Success   200   {object}  response.Response{data=exampleRes.ExaFileResponse,msg=string}  "上传文件示例,返回包括文件详情"
//// @Router    /fileUploadAndDownload/uploadMultipartFile [post]
//func (b *FileUploadAndDownloadApi) UploadMultipartFile(c *gin.Context) {
//	//var file example.ExaFileUploadAndDownload
//	noSave := c.DefaultQuery("noSave", "0")
//	//_, header, err := c.Request.FormFile("file")
//	//if err != nil {
//	//	global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
//	//	response.FailWithMessage("接收文件失败", c)
//	//	return
//	//}
//	//file, err = fileUploadAndDownloadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
//	//if err != nil {
//	//	global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
//	//	response.FailWithMessage("修改数据库链接失败", c)
//	//	return
//	//}
//
//	// 获取上传的文件
//	formFiles := c.Request.MultipartForm.File["file"]
//	if formFiles == nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "没有上传文件"})
//		return
//	}
//
//	contact_id_str := c.PostForm("contact_id")
//	contact_id, _ := strconv.Atoi(contact_id_str)
//	result := make([]crm.CrmContactFileUploadAndDownloads, 0)
//	for k, fileHeader := range formFiles {
//		// 生成保存的文件名，这里使用时间戳加原始文件名组合
//		file, err := crmContactFileUploadAndDownloadsService.UploadFile(fileHeader, noSave, &contact_id, &k) // 文件上传后拿到文件路径
//		if err != nil {
//			global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
//			response.FailWithMessage("修改数据库链接失败", c)
//			return
//		}
//		result = append(result, file)
//	}
//	response.OkWithDetailed(crmRes.CrmFileResponse{result}, "上传成功", c)
//}

// EditFileName 编辑文件名或者备注
func (b *FileUploadAndDownloadApi) EditFileName(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = fileUploadAndDownloadService.EditFileName(file)
	if err != nil {
		global.GVA_LOG.Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

// DeleteFile
// @Tags      ExaFileUploadAndDownload
// @Summary   删除文件
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      example.ExaFileUploadAndDownload  true  "传入文件里面id即可"
// @Success   200   {object}  response.Response{msg=string}     "删除文件"
// @Router    /fileUploadAndDownload/deleteFile [post]
func (b *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileUploadAndDownloadService.DeleteFile(file); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetFileList
// @Tags      ExaFileUploadAndDownload
// @Summary   分页文件列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页文件列表,返回包括列表,总数,页码,每页数量"
// @Router    /fileUploadAndDownload/getFileList [post]
func (b *FileUploadAndDownloadApi) GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := fileUploadAndDownloadService.GetFileRecordInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetFileIdsList
// @Tags      ExaFileUploadAndDownload
// @Summary   分页文件列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页文件列表,返回包括列表,总数,页码,每页数量"
// @Router    /fileUploadAndDownload/getFileList [get]
func (b *FileUploadAndDownloadApi) GetFileIdsList(c *gin.Context) {
	IDs := c.Query("ids")

	list, _, err := fileUploadAndDownloadService.GetFileRecordInfoIdsList(IDs)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(list, "获取成功", c)
}

// DownloadFile
// @Tags      CrmContactFileUploadAndDownloads
// @Summary   合同文件下载
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                                                           true  "上传文件示例"
// @Success   200   "file"  "上传文件示例,返回包括文件详情"
// @Router    /fileUploadAndDownload/upload [post]
func (b *FileUploadAndDownloadApi) DownloadFile(c *gin.Context) {

	id, ok := c.GetQuery("id")
	if !ok {
		response.FailWithMessage("参数错误", c)
		return
	}
	file, file_name, err := fileUploadAndDownloadService.DownloadFile(id) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file_name)) // 对下载的文件重命名
	c.Header("success", "true")
	c.Data(http.StatusOK, "application/octet-stream", file)

}