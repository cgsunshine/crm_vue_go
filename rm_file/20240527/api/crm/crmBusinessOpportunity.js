import service from '@/utils/request'

// @Tags CrmBusinessOpportunity
// @Summary 创建crmBusinessOpportunity表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunity true "创建crmBusinessOpportunity表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunity/createCrmBusinessOpportunity [post]
export const createCrmBusinessOpportunity = (data) => {
  return service({
    url: '/crmBusinessOpportunity/createCrmBusinessOpportunity',
    method: 'post',
    data
  })
}

// @Tags CrmBusinessOpportunity
// @Summary 删除crmBusinessOpportunity表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunity true "删除crmBusinessOpportunity表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunity/deleteCrmBusinessOpportunity [delete]
export const deleteCrmBusinessOpportunity = (params) => {
  return service({
    url: '/crmBusinessOpportunity/deleteCrmBusinessOpportunity',
    method: 'delete',
    params
  })
}

// @Tags CrmBusinessOpportunity
// @Summary 批量删除crmBusinessOpportunity表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmBusinessOpportunity表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunity/deleteCrmBusinessOpportunity [delete]
export const deleteCrmBusinessOpportunityByIds = (params) => {
  return service({
    url: '/crmBusinessOpportunity/deleteCrmBusinessOpportunityByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmBusinessOpportunity
// @Summary 更新crmBusinessOpportunity表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunity true "更新crmBusinessOpportunity表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBusinessOpportunity/updateCrmBusinessOpportunity [put]
export const updateCrmBusinessOpportunity = (data) => {
  return service({
    url: '/crmBusinessOpportunity/updateCrmBusinessOpportunity',
    method: 'put',
    data
  })
}

// @Tags CrmBusinessOpportunity
// @Summary 用id查询crmBusinessOpportunity表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmBusinessOpportunity true "用id查询crmBusinessOpportunity表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunity/findCrmBusinessOpportunity [get]
export const findCrmBusinessOpportunity = (params) => {
  return service({
    url: '/crmBusinessOpportunity/findCrmBusinessOpportunity',
    method: 'get',
    params
  })
}

// @Tags CrmBusinessOpportunity
// @Summary 分页获取crmBusinessOpportunity表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmBusinessOpportunity表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunity/getCrmBusinessOpportunityList [get]
export const getCrmBusinessOpportunityList = (params) => {
  return service({
    url: '/crmBusinessOpportunity/getCrmBusinessOpportunityList',
    method: 'get',
    params
  })
}
