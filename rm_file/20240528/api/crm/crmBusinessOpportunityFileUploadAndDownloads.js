import service from '@/utils/request'

// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 创建上传商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityFileUploadAndDownloads true "创建上传商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/createCrmBusinessOpportunityFileUploadAndDownloads [post]
export const createCrmBusinessOpportunityFileUploadAndDownloads = (data) => {
  return service({
    url: '/crmBusinessOpportunityFileUploadAndDownloads/createCrmBusinessOpportunityFileUploadAndDownloads',
    method: 'post',
    data
  })
}

// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 删除上传商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityFileUploadAndDownloads true "删除上传商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/deleteCrmBusinessOpportunityFileUploadAndDownloads [delete]
export const deleteCrmBusinessOpportunityFileUploadAndDownloads = (params) => {
  return service({
    url: '/crmBusinessOpportunityFileUploadAndDownloads/deleteCrmBusinessOpportunityFileUploadAndDownloads',
    method: 'delete',
    params
  })
}

// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 批量删除上传商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除上传商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/deleteCrmBusinessOpportunityFileUploadAndDownloads [delete]
export const deleteCrmBusinessOpportunityFileUploadAndDownloadsByIds = (params) => {
  return service({
    url: '/crmBusinessOpportunityFileUploadAndDownloads/deleteCrmBusinessOpportunityFileUploadAndDownloadsByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 更新上传商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityFileUploadAndDownloads true "更新上传商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/updateCrmBusinessOpportunityFileUploadAndDownloads [put]
export const updateCrmBusinessOpportunityFileUploadAndDownloads = (data) => {
  return service({
    url: '/crmBusinessOpportunityFileUploadAndDownloads/updateCrmBusinessOpportunityFileUploadAndDownloads',
    method: 'put',
    data
  })
}

// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 用id查询上传商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmBusinessOpportunityFileUploadAndDownloads true "用id查询上传商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/findCrmBusinessOpportunityFileUploadAndDownloads [get]
export const findCrmBusinessOpportunityFileUploadAndDownloads = (params) => {
  return service({
    url: '/crmBusinessOpportunityFileUploadAndDownloads/findCrmBusinessOpportunityFileUploadAndDownloads',
    method: 'get',
    params
  })
}

// @Tags CrmBusinessOpportunityFileUploadAndDownloads
// @Summary 分页获取上传商机文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取上传商机文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityFileUploadAndDownloads/getCrmBusinessOpportunityFileUploadAndDownloadsList [get]
export const getCrmBusinessOpportunityFileUploadAndDownloadsList = (params) => {
  return service({
    url: '/crmBusinessOpportunityFileUploadAndDownloads/getCrmBusinessOpportunityFileUploadAndDownloadsList',
    method: 'get',
    params
  })
}
