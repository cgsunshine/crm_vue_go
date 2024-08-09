import service from '@/utils/request'

// @Tags CrmRefundTasks
// @Summary 创建退款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmRefundTasks true "创建退款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmRefundTasks/createCrmRefundTasks [post]
export const createCrmRefundTasks = (data) => {
  return service({
    url: '/crmRefundTasks/createCrmRefundTasks',
    method: 'post',
    data
  })
}

// @Tags CrmRefundTasks
// @Summary 删除退款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmRefundTasks true "删除退款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmRefundTasks/deleteCrmRefundTasks [delete]
export const deleteCrmRefundTasks = (params) => {
  return service({
    url: '/crmRefundTasks/deleteCrmRefundTasks',
    method: 'delete',
    params
  })
}

// @Tags CrmRefundTasks
// @Summary 批量删除退款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除退款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmRefundTasks/deleteCrmRefundTasks [delete]
export const deleteCrmRefundTasksByIds = (params) => {
  return service({
    url: '/crmRefundTasks/deleteCrmRefundTasksByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmRefundTasks
// @Summary 更新退款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmRefundTasks true "更新退款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmRefundTasks/updateCrmRefundTasks [put]
export const updateCrmRefundTasks = (data) => {
  return service({
    url: '/crmRefundTasks/updateCrmRefundTasks',
    method: 'put',
    data
  })
}

// @Tags CrmRefundTasks
// @Summary 用id查询退款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmRefundTasks true "用id查询退款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmRefundTasks/findCrmRefundTasks [get]
export const findCrmRefundTasks = (params) => {
  return service({
    url: '/crmRefundTasks/findCrmRefundTasks',
    method: 'get',
    params
  })
}

// @Tags CrmRefundTasks
// @Summary 分页获取退款管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取退款管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmRefundTasks/getCrmRefundTasksList [get]
export const getCrmRefundTasksList = (params) => {
  return service({
    url: '/crmRefundTasks/getCrmRefundTasksList',
    method: 'get',
    params
  })
}
