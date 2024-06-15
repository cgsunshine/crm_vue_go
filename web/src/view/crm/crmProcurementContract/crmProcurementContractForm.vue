<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="合同金额:" prop="contractAmount">
          <el-input-number v-model="formData.contractAmount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="合同文件:" prop="contractFile">
          <el-input v-model="formData.contractFile" :clearable="true"  placeholder="请输入合同文件" />
       </el-form-item>
        <el-form-item label="合同名称:" prop="contractName">
          <el-input v-model="formData.contractName" :clearable="true"  placeholder="请输入合同名称" />
       </el-form-item>
        <el-form-item label="合同状态:" prop="contractStatus">
           <el-select v-model="formData.contractStatus" placeholder="请选择合同状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in contract_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="创建时间:" prop="creationTime">
          <el-date-picker v-model="formData.creationTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="到期时间:" prop="expirationTime">
          <el-date-picker v-model="formData.expirationTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="负责人:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createCrmProcurementContract,
  updateCrmProcurementContract,
  findCrmProcurementContract
} from '@/api/crm/crmProcurementContract'

defineOptions({
    name: 'CrmProcurementContractForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const contract_statusOptions = ref([])
const formData = ref({
            contractAmount: 0,
            contractFile: '',
            contractName: '',
            contractStatus: '',
            creationTime: new Date(),
            expirationTime: new Date(),
            userId: 0,
        })
// 验证规则
const rule = reactive({
               contractAmount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               contractName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmProcurementContract({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmProcurementContract
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    contract_statusOptions.value = await getDictFunc('contract_status')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCrmProcurementContract(formData.value)
               break
             case 'update':
               res = await updateCrmProcurementContract(formData.value)
               break
             default:
               res = await createCrmProcurementContract(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
