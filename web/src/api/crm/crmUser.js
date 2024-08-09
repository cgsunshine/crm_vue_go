import service from '@/utils/request'

// @Tags CrmUser
// @Summary 创建crmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmUser true "创建crmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmUser/createCrmUser [post]
export const createCrmUser = (data) => {
  return service({
    url: '/crmUser/createCrmUser',
    method: 'post',
    data
  })
}

// @Tags CrmUser
// @Summary 删除crmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmUser true "删除crmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmUser/deleteCrmUser [delete]
export const deleteCrmUser = (params) => {
  return service({
    url: '/crmUser/deleteCrmUser',
    method: 'delete',
    params
  })
}

// @Tags CrmUser
// @Summary 批量删除crmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmUser/deleteCrmUser [delete]
export const deleteCrmUserByIds = (params) => {
  return service({
    url: '/crmUser/deleteCrmUserByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmUser
// @Summary 更新crmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmUser true "更新crmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmUser/updateCrmUser [put]
export const updateCrmUser = (data) => {
  return service({
    url: '/crmUser/updateCrmUser',
    method: 'put',
    data
  })
}

// @Tags CrmUser
// @Summary 用id查询crmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmUser true "用id查询crmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmUser/findCrmUser [get]
export const findCrmUser = (params) => {
  return service({
    url: '/crmUser/findCrmUser',
    method: 'get',
    params
  })
}

// @Tags CrmUser
// @Summary 分页获取crmUser表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmUser表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmUser/getCrmUserList [get]
export const getCrmUserList = (params) => {
  return service({
    url: '/crmUser/getCrmUserList',
    method: 'get',
    params
  })
}
