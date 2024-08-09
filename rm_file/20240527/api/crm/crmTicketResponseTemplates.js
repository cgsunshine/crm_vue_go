import service from '@/utils/request'

// @Tags CrmTicketResponseTemplates
// @Summary 创建crmTicketResponseTemplates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTicketResponseTemplates true "创建crmTicketResponseTemplates表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmTicketResponseTemplates/createCrmTicketResponseTemplates [post]
export const createCrmTicketResponseTemplates = (data) => {
  return service({
    url: '/crmTicketResponseTemplates/createCrmTicketResponseTemplates',
    method: 'post',
    data
  })
}

// @Tags CrmTicketResponseTemplates
// @Summary 删除crmTicketResponseTemplates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTicketResponseTemplates true "删除crmTicketResponseTemplates表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTicketResponseTemplates/deleteCrmTicketResponseTemplates [delete]
export const deleteCrmTicketResponseTemplates = (params) => {
  return service({
    url: '/crmTicketResponseTemplates/deleteCrmTicketResponseTemplates',
    method: 'delete',
    params
  })
}

// @Tags CrmTicketResponseTemplates
// @Summary 批量删除crmTicketResponseTemplates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmTicketResponseTemplates表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTicketResponseTemplates/deleteCrmTicketResponseTemplates [delete]
export const deleteCrmTicketResponseTemplatesByIds = (params) => {
  return service({
    url: '/crmTicketResponseTemplates/deleteCrmTicketResponseTemplatesByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmTicketResponseTemplates
// @Summary 更新crmTicketResponseTemplates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTicketResponseTemplates true "更新crmTicketResponseTemplates表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmTicketResponseTemplates/updateCrmTicketResponseTemplates [put]
export const updateCrmTicketResponseTemplates = (data) => {
  return service({
    url: '/crmTicketResponseTemplates/updateCrmTicketResponseTemplates',
    method: 'put',
    data
  })
}

// @Tags CrmTicketResponseTemplates
// @Summary 用id查询crmTicketResponseTemplates表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmTicketResponseTemplates true "用id查询crmTicketResponseTemplates表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmTicketResponseTemplates/findCrmTicketResponseTemplates [get]
export const findCrmTicketResponseTemplates = (params) => {
  return service({
    url: '/crmTicketResponseTemplates/findCrmTicketResponseTemplates',
    method: 'get',
    params
  })
}

// @Tags CrmTicketResponseTemplates
// @Summary 分页获取crmTicketResponseTemplates表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmTicketResponseTemplates表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTicketResponseTemplates/getCrmTicketResponseTemplatesList [get]
export const getCrmTicketResponseTemplatesList = (params) => {
  return service({
    url: '/crmTicketResponseTemplates/getCrmTicketResponseTemplatesList',
    method: 'get',
    params
  })
}
