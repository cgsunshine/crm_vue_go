import service from '@/utils/request'

// @Tags CrmCustomerGroup
// @Summary 创建crmCustomerGroup表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCustomerGroup true "创建crmCustomerGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmCustomerGroup/createCrmCustomerGroup [post]
export const createCrmCustomerGroup = (data) => {
  return service({
    url: '/crmCustomerGroup/createCrmCustomerGroup',
    method: 'post',
    data
  })
}

// @Tags CrmCustomerGroup
// @Summary 删除crmCustomerGroup表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCustomerGroup true "删除crmCustomerGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCustomerGroup/deleteCrmCustomerGroup [delete]
export const deleteCrmCustomerGroup = (params) => {
  return service({
    url: '/crmCustomerGroup/deleteCrmCustomerGroup',
    method: 'delete',
    params
  })
}

// @Tags CrmCustomerGroup
// @Summary 批量删除crmCustomerGroup表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmCustomerGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmCustomerGroup/deleteCrmCustomerGroup [delete]
export const deleteCrmCustomerGroupByIds = (params) => {
  return service({
    url: '/crmCustomerGroup/deleteCrmCustomerGroupByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmCustomerGroup
// @Summary 更新crmCustomerGroup表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmCustomerGroup true "更新crmCustomerGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmCustomerGroup/updateCrmCustomerGroup [put]
export const updateCrmCustomerGroup = (data) => {
  return service({
    url: '/crmCustomerGroup/updateCrmCustomerGroup',
    method: 'put',
    data
  })
}

// @Tags CrmCustomerGroup
// @Summary 用id查询crmCustomerGroup表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmCustomerGroup true "用id查询crmCustomerGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmCustomerGroup/findCrmCustomerGroup [get]
export const findCrmCustomerGroup = (params) => {
  return service({
    url: '/crmCustomerGroup/findCrmCustomerGroup',
    method: 'get',
    params
  })
}

// @Tags CrmCustomerGroup
// @Summary 分页获取crmCustomerGroup表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmCustomerGroup表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmCustomerGroup/getCrmCustomerGroupList [get]
export const getCrmCustomerGroupList = (params) => {
  return service({
    url: '/crmCustomerGroup/getCrmCustomerGroupList',
    method: 'get',
    params
  })
}
