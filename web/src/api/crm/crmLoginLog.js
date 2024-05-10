import service from '@/utils/request'

// @Tags CrmLoginLog
// @Summary 创建登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmLoginLog true "创建登录日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmLoginLog/createCrmLoginLog [post]
export const createCrmLoginLog = (data) => {
  return service({
    url: '/crmLoginLog/createCrmLoginLog',
    method: 'post',
    data
  })
}

// @Tags CrmLoginLog
// @Summary 删除登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmLoginLog true "删除登录日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmLoginLog/deleteCrmLoginLog [delete]
export const deleteCrmLoginLog = (params) => {
  return service({
    url: '/crmLoginLog/deleteCrmLoginLog',
    method: 'delete',
    params
  })
}

// @Tags CrmLoginLog
// @Summary 批量删除登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除登录日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmLoginLog/deleteCrmLoginLog [delete]
export const deleteCrmLoginLogByIds = (params) => {
  return service({
    url: '/crmLoginLog/deleteCrmLoginLogByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmLoginLog
// @Summary 更新登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmLoginLog true "更新登录日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmLoginLog/updateCrmLoginLog [put]
export const updateCrmLoginLog = (data) => {
  return service({
    url: '/crmLoginLog/updateCrmLoginLog',
    method: 'put',
    data
  })
}

// @Tags CrmLoginLog
// @Summary 用id查询登录日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmLoginLog true "用id查询登录日志"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmLoginLog/findCrmLoginLog [get]
export const findCrmLoginLog = (params) => {
  return service({
    url: '/crmLoginLog/findCrmLoginLog',
    method: 'get',
    params
  })
}

// @Tags CrmLoginLog
// @Summary 分页获取登录日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取登录日志列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmLoginLog/getCrmLoginLogList [get]
export const getCrmLoginLogList = (params) => {
  return service({
    url: '/crmLoginLog/getCrmLoginLogList',
    method: 'get',
    params
  })
}
