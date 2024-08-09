import service from '@/utils/request'

// @Tags CrmAuthorities
// @Summary 创建数据权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmAuthorities true "创建数据权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmAuthorities/createCrmAuthorities [post]
export const createCrmAuthorities = (data) => {
  return service({
    url: '/crmAuthorities/createCrmAuthorities',
    method: 'post',
    data
  })
}

// @Tags CrmAuthorities
// @Summary 删除数据权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmAuthorities true "删除数据权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmAuthorities/deleteCrmAuthorities [delete]
export const deleteCrmAuthorities = (params) => {
  return service({
    url: '/crmAuthorities/deleteCrmAuthorities',
    method: 'delete',
    params
  })
}

// @Tags CrmAuthorities
// @Summary 批量删除数据权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除数据权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmAuthorities/deleteCrmAuthorities [delete]
export const deleteCrmAuthoritiesByIds = (params) => {
  return service({
    url: '/crmAuthorities/deleteCrmAuthoritiesByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmAuthorities
// @Summary 更新数据权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmAuthorities true "更新数据权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmAuthorities/updateCrmAuthorities [put]
export const updateCrmAuthorities = (data) => {
  return service({
    url: '/crmAuthorities/updateCrmAuthorities',
    method: 'put',
    data
  })
}

// @Tags CrmAuthorities
// @Summary 用id查询数据权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmAuthorities true "用id查询数据权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmAuthorities/findCrmAuthorities [get]
export const findCrmAuthorities = (params) => {
  return service({
    url: '/crmAuthorities/findCrmAuthorities',
    method: 'get',
    params
  })
}

// @Tags CrmAuthorities
// @Summary 分页获取数据权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取数据权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmAuthorities/getCrmAuthoritiesList [get]
export const getCrmAuthoritiesList = (params) => {
  return service({
    url: '/crmAuthorities/getCrmAuthoritiesList',
    method: 'get',
    params
  })
}
