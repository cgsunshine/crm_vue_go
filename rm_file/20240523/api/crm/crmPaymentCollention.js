import service from '@/utils/request'

// @Tags CrmPaymentCollention
// @Summary 创建回款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPaymentCollention true "创建回款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmPaymentCollention/createCrmPaymentCollention [post]
export const createCrmPaymentCollention = (data) => {
  return service({
    url: '/crmPaymentCollention/createCrmPaymentCollention',
    method: 'post',
    data
  })
}

// @Tags CrmPaymentCollention
// @Summary 删除回款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPaymentCollention true "删除回款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPaymentCollention/deleteCrmPaymentCollention [delete]
export const deleteCrmPaymentCollention = (params) => {
  return service({
    url: '/crmPaymentCollention/deleteCrmPaymentCollention',
    method: 'delete',
    params
  })
}

// @Tags CrmPaymentCollention
// @Summary 批量删除回款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除回款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmPaymentCollention/deleteCrmPaymentCollention [delete]
export const deleteCrmPaymentCollentionByIds = (params) => {
  return service({
    url: '/crmPaymentCollention/deleteCrmPaymentCollentionByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmPaymentCollention
// @Summary 更新回款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmPaymentCollention true "更新回款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmPaymentCollention/updateCrmPaymentCollention [put]
export const updateCrmPaymentCollention = (data) => {
  return service({
    url: '/crmPaymentCollention/updateCrmPaymentCollention',
    method: 'put',
    data
  })
}

// @Tags CrmPaymentCollention
// @Summary 用id查询回款管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmPaymentCollention true "用id查询回款管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmPaymentCollention/findCrmPaymentCollention [get]
export const findCrmPaymentCollention = (params) => {
  return service({
    url: '/crmPaymentCollention/findCrmPaymentCollention',
    method: 'get',
    params
  })
}

// @Tags CrmPaymentCollention
// @Summary 分页获取回款管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取回款管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmPaymentCollention/getCrmPaymentCollentionList [get]
export const getCrmPaymentCollentionList = (params) => {
  return service({
    url: '/crmPaymentCollention/getCrmPaymentCollentionList',
    method: 'get',
    params
  })
}
