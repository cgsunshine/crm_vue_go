import service from '@/utils/request'

// @Tags CrmOperationRecords
// @Summary 创建操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmOperationRecords true "创建操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmOperationRecords/createCrmOperationRecords [post]
export const createCrmOperationRecords = (data) => {
  return service({
    url: '/crmOperationRecords/createCrmOperationRecords',
    method: 'post',
    data
  })
}

// @Tags CrmOperationRecords
// @Summary 删除操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmOperationRecords true "删除操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmOperationRecords/deleteCrmOperationRecords [delete]
export const deleteCrmOperationRecords = (params) => {
  return service({
    url: '/crmOperationRecords/deleteCrmOperationRecords',
    method: 'delete',
    params
  })
}

// @Tags CrmOperationRecords
// @Summary 批量删除操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmOperationRecords/deleteCrmOperationRecords [delete]
export const deleteCrmOperationRecordsByIds = (params) => {
  return service({
    url: '/crmOperationRecords/deleteCrmOperationRecordsByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmOperationRecords
// @Summary 更新操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmOperationRecords true "更新操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmOperationRecords/updateCrmOperationRecords [put]
export const updateCrmOperationRecords = (data) => {
  return service({
    url: '/crmOperationRecords/updateCrmOperationRecords',
    method: 'put',
    data
  })
}

// @Tags CrmOperationRecords
// @Summary 用id查询操作记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmOperationRecords true "用id查询操作记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmOperationRecords/findCrmOperationRecords [get]
export const findCrmOperationRecords = (params) => {
  return service({
    url: '/crmOperationRecords/findCrmOperationRecords',
    method: 'get',
    params
  })
}

// @Tags CrmOperationRecords
// @Summary 分页获取操作记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取操作记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmOperationRecords/getCrmOperationRecordsList [get]
export const getCrmOperationRecordsList = (params) => {
  return service({
    url: '/crmOperationRecords/getCrmOperationRecordsList',
    method: 'get',
    params
  })
}
