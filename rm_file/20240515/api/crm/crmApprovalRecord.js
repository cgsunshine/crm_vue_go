import service from '@/utils/request'

// @Tags CrmApprovalRecord
// @Summary 创建crmApprovalRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalRecord true "创建crmApprovalRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmApprovalRecord/createCrmApprovalRecord [post]
export const createCrmApprovalRecord = (data) => {
  return service({
    url: '/crmApprovalRecord/createCrmApprovalRecord',
    method: 'post',
    data
  })
}

// @Tags CrmApprovalRecord
// @Summary 删除crmApprovalRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalRecord true "删除crmApprovalRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalRecord/deleteCrmApprovalRecord [delete]
export const deleteCrmApprovalRecord = (params) => {
  return service({
    url: '/crmApprovalRecord/deleteCrmApprovalRecord',
    method: 'delete',
    params
  })
}

// @Tags CrmApprovalRecord
// @Summary 批量删除crmApprovalRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmApprovalRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmApprovalRecord/deleteCrmApprovalRecord [delete]
export const deleteCrmApprovalRecordByIds = (params) => {
  return service({
    url: '/crmApprovalRecord/deleteCrmApprovalRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmApprovalRecord
// @Summary 更新crmApprovalRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmApprovalRecord true "更新crmApprovalRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmApprovalRecord/updateCrmApprovalRecord [put]
export const updateCrmApprovalRecord = (data) => {
  return service({
    url: '/crmApprovalRecord/updateCrmApprovalRecord',
    method: 'put',
    data
  })
}

// @Tags CrmApprovalRecord
// @Summary 用id查询crmApprovalRecord表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmApprovalRecord true "用id查询crmApprovalRecord表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmApprovalRecord/findCrmApprovalRecord [get]
export const findCrmApprovalRecord = (params) => {
  return service({
    url: '/crmApprovalRecord/findCrmApprovalRecord',
    method: 'get',
    params
  })
}

// @Tags CrmApprovalRecord
// @Summary 分页获取crmApprovalRecord表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmApprovalRecord表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmApprovalRecord/getCrmApprovalRecordList [get]
export const getCrmApprovalRecordList = (params) => {
  return service({
    url: '/crmApprovalRecord/getCrmApprovalRecordList',
    method: 'get',
    params
  })
}
