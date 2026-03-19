<template>
  <div class="detail-page">
    <!-- Back button -->
    <el-button text @click="$router.back()" class="back-btn">
      <el-icon><ArrowLeft /></el-icon> 返回
    </el-button>

    <!-- Header -->
    <div class="detail-header">
      <div class="header-icon">
        <el-icon :size="48" color="#409eff"><Box /></el-icon>
      </div>
      <div class="header-info">
        <h1 class="software-name">{{ software.name }}</h1>
        <p class="software-developer">
          <el-icon><User /></el-icon> {{ software.developer }}
        </p>
        <div class="software-tags">
          <el-tag size="small" type="info">{{ software.category }}</el-tag>
          <el-tag v-for="tag in software.tags" :key="tag" size="small" effect="plain">{{ tag }}</el-tag>
        </div>
        <div class="software-stats">
          <div class="stat-item">
            <el-rate v-model="software.rating" disabled allow-half :size="'small'" />
            <span class="stat-text">{{ software.rating }}</span>
          </div>
          <div class="stat-item">
            <el-icon><Download /></el-icon>
            <span class="stat-text">{{ formatCount(software.downloadCount) }} 次下载</span>
          </div>
          <div class="stat-item">
            <el-icon><Coin /></el-icon>
            <span class="stat-text">{{ software.size }}</span>
          </div>
          <div class="stat-item">
            <el-icon><InfoFilled /></el-icon>
            <span class="stat-text">v{{ software.version }}</span>
          </div>
        </div>
      </div>
      <div class="header-actions">
        <el-button type="primary" size="large" round @click="handleDownload">
          <el-icon><Download /></el-icon> 下载安装
        </el-button>
        <el-button size="large" round>
          <el-icon><Share /></el-icon> 分享
        </el-button>
      </div>
    </div>

    <!-- Main Content -->
    <div class="detail-body">
      <!-- Left: Main area -->
      <div class="detail-main">
        <!-- Screenshots -->
        <div class="section">
          <h3 class="section-title">软件截图</h3>
          <div class="screenshots">
            <div v-for="(_, i) in 4" :key="i" class="screenshot-item">
              <div class="screenshot-placeholder">
                <el-icon :size="32" color="#c0c4cc"><Picture /></el-icon>
                <span>截图 {{ i + 1 }}</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Description -->
        <div class="section">
          <h3 class="section-title">软件介绍</h3>
          <div class="description-content">
            <p>{{ software.description }}</p>
            <h4>主要特性</h4>
            <ul>
              <li>强大的代码编辑和智能感知功能</li>
              <li>丰富的扩展市场，支持数千个插件</li>
              <li>内置 Git 版本控制支持</li>
              <li>集成终端和调试工具</li>
              <li>跨平台运行（Windows / macOS / Linux）</li>
              <li>定期更新，持续改进</li>
            </ul>
          </div>
        </div>

        <!-- Reviews -->
        <div class="section">
          <h3 class="section-title">用户评价</h3>
          <div class="reviews-list">
            <div v-for="review in reviews" :key="review.id" class="review-item">
              <div class="review-header">
                <el-avatar :size="36">{{ review.user.charAt(0) }}</el-avatar>
                <div class="review-user-info">
                  <span class="review-user">{{ review.user }}</span>
                  <span class="review-date">{{ review.date }}</span>
                </div>
                <el-rate v-model="review.rating" disabled :size="'small'" allow-half class="review-rating" />
              </div>
              <p class="review-content">{{ review.content }}</p>
            </div>
          </div>
        </div>
      </div>

      <!-- Right: Info sidebar -->
      <div class="detail-sidebar">
        <!-- Software Info Card -->
        <div class="info-card">
          <h4 class="info-card-title">软件信息</h4>
          <div class="info-rows">
            <div class="info-row">
              <span class="info-label">版本</span>
              <span class="info-value">{{ software.version }}</span>
            </div>
            <div class="info-row">
              <span class="info-label">大小</span>
              <span class="info-value">{{ software.size }}</span>
            </div>
            <div class="info-row">
              <span class="info-label">开发者</span>
              <span class="info-value">{{ software.developer }}</span>
            </div>
            <div class="info-row">
              <span class="info-label">分类</span>
              <span class="info-value">{{ software.category }}</span>
            </div>
            <div class="info-row">
              <span class="info-label">系统要求</span>
              <span class="info-value">{{ software.systemRequirement }}</span>
            </div>
            <div class="info-row">
              <span class="info-label">更新日期</span>
              <span class="info-value">{{ software.updatedAt }}</span>
            </div>
            <div class="info-row">
              <span class="info-label">许可证</span>
              <span class="info-value">{{ software.license }}</span>
            </div>
          </div>
        </div>

        <!-- Changelog Card -->
        <div class="info-card">
          <h4 class="info-card-title">更新日志</h4>
          <div class="changelog">
            <div class="changelog-version">
              <span class="version-tag">v{{ software.version }}</span>
              <span class="version-date">{{ software.updatedAt }}</span>
            </div>
            <ul class="changelog-list">
              <li>修复了已知的稳定性问题</li>
              <li>优化了启动性能</li>
              <li>新增多项编辑器功能</li>
              <li>更新了内置终端体验</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'
import { useDownloadStore } from '@/stores/download'

const route = useRoute()
const downloadStore = useDownloadStore()

const software = reactive({
  id: Number(route.params.id) || 1,
  name: 'Visual Studio Code',
  icon: '',
  description: 'Visual Studio Code 是一款由微软开发的免费、开源的代码编辑器。它支持数百种编程语言的语法高亮、智能代码补全、代码片段、代码重构等功能。VS Code 拥有丰富的扩展市场，用户可以安装数千个扩展来增强编辑器的功能。它还内置了对 Git 的支持，以及集成终端和调试工具，是当今最受欢迎的代码编辑器之一。',
  shortDesc: '强大的代码编辑器，支持数百种语言',
  developer: 'Microsoft Corporation',
  category: '开发工具',
  version: '1.87.0',
  size: '98.5 MB',
  rating: 4.8,
  downloadCount: 1250000,
  screenshots: [],
  tags: ['免费', '跨平台', '开源'],
  systemRequirement: 'Windows 10+, macOS 10.15+, Linux',
  updatedAt: '2024-03-08',
  createdAt: '2015-04-29',
  license: 'MIT License',
  changelog: '',
  price: 0
})

const reviews = ref([
  { id: 1, user: '张三', date: '2024-03-05', rating: 5, content: '非常好用的编辑器，扩展丰富，启动速度也很快。作为日常开发的主力工具已经用了三年，越来越好用。' },
  { id: 2, user: '李四', date: '2024-02-28', rating: 4.5, content: '功能强大，界面美观，插件生态系统非常完善。唯一不足是内存占用稍微有点大。' },
  { id: 3, user: '王五', date: '2024-02-15', rating: 4, content: '轻量级但功能全面，比起传统 IDE 更加灵活。Git 集成做得很好，调试功能也够用。' },
  { id: 4, user: '赵六', date: '2024-01-20', rating: 5, content: '从 Sublime Text 迁移过来，体验非常好。Remote SSH 插件让远程开发变得超级方便。' }
])

function formatCount(count: number): string {
  if (count >= 1000000) return (count / 1000000).toFixed(1) + 'M'
  if (count >= 1000) return (count / 1000).toFixed(0) + 'K'
  return count.toString()
}

function handleDownload() {
  ElMessage.success(`开始下载 ${software.name}`)
}
</script>

<style scoped>
.detail-page {
  max-width: 1200px;
  margin: 0 auto;
}

.back-btn {
  margin-bottom: 16px;
  font-size: 14px;
  color: var(--text-color-secondary);
}

.detail-header {
  display: flex;
  align-items: flex-start;
  gap: 24px;
  background: var(--bg-color-card);
  padding: 32px;
  border-radius: var(--radius-lg);
  border: 1px solid var(--border-color-light);
  margin-bottom: 24px;
}

.header-icon {
  width: 88px;
  height: 88px;
  border-radius: 20px;
  background: linear-gradient(135deg, #ecf5ff 0%, #d9ecff 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.header-info {
  flex: 1;
}

.software-name {
  font-size: 24px;
  font-weight: 700;
  color: var(--text-color-primary);
  margin-bottom: 8px;
}

.software-developer {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 14px;
  color: var(--text-color-secondary);
  margin-bottom: 12px;
}

.software-tags {
  display: flex;
  gap: 6px;
  margin-bottom: 12px;
}

.software-stats {
  display: flex;
  gap: 24px;
  flex-wrap: wrap;
}

.stat-item {
  display: flex;
  align-items: center;
  gap: 4px;
  color: var(--text-color-secondary);
  font-size: 13px;
}

.stat-item :deep(.el-rate__icon) {
  font-size: 14px !important;
}

.header-actions {
  display: flex;
  flex-direction: column;
  gap: 10px;
  flex-shrink: 0;
}

.detail-body {
  display: flex;
  gap: 24px;
}

.detail-main {
  flex: 1;
  min-width: 0;
}

.detail-sidebar {
  width: 320px;
  flex-shrink: 0;
}

.section {
  background: var(--bg-color-card);
  border-radius: var(--radius-lg);
  padding: 24px;
  margin-bottom: 20px;
  border: 1px solid var(--border-color-light);
}

.section-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-color-primary);
  margin-bottom: 16px;
}

.screenshots {
  display: flex;
  gap: 12px;
  overflow-x: auto;
  padding-bottom: 8px;
}

.screenshot-item {
  flex-shrink: 0;
  width: 240px;
  height: 150px;
}

.screenshot-placeholder {
  width: 100%;
  height: 100%;
  background: var(--bg-color);
  border-radius: var(--radius-md);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 8px;
  color: var(--text-color-placeholder);
  font-size: 12px;
}

.description-content {
  font-size: 14px;
  color: var(--text-color-regular);
  line-height: 1.8;
}

.description-content h4 {
  font-size: 15px;
  color: var(--text-color-primary);
  margin: 16px 0 8px;
}

.description-content ul {
  padding-left: 20px;
}

.description-content li {
  margin-bottom: 6px;
}

.reviews-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.review-item {
  padding: 16px;
  background: var(--bg-color);
  border-radius: var(--radius-md);
}

.review-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 10px;
}

.review-user-info {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.review-user {
  font-weight: 500;
  font-size: 14px;
  color: var(--text-color-primary);
}

.review-date {
  font-size: 12px;
  color: var(--text-color-placeholder);
}

.review-rating :deep(.el-rate__icon) {
  font-size: 14px !important;
}

.review-content {
  font-size: 13px;
  color: var(--text-color-regular);
  line-height: 1.6;
}

.info-card {
  background: var(--bg-color-card);
  border-radius: var(--radius-lg);
  padding: 20px;
  margin-bottom: 16px;
  border: 1px solid var(--border-color-light);
}

.info-card-title {
  font-size: 15px;
  font-weight: 600;
  color: var(--text-color-primary);
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid var(--border-color-light);
}

.info-rows {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.info-row {
  display: flex;
  justify-content: space-between;
  font-size: 13px;
}

.info-label {
  color: var(--text-color-secondary);
}

.info-value {
  color: var(--text-color-primary);
  font-weight: 500;
}

.changelog-version {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 12px;
}

.version-tag {
  background: var(--color-primary);
  color: #fff;
  padding: 2px 8px;
  border-radius: var(--radius-sm);
  font-size: 12px;
  font-weight: 500;
}

.version-date {
  font-size: 12px;
  color: var(--text-color-secondary);
}

.changelog-list {
  padding-left: 16px;
  font-size: 13px;
  color: var(--text-color-regular);
  line-height: 1.8;
}
</style>
