import service from '@/utils/request'

// @Tags CrmDeposits
// @Summary 创建押金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmDeposits true "创建押金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmDeposits/createCrmDeposits [post]
export const createCrmDeposits = (data) => {
  return service({
    url: '/crmDeposits/createCrmDeposits',
    method: 'post',
    data
  })
}

// @Tags CrmDeposits
// @Summary 删除押金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmDeposits true "删除押金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmDeposits/deleteCrmDeposits [delete]
export const deleteCrmDeposits = (params) => {
  return service({
    url: '/crmDeposits/deleteCrmDeposits',
    method: 'delete',
    params
  })
}

// @Tags CrmDeposits
// @Summary 批量删除押金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除押金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmDeposits/deleteCrmDeposits [delete]
export const deleteCrmDepositsByIds = (params) => {
  return service({
    url: '/crmDeposits/deleteCrmDepositsByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmDeposits
// @Summary 更新押金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmDeposits true "更新押金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmDeposits/updateCrmDeposits [put]
export const updateCrmDeposits = (data) => {
  return service({
    url: '/crmDeposits/updateCrmDeposits',
    method: 'put',
    data
  })
}

// @Tags CrmDeposits
// @Summary 用id查询押金管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmDeposits true "用id查询押金管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmDeposits/findCrmDeposits [get]
export const findCrmDeposits = (params) => {
  return service({
    url: '/crmDeposits/findCrmDeposits',
    method: 'get',
    params
  })
}

// @Tags CrmDeposits
// @Summary 分页获取押金管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取押金管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmDeposits/getCrmDepositsList [get]
export const getCrmDepositsList = (params) => {
  return service({
    url: '/crmDeposits/getCrmDepositsList',
    method: 'get',
    params
  })
}
