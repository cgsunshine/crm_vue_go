<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="审批人:" prop="approvedById">
          <el-input v-model="formData.approvedById" :clearable="true"  placeholder="请输入审批人" />
       </el-form-item>
        <el-form-item label="审核状态:" prop="auditingStatus">
           <el-select v-model="formData.auditingStatus" placeholder="请选择审核状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in review_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="审核时间:" prop="auditingTime">
          <el-date-picker v-model="formData.auditingTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="账单ID:" prop="billId">
          <el-input v-model.number="formData.billId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="币种:" prop="currency">
          <el-input v-model="formData.currency" :clearable="true"  placeholder="请输入币种" />
       </el-form-item>
        <el-form-item label="客户ID:" prop="customerId">
          <el-input v-model.number="formData.customerId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item label="凭证:" prop="proof">
          <el-input v-model="formData.proof" :clearable="true"  placeholder="请输入凭证" />
       </el-form-item>
        <el-form-item label="管理ID 销售代表:" prop="userId">
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
  createCrmPaymentCollention,
  updateCrmPaymentCollention,
  findCrmPaymentCollention
} from '@/api/crm/crmPaymentCollention'

defineOptions({
    name: 'CrmPaymentCollentionForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const review_statusOptions = ref([])
const formData = ref({
            approvedById: '',
            auditingStatus: '',
            auditingTime: new Date(),
            billId: 0,
            currency: '',
            customerId: 0,
            description: '',
            proof: '',
            userId: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmPaymentCollention({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmPaymentCollention
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    review_statusOptions.value = await getDictFunc('review_status')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCrmPaymentCollention(formData.value)
               break
             case 'update':
               res = await updateCrmPaymentCollention(formData.value)
               break
             default:
               res = await createCrmPaymentCollention(formData.value)
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
