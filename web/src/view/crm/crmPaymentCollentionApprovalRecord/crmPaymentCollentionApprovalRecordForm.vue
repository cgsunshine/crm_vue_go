<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="审批意见:" prop="approveOpinion">
          <el-input v-model="formData.approveOpinion" :clearable="true"  placeholder="请输入审批意见" />
       </el-form-item>
        <el-form-item label="当前审批人ID:" prop="approverId">
          <el-input v-model.number="formData.approverId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="回款id:" prop="paymentCollentionId">
          <el-input v-model.number="formData.paymentCollentionId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="审批状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true"  placeholder="请输入审批状态" />
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
  createCrmPaymentCollentionApprovalRecord,
  updateCrmPaymentCollentionApprovalRecord,
  findCrmPaymentCollentionApprovalRecord
} from '@/api/crm/crmPaymentCollentionApprovalRecord'

defineOptions({
    name: 'CrmPaymentCollentionApprovalRecordForm'
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
            approveOpinion: '',
            approverId: 0,
            paymentCollentionId: 0,
            status: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmPaymentCollentionApprovalRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmPaymentCollentionApprovalRecord
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
               res = await createCrmPaymentCollentionApprovalRecord(formData.value)
               break
             case 'update':
               res = await updateCrmPaymentCollentionApprovalRecord(formData.value)
               break
             default:
               res = await createCrmPaymentCollentionApprovalRecord(formData.value)
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
