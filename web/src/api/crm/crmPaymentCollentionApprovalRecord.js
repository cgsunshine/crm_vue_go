import service from '@/utils/request'

// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 创建回款审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPaymentCollentionApprovalRecord true "创建回款审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPaymentCollentionApprovalRecord/createCrmPaymentCollentionApprovalRecord [post]
export const createCrmPaymentCollentionApprovalRecord = (data) => {
  return service({
    url: '/crmPaymentCollentionApprovalRecord/createCrmPaymentCollentionApprovalRecord',
    method: 'post',
    data
  })
}

// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 删除回款审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPaymentCollentionApprovalRecord true "删除回款审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPaymentCollentionApprovalRecord/deleteCrmPaymentCollentionApprovalRecord [delete]
export const deleteCrmPaymentCollentionApprovalRecord = (params) => {
  return service({
    url: '/crmPaymentCollentionApprovalRecord/deleteCrmPaymentCollentionApprovalRecord',
    method: 'delete',
    params
  })
}

// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 批量删除回款审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除回款审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPaymentCollentionApprovalRecord/deleteCrmPaymentCollentionApprovalRecord [delete]
export const deleteCrmPaymentCollentionApprovalRecordByIds = (params) => {
  return service({
    url: '/crmPaymentCollentionApprovalRecord/deleteCrmPaymentCollentionApprovalRecordByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 更新回款审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPaymentCollentionApprovalRecord true "更新回款审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmPaymentCollentionApprovalRecord/updateCrmPaymentCollentionApprovalRecord [put]
export const updateCrmPaymentCollentionApprovalRecord = (data) => {
  return service({
    url: '/crmPaymentCollentionApprovalRecord/updateCrmPaymentCollentionApprovalRecord',
    method: 'put',
    data
  })
}

// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 用id查询回款审批记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmPaymentCollentionApprovalRecord true "用id查询回款审批记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPaymentCollentionApprovalRecord/findCrmPaymentCollentionApprovalRecord [get]
export const findCrmPaymentCollentionApprovalRecord = (params) => {
  return service({
    url: '/crmPaymentCollentionApprovalRecord/findCrmPaymentCollentionApprovalRecord',
    method: 'get',
    params
  })
}

// @Tags CrmPaymentCollentionApprovalRecord
// @Summary 分页获取回款审批记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取回款审批记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPaymentCollentionApprovalRecord/getCrmPaymentCollentionApprovalRecordList [get]
export const getCrmPaymentCollentionApprovalRecordList = (params) => {
  return service({
    url: '/crmPaymentCollentionApprovalRecord/getCrmPaymentCollentionApprovalRecordList',
    method: 'get',
    params
  })
}
