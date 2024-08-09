<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="指派审批人id:" prop="assigneeId">
          <el-input v-model.number="formData.assigneeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="退款状态:" prop="refundStatus">
           <el-select v-model="formData.refundStatus" placeholder="请选择退款状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in refund_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="补充意见:" prop="comments">
          <el-input v-model="formData.comments" :clearable="true"  placeholder="请输入补充意见" />
       </el-form-item>
        <el-form-item label="审批是否有效 1 有效 2 失效（多人审批中，有人拒绝）:" prop="valid">
          <el-input v-model.number="formData.valid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="关联id 合同 商机 回款:" prop="associatedId">
          <el-input v-model.number="formData.associatedId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="退款类型 1押金:" prop="refundType">
          <el-input v-model.number="formData.refundType" :clearable="true" placeholder="请输入" />
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
  createCrmRefundTasks,
  updateCrmRefundTasks,
  findCrmRefundTasks
} from '@/api/crm/crmRefundTasks'

defineOptions({
    name: 'CrmRefundTasksForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const refund_statusOptions = ref([])
const formData = ref({
            assigneeId: 0,
            refundStatus: '',
            comments: '',
            valid: 0,
            associatedId: 0,
            refundType: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmRefundTasks({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmRefundTasks
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    refund_statusOptions.value = await getDictFunc('refund_status')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createCrmRefundTasks(formData.value)
               break
             case 'update':
               res = await updateCrmRefundTasks(formData.value)
               break
             default:
               res = await createCrmRefundTasks(formData.value)
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
