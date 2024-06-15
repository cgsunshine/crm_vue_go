import service from '@/utils/request'

// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 创建对账单上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmStatementAccountFileUploadAndDownloads true "创建对账单上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/createCrmStatementAccountFileUploadAndDownloads [post]
export const createCrmStatementAccountFileUploadAndDownloads = (data) => {
  return service({
    url: '/crmStatementAccountFileUploadAndDownloads/createCrmStatementAccountFileUploadAndDownloads',
    method: 'post',
    data
  })
}

// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 删除对账单上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmStatementAccountFileUploadAndDownloads true "删除对账单上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/deleteCrmStatementAccountFileUploadAndDownloads [delete]
export const deleteCrmStatementAccountFileUploadAndDownloads = (params) => {
  return service({
    url: '/crmStatementAccountFileUploadAndDownloads/deleteCrmStatementAccountFileUploadAndDownloads',
    method: 'delete',
    params
  })
}

// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 批量删除对账单上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除对账单上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/deleteCrmStatementAccountFileUploadAndDownloads [delete]
export const deleteCrmStatementAccountFileUploadAndDownloadsByIds = (params) => {
  return service({
    url: '/crmStatementAccountFileUploadAndDownloads/deleteCrmStatementAccountFileUploadAndDownloadsByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 更新对账单上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmStatementAccountFileUploadAndDownloads true "更新对账单上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/updateCrmStatementAccountFileUploadAndDownloads [put]
export const updateCrmStatementAccountFileUploadAndDownloads = (data) => {
  return service({
    url: '/crmStatementAccountFileUploadAndDownloads/updateCrmStatementAccountFileUploadAndDownloads',
    method: 'put',
    data
  })
}

// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 用id查询对账单上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmStatementAccountFileUploadAndDownloads true "用id查询对账单上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/findCrmStatementAccountFileUploadAndDownloads [get]
export const findCrmStatementAccountFileUploadAndDownloads = (params) => {
  return service({
    url: '/crmStatementAccountFileUploadAndDownloads/findCrmStatementAccountFileUploadAndDownloads',
    method: 'get',
    params
  })
}

// @Tags CrmStatementAccountFileUploadAndDownloads
// @Summary 分页获取对账单上传文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取对账单上传文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmStatementAccountFileUploadAndDownloads/getCrmStatementAccountFileUploadAndDownloadsList [get]
export const getCrmStatementAccountFileUploadAndDownloadsList = (params) => {
  return service({
    url: '/crmStatementAccountFileUploadAndDownloads/getCrmStatementAccountFileUploadAndDownloadsList',
    method: 'get',
    params
  })
}
