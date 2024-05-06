import service from '@/utils/request'

// @Tags CrmTest
// @Summary 创建crmTest表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTest true "创建crmTest表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmTest/createCrmTest [post]
export const createCrmTest = (data) => {
  return service({
    url: '/crmTest/createCrmTest',
    method: 'post',
    data
  })
}

// @Tags CrmTest
// @Summary 删除crmTest表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTest true "删除crmTest表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTest/deleteCrmTest [delete]
export const deleteCrmTest = (params) => {
  return service({
    url: '/crmTest/deleteCrmTest',
    method: 'delete',
    params
  })
}

// @Tags CrmTest
// @Summary 批量删除crmTest表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmTest表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmTest/deleteCrmTest [delete]
export const deleteCrmTestByIds = (params) => {
  return service({
    url: '/crmTest/deleteCrmTestByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmTest
// @Summary 更新crmTest表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmTest true "更新crmTest表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmTest/updateCrmTest [put]
export const updateCrmTest = (data) => {
  return service({
    url: '/crmTest/updateCrmTest',
    method: 'put',
    data
  })
}

// @Tags CrmTest
// @Summary 用id查询crmTest表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmTest true "用id查询crmTest表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmTest/findCrmTest [get]
export const findCrmTest = (params) => {
  return service({
    url: '/crmTest/findCrmTest',
    method: 'get',
    params
  })
}

// @Tags CrmTest
// @Summary 分页获取crmTest表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmTest表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmTest/getCrmTestList [get]
export const getCrmTestList = (params) => {
  return service({
    url: '/crmTest/getCrmTestList',
    method: 'get',
    params
  })
}
