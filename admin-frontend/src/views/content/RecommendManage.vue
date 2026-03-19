<template>
  <div class="page-container">
    <!-- Hot Recommendations -->
    <el-card shadow="never" class="section-card">
      <template #header>
        <div class="card-header">
          <span class="section-title">热门推荐列表</span>
          <el-button type="primary" size="small" @click="addDialogVisible = true"><el-icon><Plus /></el-icon>添加推荐</el-button>
        </div>
      </template>
      <el-table :data="recommendList" stripe border row-key="id">
        <el-table-column width="60" align="center" label="排序">
          <template #default="{ $index }">
            <div class="sort-btns">
              <el-button link size="small" :disabled="$index === 0" @click="moveUp($index)"><el-icon><Top /></el-icon></el-button>
              <el-button link size="small" :disabled="$index === recommendList.length - 1" @click="moveDown($index)"><el-icon><Bottom /></el-icon></el-button>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="rank" label="排名" width="70" align="center">
          <template #default="{ $index }">
            <el-tag :type="$index < 3 ? 'danger' : 'info'" size="small" round>{{ $index + 1 }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="软件信息" min-width="200">
          <template #default="{ row }">
            <div style="display: flex; align-items: center; gap: 10px;">
              <el-avatar :size="36" shape="square">{{ row.name[0] }}</el-avatar>
              <div>
                <div style="font-weight: 500;">{{ row.name }}</div>
                <div style="font-size: 12px; color: #909399;">{{ row.category }}</div>
              </div>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="downloads" label="下载量" width="100" align="right" />
        <el-table-column prop="rating" label="评分" width="80" align="center" />
        <el-table-column prop="addTime" label="添加时间" width="170" />
        <el-table-column label="操作" width="100">
          <template #default="{ row }">
            <el-popconfirm title="确定移除此推荐？" @confirm="removeRecommend(row.id)">
              <template #reference>
                <el-button link type="danger" size="small">移除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- Ranking Config -->
    <el-card shadow="never" class="section-card">
      <template #header><span class="section-title">排行榜配置</span></template>
      <el-form label-width="120px" style="max-width: 500px;">
        <el-form-item label="排行周期">
          <el-select v-model="rankingConfig.period" style="width: 100%">
            <el-option label="实时" value="realtime" />
            <el-option label="日榜" value="daily" />
            <el-option label="周榜" value="weekly" />
            <el-option label="月榜" value="monthly" />
          </el-select>
        </el-form-item>
        <el-form-item label="展示数量">
          <el-input-number v-model="rankingConfig.count" :min="5" :max="100" :step="5" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="saveRankingConfig">保存配置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- Add Dialog -->
    <el-dialog v-model="addDialogVisible" title="添加推荐软件" width="500px">
      <el-form label-width="80px">
        <el-form-item label="搜索软件">
          <el-select v-model="selectedSoftware" filterable remote placeholder="输入软件名称搜索" style="width: 100%">
            <el-option v-for="sw in softwareOptions" :key="sw.id" :label="sw.name" :value="sw.id" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="addDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="addRecommend">添加</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

const addDialogVisible = ref(false)
const selectedSoftware = ref(null)

const softwareOptions = [
  { id: 101, name: '超级文件管理器' },
  { id: 102, name: '极速浏览器' },
  { id: 103, name: '智能笔记' },
  { id: 104, name: '密码管理器' },
  { id: 105, name: '视频编辑器' }
]

const recommendList = ref([
  { id: 1, name: '超级文件管理器', category: '工具软件', downloads: 52340, rating: 4.5, addTime: '2024-03-15 10:00:00' },
  { id: 2, name: '极速浏览器', category: '工具软件', downloads: 48210, rating: 4.3, addTime: '2024-03-14 09:30:00' },
  { id: 3, name: '社交聊天', category: '社交通讯', downloads: 65100, rating: 4.6, addTime: '2024-03-13 15:20:00' },
  { id: 4, name: '音乐播放器Pro', category: '影音播放', downloads: 38920, rating: 4.2, addTime: '2024-03-12 11:00:00' },
  { id: 5, name: '视频编辑器', category: '影音播放', downloads: 31240, rating: 4.4, addTime: '2024-03-11 14:00:00' }
])

const rankingConfig = reactive({ period: 'weekly', count: 20 })

function moveUp(index: number) {
  if (index > 0) {
    const arr = recommendList.value
    ;[arr[index - 1], arr[index]] = [arr[index], arr[index - 1]]
  }
}

function moveDown(index: number) {
  const arr = recommendList.value
  if (index < arr.length - 1) {
    ;[arr[index], arr[index + 1]] = [arr[index + 1], arr[index]]
  }
}

function removeRecommend(id: number) {
  recommendList.value = recommendList.value.filter((item) => item.id !== id)
  ElMessage.success('已移除')
}

function addRecommend() {
  if (!selectedSoftware.value) {
    ElMessage.warning('请选择软件')
    return
  }
  const sw = softwareOptions.find((s) => s.id === selectedSoftware.value)
  if (sw) {
    recommendList.value.push({
      id: Date.now(), name: sw.name, category: '工具软件', downloads: 0, rating: 0, addTime: new Date().toLocaleString()
    })
  }
  addDialogVisible.value = false
  selectedSoftware.value = null
  ElMessage.success('添加成功')
}

function saveRankingConfig() {
  ElMessage.success('排行榜配置已保存')
}
</script>

<style scoped>
.section-card { margin-bottom: 20px; }
.card-header { display: flex; justify-content: space-between; align-items: center; }
.section-title { font-size: 16px; font-weight: 600; }
.sort-btns { display: flex; flex-direction: column; align-items: center; }
</style>
