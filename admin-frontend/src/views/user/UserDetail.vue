<template>
  <div class="page-container">
    <el-page-header @back="$router.back()" title="返回列表" content="用户详情" />

    <div class="detail-content">
      <!-- Basic Info Card -->
      <el-card shadow="never" class="section-card">
        <template #header><span class="section-title">基本信息</span></template>
        <div class="user-profile">
          <el-avatar :size="80">科</el-avatar>
          <div class="user-basic">
            <el-descriptions :column="2" border>
              <el-descriptions-item label="用户ID">10001</el-descriptions-item>
              <el-descriptions-item label="昵称">科技达人</el-descriptions-item>
              <el-descriptions-item label="邮箱">tech@example.com</el-descriptions-item>
              <el-descriptions-item label="注册时间">2024-01-15 08:30:00</el-descriptions-item>
              <el-descriptions-item label="最后登录">2024-03-15 20:30:00</el-descriptions-item>
              <el-descriptions-item label="账户状态">
                <el-tag type="success" size="small">正常</el-tag>
              </el-descriptions-item>
            </el-descriptions>
          </div>
        </div>
      </el-card>

      <!-- Space Info -->
      <el-card shadow="never" class="section-card">
        <template #header>
          <div class="card-header">
            <span class="section-title">空间信息</span>
            <el-button size="small" @click="adjustSpaceDialog = true">调整空间</el-button>
          </div>
        </template>
        <div class="space-info">
          <el-progress :percentage="68" :stroke-width="20" :text-inside="true" style="max-width: 600px;" />
          <div class="space-text">已使用 680 MB / 1 GB</div>
        </div>
      </el-card>

      <!-- Tabs -->
      <el-card shadow="never" class="section-card">
        <el-tabs v-model="activeTab">
          <el-tab-pane label="上传软件" name="software">
            <el-table :data="softwareList" stripe border>
              <el-table-column prop="name" label="软件名称" />
              <el-table-column prop="version" label="版本" width="100" />
              <el-table-column prop="downloads" label="下载量" width="100" align="right" />
              <el-table-column prop="uploadTime" label="上传时间" width="170" />
              <el-table-column label="状态" width="100" align="center">
                <template #default="{ row }">
                  <el-tag :type="row.status === 'approved' ? 'success' : 'warning'" size="small">
                    {{ row.status === 'approved' ? '已通过' : '审核中' }}
                  </el-tag>
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>
          <el-tab-pane label="评论记录" name="reviews">
            <el-table :data="reviewList" stripe border>
              <el-table-column prop="software" label="评论软件" width="150" />
              <el-table-column prop="content" label="评论内容" show-overflow-tooltip />
              <el-table-column label="评分" width="140" align="center">
                <template #default="{ row }">
                  <el-rate v-model="row.rating" disabled size="small" />
                </template>
              </el-table-column>
              <el-table-column prop="time" label="评论时间" width="170" />
            </el-table>
          </el-tab-pane>
          <el-tab-pane label="操作记录" name="logs">
            <el-table :data="logList" stripe border>
              <el-table-column prop="action" label="操作" width="150" />
              <el-table-column prop="detail" label="详情" />
              <el-table-column prop="time" label="时间" width="170" />
              <el-table-column prop="ip" label="IP地址" width="140" />
            </el-table>
          </el-tab-pane>
        </el-tabs>
      </el-card>

      <!-- Management Actions -->
      <el-card shadow="never" class="section-card">
        <template #header><span class="section-title">管理操作</span></template>
        <div class="action-buttons">
          <el-button type="warning" @click="handleDisable">禁用账户</el-button>
          <el-button type="danger" @click="handleResetPassword">重置密码</el-button>
        </div>
      </el-card>
    </div>

    <!-- Adjust Space Dialog -->
    <el-dialog v-model="adjustSpaceDialog" title="调整用户空间" width="400px">
      <el-form label-width="80px">
        <el-form-item label="调整为">
          <el-select v-model="newSpace" style="width: 100%">
            <el-option label="500 MB" value="500MB" />
            <el-option label="1 GB" value="1GB" />
            <el-option label="2 GB" value="2GB" />
            <el-option label="5 GB" value="5GB" />
            <el-option label="10 GB" value="10GB" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="adjustSpaceDialog = false">取消</el-button>
        <el-button type="primary" @click="confirmAdjust">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const activeTab = ref('software')
const adjustSpaceDialog = ref(false)
const newSpace = ref('1GB')

const softwareList = [
  { name: '超级文件管理器', version: 'v3.2.1', downloads: 52340, uploadTime: '2024-03-15 10:30:00', status: 'approved' },
  { name: '文件加密工具', version: 'v1.0.0', downloads: 8200, uploadTime: '2024-02-20 14:00:00', status: 'approved' },
  { name: '批量重命名', version: 'v2.1.0', downloads: 3500, uploadTime: '2024-01-28 09:00:00', status: 'approved' }
]

const reviewList = [
  { software: '极速浏览器', content: '非常好用，速度快！', rating: 5, time: '2024-03-14 15:00:00' },
  { software: '智能笔记', content: '功能很全面。', rating: 4, time: '2024-03-10 10:30:00' },
  { software: '天气预报', content: '准确率一般。', rating: 3, time: '2024-03-05 18:00:00' }
]

const logList = [
  { action: '上传软件', detail: '上传了超级文件管理器 v3.2.1', time: '2024-03-15 10:30:00', ip: '192.168.1.100' },
  { action: '登录', detail: '用户登录', time: '2024-03-15 08:00:00', ip: '192.168.1.100' },
  { action: '发表评论', detail: '评论了极速浏览器', time: '2024-03-14 15:00:00', ip: '192.168.1.100' },
  { action: '修改资料', detail: '修改了头像和昵称', time: '2024-03-12 11:00:00', ip: '10.0.0.1' }
]

function handleDisable() {
  ElMessageBox.confirm('确定要禁用此用户账户吗？', '警告', { type: 'warning' }).then(() => {
    ElMessage.success('账户已禁用')
  })
}

function handleResetPassword() {
  ElMessageBox.confirm('确定要重置此用户的密码吗？新密码将发送至用户邮箱。', '确认', { type: 'warning' }).then(() => {
    ElMessage.success('密码已重置')
  })
}

function confirmAdjust() {
  adjustSpaceDialog.value = false
  ElMessage.success('空间调整成功')
}
</script>

<style scoped>
.detail-content { margin-top: 20px; }
.section-card { margin-bottom: 20px; }
.section-title { font-size: 16px; font-weight: 600; }
.card-header { display: flex; justify-content: space-between; align-items: center; }
.user-profile { display: flex; gap: 20px; align-items: flex-start; }
.user-basic { flex: 1; }
.space-info { text-align: center; }
.space-text { margin-top: 8px; font-size: 14px; color: #606266; }
.action-buttons { display: flex; gap: 12px; }
</style>
