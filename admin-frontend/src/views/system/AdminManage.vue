<template>
  <div class="page-container">
    <el-card shadow="never">
      <div class="action-bar">
        <el-button type="primary" @click="openDialog()"><el-icon><Plus /></el-icon>新增管理员</el-button>
      </div>

      <el-table :data="tableData" stripe border>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column label="管理员信息" min-width="220">
          <template #default="{ row }">
            <div style="display: flex; align-items: center; gap: 10px;">
              <el-avatar :size="36" :src="row.avatar">{{ row.name[0] }}</el-avatar>
              <div>
                <div style="font-weight: 500;">{{ row.name }}</div>
                <div style="font-size: 12px; color: #909399;">{{ row.email }}</div>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="角色" width="120" align="center">
          <template #default="{ row }">
            <el-tag :type="row.role === '超级管理员' ? 'danger' : row.role === '运营管理员' ? 'primary' : 'info'" size="small">{{ row.role }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="状态" width="80" align="center">
          <template #default="{ row }">
            <el-tag :type="row.enabled ? 'success' : 'danger'" size="small">{{ row.enabled ? '正常' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="lastLogin" label="最后登录" width="170" />
        <el-table-column prop="createdAt" label="创建时间" width="170" />
        <el-table-column label="操作" width="250">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="openDialog(row)">编辑</el-button>
            <el-button link type="warning" size="small" @click="resetPassword(row)">重置密码</el-button>
            <el-button link :type="row.enabled ? 'info' : 'success'" size="small" @click="toggleStatus(row)">
              {{ row.enabled ? '禁用' : '启用' }}
            </el-button>
            <el-popconfirm title="确定删除此管理员？" @confirm="handleDelete(row.id)">
              <template #reference>
                <el-button link type="danger" size="small" :disabled="row.role === '超级管理员'">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="editingItem ? '编辑管理员' : '新增管理员'" width="520px">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="姓名" prop="name">
          <el-input v-model="form.name" placeholder="请输入姓名" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item v-if="!editingItem" label="密码" prop="password">
          <el-input v-model="form.password" type="password" show-password placeholder="请输入初始密码" />
        </el-form-item>
        <el-form-item label="角色" prop="role">
          <el-select v-model="form.role" style="width: 100%">
            <el-option label="超级管理员" value="超级管理员" />
            <el-option label="运营管理员" value="运营管理员" />
            <el-option label="审核员" value="审核员" />
            <el-option label="客服" value="客服" />
          </el-select>
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
import { ElMessage, ElMessageBox } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'

const dialogVisible = ref(false)
const editingItem = ref<any>(null)
const formRef = ref<FormInstance>()

const form = reactive({ name: '', email: '', password: '', role: '', enabled: true })

const rules: FormRules = {
  name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱', trigger: 'blur' }
  ],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }, { min: 6, message: '密码至少6位', trigger: 'blur' }],
  role: [{ required: true, message: '请选择角色', trigger: 'change' }]
}

const tableData = ref([
  { id: 1, name: '超级管理员', email: 'admin@tbox.com', avatar: '', role: '超级管理员', enabled: true, lastLogin: '2024-03-15 20:30:00', createdAt: '2024-01-01 00:00:00' },
  { id: 2, name: '运营小王', email: 'wang@tbox.com', avatar: '', role: '运营管理员', enabled: true, lastLogin: '2024-03-15 18:00:00', createdAt: '2024-01-15 10:00:00' },
  { id: 3, name: '审核员张', email: 'zhang@tbox.com', avatar: '', role: '审核员', enabled: true, lastLogin: '2024-03-15 17:30:00', createdAt: '2024-02-01 09:00:00' },
  { id: 4, name: '客服小李', email: 'li@tbox.com', avatar: '', role: '客服', enabled: true, lastLogin: '2024-03-14 16:00:00', createdAt: '2024-02-10 14:00:00' },
  { id: 5, name: '前运营', email: 'ex@tbox.com', avatar: '', role: '运营管理员', enabled: false, lastLogin: '2024-02-01 10:00:00', createdAt: '2024-01-20 08:00:00' }
])

function openDialog(item?: any) {
  editingItem.value = item || null
  if (item) {
    Object.assign(form, { name: item.name, email: item.email, password: '', role: item.role, enabled: item.enabled })
  } else {
    Object.assign(form, { name: '', email: '', password: '', role: '', enabled: true })
  }
  dialogVisible.value = true
}

async function handleSubmit() {
  await formRef.value?.validate()
  if (editingItem.value) {
    Object.assign(editingItem.value, { name: form.name, email: form.email, role: form.role, enabled: form.enabled })
    ElMessage.success('编辑成功')
  } else {
    tableData.value.push({
      id: Date.now(), name: form.name, email: form.email, avatar: '', role: form.role,
      enabled: form.enabled, lastLogin: '-', createdAt: new Date().toLocaleString()
    })
    ElMessage.success('新增成功')
  }
  dialogVisible.value = false
}

function resetPassword(row: any) {
  ElMessageBox.confirm(`确定要重置 ${row.name} 的密码吗？`, '确认', { type: 'warning' }).then(() => {
    ElMessage.success('密码已重置')
  })
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
.action-bar { margin-bottom: 16px; }
</style>
