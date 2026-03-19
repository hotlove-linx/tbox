<template>
  <div class="home-page">
    <!-- Banner Carousel -->
    <BannerCarousel />

    <!-- Category Tags -->
    <div class="category-tags">
      <el-tag
        v-for="cat in categories"
        :key="cat"
        :effect="activeCategory === cat ? 'dark' : 'plain'"
        :type="activeCategory === cat ? '' : 'info'"
        size="large"
        class="category-tag"
        @click="activeCategory = cat"
      >
        {{ cat }}
      </el-tag>
    </div>

    <!-- Hot Recommendations -->
    <div class="section">
      <div class="section-header">
        <h3 class="section-title">
          <el-icon><Star /></el-icon>
          热门推荐
        </h3>
        <el-link type="primary" :underline="false" @click="$router.push('/category')">
          查看全部 <el-icon><ArrowRight /></el-icon>
        </el-link>
      </div>
      <div class="software-grid">
        <SoftwareCard
          v-for="item in filteredSoftware"
          :key="item.id"
          :software="item"
        />
      </div>
    </div>

    <!-- Download Rankings -->
    <div class="section">
      <div class="section-header">
        <h3 class="section-title">
          <el-icon><TrophyBase /></el-icon>
          下载排行榜
        </h3>
        <el-link type="primary" :underline="false" @click="$router.push('/category')">
          完整榜单 <el-icon><ArrowRight /></el-icon>
        </el-link>
      </div>
      <div class="ranking-list">
        <div
          v-for="(item, index) in rankingList"
          :key="item.id"
          class="ranking-item"
          @click="$router.push(`/software/${item.id}`)"
        >
          <span class="ranking-number" :class="{ 'top-3': index < 3 }">{{ index + 1 }}</span>
          <div class="ranking-icon">
            <el-icon :size="28" color="#409eff"><Box /></el-icon>
          </div>
          <div class="ranking-info">
            <h4 class="ranking-name">{{ item.name }}</h4>
            <p class="ranking-desc text-ellipsis">{{ item.shortDesc }}</p>
          </div>
          <div class="ranking-downloads">
            <el-icon><Download /></el-icon>
            <span>{{ formatDownloads(item.downloadCount) }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import BannerCarousel from '@/components/software/BannerCarousel.vue'
import SoftwareCard from '@/components/software/SoftwareCard.vue'

const categories = ['全部', '办公工具', '开发工具', '影音娱乐', '系统工具', '安全防护', '设计创意', '网络通讯']
const activeCategory = ref('全部')

// Mock data
const softwareList = ref([
  { id: 1, name: 'Visual Studio Code', icon: '', shortDesc: '强大的代码编辑器，支持数百种语言和丰富的扩展', rating: 4.8, size: '98.5 MB', category: '开发工具', downloadCount: 1250000 },
  { id: 2, name: 'Node.js', icon: '', shortDesc: 'JavaScript 运行时，构建高性能服务端应用', rating: 4.7, size: '32.1 MB', category: '开发工具', downloadCount: 980000 },
  { id: 3, name: 'Git', icon: '', shortDesc: '分布式版本控制系统，高效管理代码协作', rating: 4.9, size: '55.2 MB', category: '开发工具', downloadCount: 1500000 },
  { id: 4, name: 'Python', icon: '', shortDesc: '简洁优雅的编程语言，适合AI和数据科学开发', rating: 4.8, size: '25.8 MB', category: '开发工具', downloadCount: 1100000 },
  { id: 5, name: 'Docker Desktop', icon: '', shortDesc: '容器化开发平台，简化应用部署与管理流程', rating: 4.6, size: '580 MB', category: '开发工具', downloadCount: 850000 },
  { id: 6, name: 'Figma', icon: '', shortDesc: '专业级协作设计工具，实时协作打造精美界面', rating: 4.7, size: '120 MB', category: '设计创意', downloadCount: 720000 },
  { id: 7, name: 'Notion', icon: '', shortDesc: '全能工作空间，集笔记、任务管理、知识库于一体', rating: 4.5, size: '85 MB', category: '办公工具', downloadCount: 650000 },
  { id: 8, name: 'Slack', icon: '', shortDesc: '团队协作沟通利器，整合工作流提升团队效率', rating: 4.4, size: '145 MB', category: '网络通讯', downloadCount: 580000 },
  { id: 9, name: 'VLC Player', icon: '', shortDesc: '开源免费的多媒体播放器，支持几乎所有格式', rating: 4.6, size: '42 MB', category: '影音娱乐', downloadCount: 920000 },
  { id: 10, name: '火绒安全', icon: '', shortDesc: '轻巧高效的安全软件，低资源占用保护系统安全', rating: 4.5, size: '18 MB', category: '安全防护', downloadCount: 480000 },
  { id: 11, name: '7-Zip', icon: '', shortDesc: '开源免费的解压缩工具，支持多种格式高压缩率', rating: 4.7, size: '1.5 MB', category: '系统工具', downloadCount: 1300000 },
  { id: 12, name: 'OBS Studio', icon: '', shortDesc: '专业级开源直播和录屏软件，功能强大灵活', rating: 4.6, size: '280 MB', category: '影音娱乐', downloadCount: 750000 }
])

const filteredSoftware = computed(() => {
  if (activeCategory.value === '全部') return softwareList.value
  return softwareList.value.filter(s => s.category === activeCategory.value)
})

const rankingList = computed(() => {
  return [...softwareList.value].sort((a, b) => b.downloadCount - a.downloadCount).slice(0, 8)
})

function formatDownloads(count: number): string {
  if (count >= 1000000) return (count / 1000000).toFixed(1) + 'M'
  if (count >= 1000) return (count / 1000).toFixed(0) + 'K'
  return count.toString()
}
</script>

<style scoped>
.home-page {
  max-width: 1200px;
  margin: 0 auto;
}

.category-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 32px;
}

.category-tag {
  cursor: pointer;
  border-radius: var(--radius-round);
  padding: 0 16px;
  transition: all 0.2s;
}

.category-tag:hover {
  opacity: 0.85;
}

.section {
  margin-bottom: 40px;
}

.section-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.section-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 20px;
  font-weight: 600;
  color: var(--text-color-primary);
}

.software-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
}

.ranking-list {
  background: var(--bg-color-card);
  border-radius: var(--radius-lg);
  border: 1px solid var(--border-color-light);
  overflow: hidden;
}

.ranking-item {
  display: flex;
  align-items: center;
  padding: 16px 20px;
  gap: 16px;
  cursor: pointer;
  transition: background-color 0.2s;
  border-bottom: 1px solid var(--border-color-light);
}

.ranking-item:last-child {
  border-bottom: none;
}

.ranking-item:hover {
  background-color: var(--bg-color);
}

.ranking-number {
  width: 28px;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  font-size: 14px;
  font-weight: 700;
  color: var(--text-color-secondary);
  background: var(--bg-color);
  flex-shrink: 0;
}

.ranking-number.top-3 {
  background: linear-gradient(135deg, #f6d365 0%, #fda085 100%);
  color: #fff;
}

.ranking-icon {
  width: 44px;
  height: 44px;
  border-radius: var(--radius-md);
  background: linear-gradient(135deg, #ecf5ff 0%, #d9ecff 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.ranking-info {
  flex: 1;
  min-width: 0;
}

.ranking-name {
  font-size: 14px;
  font-weight: 600;
  color: var(--text-color-primary);
  margin-bottom: 4px;
}

.ranking-desc {
  font-size: 12px;
  color: var(--text-color-secondary);
}

.ranking-downloads {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 13px;
  color: var(--text-color-secondary);
  flex-shrink: 0;
}
</style>
