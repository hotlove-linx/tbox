<template>
  <div class="page-container">
    <el-card shadow="never">
      <div class="top-bar">
        <el-tabs v-model="activeTab" @tab-change="handleTabChange">
          <el-tab-pane label="下载排行" name="downloads" />
          <el-tab-pane label="评分排行" name="rating" />
          <el-tab-pane label="新增排行" name="new" />
          <el-tab-pane label="增长排行" name="growth" />
        </el-tabs>
        <div class="top-actions">
          <el-date-picker v-model="dateRange" type="daterange" range-separator="~" start-placeholder="开始" end-placeholder="结束" size="default" style="width: 260px" />
          <el-button type="success" @click="handleExport"><el-icon><Download /></el-icon>导出</el-button>
        </div>
      </div>

      <el-table :data="rankingData" stripe border>
        <el-table-column label="排名" width="80" align="center">
          <template #default="{ $index }">
            <div :class="['rank-badge', { top: $index < 3 }]">{{ $index + 1 }}</div>
          </template>
        </el-table-column>
        <el-table-column label="软件信息" min-width="250">
          <template #default="{ row }">
            <div style="display: flex; align-items: center; gap: 10px;">
              <el-avatar :size="40" shape="square">{{ row.name[0] }}</el-avatar>
              <div>
                <div style="font-weight: 500;">{{ row.name }}</div>
                <div style="font-size: 12px; color: #909399;">{{ row.category }} | v{{ row.version }}</div>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="developer" label="开发者" width="120" />
        <el-table-column :label="metricLabel" width="120" align="right" sortable>
          <template #default="{ row }">
            <span class="metric-value">{{ row.metricValue }}</span>
          </template>
        </el-table-column>
        <el-table-column label="评分" width="100" align="center" v-if="activeTab !== 'rating'">
          <template #default="{ row }">{{ row.rating }}</template>
        </el-table-column>
        <el-table-column label="变化" width="120" align="center">
          <template #default="{ row }">
            <span :class="['change-text', row.changeType]">
              <el-icon v-if="row.changeType === 'up'"><Top /></el-icon>
              <el-icon v-else-if="row.changeType === 'down'"><Bottom /></el-icon>
              {{ row.change }}
            </span>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination-wrap">
        <el-pagination v-model:current-page="currentPage" :page-size="20" :total="100" layout="total, prev, pager, next" background />
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'

const activeTab = ref('downloads')
const dateRange = ref(null)
const currentPage = ref(1)

const metricLabel = computed(() => {
  const map: Record<string, string> = { downloads: '下载量', rating: '评分', new: '新增下载', growth: '增长率' }
  return map[activeTab.value] || '指标'
})

const allData: Record<string, any[]> = {
  downloads: [
    { name: '社交聊天', category: '社交通讯', version: '5.0', developer: '孙七', metricValue: '65,100', rating: '4.6', change: '+2,300', changeType: 'up' },
    { name: '超级文件管理器', category: '工具软件', version: '3.2.1', developer: '张三', metricValue: '52,340', rating: '4.5', change: '+1,800', changeType: 'up' },
    { name: '极速浏览器', category: '工具软件', version: '5.0.0', developer: '李四', metricValue: '48,210', rating: '4.3', change: '+1,200', changeType: 'up' },
    { name: '音乐播放器Pro', category: '影音播放', version: '4.5.2', developer: '赵六', metricValue: '38,920', rating: '4.2', change: '-500', changeType: 'down' },
    { name: '视频编辑器', category: '影音播放', version: '2.0.1', developer: '张三', metricValue: '31,240', rating: '4.4', change: '+800', changeType: 'up' },
    { name: '天气预报', category: '工具软件', version: '3.0.0', developer: '吴九', metricValue: '28,700', rating: '4.1', change: '+600', changeType: 'up' },
    { name: '健康计步器', category: '工具软件', version: '1.2.0', developer: '李四', metricValue: '15,600', rating: '4.0', change: '-200', changeType: 'down' },
    { name: '智能笔记', category: '效率办公', version: '2.1.0', developer: '王五', metricValue: '12,560', rating: '4.7', change: '+1,500', changeType: 'up' },
    { name: '在线学习平台', category: '教育学习', version: '2.3.1', developer: '周八', metricValue: '8,900', rating: '3.8', change: '+300', changeType: 'up' },
    { name: '密码管理器', category: '工具软件', version: '1.5.3', developer: '郑十', metricValue: '5,200', rating: '4.8', change: '+400', changeType: 'up' }
  ],
  rating: [
    { name: '密码管理器', category: '工具软件', version: '1.5.3', developer: '郑十', metricValue: '4.8', rating: '4.8', change: '+0.1', changeType: 'up' },
    { name: '智能笔记', category: '效率办公', version: '2.1.0', developer: '王五', metricValue: '4.7', rating: '4.7', change: '0', changeType: 'up' },
    { name: '社交聊天', category: '社交通讯', version: '5.0', developer: '孙七', metricValue: '4.6', rating: '4.6', change: '+0.1', changeType: 'up' },
    { name: '超级文件管理器', category: '工具软件', version: '3.2.1', developer: '张三', metricValue: '4.5', rating: '4.5', change: '0', changeType: 'up' },
    { name: '视频编辑器', category: '影音播放', version: '2.0.1', developer: '张三', metricValue: '4.4', rating: '4.4', change: '-0.1', changeType: 'down' }
  ],
  new: [
    { name: '智能笔记', category: '效率办公', version: '2.1.0', developer: '王五', metricValue: '3,200', rating: '4.7', change: '+150%', changeType: 'up' },
    { name: '社交聊天', category: '社交通讯', version: '5.0', developer: '孙七', metricValue: '2,800', rating: '4.6', change: '+80%', changeType: 'up' },
    { name: '超级文件管理器', category: '工具软件', version: '3.2.1', developer: '张三', metricValue: '2,100', rating: '4.5', change: '+45%', changeType: 'up' }
  ],
  growth: [
    { name: '智能笔记', category: '效率办公', version: '2.1.0', developer: '王五', metricValue: '+150%', rating: '4.7', change: '新上架', changeType: 'up' },
    { name: '密码管理器', category: '工具软件', version: '1.5.3', developer: '郑十', metricValue: '+120%', rating: '4.8', change: '新上架', changeType: 'up' },
    { name: '社交聊天', category: '社交通讯', version: '5.0', developer: '孙七', metricValue: '+80%', rating: '4.6', change: '版本更新', changeType: 'up' }
  ]
}

const rankingData = ref(allData.downloads)

function handleTabChange(tab: string) {
  rankingData.value = allData[tab] || []
}

function handleExport() {
  ElMessage.success('导出任务已提交')
}
</script>

<style scoped>
.top-bar { display: flex; justify-content: space-between; align-items: flex-start; margin-bottom: 16px; }
.top-actions { display: flex; gap: 12px; align-items: center; }
.rank-badge {
  width: 28px; height: 28px; border-radius: 50%; display: inline-flex; align-items: center; justify-content: center;
  font-size: 14px; font-weight: 600; background: #f0f2f5; color: #606266;
}
.rank-badge.top { background: #ff6b6b; color: #fff; }
.metric-value { font-weight: 600; color: #303133; }
.change-text { display: flex; align-items: center; gap: 2px; }
.change-text.up { color: #67c23a; }
.change-text.down { color: #f56c6c; }
.pagination-wrap { display: flex; justify-content: flex-end; margin-top: 16px; }
:deep(.el-tabs__header) { margin-bottom: 0; }
</style>
