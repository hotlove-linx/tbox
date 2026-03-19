<template>
  <div class="software-card" @click="$router.push(`/software/${software.id}`)">
    <div class="card-icon">
      <el-icon :size="40" color="#409eff"><Box /></el-icon>
    </div>
    <div class="card-info">
      <h4 class="card-name text-ellipsis">{{ software.name }}</h4>
      <p class="card-desc text-ellipsis-2">{{ software.shortDesc }}</p>
      <div class="card-meta">
        <div class="card-rating">
          <el-rate v-model="ratingValue" disabled :size="'small'" allow-half />
          <span class="rating-text">{{ software.rating }}</span>
        </div>
        <span class="card-size">{{ software.size }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface SoftwareItem {
  id: number
  name: string
  icon: string
  shortDesc: string
  rating: number
  size: string
}

const props = defineProps<{
  software: SoftwareItem
}>()

const ratingValue = computed(() => props.software.rating)
</script>

<style scoped>
.software-card {
  background: var(--bg-color-card);
  border-radius: var(--radius-lg);
  padding: 20px;
  cursor: pointer;
  transition: all 0.25s ease;
  border: 1px solid var(--border-color-light);
}

.software-card:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
  border-color: var(--color-primary);
}

.card-icon {
  width: 56px;
  height: 56px;
  border-radius: var(--radius-md);
  background: linear-gradient(135deg, #ecf5ff 0%, #d9ecff 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 12px;
}

.card-name {
  font-size: 15px;
  font-weight: 600;
  color: var(--text-color-primary);
  margin-bottom: 6px;
}

.card-desc {
  font-size: 12px;
  color: var(--text-color-secondary);
  line-height: 1.5;
  margin-bottom: 12px;
  min-height: 36px;
}

.card-meta {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.card-rating {
  display: flex;
  align-items: center;
  gap: 4px;
}

.card-rating :deep(.el-rate) {
  height: auto;
}

.card-rating :deep(.el-rate__icon) {
  font-size: 12px !important;
  margin-right: 1px;
}

.rating-text {
  font-size: 12px;
  color: var(--text-color-secondary);
  font-weight: 500;
}

.card-size {
  font-size: 12px;
  color: var(--text-color-placeholder);
}
</style>
