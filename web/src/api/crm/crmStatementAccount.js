import service from '@/utils/request'

// @Tags CrmStatementAccount
// @Summary 创建crmStatementAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmStatementAccount true "创建crmStatementAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmStatementAccount/createCrmStatementAccount [post]
export const createCrmStatementAccount = (data) => {
  return service({
    url: '/crmStatementAccount/createCrmStatementAccount',
    method: 'post',
    data
  })
}

// @Tags CrmStatementAccount
// @Summary 删除crmStatementAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmStatementAccount true "删除crmStatementAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmStatementAccount/deleteCrmStatementAccount [delete]
export const deleteCrmStatementAccount = (params) => {
  return service({
    url: '/crmStatementAccount/deleteCrmStatementAccount',
    method: 'delete',
    params
  })
}

// @Tags CrmStatementAccount
// @Summary 批量删除crmStatementAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmStatementAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmStatementAccount/deleteCrmStatementAccount [delete]
export const deleteCrmStatementAccountByIds = (params) => {
  return service({
    url: '/crmStatementAccount/deleteCrmStatementAccountByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmStatementAccount
// @Summary 更新crmStatementAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmStatementAccount true "更新crmStatementAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmStatementAccount/updateCrmStatementAccount [put]
export const updateCrmStatementAccount = (data) => {
  return service({
    url: '/crmStatementAccount/updateCrmStatementAccount',
    method: 'put',
    data
  })
}

// @Tags CrmStatementAccount
// @Summary 用id查询crmStatementAccount表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmStatementAccount true "用id查询crmStatementAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmStatementAccount/findCrmStatementAccount [get]
export const findCrmStatementAccount = (params) => {
  return service({
    url: '/crmStatementAccount/findCrmStatementAccount',
    method: 'get',
    params
  })
}

// @Tags CrmStatementAccount
// @Summary 分页获取crmStatementAccount表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmStatementAccount表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmStatementAccount/getCrmStatementAccountList [get]
export const getCrmStatementAccountList = (params) => {
  return service({
    url: '/crmStatementAccount/getCrmStatementAccountList',
    method: 'get',
    params
  })
}
