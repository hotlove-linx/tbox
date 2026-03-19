<template>
  <div class="category-page">
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

    <!-- Search & Filter Bar -->
    <div class="filter-bar">
      <el-input
        v-model="searchKeyword"
        placeholder="搜索软件..."
        :prefix-icon="Search"
        clearable
        class="filter-search"
      />
      <div class="filter-actions">
        <el-select v-model="sortBy" placeholder="排序方式" style="width: 140px">
          <el-option label="热度优先" value="hot" />
          <el-option label="评分最高" value="rating" />
          <el-option label="最近更新" value="updated" />
          <el-option label="文件大小" value="size" />
        </el-select>
        <el-select v-model="priceFilter" placeholder="价格筛选" style="width: 120px">
          <el-option label="全部" value="all" />
          <el-option label="免费" value="free" />
          <el-option label="付费" value="paid" />
        </el-select>
      </div>
    </div>

    <!-- Software Grid -->
    <div class="software-grid">
      <SoftwareCard
        v-for="item in filteredList"
        :key="item.id"
        :software="item"
      />
    </div>

    <!-- Empty state -->
    <el-empty v-if="filteredList.length === 0" description="未找到匹配的软件" />

    <!-- Pagination -->
    <div class="pagination-wrapper" v-if="filteredList.length > 0">
      <el-pagination
        v-model:current-page="currentPage"
        :page-size="pageSize"
        :total="totalItems"
        layout="prev, pager, next, jumper, total"
        background
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { Search } from '@element-plus/icons-vue'
import SoftwareCard from '@/components/software/SoftwareCard.vue'

const route = useRoute()

const categories = ['全部', '办公工具', '开发工具', '影音娱乐', '系统工具', '安全防护', '设计创意', '网络通讯']
const activeCategory = ref('全部')
const searchKeyword = ref('')
const sortBy = ref('hot')
const priceFilter = ref('all')
const currentPage = ref(1)
const pageSize = 12

// Mock data
const softwareList = ref([
  { id: 1, name: 'Visual Studio Code', icon: '', shortDesc: '强大的代码编辑器，支持数百种语言和扩展', rating: 4.8, size: '98.5 MB', category: '开发工具', downloadCount: 1250000, price: 0 },
  { id: 2, name: 'Node.js', icon: '', shortDesc: 'JavaScript 运行时，构建高性能服务端应用', rating: 4.7, size: '32.1 MB', category: '开发工具', downloadCount: 980000, price: 0 },
  { id: 3, name: 'Git', icon: '', shortDesc: '分布式版本控制系统', rating: 4.9, size: '55.2 MB', category: '开发工具', downloadCount: 1500000, price: 0 },
  { id: 4, name: 'Python', icon: '', shortDesc: '简洁优雅的编程语言', rating: 4.8, size: '25.8 MB', category: '开发工具', downloadCount: 1100000, price: 0 },
  { id: 5, name: 'Docker Desktop', icon: '', shortDesc: '容器化开发平台', rating: 4.6, size: '580 MB', category: '开发工具', downloadCount: 850000, price: 0 },
  { id: 6, name: 'Figma', icon: '', shortDesc: '专业级协作设计工具', rating: 4.7, size: '120 MB', category: '设计创意', downloadCount: 720000, price: 0 },
  { id: 7, name: 'Notion', icon: '', shortDesc: '全能工作空间', rating: 4.5, size: '85 MB', category: '办公工具', downloadCount: 650000, price: 0 },
  { id: 8, name: 'Slack', icon: '', shortDesc: '团队协作沟通利器', rating: 4.4, size: '145 MB', category: '网络通讯', downloadCount: 580000, price: 0 },
  { id: 9, name: 'VLC Player', icon: '', shortDesc: '开源多媒体播放器', rating: 4.6, size: '42 MB', category: '影音娱乐', downloadCount: 920000, price: 0 },
  { id: 10, name: '火绒安全', icon: '', shortDesc: '轻巧高效的安全软件', rating: 4.5, size: '18 MB', category: '安全防护', downloadCount: 480000, price: 0 },
  { id: 11, name: '7-Zip', icon: '', shortDesc: '开源免费的解压缩工具', rating: 4.7, size: '1.5 MB', category: '系统工具', downloadCount: 1300000, price: 0 },
  { id: 12, name: 'OBS Studio', icon: '', shortDesc: '专业级开源直播和录屏软件', rating: 4.6, size: '280 MB', category: '影音娱乐', downloadCount: 750000, price: 0 },
  { id: 13, name: 'Postman', icon: '', shortDesc: 'API 开发和测试工具', rating: 4.5, size: '200 MB', category: '开发工具', downloadCount: 690000, price: 0 },
  { id: 14, name: 'Typora', icon: '', shortDesc: '优雅的 Markdown 编辑器', rating: 4.6, size: '68 MB', category: '办公工具', downloadCount: 520000, price: 89 },
  { id: 15, name: 'Adobe Photoshop', icon: '', shortDesc: '专业图像处理软件', rating: 4.8, size: '2.1 GB', category: '设计创意', downloadCount: 1450000, price: 298 },
  { id: 16, name: 'WinRAR', icon: '', shortDesc: '经典解压缩工具', rating: 4.3, size: '3.5 MB', category: '系统工具', downloadCount: 1800000, price: 29 },
  { id: 17, name: 'PotPlayer', icon: '', shortDesc: '功能强大的视频播放器', rating: 4.7, size: '36 MB', category: '影音娱乐', downloadCount: 880000, price: 0 },
  { id: 18, name: '微信', icon: '', shortDesc: '国民级社交通讯应用', rating: 4.2, size: '165 MB', category: '网络通讯', downloadCount: 2500000, price: 0 }
])

const filteredList = computed(() => {
  let list = [...softwareList.value]

  if (activeCategory.value !== '全部') {
    list = list.filter(s => s.category === activeCategory.value)
  }

  if (searchKeyword.value.trim()) {
    const kw = searchKeyword.value.toLowerCase()
    list = list.filter(s => s.name.toLowerCase().includes(kw) || s.shortDesc.toLowerCase().includes(kw))
  }

  if (priceFilter.value === 'free') {
    list = list.filter(s => s.price === 0)
  } else if (priceFilter.value === 'paid') {
    list = list.filter(s => s.price > 0)
  }

  switch (sortBy.value) {
    case 'hot': list.sort((a, b) => b.downloadCount - a.downloadCount); break
    case 'rating': list.sort((a, b) => b.rating - a.rating); break
    case 'updated': break
    case 'size': break
  }

  return list
})

const totalItems = computed(() => filteredList.value.length)

onMounted(() => {
  if (route.query.keyword) {
    searchKeyword.value = route.query.keyword as string
  }
})
</script>

<style scoped>
.category-page {
  max-width: 1200px;
  margin: 0 auto;
}

.category-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 20px;
}

.category-tag {
  cursor: pointer;
  border-radius: var(--radius-round);
  padding: 0 16px;
  transition: all 0.2s;
}

.filter-bar {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 24px;
  background: var(--bg-color-card);
  padding: 16px;
  border-radius: var(--radius-lg);
  border: 1px solid var(--border-color-light);
}

.filter-search {
  width: 280px;
}

.filter-actions {
  display: flex;
  gap: 12px;
  margin-left: auto;
}

.software-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
  margin-bottom: 32px;
}

.pagination-wrapper {
  display: flex;
  justify-content: center;
  padding: 20px 0;
}
</style>
