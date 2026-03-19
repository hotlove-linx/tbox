<template>
  <div class="page-container">
    <el-card shadow="never">
      <div class="filter-bar">
        <el-input v-model="searchText" placeholder="搜索用户昵称/邮箱" prefix-icon="Search" clearable style="width: 220px" />
        <el-date-picker v-model="dateRange" type="daterange" range-separator="至" start-placeholder="注册开始" end-placeholder="注册结束" style="width: 260px" />
        <el-select v-model="statusFilter" placeholder="账户状态" clearable style="width: 120px">
          <el-option label="正常" value="active" />
          <el-option label="禁用" value="disabled" />
        </el-select>
        <el-button type="primary" @click="handleSearch"><el-icon><Search /></el-icon>搜索</el-button>
        <el-button @click="resetFilters"><el-icon><Refresh /></el-icon>重置</el-button>
        <el-button type="success" @click="handleExport"><el-icon><Download /></el-icon>导出</el-button>
      </div>

      <el-table :data="tableData" stripe border>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column label="用户信息" min-width="220">
          <template #default="{ row }">
            <div style="display: flex; align-items: center; gap: 10px;">
              <el-avatar :size="36" :src="row.avatar">{{ row.nickname[0] }}</el-avatar>
              <div>
                <div style="font-weight: 500;">{{ row.nickname }}</div>
                <div style="font-size: 12px; color: #909399;">{{ row.email }}</div>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="registerTime" label="注册时间" width="170" />
        <el-table-column prop="uploadCount" label="上传软件数" width="100" align="center" />
        <el-table-column prop="totalDownloads" label="总下载量" width="100" align="right" />
        <el-table-column label="空间使用" width="160">
          <template #default="{ row }">
            <el-progress :percentage="row.spacePercent" :color="row.spacePercent > 80 ? '#f56c6c' : '#409eff'" :stroke-width="10" />
            <div style="font-size: 12px; color: #909399; text-align: center;">{{ row.spaceUsed }}/{{ row.spaceTotal }}</div>
          </template>
        </el-table-column>
        <el-table-column label="账户状态" width="90" align="center">
          <template #default="{ row }">
            <el-tag :type="row.status === 'active' ? 'success' : 'danger'" size="small">{{ row.status === 'active' ? '正常' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="lastLogin" label="最后登录" width="170" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="$router.push(`/user/detail/${row.id}`)">详情</el-button>
            <el-button link :type="row.status === 'active' ? 'warning' : 'success'" size="small" @click="toggleStatus(row)">
              {{ row.status === 'active' ? '禁用' : '启用' }}
            </el-button>
            <el-button link type="info" size="small" @click="adjustSpace(row)">调整空间</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination-wrap">
        <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :page-sizes="[10, 20, 50, 100]" :total="total" layout="total, sizes, prev, pager, next, jumper" background />
      </div>
    </el-card>

    <!-- Adjust Space Dialog -->
    <el-dialog v-model="spaceDialogVisible" title="调整用户空间" width="420px">
      <el-form label-width="80px" v-if="currentUser">
        <el-form-item label="用户">{{ currentUser.nickname }}</el-form-item>
        <el-form-item label="当前空间">{{ currentUser.spaceTotal }}</el-form-item>
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
        <el-button @click="spaceDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmAdjustSpace">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const searchText = ref('')
const dateRange = ref(null)
const statusFilter = ref('')
const currentPage = ref(1)
const pageSize = ref(20)
const total = ref(128365)
const spaceDialogVisible = ref(false)
const currentUser = ref<any>(null)
const newSpace = ref('1GB')

const tableData = ref([
  { id: 10001, nickname: '科技达人', email: 'tech@example.com', avatar: '', registerTime: '2024-01-15 08:30:00', uploadCount: 15, totalDownloads: 52340, spaceUsed: '680MB', spaceTotal: '1GB', spacePercent: 68, status: 'active', lastLogin: '2024-03-15 20:30:00' },
  { id: 10002, nickname: '文艺青年', email: 'art@example.com', avatar: '', registerTime: '2024-01-20 14:00:00', uploadCount: 8, totalDownloads: 28100, spaceUsed: '420MB', spaceTotal: '1GB', spacePercent: 42, status: 'active', lastLogin: '2024-03-15 18:00:00' },
  { id: 10003, nickname: '效率狂人', email: 'eff@example.com', avatar: '', registerTime: '2024-02-01 10:00:00', uploadCount: 22, totalDownloads: 95200, spaceUsed: '950MB', spaceTotal: '1GB', spacePercent: 95, status: 'active', lastLogin: '2024-03-15 16:45:00' },
  { id: 10004, nickname: '游戏玩家', email: 'gamer@example.com', avatar: '', registerTime: '2024-02-10 16:30:00', uploadCount: 3, totalDownloads: 5800, spaceUsed: '150MB', spaceTotal: '500MB', spacePercent: 30, status: 'active', lastLogin: '2024-03-14 22:00:00' },
  { id: 10005, nickname: '违规用户', email: 'bad@example.com', avatar: '', registerTime: '2024-02-15 09:00:00', uploadCount: 1, totalDownloads: 200, spaceUsed: '50MB', spaceTotal: '500MB', spacePercent: 10, status: 'disabled', lastLogin: '2024-03-01 10:00:00' },
  { id: 10006, nickname: '学习达人', email: 'learn@example.com', avatar: '', registerTime: '2024-02-20 11:00:00', uploadCount: 6, totalDownloads: 12500, spaceUsed: '380MB', spaceTotal: '1GB', spacePercent: 38, status: 'active', lastLogin: '2024-03-15 14:00:00' },
  { id: 10007, nickname: '设计师小王', email: 'design@example.com', avatar: '', registerTime: '2024-03-01 13:00:00', uploadCount: 10, totalDownloads: 34200, spaceUsed: '720MB', spaceTotal: '1GB', spacePercent: 72, status: 'active', lastLogin: '2024-03-15 09:30:00' },
  { id: 10008, nickname: '新手用户', email: 'newbie@example.com', avatar: '', registerTime: '2024-03-10 15:00:00', uploadCount: 0, totalDownloads: 0, spaceUsed: '0MB', spaceTotal: '500MB', spacePercent: 0, status: 'active', lastLogin: '2024-03-14 08:00:00' }
])

function handleSearch() {
  ElMessage.info('搜索功能已触发')
}

function resetFilters() {
  searchText.value = ''
  dateRange.value = null
  statusFilter.value = ''
}

function handleExport() {
  ElMessage.success('导出任务已提交')
}

function toggleStatus(row: any) {
  const action = row.status === 'active' ? '禁用' : '启用'
  ElMessageBox.confirm(`确定要${action}用户 ${row.nickname} 吗？`, '提示', { type: 'warning' }).then(() => {
    row.status = row.status === 'active' ? 'disabled' : 'active'
    ElMessage.success(`已${action}`)
  })
}

function adjustSpace(row: any) {
  currentUser.value = row
  newSpace.value = row.spaceTotal
  spaceDialogVisible.value = true
}

function confirmAdjustSpace() {
  if (currentUser.value) {
    currentUser.value.spaceTotal = newSpace.value
  }
  spaceDialogVisible.value = false
  ElMessage.success('空间调整成功')
}
</script>

<style scoped>
.filter-bar { display: flex; gap: 12px; margin-bottom: 16px; flex-wrap: wrap; }
.pagination-wrap { display: flex; justify-content: flex-end; margin-top: 16px; }
</style>
