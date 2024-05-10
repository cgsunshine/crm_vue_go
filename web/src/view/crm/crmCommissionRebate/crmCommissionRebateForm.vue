<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="age字段:" prop="age">
          <el-input v-model.number="formData.age" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="账单ID:" prop="orderId">
          <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="员工ID 负责人:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="收款人:" prop="payee">
          <el-input v-model="formData.payee" :clearable="true"  placeholder="请输入收款人" />
       </el-form-item>
        <el-form-item label="收款方式:" prop="paymentMethod">
          <el-input v-model="formData.paymentMethod" :clearable="true"  placeholder="请输入收款方式" />
       </el-form-item>
        <el-form-item label="账户:" prop="account">
          <el-input v-model="formData.account" :clearable="true"  placeholder="请输入账户" />
       </el-form-item>
        <el-form-item label="金额:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="返利详情:" prop="rebateDetails">
          <el-input v-model="formData.rebateDetails" :clearable="true"  placeholder="请输入返利详情" />
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
  createCrmCommissionRebate,
  updateCrmCommissionRebate,
  findCrmCommissionRebate
} from '@/api/crm/crmCommissionRebate'

defineOptions({
    name: 'CrmCommissionRebateForm'
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
            age: 0,
            orderId: 0,
            userId: 0,
            payee: '',
            paymentMethod: '',
            account: '',
            amount: 0,
            rebateDetails: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmCommissionRebate({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmCommissionRebate
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
               res = await createCrmCommissionRebate(formData.value)
               break
             case 'update':
               res = await updateCrmCommissionRebate(formData.value)
               break
             default:
               res = await createCrmCommissionRebate(formData.value)
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
