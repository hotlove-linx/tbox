<template>
  <div class="page-container">
    <el-card shadow="never">
      <div class="action-bar">
        <el-button type="primary" @click="openDialog()"><el-icon><Plus /></el-icon>新增公告</el-button>
      </div>
      <el-table :data="tableData" stripe border>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column prop="title" label="标题" min-width="200" show-overflow-tooltip />
        <el-table-column label="类型" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="typeMap[row.type]?.type || 'info'" size="small">{{ typeMap[row.type]?.label || row.type }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="发布状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="statusMap[row.status]?.type || 'info'" size="small">{{ statusMap[row.status]?.label || row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="pushMethod" label="推送方式" width="100" />
        <el-table-column prop="targetUsers" label="目标用户" width="100" />
        <el-table-column prop="publishTime" label="发布时间" width="170" />
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="openDialog(row)">编辑</el-button>
            <el-button v-if="row.status === 'draft'" link type="success" size="small" @click="publishItem(row)">发布</el-button>
            <el-popconfirm title="确定删除此公告？" @confirm="handleDelete(row.id)">
              <template #reference>
                <el-button link type="danger" size="small">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible" :title="editingItem ? '编辑公告' : '新增公告'" width="700px" top="5vh">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="标题" prop="title">
          <el-input v-model="form.title" placeholder="请输入公告标题" />
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select v-model="form.type" style="width: 100%">
            <el-option label="系统公告" value="system" />
            <el-option label="活动通知" value="activity" />
            <el-option label="更新日志" value="update" />
            <el-option label="维护通知" value="maintenance" />
          </el-select>
        </el-form-item>
        <el-form-item label="公告内容" prop="content">
          <el-input v-model="form.content" type="textarea" :rows="8" placeholder="请输入公告内容（支持富文本）" />
        </el-form-item>
        <el-form-item label="推送方式">
          <el-checkbox-group v-model="form.pushMethods">
            <el-checkbox value="app">App推送</el-checkbox>
            <el-checkbox value="email">邮件</el-checkbox>
            <el-checkbox value="inapp">站内信</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-form-item label="目标范围">
          <el-radio-group v-model="form.targetScope">
            <el-radio value="all">全部用户</el-radio>
            <el-radio value="active">活跃用户</el-radio>
            <el-radio value="new">新用户</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="定时发布">
          <el-switch v-model="form.scheduled" />
          <el-date-picker v-if="form.scheduled" v-model="form.scheduledTime" type="datetime" placeholder="选择发布时间" style="margin-left: 12px;" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button @click="saveAsDraft">保存草稿</el-button>
        <el-button type="primary" @click="handleSubmit">{{ form.scheduled ? '定时发布' : '立即发布' }}</el-button>
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

const typeMap: Record<string, any> = {
  system: { label: '系统公告', type: 'primary' },
  activity: { label: '活动通知', type: 'success' },
  update: { label: '更新日志', type: 'warning' },
  maintenance: { label: '维护通知', type: 'danger' }
}

const statusMap: Record<string, any> = {
  published: { label: '已发布', type: 'success' },
  draft: { label: '草稿', type: 'info' },
  scheduled: { label: '定时发布', type: 'warning' }
}

const form = reactive({
  title: '', type: 'system', content: '', pushMethods: ['inapp'] as string[],
  targetScope: 'all', scheduled: false, scheduledTime: null as any
})

const rules: FormRules = {
  title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
  type: [{ required: true, message: '请选择类型', trigger: 'change' }],
  content: [{ required: true, message: '请输入公告内容', trigger: 'blur' }]
}

const tableData = ref([
  { id: 1, title: 'TBox v3.0 正式发布，带来全新体验！', type: 'update', status: 'published', pushMethod: 'App+站内信', targetUsers: '全部用户', publishTime: '2024-03-15 10:00:00' },
  { id: 2, title: '春节活动：下载有礼', type: 'activity', status: 'published', pushMethod: '全渠道', targetUsers: '全部用户', publishTime: '2024-03-10 09:00:00' },
  { id: 3, title: '服务器维护通知 3月20日', type: 'maintenance', status: 'scheduled', pushMethod: '站内信', targetUsers: '全部用户', publishTime: '2024-03-20 02:00:00' },
  { id: 4, title: '新手引导优化说明', type: 'system', status: 'draft', pushMethod: '-', targetUsers: '新用户', publishTime: '-' }
])

function openDialog(item?: any) {
  editingItem.value = item || null
  if (item) {
    Object.assign(form, { title: item.title, type: item.type, content: '', pushMethods: ['inapp'], targetScope: 'all', scheduled: false, scheduledTime: null })
  } else {
    Object.assign(form, { title: '', type: 'system', content: '', pushMethods: ['inapp'], targetScope: 'all', scheduled: false, scheduledTime: null })
  }
  dialogVisible.value = true
}

function saveAsDraft() {
  ElMessage.success('已保存为草稿')
  dialogVisible.value = false
}

async function handleSubmit() {
  await formRef.value?.validate()
  ElMessage.success(form.scheduled ? '定时发布设置成功' : '发布成功')
  dialogVisible.value = false
}

function publishItem(row: any) {
  row.status = 'published'
  row.publishTime = new Date().toLocaleString()
  ElMessage.success('发布成功')
}

function handleDelete(id: number) {
  tableData.value = tableData.value.filter((item) => item.id !== id)
  ElMessage.success('删除成功')
}
</script>

<style scoped>
.action-bar { margin-bottom: 16px; }
</style>
