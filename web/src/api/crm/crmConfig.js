import service from '@/utils/request'

// @Tags CrmConfig
// @Summary 创建系统配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmConfig true "创建系统配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmConfig/createCrmConfig [post]
export const createCrmConfig = (data) => {
  return service({
    url: '/crmConfig/createCrmConfig',
    method: 'post',
    data
  })
}

// @Tags CrmConfig
// @Summary 删除系统配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmConfig true "删除系统配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmConfig/deleteCrmConfig [delete]
export const deleteCrmConfig = (params) => {
  return service({
    url: '/crmConfig/deleteCrmConfig',
    method: 'delete',
    params
  })
}

// @Tags CrmConfig
// @Summary 批量删除系统配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除系统配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmConfig/deleteCrmConfig [delete]
export const deleteCrmConfigByIds = (params) => {
  return service({
    url: '/crmConfig/deleteCrmConfigByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmConfig
// @Summary 更新系统配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmConfig true "更新系统配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmConfig/updateCrmConfig [put]
export const updateCrmConfig = (data) => {
  return service({
    url: '/crmConfig/updateCrmConfig',
    method: 'put',
    data
  })
}

// @Tags CrmConfig
// @Summary 用id查询系统配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmConfig true "用id查询系统配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmConfig/findCrmConfig [get]
export const findCrmConfig = (params) => {
  return service({
    url: '/crmConfig/findCrmConfig',
    method: 'get',
    params
  })
}

// @Tags CrmConfig
// @Summary 分页获取系统配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取系统配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmConfig/getCrmConfigList [get]
export const getCrmConfigList = (params) => {
  return service({
    url: '/crmConfig/getCrmConfigList',
    method: 'get',
    params
  })
}
