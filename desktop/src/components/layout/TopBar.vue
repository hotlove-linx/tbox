<template>
  <div class="topbar drag-region">
    <!-- Search -->
    <div class="topbar-search no-drag">
      <SearchBox />
    </div>

    <!-- Right actions -->
    <div class="topbar-actions no-drag">
      <!-- Download shortcut -->
      <div class="action-item" @click="router.push('/downloads')" title="下载管理">
        <el-badge :value="downloadStore.downloadingCount" :hidden="downloadStore.downloadingCount === 0" :max="99">
          <el-icon :size="20"><Download /></el-icon>
        </el-badge>
      </div>

      <!-- Notifications -->
      <div class="action-item" title="通知">
        <el-badge :value="3" :max="99" is-dot>
          <el-icon :size="20"><Bell /></el-icon>
        </el-badge>
      </div>

      <!-- User info -->
      <el-dropdown trigger="click" @command="handleUserCommand">
        <div class="user-info">
          <el-avatar :size="32" :src="authStore.userAvatar || undefined">
            {{ authStore.userName?.charAt(0) || 'U' }}
          </el-avatar>
          <div class="user-text">
            <span class="user-name">{{ authStore.userName }}</span>
            <span class="user-email">{{ authStore.userEmail }}</span>
          </div>
          <el-icon class="user-arrow"><ArrowDown /></el-icon>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="profile">
              <el-icon><User /></el-icon>个人中心
            </el-dropdown-item>
            <el-dropdown-item command="space">
              <el-icon><FolderOpened /></el-icon>个人空间
            </el-dropdown-item>
            <el-dropdown-item divided command="logout">
              <el-icon><SwitchButton /></el-icon>退出登录
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>

      <!-- Window controls -->
      <div class="window-controls">
        <div class="win-btn" @click="handleMinimize" title="最小化">
          <el-icon :size="14"><Minus /></el-icon>
        </div>
        <div class="win-btn" @click="handleMaximize" title="最大化">
          <el-icon :size="14"><FullScreen /></el-icon>
        </div>
        <div class="win-btn win-btn-close" @click="handleClose" title="关闭">
          <el-icon :size="14"><Close /></el-icon>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useDownloadStore } from '@/stores/download'
import SearchBox from '@/components/common/SearchBox.vue'

const router = useRouter()
const authStore = useAuthStore()
const downloadStore = useDownloadStore()

function handleUserCommand(command: string) {
  switch (command) {
    case 'profile':
      router.push('/profile')
      break
    case 'space':
      router.push('/space')
      break
    case 'logout':
      authStore.logout()
      break
  }
}

function handleMinimize() {
  window.electronAPI?.minimize()
}

function handleMaximize() {
  window.electronAPI?.maximize()
}

function handleClose() {
  window.electronAPI?.close()
}
</script>

<style scoped>
.topbar {
  height: var(--topbar-height);
  min-height: var(--topbar-height);
  background-color: var(--bg-color-card);
  border-bottom: 1px solid var(--border-color-light);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 16px;
  gap: 16px;
}

.topbar-search {
  flex: 1;
  max-width: 480px;
}

.topbar-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.action-item {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  border-radius: var(--radius-md);
  cursor: pointer;
  color: var(--text-color-regular);
  transition: all 0.2s;
}

.action-item:hover {
  background-color: var(--bg-color);
  color: var(--color-primary);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 4px 8px;
  border-radius: var(--radius-md);
  cursor: pointer;
  transition: all 0.2s;
}

.user-info:hover {
  background-color: var(--bg-color);
}

.user-text {
  display: flex;
  flex-direction: column;
}

.user-name {
  font-size: 13px;
  font-weight: 500;
  color: var(--text-color-primary);
  line-height: 1.3;
}

.user-email {
  font-size: 11px;
  color: var(--text-color-secondary);
  line-height: 1.3;
}

.user-arrow {
  color: var(--text-color-secondary);
  font-size: 12px;
}

.window-controls {
  display: flex;
  align-items: center;
  margin-left: 8px;
  gap: 4px;
}

.win-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: var(--radius-sm);
  cursor: pointer;
  color: var(--text-color-secondary);
  transition: all 0.15s;
}

.win-btn:hover {
  background-color: var(--bg-color);
  color: var(--text-color-primary);
}

.win-btn-close:hover {
  background-color: #e81123;
  color: #fff;
}
</style>
