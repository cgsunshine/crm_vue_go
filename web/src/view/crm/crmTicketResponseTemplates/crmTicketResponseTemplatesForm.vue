<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="适用工单类别ID:" prop="categoryId">
          <el-input v-model.number="formData.categoryId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="回复内容:" prop="content">
          <el-input v-model="formData.content" :clearable="true"  placeholder="请输入回复内容" />
       </el-form-item>
        <el-form-item label="创建者ID:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="回复模板名称:" prop="templateName">
          <el-input v-model="formData.templateName" :clearable="true"  placeholder="请输入回复模板名称" />
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
  createCrmTicketResponseTemplates,
  updateCrmTicketResponseTemplates,
  findCrmTicketResponseTemplates
} from '@/api/crm/crmTicketResponseTemplates'

defineOptions({
    name: 'CrmTicketResponseTemplatesForm'
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
            categoryId: 0,
            content: '',
            createdBy: 0,
            templateName: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmTicketResponseTemplates({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmTicketResponseTemplates
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
               res = await createCrmTicketResponseTemplates(formData.value)
               break
             case 'update':
               res = await updateCrmTicketResponseTemplates(formData.value)
               break
             default:
               res = await createCrmTicketResponseTemplates(formData.value)
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
