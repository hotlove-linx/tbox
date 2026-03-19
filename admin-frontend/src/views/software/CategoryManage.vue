<template>
  <div class="page-container">
    <el-card shadow="never">
      <div class="action-bar">
        <el-button type="primary" @click="openDialog()"><el-icon><Plus /></el-icon>新增分类</el-button>
      </div>
      <el-table :data="tableData" stripe border>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column prop="name" label="分类名称" width="150" />
        <el-table-column label="分类图标" width="100" align="center">
          <template #default="{ row }">
            <el-icon :size="24"><component :is="row.icon" /></el-icon>
          </template>
        </el-table-column>
        <el-table-column prop="sort" label="排序权重" width="100" align="center" sortable />
        <el-table-column prop="softwareCount" label="软件数量" width="100" align="center" />
        <el-table-column label="状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="row.enabled ? 'success' : 'danger'" size="small">{{ row.enabled ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" label="创建时间" width="170" />
        <el-table-column label="操作" width="220">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="openDialog(row)">编辑</el-button>
            <el-button link :type="row.enabled ? 'warning' : 'success'" size="small" @click="toggleStatus(row)">
              {{ row.enabled ? '禁用' : '启用' }}
            </el-button>
            <el-popconfirm title="确定删除此分类？" @confirm="handleDelete(row.id)">
              <template #reference>
                <el-button link type="danger" size="small">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- Dialog -->
    <el-dialog v-model="dialogVisible" :title="editingItem ? '编辑分类' : '新增分类'" width="500px">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入分类名称" />
        </el-form-item>
        <el-form-item label="图标" prop="icon">
          <el-select v-model="form.icon" placeholder="选择图标" style="width: 100%">
            <el-option v-for="icon in iconOptions" :key="icon" :label="icon" :value="icon">
              <div style="display: flex; align-items: center; gap: 8px;">
                <el-icon><component :is="icon" /></el-icon>
                <span>{{ icon }}</span>
              </div>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="form.sort" :min="0" :max="999" style="width: 100%" />
        </el-form-item>
        <el-form-item label="状态" prop="enabled">
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
const iconOptions = ['Box', 'Monitor', 'Headset', 'ChatDotRound', 'Film', 'Reading', 'Tools', 'Folder', 'Star', 'Trophy']

const form = reactive({ name: '', icon: '', sort: 0, enabled: true })

const rules: FormRules = {
  name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }],
  icon: [{ required: true, message: '请选择图标', trigger: 'change' }]
}

const tableData = ref([
  { id: 1, name: '工具软件', icon: 'Tools', sort: 10, softwareCount: 580, enabled: true, createdAt: '2024-01-01 00:00:00' },
  { id: 2, name: '游戏娱乐', icon: 'Monitor', sort: 20, softwareCount: 420, enabled: true, createdAt: '2024-01-01 00:00:00' },
  { id: 3, name: '效率办公', icon: 'Folder', sort: 30, softwareCount: 350, enabled: true, createdAt: '2024-01-01 00:00:00' },
  { id: 4, name: '社交通讯', icon: 'ChatDotRound', sort: 40, softwareCount: 280, enabled: true, createdAt: '2024-01-01 00:00:00' },
  { id: 5, name: '影音播放', icon: 'Film', sort: 50, softwareCount: 220, enabled: true, createdAt: '2024-01-01 00:00:00' },
  { id: 6, name: '教育学习', icon: 'Reading', sort: 60, softwareCount: 180, enabled: true, createdAt: '2024-01-01 00:00:00' },
  { id: 7, name: '系统工具', icon: 'Box', sort: 70, softwareCount: 150, enabled: false, createdAt: '2024-01-01 00:00:00' }
])

function openDialog(item?: any) {
  editingItem.value = item || null
  if (item) {
    form.name = item.name
    form.icon = item.icon
    form.sort = item.sort
    form.enabled = item.enabled
  } else {
    form.name = ''
    form.icon = ''
    form.sort = 0
    form.enabled = true
  }
  dialogVisible.value = true
}

async function handleSubmit() {
  await formRef.value?.validate()
  if (editingItem.value) {
    Object.assign(editingItem.value, { name: form.name, icon: form.icon, sort: form.sort, enabled: form.enabled })
    ElMessage.success('编辑成功')
  } else {
    tableData.value.push({
      id: Date.now(),
      name: form.name,
      icon: form.icon,
      sort: form.sort,
      softwareCount: 0,
      enabled: form.enabled,
      createdAt: new Date().toLocaleString()
    })
    ElMessage.success('新增成功')
  }
  dialogVisible.value = false
}

function toggleStatus(row: any) {
  row.enabled = !row.enabled
  ElMessage.success(row.enabled ? '已启用' : '已禁用')
}

function handleDelete(id: number) {
  tableData.value = tableData.value.filter((item) => item.id !== id)
  ElMessage.success('删除成功')
}
</script>

<style scoped>
.action-bar {
  margin-bottom: 16px;
}
</style>
