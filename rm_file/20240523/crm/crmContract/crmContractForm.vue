<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="产品名称:" prop="orderId">
          <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产品ID:" prop="customerId">
          <el-input v-model.number="formData.customerId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="管理ID 销售代表:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item label="产品原价:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="产品折扣价:" prop="salesPrice">
          <el-input-number v-model="formData.salesPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="币种:" prop="currency">
          <el-input v-model="formData.currency" :clearable="true"  placeholder="请输入币种" />
       </el-form-item>
        <el-form-item label="折扣率:" prop="discountRate">
          <el-input v-model.number="formData.discountRate" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="合同类型:" prop="contractTypeId">
          <el-input v-model.number="formData.contractTypeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="申请时间:" prop="applicationTime">
          <el-date-picker v-model="formData.applicationTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="到期时间:" prop="expirationTime">
          <el-date-picker v-model="formData.expirationTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="合同状态:" prop="contractStatus">
          <el-input v-model="formData.contractStatus" :clearable="true"  placeholder="请输入合同状态" />
       </el-form-item>
        <el-form-item label="审核状态:" prop="reviewStatus">
          <el-input v-model="formData.reviewStatus" :clearable="true"  placeholder="请输入审核状态" />
       </el-form-item>
        <el-form-item label="审核结果:" prop="reviewResult">
          <el-input v-model="formData.reviewResult" :clearable="true"  placeholder="请输入审核结果" />
       </el-form-item>
        <el-form-item label="合同文件:" prop="contractDocument">
          <el-input v-model="formData.contractDocument" :clearable="true"  placeholder="请输入合同文件" />
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
  createCrmContract,
  updateCrmContract,
  findCrmContract
} from '@/api/crm/crmContract'

defineOptions({
    name: 'CrmContractForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            orderId: 0,
            customerId: 0,
            userId: 0,
            description: '',
            price: 0,
            salesPrice: 0,
            currency: '',
            discountRate: 0,
            contractTypeId: 0,
            applicationTime: new Date(),
            expirationTime: new Date(),
            contractStatus: '',
            reviewStatus: '',
            reviewResult: '',
            contractDocument: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmContract({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmContract
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
