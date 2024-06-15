import service from '@/utils/request'

// @Tags CrmContactApprovalRecord
// @Summary 创建合同审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContactApprovalRecord true "创建合同审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContactApprovalRecord/createCrmContactApprovalRecord [post]
export const createCrmContactApprovalRecord = (data) => {
  return service({
    url: '/crmContactApprovalRecord/createCrmContactApprovalRecord',
    method: 'post',
    data
  })
}

// @Tags CrmContactApprovalRecord
// @Summary 删除合同审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContactApprovalRecord true "删除合同审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContactApprovalRecord/deleteCrmContactApprovalRecord [delete]
export const deleteCrmContactApprovalRecord = (params) => {
  return service({
    url: '/crmContactApprovalRecord/deleteCrmContactApprovalRecord',
    method: 'delete',
    params
  })
}

// @Tags CrmContactApprovalRecord
// @Summary 批量删除合同审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除合同审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContactApprovalRecord/deleteCrmContactApprovalRecord [delete]
export const deleteCrmContactApprovalRecordByIds = (params) => {
  return service({
    url: '/crmContactApprovalRecord/deleteCrmContactApprovalRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmContactApprovalRecord
// @Summary 更新合同审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContactApprovalRecord true "更新合同审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactApprovalRecord/updateCrmContactApprovalRecord [put]
export const updateCrmContactApprovalRecord = (data) => {
  return service({
    url: '/crmContactApprovalRecord/updateCrmContactApprovalRecord',
    method: 'put',
    data
  })
}

// @Tags CrmContactApprovalRecord
// @Summary 用id查询合同审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmContactApprovalRecord true "用id查询合同审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactApprovalRecord/findCrmContactApprovalRecord [get]
export const findCrmContactApprovalRecord = (params) => {
  return service({
    url: '/crmContactApprovalRecord/findCrmContactApprovalRecord',
    method: 'get',
    params
  })
}

// @Tags CrmContactApprovalRecord
// @Summary 分页获取合同审批记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取合同审批记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactApprovalRecord/getCrmContactApprovalRecordList [get]
export const getCrmContactApprovalRecordList = (params) => {
  return service({
    url: '/crmContactApprovalRecord/getCrmContactApprovalRecordList',
    method: 'get',
    params
  })
}
