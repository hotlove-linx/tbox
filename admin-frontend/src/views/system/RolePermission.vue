<template>
  <div class="page-container">
    <el-card shadow="never">
      <div class="action-bar">
        <el-button type="primary" @click="openRoleDialog()"><el-icon><Plus /></el-icon>新增角色</el-button>
      </div>

      <el-table :data="roleList" stripe border>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column prop="name" label="角色名称" width="150" />
        <el-table-column prop="description" label="描述" min-width="200" />
        <el-table-column prop="adminCount" label="管理员数" width="100" align="center" />
        <el-table-column label="状态" width="80" align="center">
          <template #default="{ row }">
            <el-tag :type="row.enabled ? 'success' : 'danger'" size="small">{{ row.enabled ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createdAt" label="创建时间" width="170" />
        <el-table-column label="操作" width="250">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="openRoleDialog(row)">编辑</el-button>
            <el-button link type="success" size="small" @click="openPermDialog(row)">权限配置</el-button>
            <el-popconfirm v-if="!row.isSystem" title="确定删除此角色？" @confirm="handleDelete(row.id)">
              <template #reference>
                <el-button link type="danger" size="small">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- Role Dialog -->
    <el-dialog v-model="roleDialogVisible" :title="editingRole ? '编辑角色' : '新增角色'" width="480px">
      <el-form :model="roleForm" :rules="roleRules" ref="roleFormRef" label-width="80px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="roleForm.name" placeholder="请输入角色名称" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="roleForm.description" type="textarea" :rows="3" placeholder="角色描述" />
        </el-form-item>
        <el-form-item label="状态">
          <el-switch v-model="roleForm.enabled" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="roleDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitRole">确定</el-button>
      </template>
    </el-dialog>

    <!-- Permission Matrix Dialog -->
    <el-dialog v-model="permDialogVisible" :title="`权限配置 - ${currentRole?.name || ''}`" width="800px" top="5vh">
      <el-table :data="permissionMatrix" stripe border>
        <el-table-column prop="module" label="功能模块" width="150" fixed />
        <el-table-column label="查看" width="80" align="center">
          <template #default="{ row }">
            <el-checkbox v-model="row.view" />
          </template>
        </el-table-column>
        <el-table-column label="新增" width="80" align="center">
          <template #default="{ row }">
            <el-checkbox v-model="row.create" :disabled="!row.hasCreate" />
          </template>
        </el-table-column>
        <el-table-column label="编辑" width="80" align="center">
          <template #default="{ row }">
            <el-checkbox v-model="row.edit" :disabled="!row.hasEdit" />
          </template>
        </el-table-column>
        <el-table-column label="删除" width="80" align="center">
          <template #default="{ row }">
            <el-checkbox v-model="row.delete" :disabled="!row.hasDelete" />
          </template>
        </el-table-column>
        <el-table-column label="审核" width="80" align="center">
          <template #default="{ row }">
            <el-checkbox v-model="row.audit" :disabled="!row.hasAudit" />
          </template>
        </el-table-column>
        <el-table-column label="导出" width="80" align="center">
          <template #default="{ row }">
            <el-checkbox v-model="row.export" :disabled="!row.hasExport" />
          </template>
        </el-table-column>
      </el-table>
      <template #footer>
        <el-button @click="permDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="savePermissions">保存权限</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'

const roleDialogVisible = ref(false)
const permDialogVisible = ref(false)
const editingRole = ref<any>(null)
const currentRole = ref<any>(null)
const roleFormRef = ref<FormInstance>()

const roleForm = reactive({ name: '', description: '', enabled: true })
const roleRules: FormRules = {
  name: [{ required: true, message: '请输入角色名称', trigger: 'blur' }]
}

const roleList = ref([
  { id: 1, name: '超级管理员', description: '拥有系统所有权限', adminCount: 1, enabled: true, isSystem: true, createdAt: '2024-01-01 00:00:00' },
  { id: 2, name: '运营管理员', description: '负责内容运营、推荐管理等', adminCount: 2, enabled: true, isSystem: false, createdAt: '2024-01-01 00:00:00' },
  { id: 3, name: '审核员', description: '负责软件审核、评论审核、举报处理', adminCount: 3, enabled: true, isSystem: false, createdAt: '2024-01-01 00:00:00' },
  { id: 4, name: '客服', description: '处理用户反馈和咨询', adminCount: 2, enabled: true, isSystem: false, createdAt: '2024-01-01 00:00:00' }
])

const permissionMatrix = ref([
  { module: '仪表盘', view: true, create: false, edit: false, delete: false, audit: false, export: false, hasCreate: false, hasEdit: false, hasDelete: false, hasAudit: false, hasExport: false },
  { module: '软件管理', view: true, create: true, edit: true, delete: true, audit: false, export: true, hasCreate: true, hasEdit: true, hasDelete: true, hasAudit: false, hasExport: true },
  { module: '分类管理', view: true, create: true, edit: true, delete: true, audit: false, export: false, hasCreate: true, hasEdit: true, hasDelete: true, hasAudit: false, hasExport: false },
  { module: '标签管理', view: true, create: true, edit: true, delete: true, audit: false, export: false, hasCreate: true, hasEdit: true, hasDelete: true, hasAudit: false, hasExport: false },
  { module: 'Banner管理', view: true, create: true, edit: true, delete: true, audit: false, export: false, hasCreate: true, hasEdit: true, hasDelete: true, hasAudit: false, hasExport: false },
  { module: '推荐管理', view: true, create: true, edit: true, delete: true, audit: false, export: false, hasCreate: true, hasEdit: true, hasDelete: true, hasAudit: false, hasExport: false },
  { module: '公告管理', view: true, create: true, edit: true, delete: true, audit: false, export: false, hasCreate: true, hasEdit: true, hasDelete: true, hasAudit: false, hasExport: false },
  { module: '软件审核', view: true, create: false, edit: false, delete: false, audit: true, export: false, hasCreate: false, hasEdit: false, hasDelete: false, hasAudit: true, hasExport: false },
  { module: '评论审核', view: true, create: false, edit: false, delete: true, audit: true, export: false, hasCreate: false, hasEdit: false, hasDelete: true, hasAudit: true, hasExport: false },
  { module: '举报处理', view: true, create: false, edit: false, delete: false, audit: true, export: false, hasCreate: false, hasEdit: false, hasDelete: false, hasAudit: true, hasExport: false },
  { module: '用户管理', view: true, create: false, edit: true, delete: false, audit: false, export: true, hasCreate: false, hasEdit: true, hasDelete: false, hasAudit: false, hasExport: true },
  { module: '数据统计', view: true, create: false, edit: false, delete: false, audit: false, export: true, hasCreate: false, hasEdit: false, hasDelete: false, hasAudit: false, hasExport: true },
  { module: '系统管理', view: true, create: true, edit: true, delete: true, audit: false, export: false, hasCreate: true, hasEdit: true, hasDelete: true, hasAudit: false, hasExport: false }
])

function openRoleDialog(item?: any) {
  editingRole.value = item || null
  if (item) {
    Object.assign(roleForm, { name: item.name, description: item.description, enabled: item.enabled })
  } else {
    Object.assign(roleForm, { name: '', description: '', enabled: true })
  }
  roleDialogVisible.value = true
}

async function submitRole() {
  await roleFormRef.value?.validate()
  if (editingRole.value) {
    Object.assign(editingRole.value, roleForm)
    ElMessage.success('编辑成功')
  } else {
    roleList.value.push({
      id: Date.now(), name: roleForm.name, description: roleForm.description,
      adminCount: 0, enabled: roleForm.enabled, isSystem: false, createdAt: new Date().toLocaleString()
    })
    ElMessage.success('新增成功')
  }
  roleDialogVisible.value = false
}

function openPermDialog(role: any) {
  currentRole.value = role
  permDialogVisible.value = true
}

function savePermissions() {
  permDialogVisible.value = false
  ElMessage.success('权限配置已保存')
}

function handleDelete(id: number) {
  roleList.value = roleList.value.filter((r) => r.id !== id)
  ElMessage.success('删除成功')
}
</script>

<style scoped>
.action-bar { margin-bottom: 16px; }
</style>
