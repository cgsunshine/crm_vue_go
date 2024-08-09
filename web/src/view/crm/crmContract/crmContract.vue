<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="申请时间" prop="applicationTime">
            
            <template #label>
            <span>
              申请时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startApplicationTime" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endApplicationTime ? time.getTime() > searchInfo.endApplicationTime.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endApplicationTime" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startApplicationTime ? time.getTime() < searchInfo.startApplicationTime.getTime() : false"></el-date-picker>

        </el-form-item>
        <el-form-item label="合同名称" prop="contractName">
         <el-input v-model="searchInfo.contractName" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="合同状态" prop="contractStatus">
            <el-select v-model="searchInfo.contractStatus" clearable placeholder="请选择" @clear="()=>{searchInfo.contractStatus=undefined}">
              <el-option v-for="(item,key) in contract_statusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="客户ID" prop="customerId">
            
             <el-input v-model.number="searchInfo.customerId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="到期时间" prop="expirationTime">
            
            <template #label>
            <span>
              到期时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startExpirationTime" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endExpirationTime ? time.getTime() > searchInfo.endExpirationTime.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endExpirationTime" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startExpirationTime ? time.getTime() < searchInfo.startExpirationTime.getTime() : false"></el-date-picker>

        </el-form-item>
        <el-form-item label="订单ID" prop="orderId">
            
             <el-input v-model.number="searchInfo.orderId" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="审核结果" prop="reviewResult">
            <el-select v-model="searchInfo.reviewResult" clearable placeholder="请选择" @clear="()=>{searchInfo.reviewResult=undefined}">
              <el-option v-for="(item,key) in review_resultOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="审核状态" prop="reviewStatus">
            <el-select v-model="searchInfo.reviewStatus" clearable placeholder="请选择" @clear="()=>{searchInfo.reviewStatus=undefined}">
              <el-option v-for="(item,key) in review_statusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
         <el-table-column align="left" label="申请时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.applicationTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="合同文件路径" prop="contractFile" width="120" />
        <el-table-column align="left" label="合同名称" prop="contractName" width="120" />
        <el-table-column align="left" label="合同状态" prop="contractStatus" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.contractStatus,contract_statusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="合同类型" prop="contractTypeId" width="120" />
        <el-table-column align="left" label="客户ID" prop="customerId" width="120" />
        <el-table-column align="left" label="备注" prop="description" width="120" />
         <el-table-column align="left" label="到期时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.expirationTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="订单ID" prop="orderId" width="120" />
        <el-table-column align="left" label="审核结果" prop="reviewResult" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.reviewResult,review_resultOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="审核状态" prop="reviewStatus" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.reviewStatus,review_statusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="管理ID 销售代表" prop="userId" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateCrmContractFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #title>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="申请时间:"  prop="applicationTime" >
              <el-date-picker v-model="formData.applicationTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="合同文件路径:"  prop="contractFile" >
              <el-input v-model="formData.contractFile" :clearable="true"  placeholder="请输入合同文件路径" />
            </el-form-item>
            <el-form-item label="合同名称:"  prop="contractName" >
              <el-input v-model="formData.contractName" :clearable="true"  placeholder="请输入合同名称" />
            </el-form-item>
            <el-form-item label="合同状态:"  prop="contractStatus" >
              <el-select v-model="formData.contractStatus" placeholder="请选择合同状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in contract_statusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="合同类型:"  prop="contractTypeId" >
              <el-input v-model.number="formData.contractTypeId" :clearable="true" placeholder="请输入合同类型" />
            </el-form-item>
            <el-form-item label="客户ID:"  prop="customerId" >
              <el-input v-model.number="formData.customerId" :clearable="true" placeholder="请输入客户ID" />
            </el-form-item>
            <el-form-item label="备注:"  prop="description" >
              <el-input v-model="formData.description" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>
            <el-form-item label="到期时间:"  prop="expirationTime" >
              <el-date-picker v-model="formData.expirationTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="订单ID:"  prop="orderId" >
              <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入订单ID" />
            </el-form-item>
            <el-form-item label="审核结果:"  prop="reviewResult" >
              <el-select v-model="formData.reviewResult" placeholder="请选择审核结果" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in review_resultOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="审核状态:"  prop="reviewStatus" >
              <el-select v-model="formData.reviewStatus" placeholder="请选择审核状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in review_statusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="管理ID 销售代表:"  prop="userId" >
              <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入管理ID 销售代表" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer size="800" v-model="detailShow" :before-close="closeDetailShow" title="查看详情" destroy-on-close>
          <template #title>
             <div class="flex justify-between items-center">
               <span class="text-lg">查看详情</span>
             </div>
         </template>
        <el-descriptions :column="1" border>
                <el-descriptions-item label="申请时间">
                      {{ formatDate(formData.applicationTime) }}
                </el-descriptions-item>
                <el-descriptions-item label="合同文件路径">
                        {{ formData.contractFile }}
                </el-descriptions-item>
                <el-descriptions-item label="合同名称">
                        {{ formData.contractName }}
                </el-descriptions-item>
                <el-descriptions-item label="合同状态">
                        {{ filterDict(formData.contractStatus,contract_statusOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="合同类型">
                        {{ formData.contractTypeId }}
                </el-descriptions-item>
                <el-descriptions-item label="客户ID">
                        {{ formData.customerId }}
                </el-descriptions-item>
                <el-descriptions-item label="备注">
                        {{ formData.description }}
                </el-descriptions-item>
                <el-descriptions-item label="到期时间">
                      {{ formatDate(formData.expirationTime) }}
                </el-descriptions-item>
                <el-descriptions-item label="订单ID">
                        {{ formData.orderId }}
                </el-descriptions-item>
                <el-descriptions-item label="审核结果">
                        {{ filterDict(formData.reviewResult,review_resultOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="审核状态">
                        {{ filterDict(formData.reviewStatus,review_statusOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="管理ID 销售代表">
                        {{ formData.userId }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createCrmContract,
  deleteCrmContract,
  deleteCrmContractByIds,
  updateCrmContract,
  findCrmContract,
  getCrmContractList
} from '@/api/crm/crmContract'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'CrmContract'
})

// 自动化生成的字典（可能为空）以及字段
const contract_statusOptions = ref([])
const review_resultOptions = ref([])
const review_statusOptions = ref([])
const formData = ref({
        applicationTime: new Date(),
        contractFile: '',
        contractName: '',
        contractStatus: '',
        contractTypeId: 0,
        customerId: 0,
        description: '',
        expirationTime: new Date(),
        orderId: 0,
        reviewResult: '',
        reviewStatus: '',
        userId: 0,
        })


// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
        applicationTime : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startApplicationTime && !searchInfo.value.endApplicationTime) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startApplicationTime && searchInfo.value.endApplicationTime) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startApplicationTime && searchInfo.value.endApplicationTime && (searchInfo.value.startApplicationTime.getTime() === searchInfo.value.endApplicationTime.getTime() || searchInfo.value.startApplicationTime.getTime() > searchInfo.value.endApplicationTime.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
        expirationTime : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startExpirationTime && !searchInfo.value.endExpirationTime) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startExpirationTime && searchInfo.value.endExpirationTime) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startExpirationTime && searchInfo.value.endExpirationTime && (searchInfo.value.startExpirationTime.getTime() === searchInfo.value.endExpirationTime.getTime() || searchInfo.value.startExpirationTime.getTime() > searchInfo.value.endExpirationTime.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getCrmContractList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    contract_statusOptions.value = await getDictFunc('contract_status')
    review_resultOptions.value = await getDictFunc('review_result')
    review_statusOptions.value = await getDictFunc('review_status')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteCrmContractFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteCrmContractByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateCrmContractFunc = async(row) => {
    const res = await findCrmContract({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recrmContract
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCrmContractFunc = async (row) => {
    const res = await deleteCrmContract({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findCrmContract({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recrmContract
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          applicationTime: new Date(),
          contractFile: '',
          contractName: '',
          contractStatus: '',
          contractTypeId: 0,
          customerId: 0,
          description: '',
          expirationTime: new Date(),
          orderId: 0,
          reviewResult: '',
          reviewStatus: '',
          userId: 0,
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        applicationTime: new Date(),
        contractFile: '',
        contractName: '',
        contractStatus: '',
        contractTypeId: 0,
        customerId: 0,
        description: '',
        expirationTime: new Date(),
        orderId: 0,
        reviewResult: '',
        reviewStatus: '',
        userId: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCrmContract(formData.value)
                  break
                case 'update':
                  res = await updateCrmContract(formData.value)
                  break
                default:
                  res = await createCrmContract(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
