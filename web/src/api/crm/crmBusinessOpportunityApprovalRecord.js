import service from '@/utils/request'

// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 创建商机审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityApprovalRecord true "创建商机审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/createCrmBusinessOpportunityApprovalRecord [post]
export const createCrmBusinessOpportunityApprovalRecord = (data) => {
  return service({
    url: '/crmBusinessOpportunityApprovalRecord/createCrmBusinessOpportunityApprovalRecord',
    method: 'post',
    data
  })
}

// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 删除商机审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityApprovalRecord true "删除商机审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/deleteCrmBusinessOpportunityApprovalRecord [delete]
export const deleteCrmBusinessOpportunityApprovalRecord = (params) => {
  return service({
    url: '/crmBusinessOpportunityApprovalRecord/deleteCrmBusinessOpportunityApprovalRecord',
    method: 'delete',
    params
  })
}

// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 批量删除商机审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商机审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/deleteCrmBusinessOpportunityApprovalRecord [delete]
export const deleteCrmBusinessOpportunityApprovalRecordByIds = (params) => {
  return service({
    url: '/crmBusinessOpportunityApprovalRecord/deleteCrmBusinessOpportunityApprovalRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 更新商机审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityApprovalRecord true "更新商机审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/updateCrmBusinessOpportunityApprovalRecord [put]
export const updateCrmBusinessOpportunityApprovalRecord = (data) => {
  return service({
    url: '/crmBusinessOpportunityApprovalRecord/updateCrmBusinessOpportunityApprovalRecord',
    method: 'put',
    data
  })
}

// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 用id查询商机审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmBusinessOpportunityApprovalRecord true "用id查询商机审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/findCrmBusinessOpportunityApprovalRecord [get]
export const findCrmBusinessOpportunityApprovalRecord = (params) => {
  return service({
    url: '/crmBusinessOpportunityApprovalRecord/findCrmBusinessOpportunityApprovalRecord',
    method: 'get',
    params
  })
}

// @Tags CrmBusinessOpportunityApprovalRecord
// @Summary 分页获取商机审批记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商机审批记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityApprovalRecord/getCrmBusinessOpportunityApprovalRecordList [get]
export const getCrmBusinessOpportunityApprovalRecordList = (params) => {
  return service({
    url: '/crmBusinessOpportunityApprovalRecord/getCrmBusinessOpportunityApprovalRecordList',
    method: 'get',
    params
  })
}
