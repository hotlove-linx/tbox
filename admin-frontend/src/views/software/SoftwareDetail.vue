<template>
  <div class="page-container">
    <el-page-header @back="$router.back()" title="返回列表" :content="isEdit ? '编辑软件' : '软件详情'" />

    <div class="detail-content">
      <!-- Basic Info -->
      <el-card shadow="never" class="section-card">
        <template #header>
          <span class="section-title">基本信息</span>
          <el-button v-if="!isEdit" type="primary" size="small" @click="isEdit = true">编辑</el-button>
        </template>
        <el-form :model="form" :rules="rules" ref="formRef" label-width="100px" :disabled="!isEdit">
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="软件名称" prop="name">
                <el-input v-model="form.name" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="版本号" prop="version">
                <el-input v-model="form.version" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="分类" prop="category">
                <el-select v-model="form.category" style="width: 100%">
                  <el-option label="工具软件" value="tools" />
                  <el-option label="游戏娱乐" value="games" />
                  <el-option label="效率办公" value="productivity" />
                  <el-option label="社交通讯" value="social" />
                  <el-option label="影音播放" value="media" />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="开发者" prop="developer">
                <el-input v-model="form.developer" />
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="软件描述" prop="description">
                <el-input v-model="form.description" type="textarea" :rows="4" />
              </el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="标签">
                <el-tag v-for="tag in form.tags" :key="tag" closable :disable-transitions="false" @close="removeTag(tag)" style="margin-right: 8px">{{ tag }}</el-tag>
                <el-input v-if="tagInputVisible" ref="tagInputRef" v-model="tagInputValue" size="small" style="width: 100px" @keyup.enter="addTag" @blur="addTag" />
                <el-button v-else size="small" @click="tagInputVisible = true">+ 新标签</el-button>
              </el-form-item>
            </el-col>
          </el-row>
          <div v-if="isEdit" class="form-actions">
            <el-button type="primary" @click="handleSave">保存</el-button>
            <el-button @click="isEdit = false">取消</el-button>
          </div>
        </el-form>
      </el-card>

      <!-- File Info -->
      <el-card shadow="never" class="section-card">
        <template #header><span class="section-title">文件信息</span></template>
        <el-descriptions :column="3" border>
          <el-descriptions-item label="文件名">super-file-manager-3.2.1.apk</el-descriptions-item>
          <el-descriptions-item label="文件大小">28.5 MB</el-descriptions-item>
          <el-descriptions-item label="MD5">a1b2c3d4e5f6g7h8i9j0</el-descriptions-item>
          <el-descriptions-item label="包名">com.example.filemanager</el-descriptions-item>
          <el-descriptions-item label="最低SDK">Android 8.0</el-descriptions-item>
          <el-descriptions-item label="上传时间">2024-03-15 10:30:00</el-descriptions-item>
        </el-descriptions>
      </el-card>

      <!-- Screenshots -->
      <el-card shadow="never" class="section-card">
        <template #header>
          <span class="section-title">展示资源</span>
        </template>
        <div class="screenshots-section">
          <div class="label-row">
            <span>应用图标</span>
          </div>
          <el-avatar :size="80" shape="square">FM</el-avatar>

          <div class="label-row" style="margin-top: 20px">
            <span>截图 ({{ screenshots.length }})</span>
          </div>
          <div class="screenshot-list">
            <div v-for="(img, idx) in screenshots" :key="idx" class="screenshot-item">
              <div class="screenshot-placeholder">截图 {{ idx + 1 }}</div>
              <el-button link type="danger" size="small" @click="screenshots.splice(idx, 1)">删除</el-button>
            </div>
            <div class="screenshot-add" @click="screenshots.push(`screenshot-${Date.now()}`)">
              <el-icon :size="24"><Plus /></el-icon>
              <span>添加截图</span>
            </div>
          </div>
        </div>
      </el-card>

      <!-- Operation Attributes -->
      <el-card shadow="never" class="section-card">
        <template #header><span class="section-title">运营属性</span></template>
        <el-form label-width="100px">
          <el-form-item label="推荐权重">
            <el-input-number v-model="opAttrs.weight" :min="0" :max="100" />
          </el-form-item>
          <el-form-item label="首页推荐">
            <el-switch v-model="opAttrs.recommended" />
          </el-form-item>
          <el-form-item label="精品标记">
            <el-switch v-model="opAttrs.featured" />
          </el-form-item>
        </el-form>
      </el-card>

      <!-- Audit Info -->
      <el-card shadow="never" class="section-card">
        <template #header><span class="section-title">审核信息</span></template>
        <el-descriptions :column="2" border>
          <el-descriptions-item label="审核状态">
            <el-tag type="success">已通过</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="审核人">管理员A</el-descriptions-item>
          <el-descriptions-item label="审核时间">2024-03-15 14:00:00</el-descriptions-item>
          <el-descriptions-item label="审核备注">内容合规，允许上架</el-descriptions-item>
        </el-descriptions>
        <div style="margin-top: 16px">
          <div class="section-subtitle">审核记录</div>
          <el-timeline>
            <el-timeline-item timestamp="2024-03-15 14:00:00" placement="top" type="success">
              <span>管理员A 通过审核：内容合规，允许上架</span>
            </el-timeline-item>
            <el-timeline-item timestamp="2024-03-15 10:30:00" placement="top" type="primary">
              <span>开发者 张三 提交审核</span>
            </el-timeline-item>
          </el-timeline>
        </div>
      </el-card>

      <!-- Version History -->
      <el-card shadow="never" class="section-card">
        <template #header><span class="section-title">版本历史</span></template>
        <el-table :data="versionHistory" stripe border>
          <el-table-column prop="version" label="版本号" width="100" />
          <el-table-column prop="updateLog" label="更新日志" />
          <el-table-column prop="fileSize" label="文件大小" width="100" />
          <el-table-column prop="uploadTime" label="上传时间" width="170" />
          <el-table-column label="状态" width="100">
            <template #default="{ row }">
              <el-tag :type="row.current ? 'success' : 'info'" size="small">{{ row.current ? '当前版本' : '历史版本' }}</el-tag>
            </template>
          </el-table-column>
        </el-table>
      </el-card>

      <!-- Statistics -->
      <el-card shadow="never" class="section-card">
        <template #header><span class="section-title">统计数据</span></template>
        <el-row :gutter="20">
          <el-col :span="6">
            <el-statistic title="总下载量" :value="52340" />
          </el-col>
          <el-col :span="6">
            <el-statistic title="今日下载" :value="326" />
          </el-col>
          <el-col :span="6">
            <el-statistic title="评价数" :value="1280" />
          </el-col>
          <el-col :span="6">
            <el-statistic title="平均评分" :value="4.5" :precision="1" />
          </el-col>
        </el-row>
      </el-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'

const route = useRoute()
const isEdit = ref(false)
const formRef = ref<FormInstance>()
const tagInputVisible = ref(false)
const tagInputValue = ref('')
const tagInputRef = ref()

const form = reactive({
  name: '超级文件管理器',
  version: '3.2.1',
  category: 'tools',
  developer: '张三',
  description: '一款功能强大的文件管理工具，支持多种文件格式预览、压缩解压、云存储同步等功能。界面简洁直观，操作流畅，是手机上必备的文件管理利器。',
  tags: ['文件管理', '工具', '效率']
})

const rules: FormRules = {
  name: [{ required: true, message: '请输入软件名称', trigger: 'blur' }],
  version: [{ required: true, message: '请输入版本号', trigger: 'blur' }],
  category: [{ required: true, message: '请选择分类', trigger: 'change' }],
  developer: [{ required: true, message: '请输入开发者', trigger: 'blur' }],
  description: [{ required: true, message: '请输入软件描述', trigger: 'blur' }]
}

const screenshots = ref(['s1', 's2', 's3', 's4'])

const opAttrs = reactive({
  weight: 50,
  recommended: true,
  featured: false
})

const versionHistory = [
  { version: 'v3.2.1', updateLog: '修复已知问题，优化性能', fileSize: '28.5MB', uploadTime: '2024-03-15 10:30:00', current: true },
  { version: 'v3.2.0', updateLog: '新增云存储同步功能', fileSize: '27.8MB', uploadTime: '2024-02-20 09:00:00', current: false },
  { version: 'v3.1.0', updateLog: '支持RAR格式解压', fileSize: '26.2MB', uploadTime: '2024-01-10 14:30:00', current: false }
]

function removeTag(tag: string) {
  form.tags = form.tags.filter((t) => t !== tag)
}

function addTag() {
  if (tagInputValue.value) {
    form.tags.push(tagInputValue.value)
  }
  tagInputVisible.value = false
  tagInputValue.value = ''
}

function handleSave() {
  ElMessage.success('保存成功')
  isEdit.value = false
}
</script>

<style scoped>
.detail-content {
  margin-top: 20px;
}

.section-card {
  margin-bottom: 20px;
}

.section-card :deep(.el-card__header) {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.section-title {
  font-size: 16px;
  font-weight: 600;
}

.section-subtitle {
  font-size: 14px;
  font-weight: 500;
  color: #606266;
  margin-bottom: 12px;
}

.form-actions {
  text-align: center;
  margin-top: 20px;
}

.label-row {
  font-size: 14px;
  color: #606266;
  margin-bottom: 10px;
}

.screenshot-list {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
}

.screenshot-item {
  text-align: center;
}

.screenshot-placeholder {
  width: 120px;
  height: 200px;
  background: #f5f7fa;
  border: 1px solid #e4e7ed;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #909399;
  font-size: 13px;
  margin-bottom: 4px;
}

.screenshot-add {
  width: 120px;
  height: 200px;
  border: 2px dashed #dcdfe6;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 8px;
  color: #909399;
  cursor: pointer;
  transition: border-color 0.2s;
}

.screenshot-add:hover {
  border-color: #409eff;
  color: #409eff;
}
</style>
