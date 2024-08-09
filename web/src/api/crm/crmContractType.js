import service from '@/utils/request'

// @Tags CrmContractType
// @Summary 创建合同类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContractType true "创建合同类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContractType/createCrmContractType [post]
export const createCrmContractType = (data) => {
  return service({
    url: '/crmContractType/createCrmContractType',
    method: 'post',
    data
  })
}

// @Tags CrmContractType
// @Summary 删除合同类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContractType true "删除合同类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContractType/deleteCrmContractType [delete]
export const deleteCrmContractType = (params) => {
  return service({
    url: '/crmContractType/deleteCrmContractType',
    method: 'delete',
    params
  })
}

// @Tags CrmContractType
// @Summary 批量删除合同类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除合同类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContractType/deleteCrmContractType [delete]
export const deleteCrmContractTypeByIds = (params) => {
  return service({
    url: '/crmContractType/deleteCrmContractTypeByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmContractType
// @Summary 更新合同类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContractType true "更新合同类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContractType/updateCrmContractType [put]
export const updateCrmContractType = (data) => {
  return service({
    url: '/crmContractType/updateCrmContractType',
    method: 'put',
    data
  })
}

// @Tags CrmContractType
// @Summary 用id查询合同类型
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmContractType true "用id查询合同类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContractType/findCrmContractType [get]
export const findCrmContractType = (params) => {
  return service({
    url: '/crmContractType/findCrmContractType',
    method: 'get',
    params
  })
}

// @Tags CrmContractType
// @Summary 分页获取合同类型列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取合同类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContractType/getCrmContractTypeList [get]
export const getCrmContractTypeList = (params) => {
  return service({
    url: '/crmContractType/getCrmContractTypeList',
    method: 'get',
    params
  })
}
