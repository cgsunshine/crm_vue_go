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
        
        <el-table-column align="left" label="商机名称" prop="name" width="120" />
        <el-table-column align="left" label="备注" prop="description" width="120" />
        <el-table-column align="left" label="单位" prop="unit" width="120" />
        <el-table-column align="left" label="客户ID" prop="customerId" width="120" />
        <el-table-column align="left" label="员工id" prop="userId" width="120" />
        <el-table-column align="left" label="产品id" prop="productId" width="120" />
        <el-table-column align="left" label="商机金额" prop="price" width="120" />
        <el-table-column align="left" label="商机状态" prop="status" width="120" />
         <el-table-column align="left" label="商机录入时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.inputTime) }}</template>
         </el-table-column>
         <el-table-column align="left" label="报价有效期" width="180">
            <template #default="scope">{{ formatDate(scope.row.offerValidityPeriod) }}</template>
         </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateCrmBusinessOpportunityFunc(scope.row)">变更</el-button>
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
            <el-form-item label="商机名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入商机名称" />
            </el-form-item>
            <el-form-item label="备注:"  prop="description" >
              <el-input v-model="formData.description" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>
            <el-form-item label="单位:"  prop="unit" >
              <el-input v-model="formData.unit" :clearable="true"  placeholder="请输入单位" />
            </el-form-item>
            <el-form-item label="客户ID:"  prop="customerId" >
              <el-input v-model.number="formData.customerId" :clearable="true" placeholder="请输入客户ID" />
            </el-form-item>
            <el-form-item label="员工id:"  prop="userId" >
              <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入员工id" />
            </el-form-item>
            <el-form-item label="产品id:"  prop="productId" >
              <el-input v-model.number="formData.productId" :clearable="true" placeholder="请输入产品id" />
            </el-form-item>
            <el-form-item label="商机金额:"  prop="price" >
              <el-input-number v-model="formData.price"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="商机状态:"  prop="status" >
              <el-input v-model="formData.status" :clearable="true"  placeholder="请输入商机状态" />
            </el-form-item>
            <el-form-item label="商机录入时间:"  prop="inputTime" >
              <el-date-picker v-model="formData.inputTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="报价有效期:"  prop="offerValidityPeriod" >
              <el-date-picker v-model="formData.offerValidityPeriod" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
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
                <el-descriptions-item label="商机名称">
                        {{ formData.name }}
                </el-descriptions-item>
                <el-descriptions-item label="备注">
                        {{ formData.description }}
                </el-descriptions-item>
                <el-descriptions-item label="单位">
                        {{ formData.unit }}
                </el-descriptions-item>
                <el-descriptions-item label="客户ID">
                        {{ formData.customerId }}
                </el-descriptions-item>
                <el-descriptions-item label="员工id">
                        {{ formData.userId }}
                </el-descriptions-item>
                <el-descriptions-item label="产品id">
                        {{ formData.productId }}
                </el-descriptions-item>
                <el-descriptions-item label="商机金额">
                        {{ formData.price }}
                </el-descriptions-item>
                <el-descriptions-item label="商机状态">
                        {{ formData.status }}
                </el-descriptions-item>
                <el-descriptions-item label="商机录入时间">
                      {{ formatDate(formData.inputTime) }}
                </el-descriptions-item>
                <el-descriptions-item label="报价有效期">
                      {{ formatDate(formData.offerValidityPeriod) }}
                </el-descriptions-item>
        </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createCrmBusinessOpportunity,
  deleteCrmBusinessOpportunity,
  deleteCrmBusinessOpportunityByIds,
  updateCrmBusinessOpportunity,
  findCrmBusinessOpportunity,
  getCrmBusinessOpportunityList
} from '@/api/crm/crmBusinessOpportunity'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'CrmBusinessOpportunity'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        name: '',
        description: '',
        unit: '',
        customerId: 0,
        userId: 0,
        productId: 0,
        price: 0,
        status: '',
        inputTime: new Date(),
        offerValidityPeriod: new Date(),
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
  const table = await getCrmBusinessOpportunityList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteCrmBusinessOpportunityFunc(row)
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
      const res = await deleteCrmBusinessOpportunityByIds({ IDs })
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
const updateCrmBusinessOpportunityFunc = async(row) => {
    const res = await findCrmBusinessOpportunity({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.recrmBusinessOpportunity
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteCrmBusinessOpportunityFunc = async (row) => {
    const res = await deleteCrmBusinessOpportunity({ ID: row.ID })
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
  const res = await findCrmBusinessOpportunity({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.recrmBusinessOpportunity
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          name: '',
          description: '',
          unit: '',
          customerId: 0,
          userId: 0,
          productId: 0,
          price: 0,
          status: '',
          inputTime: new Date(),
          offerValidityPeriod: new Date(),
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
        name: '',
        description: '',
        unit: '',
        customerId: 0,
        userId: 0,
        productId: 0,
        price: 0,
        status: '',
        inputTime: new Date(),
        offerValidityPeriod: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createCrmBusinessOpportunity(formData.value)
                  break
                case 'update':
                  res = await updateCrmBusinessOpportunity(formData.value)
                  break
                default:
                  res = await createCrmBusinessOpportunity(formData.value)
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
