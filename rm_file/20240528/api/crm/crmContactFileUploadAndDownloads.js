import service from '@/utils/request'

// @Tags CrmContactFileUploadAndDownloads
// @Summary 创建crmContactFileUploadAndDownloads表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContactFileUploadAndDownloads true "创建crmContactFileUploadAndDownloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContactFileUploadAndDownloads/createCrmContactFileUploadAndDownloads [post]
export const createCrmContactFileUploadAndDownloads = (data) => {
  return service({
    url: '/crmContactFileUploadAndDownloads/createCrmContactFileUploadAndDownloads',
    method: 'post',
    data
  })
}

// @Tags CrmContactFileUploadAndDownloads
// @Summary 删除crmContactFileUploadAndDownloads表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContactFileUploadAndDownloads true "删除crmContactFileUploadAndDownloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContactFileUploadAndDownloads/deleteCrmContactFileUploadAndDownloads [delete]
export const deleteCrmContactFileUploadAndDownloads = (params) => {
  return service({
    url: '/crmContactFileUploadAndDownloads/deleteCrmContactFileUploadAndDownloads',
    method: 'delete',
    params
  })
}

// @Tags CrmContactFileUploadAndDownloads
// @Summary 批量删除crmContactFileUploadAndDownloads表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmContactFileUploadAndDownloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContactFileUploadAndDownloads/deleteCrmContactFileUploadAndDownloads [delete]
export const deleteCrmContactFileUploadAndDownloadsByIds = (params) => {
  return service({
    url: '/crmContactFileUploadAndDownloads/deleteCrmContactFileUploadAndDownloadsByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmContactFileUploadAndDownloads
// @Summary 更新crmContactFileUploadAndDownloads表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContactFileUploadAndDownloads true "更新crmContactFileUploadAndDownloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactFileUploadAndDownloads/updateCrmContactFileUploadAndDownloads [put]
export const updateCrmContactFileUploadAndDownloads = (data) => {
  return service({
    url: '/crmContactFileUploadAndDownloads/updateCrmContactFileUploadAndDownloads',
    method: 'put',
    data
  })
}

// @Tags CrmContactFileUploadAndDownloads
// @Summary 用id查询crmContactFileUploadAndDownloads表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmContactFileUploadAndDownloads true "用id查询crmContactFileUploadAndDownloads表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactFileUploadAndDownloads/findCrmContactFileUploadAndDownloads [get]
export const findCrmContactFileUploadAndDownloads = (params) => {
  return service({
    url: '/crmContactFileUploadAndDownloads/findCrmContactFileUploadAndDownloads',
    method: 'get',
    params
  })
}

// @Tags CrmContactFileUploadAndDownloads
// @Summary 分页获取crmContactFileUploadAndDownloads表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmContactFileUploadAndDownloads表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactFileUploadAndDownloads/getCrmContactFileUploadAndDownloadsList [get]
export const getCrmContactFileUploadAndDownloadsList = (params) => {
  return service({
    url: '/crmContactFileUploadAndDownloads/getCrmContactFileUploadAndDownloadsList',
    method: 'get',
    params
  })
}
