<template>
  <div class="banner-carousel">
    <el-carousel :interval="5000" height="280px" indicator-position="outside" arrow="hover">
      <el-carousel-item v-for="(banner, index) in banners" :key="index">
        <div class="banner-item" :style="{ background: banner.gradient }">
          <div class="banner-content">
            <div class="banner-text">
              <h2 class="banner-title">{{ banner.name }}</h2>
              <p class="banner-desc">{{ banner.description }}</p>
              <el-button type="primary" size="large" round @click.stop="handleDownload(banner)">
                <el-icon><Download /></el-icon>
                立即下载
              </el-button>
            </div>
            <div class="banner-icon">
              <el-icon :size="80" :color="banner.iconColor"><component :is="banner.iconName" /></el-icon>
            </div>
          </div>
        </div>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const banners = [
  {
    id: 1, name: 'Visual Studio Code', description: '强大的代码编辑器，支持数百种编程语言，海量扩展插件。',
    gradient: 'linear-gradient(135deg, #1e3a5f 0%, #2d6aa0 50%, #409eff 100%)',
    iconName: 'Monitor', iconColor: 'rgba(255,255,255,0.3)'
  },
  {
    id: 2, name: 'Docker Desktop', description: '容器化开发平台，简化应用部署与管理流程。',
    gradient: 'linear-gradient(135deg, #0f4c75 0%, #1b6ca8 50%, #3282b8 100%)',
    iconName: 'Ship', iconColor: 'rgba(255,255,255,0.3)'
  },
  {
    id: 3, name: 'Figma', description: '专业级在线设计工具，实时协作，打造精美界面。',
    gradient: 'linear-gradient(135deg, #5b21b6 0%, #7c3aed 50%, #a78bfa 100%)',
    iconName: 'Brush', iconColor: 'rgba(255,255,255,0.3)'
  },
  {
    id: 4, name: 'Notion', description: '全能工作空间，集笔记、任务管理、知识库于一体。',
    gradient: 'linear-gradient(135deg, #1a1a1a 0%, #333333 50%, #4a4a4a 100%)',
    iconName: 'Notebook', iconColor: 'rgba(255,255,255,0.3)'
  },
  {
    id: 5, name: 'Slack', description: '团队协作沟通利器，整合工作流提升效率。',
    gradient: 'linear-gradient(135deg, #4a154b 0%, #611f69 50%, #8b5cf6 100%)',
    iconName: 'ChatDotRound', iconColor: 'rgba(255,255,255,0.3)'
  }
]

function handleDownload(banner: typeof banners[0]) {
  router.push(`/software/${banner.id}`)
}
</script>

<style scoped>
.banner-carousel {
  margin-bottom: 32px;
}

.banner-carousel :deep(.el-carousel__container) {
  border-radius: var(--radius-lg);
  overflow: hidden;
}

.banner-item {
  width: 100%;
  height: 100%;
  border-radius: var(--radius-lg);
  overflow: hidden;
}

.banner-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  padding: 40px 60px;
}

.banner-text {
  flex: 1;
}

.banner-title {
  font-size: 32px;
  font-weight: 700;
  color: #fff;
  margin-bottom: 12px;
}

.banner-desc {
  font-size: 16px;
  color: rgba(255, 255, 255, 0.85);
  margin-bottom: 24px;
  max-width: 400px;
  line-height: 1.6;
}

.banner-icon {
  opacity: 0.6;
  margin-left: 40px;
}

.banner-carousel :deep(.el-carousel__indicators--outside) {
  margin-top: 12px;
}
</style>
