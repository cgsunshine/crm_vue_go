<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="工单标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入工单标题" />
       </el-form-item>
        <el-form-item label="工单描述:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入工单描述" />
       </el-form-item>
        <el-form-item label="提交人ID:" prop="submitterId">
          <el-input v-model.number="formData.submitterId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="指派人ID:" prop="assigneeId">
          <el-input v-model.number="formData.assigneeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="工单类别ID:" prop="categoryId">
          <el-input v-model.number="formData.categoryId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="工单状态:" prop="status">
       </el-form-item>
        <el-form-item label="工单优先级:" prop="priority">
       </el-form-item>
        <el-form-item label="estimatedResolutionTime字段:" prop="estimatedResolutionTime">
          <el-date-picker v-model="formData.estimatedResolutionTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="actualResolutionTime字段:" prop="actualResolutionTime">
          <el-date-picker v-model="formData.actualResolutionTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createCrmTickets,
  updateCrmTickets,
  findCrmTickets
} from '@/api/crm/crmTickets'

defineOptions({
    name: 'CrmTicketsForm'
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
            title: '',
            description: '',
            submitterId: 0,
            assigneeId: 0,
            categoryId: 0,
            estimatedResolutionTime: new Date(),
            actualResolutionTime: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmTickets({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmTickets
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
               res = await createCrmTickets(formData.value)
               break
             case 'update':
               res = await updateCrmTickets(formData.value)
               break
             default:
               res = await createCrmTickets(formData.value)
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
