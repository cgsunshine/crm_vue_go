import service from '@/utils/request'

// @Tags CrmBusinessOpportunityFile
// @Summary 创建商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityFile true "创建商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /crmBusinessOpportunityFile/createCrmBusinessOpportunityFile [post]
export const createCrmBusinessOpportunityFile = (data) => {
  return service({
    url: '/crmBusinessOpportunityFile/createCrmBusinessOpportunityFile',
    method: 'post',
    data
  })
}

// @Tags CrmBusinessOpportunityFile
// @Summary 删除商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityFile true "删除商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityFile/deleteCrmBusinessOpportunityFile [delete]
export const deleteCrmBusinessOpportunityFile = (params) => {
  return service({
    url: '/crmBusinessOpportunityFile/deleteCrmBusinessOpportunityFile',
    method: 'delete',
    params
  })
}

// @Tags CrmBusinessOpportunityFile
// @Summary 批量删除商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /crmBusinessOpportunityFile/deleteCrmBusinessOpportunityFile [delete]
export const deleteCrmBusinessOpportunityFileByIds = (params) => {
  return service({
    url: '/crmBusinessOpportunityFile/deleteCrmBusinessOpportunityFileByIds',
    method: 'delete',
    params
  })
}

// @Tags CrmBusinessOpportunityFile
// @Summary 更新商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CrmBusinessOpportunityFile true "更新商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /crmBusinessOpportunityFile/updateCrmBusinessOpportunityFile [put]
export const updateCrmBusinessOpportunityFile = (data) => {
  return service({
    url: '/crmBusinessOpportunityFile/updateCrmBusinessOpportunityFile',
    method: 'put',
    data
  })
}

// @Tags CrmBusinessOpportunityFile
// @Summary 用id查询商机文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CrmBusinessOpportunityFile true "用id查询商机文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /crmBusinessOpportunityFile/findCrmBusinessOpportunityFile [get]
export const findCrmBusinessOpportunityFile = (params) => {
  return service({
    url: '/crmBusinessOpportunityFile/findCrmBusinessOpportunityFile',
    method: 'get',
    params
  })
}

// @Tags CrmBusinessOpportunityFile
// @Summary 分页获取商机文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取商机文件列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /crmBusinessOpportunityFile/getCrmBusinessOpportunityFileList [get]
export const getCrmBusinessOpportunityFileList = (params) => {
  return service({
    url: '/crmBusinessOpportunityFile/getCrmBusinessOpportunityFileList',
    method: 'get',
    params
  })
}
