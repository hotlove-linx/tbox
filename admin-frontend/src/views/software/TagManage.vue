<template>
  <div class="page-container">
    <el-card shadow="never">
      <div class="action-bar">
        <el-button type="primary" @click="openDialog()"><el-icon><Plus /></el-icon>新增标签</el-button>
      </div>
      <el-table :data="tableData" stripe border>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column prop="name" label="标签名称" width="150" />
        <el-table-column label="标签颜色" width="120" align="center">
          <template #default="{ row }">
            <div class="color-display">
              <span class="color-dot" :style="{ background: row.color }"></span>
              <span>{{ row.color }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="标签类型" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="row.type === 'system' ? 'primary' : 'warning'" size="small">
              {{ row.type === 'system' ? '系统' : '运营' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="softwareCount" label="关联软件数" width="110" align="center" />
        <el-table-column label="状态" width="80" align="center">
          <template #default="{ row }">
            <el-tag :type="row.enabled ? 'success' : 'danger'" size="small">{{ row.enabled ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="openDialog(row)">编辑</el-button>
            <el-button link :type="row.enabled ? 'warning' : 'success'" size="small" @click="row.enabled = !row.enabled">
              {{ row.enabled ? '禁用' : '启用' }}
            </el-button>
            <el-popconfirm title="确定删除此标签？" @confirm="handleDelete(row.id)">
              <template #reference>
                <el-button link type="danger" size="small">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="editingItem ? '编辑标签' : '新增标签'" width="480px">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入标签名称" />
        </el-form-item>
        <el-form-item label="颜色" prop="color">
          <el-color-picker v-model="form.color" />
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-radio-group v-model="form.type">
            <el-radio value="system">系统标签</el-radio>
            <el-radio value="operation">运营标签</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="form.enabled" active-text="启用" inactive-text="禁用" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'

const dialogVisible = ref(false)
const editingItem = ref<any>(null)
const formRef = ref<FormInstance>()

const form = reactive({ name: '', color: '#409eff', type: 'system', enabled: true })

const rules: FormRules = {
  name: [{ required: true, message: '请输入标签名称', trigger: 'blur' }],
  color: [{ required: true, message: '请选择颜色', trigger: 'change' }],
  type: [{ required: true, message: '请选择类型', trigger: 'change' }]
}

const tableData = ref([
  { id: 1, name: '热门', color: '#f56c6c', type: 'system', softwareCount: 45, enabled: true },
  { id: 2, name: '新品', color: '#67c23a', type: 'system', softwareCount: 32, enabled: true },
  { id: 3, name: '精品', color: '#e6a23c', type: 'operation', softwareCount: 28, enabled: true },
  { id: 4, name: '限免', color: '#409eff', type: 'operation', softwareCount: 15, enabled: true },
  { id: 5, name: '必备', color: '#9b59b6', type: 'system', softwareCount: 20, enabled: true },
  { id: 6, name: '独家', color: '#1abc9c', type: 'operation', softwareCount: 8, enabled: false }
])

function openDialog(item?: any) {
  editingItem.value = item || null
  if (item) {
    Object.assign(form, { name: item.name, color: item.color, type: item.type, enabled: item.enabled })
  } else {
    Object.assign(form, { name: '', color: '#409eff', type: 'system', enabled: true })
  }
  dialogVisible.value = true
}

async function handleSubmit() {
  await formRef.value?.validate()
  if (editingItem.value) {
    Object.assign(editingItem.value, form)
    ElMessage.success('编辑成功')
  } else {
    tableData.value.push({ id: Date.now(), ...form, softwareCount: 0 })
    ElMessage.success('新增成功')
  }
  dialogVisible.value = false
}

function handleDelete(id: number) {
  tableData.value = tableData.value.filter((item) => item.id !== id)
  ElMessage.success('删除成功')
}
</script>

<style scoped>
.action-bar { margin-bottom: 16px; }
.color-display { display: flex; align-items: center; gap: 6px; justify-content: center; }
.color-dot { width: 16px; height: 16px; border-radius: 50%; display: inline-block; border: 1px solid #dcdfe6; }
</style>
