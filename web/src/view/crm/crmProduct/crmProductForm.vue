<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="产品名称:" prop="productName">
          <el-input v-model="formData.productName" :clearable="true"  placeholder="请输入产品名称" />
       </el-form-item>
        <el-form-item label="管理ID 销售代表:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item label="产品分组id:" prop="productGroupId">
          <el-input v-model.number="formData.productGroupId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产品类型id:" prop="productTypeId">
          <el-input v-model.number="formData.productTypeId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="库存:" prop="inventory">
          <el-input v-model.number="formData.inventory" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产品原价:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="产品折扣价:" prop="discountPrice">
          <el-input-number v-model="formData.discountPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="产品销售价格:" prop="salesPrice">
          <el-input-number v-model="formData.salesPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="资源id:" prop="resourceId">
          <el-input v-model.number="formData.resourceId" :clearable="true" placeholder="请输入" />
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
  createCrmProduct,
  updateCrmProduct,
  findCrmProduct
} from '@/api/crm/crmProduct'

defineOptions({
    name: 'CrmProductForm'
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
            productName: '',
            userId: 0,
            description: '',
            productGroupId: 0,
            productTypeId: 0,
            inventory: 0,
            price: 0,
            discountPrice: 0,
            salesPrice: 0,
            resourceId: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCrmProduct({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recrmProduct
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
               res = await createCrmProduct(formData.value)
               break
             case 'update':
               res = await updateCrmProduct(formData.value)
               break
             default:
               res = await createCrmProduct(formData.value)
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
