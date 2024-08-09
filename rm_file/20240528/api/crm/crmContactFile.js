import service from '@/utils/request'

// @Tags CrmContactFile
// @Summary 创建crmContactFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContactFile true "创建crmContactFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmContactFile/createCrmContactFile [post]
export const createCrmContactFile = (data) => {
  return service({
    url: '/crmContactFile/createCrmContactFile',
    method: 'post',
    data
  })
}

// @Tags CrmContactFile
// @Summary 删除crmContactFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContactFile true "删除crmContactFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContactFile/deleteCrmContactFile [delete]
export const deleteCrmContactFile = (params) => {
  return service({
    url: '/crmContactFile/deleteCrmContactFile',
    method: 'delete',
    params
  })
}

// @Tags CrmContactFile
// @Summary 批量删除crmContactFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除crmContactFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmContactFile/deleteCrmContactFile [delete]
export const deleteCrmContactFileByIds = (params) => {
  return service({
    url: '/crmContactFile/deleteCrmContactFileByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmContactFile
// @Summary 更新crmContactFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmContactFile true "更新crmContactFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmContactFile/updateCrmContactFile [put]
export const updateCrmContactFile = (data) => {
  return service({
    url: '/crmContactFile/updateCrmContactFile',
    method: 'put',
    data
  })
}

// @Tags CrmContactFile
// @Summary 用id查询crmContactFile表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmContactFile true "用id查询crmContactFile表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmContactFile/findCrmContactFile [get]
export const findCrmContactFile = (params) => {
  return service({
    url: '/crmContactFile/findCrmContactFile',
    method: 'get',
    params
  })
}

// @Tags CrmContactFile
// @Summary 分页获取crmContactFile表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取crmContactFile表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmContactFile/getCrmContactFileList [get]
export const getCrmContactFileList = (params) => {
  return service({
    url: '/crmContactFile/getCrmContactFileList',
    method: 'get',
    params
  })
}
