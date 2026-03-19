<template>
  <div class="sidebar">
    <!-- Brand -->
    <div class="sidebar-brand drag-region">
      <div class="brand-logo">
        <el-icon :size="28" color="#409eff"><Box /></el-icon>
      </div>
      <div class="brand-text">
        <span class="brand-name">TBox</span>
        <span class="brand-version">v1.0</span>
      </div>
    </div>

    <!-- Navigation -->
    <nav class="sidebar-nav">
      <div class="nav-group">
        <div class="nav-group-title">发现</div>
        <router-link to="/home" class="nav-item" :class="{ active: route.path === '/home' }">
          <el-icon><HomeFilled /></el-icon>
          <span>首页</span>
        </router-link>
        <router-link to="/category" class="nav-item" :class="{ active: route.path === '/category' }">
          <el-icon><Grid /></el-icon>
          <span>分类浏览</span>
        </router-link>
      </div>

      <div class="nav-group">
        <div class="nav-group-title">管理</div>
        <router-link to="/downloads" class="nav-item" :class="{ active: route.path === '/downloads' }">
          <el-icon><Download /></el-icon>
          <span>下载管理</span>
          <el-badge v-if="downloadStore.downloadingCount > 0" :value="downloadStore.downloadingCount" class="nav-badge" />
        </router-link>
        <router-link to="/install-records" class="nav-item" :class="{ active: route.path === '/install-records' }">
          <el-icon><FolderChecked /></el-icon>
          <span>安装记录</span>
        </router-link>
      </div>

      <div class="nav-group">
        <div class="nav-group-title">空间</div>
        <router-link to="/space" class="nav-item" :class="{ active: route.path === '/space' }">
          <el-icon><User /></el-icon>
          <span>个人空间</span>
        </router-link>
      </div>

      <div class="nav-group">
        <div class="nav-group-title">系统</div>
        <router-link to="/settings" class="nav-item" :class="{ active: route.path === '/settings' }">
          <el-icon><Setting /></el-icon>
          <span>设置</span>
        </router-link>
      </div>
    </nav>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { useDownloadStore } from '@/stores/download'

const route = useRoute()
const downloadStore = useDownloadStore()
</script>

<style scoped>
.sidebar {
  width: var(--sidebar-width);
  min-width: var(--sidebar-width);
  height: 100vh;
  background-color: var(--bg-color-sidebar);
  display: flex;
  flex-direction: column;
  user-select: none;
  overflow-y: auto;
}

.sidebar-brand {
  display: flex;
  align-items: center;
  padding: 20px 20px 16px;
  gap: 12px;
}

.brand-logo {
  width: 40px;
  height: 40px;
  border-radius: 10px;
  background: linear-gradient(135deg, #409eff 0%, #66b1ff 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}

.brand-logo .el-icon {
  color: #fff !important;
}

.brand-text {
  display: flex;
  flex-direction: column;
}

.brand-name {
  font-size: 20px;
  font-weight: 700;
  color: #fff;
  letter-spacing: 1px;
}

.brand-version {
  font-size: 11px;
  color: var(--text-color-sidebar);
  margin-top: 2px;
}

.sidebar-nav {
  flex: 1;
  padding: 8px 12px;
}

.nav-group {
  margin-bottom: 8px;
}

.nav-group-title {
  font-size: 11px;
  color: #5e5e72;
  text-transform: uppercase;
  letter-spacing: 1px;
  padding: 12px 12px 6px;
  font-weight: 600;
}

.nav-item {
  display: flex;
  align-items: center;
  padding: 10px 12px;
  border-radius: var(--radius-md);
  color: var(--text-color-sidebar);
  text-decoration: none;
  font-size: 14px;
  transition: all 0.2s ease;
  gap: 10px;
  margin-bottom: 2px;
  position: relative;
}

.nav-item:hover {
  background-color: var(--bg-color-sidebar-hover);
  color: #fff;
}

.nav-item.active {
  background-color: var(--bg-color-sidebar-active);
  color: var(--text-color-sidebar-active);
  font-weight: 500;
}

.nav-item .el-icon {
  font-size: 18px;
}

.nav-badge {
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
}

.nav-badge :deep(.el-badge__content) {
  font-size: 10px;
}
</style>
