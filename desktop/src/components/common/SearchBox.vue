<template>
  <div class="search-box">
    <el-input
      v-model="keyword"
      placeholder="搜索软件、工具、插件..."
      :prefix-icon="Search"
      clearable
      @input="handleInput"
      @keyup.enter="handleSearch"
      @focus="showSuggestions = true"
      @blur="hideSuggestions"
      class="search-input"
    />
    <transition name="fade">
      <div v-if="showSuggestions && suggestions.length > 0" class="search-suggestions">
        <div
          v-for="item in suggestions"
          :key="item.id"
          class="suggestion-item"
          @mousedown.prevent="goToSoftware(item.id)"
        >
          <el-icon><Search /></el-icon>
          <span>{{ item.name }}</span>
          <span class="suggestion-category">{{ item.category }}</span>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { Search } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const keyword = ref('')
const showSuggestions = ref(false)

interface SuggestionItem {
  id: number
  name: string
  category: string
}

const suggestions = ref<SuggestionItem[]>([])

// Mock suggestions data
const mockData: SuggestionItem[] = [
  { id: 1, name: 'Visual Studio Code', category: '开发工具' },
  { id: 2, name: 'Node.js', category: '开发工具' },
  { id: 3, name: 'Git', category: '开发工具' },
  { id: 4, name: 'Python', category: '开发工具' },
  { id: 5, name: 'Docker Desktop', category: '开发工具' },
  { id: 6, name: 'Figma', category: '设计创意' },
  { id: 7, name: 'Notion', category: '办公工具' },
  { id: 8, name: 'Slack', category: '网络通讯' },
  { id: 9, name: '微信', category: '网络通讯' },
  { id: 10, name: 'VLC Player', category: '影音娱乐' }
]

let debounceTimer: ReturnType<typeof setTimeout> | null = null

function handleInput() {
  if (debounceTimer) clearTimeout(debounceTimer)
  debounceTimer = setTimeout(() => {
    if (keyword.value.trim()) {
      suggestions.value = mockData.filter(item =>
        item.name.toLowerCase().includes(keyword.value.toLowerCase())
      ).slice(0, 6)
    } else {
      suggestions.value = []
    }
  }, 300)
}

function handleSearch() {
  if (keyword.value.trim()) {
    showSuggestions.value = false
    router.push({ path: '/category', query: { keyword: keyword.value } })
  }
}

function goToSoftware(id: number) {
  showSuggestions.value = false
  router.push(`/software/${id}`)
}

function hideSuggestions() {
  setTimeout(() => { showSuggestions.value = false }, 200)
}
</script>

<style scoped>
.search-box {
  position: relative;
  width: 100%;
}

.search-input :deep(.el-input__wrapper) {
  border-radius: var(--radius-round);
  background-color: var(--bg-color);
  box-shadow: none !important;
  padding: 4px 16px;
}

.search-input :deep(.el-input__wrapper:hover),
.search-input :deep(.el-input__wrapper.is-focus) {
  background-color: var(--bg-color-card);
  box-shadow: 0 0 0 1px var(--color-primary) !important;
}

.search-suggestions {
  position: absolute;
  top: calc(100% + 4px);
  left: 0;
  right: 0;
  background: var(--bg-color-card);
  border-radius: var(--radius-md);
  box-shadow: var(--shadow-lg);
  border: 1px solid var(--border-color-light);
  z-index: 1000;
  padding: 4px;
  max-height: 320px;
  overflow-y: auto;
}

.suggestion-item {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 12px;
  border-radius: var(--radius-sm);
  cursor: pointer;
  color: var(--text-color-regular);
  font-size: 13px;
  transition: background-color 0.15s;
}

.suggestion-item:hover {
  background-color: var(--bg-color);
}

.suggestion-item .el-icon {
  color: var(--text-color-secondary);
  font-size: 14px;
}

.suggestion-category {
  margin-left: auto;
  font-size: 12px;
  color: var(--text-color-placeholder);
}
</style>
