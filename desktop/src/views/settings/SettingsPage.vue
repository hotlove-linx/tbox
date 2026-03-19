<template>
  <div class="settings-page">
    <div class="page-header">
      <h2 class="page-title">设置</h2>
    </div>

    <div class="settings-layout">
      <!-- Left: Sub navigation -->
      <div class="settings-nav">
        <div
          v-for="item in navItems"
          :key="item.key"
          class="nav-item"
          :class="{ active: activeSection === item.key }"
          @click="activeSection = item.key"
        >
          <el-icon><component :is="item.icon" /></el-icon>
          <span>{{ item.label }}</span>
        </div>
      </div>

      <!-- Right: Content -->
      <div class="settings-content">
        <!-- General Settings -->
        <div v-show="activeSection === 'general'" class="settings-section">
          <h3 class="section-title">通用设置</h3>
          <div class="setting-item">
            <div class="setting-info">
              <div class="setting-label">开机自启动</div>
              <div class="setting-desc">登录系统时自动启动 TBox</div>
            </div>
            <el-switch v-model="settingsStore.autoStart" @change="settingsStore.saveSettings()" />
          </div>
          <div class="setting-item">
            <div class="setting-info">
              <div class="setting-label">最小化到系统托盘</div>
              <div class="setting-desc">关闭窗口时最小化到系统托盘而不退出</div>
            </div>
            <el-switch v-model="settingsStore.minimizeToTray" @change="settingsStore.saveSettings()" />
          </div>
          <div class="setting-item">
            <div class="setting-info">
              <div class="setting-label">语言</div>
              <div class="setting-desc">设置应用界面语言</div>
            </div>
            <el-select v-model="settingsStore.language" style="width: 160px" @change="settingsStore.saveSettings()">
              <el-option label="简体中文" value="zh-CN" />
              <el-option label="English" value="en-US" />
            </el-select>
          </div>
          <div class="setting-item">
            <div class="setting-info">
              <div class="setting-label">主题</div>
              <div class="setting-desc">设置应用外观主题</div>
            </div>
            <el-select v-model="settingsStore.theme" style="width: 160px">
              <el-option label="浅色" value="light" />
              <el-option label="深色" value="dark" />
              <el-option label="跟随系统" value="system" />
            </el-select>
          </div>
        </div>

        <!-- Download Settings -->
        <div v-show="activeSection === 'download'" class="settings-section">
          <h3 class="section-title">下载设置</h3>
          <div class="setting-item">
            <div class="setting-info">
              <div class="setting-label">默认下载路径</div>
              <div class="setting-desc">{{ settingsStore.downloadPath }}</div>
            </div>
            <el-button @click="handleSelectPath">
              <el-icon><FolderOpened /></el-icon> 浏览
            </el-button>
          </div>
          <div class="setting-item">
            <div class="setting-info">
              <div class="setting-label">最大并发下载数</div>
              <div class="setting-desc">同时进行的最大下载任务数</div>
            </div>
            <el-select v-model="settingsStore.maxConcurrent" style="width: 120px" @change="settingsStore.saveSettings()">
              <el-option :label="1" :value="1" />
              <el-option :label="2" :value="2" />
              <el-option :label="3" :value="3" />
              <el-option :label="5" :value="5" />
            </el-select>
          </div>
          <div class="setting-item">
            <div class="setting-info">
              <div class="setting-label">下载限速</div>
              <div class="setting-desc">设置为 0 表示不限速</div>
            </div>
            <div class="speed-input">
              <el-input-number v-model="settingsStore.speedLimit" :min="0" :max="1000" :step="1" style="width: 120px" @change="settingsStore.saveSettings()" />
              <span class="unit">MB/s</span>
            </div>
          </div>
          <div class="setting-item">
            <div class="setting-info">
              <div class="setting-label">下载完成后自动安装</div>
              <div class="setting-desc">下载完成后自动运行安装程序</div>
            </div>
            <el-switch v-model="settingsStore.autoInstall" @change="settingsStore.saveSettings()" />
          </div>
        </div>

        <!-- Notification Settings -->
        <div v-show="activeSection === 'notification'" class="settings-section">
          <h3 class="section-title">通知设置</h3>
          <div class="setting-item">
            <div class="setting-info">
              <div class="setting-label">下载完成通知</div>
              <div class="setting-desc">下载任务完成时发送通知</div>
            </div>
            <el-switch v-model="settingsStore.notifyOnComplete" @change="settingsStore.saveSettings()" />
          </div>
          <div class="setting-item">
            <div class="setting-info">
              <div class="setting-label">软件更新提醒</div>
              <div class="setting-desc">已安装的软件有新版本时提醒</div>
            </div>
            <el-switch v-model="settingsStore.notifyOnUpdate" @change="settingsStore.saveSettings()" />
          </div>
          <div class="setting-item">
            <div class="setting-info">
              <div class="setting-label">通知方式</div>
              <div class="setting-desc">选择通知的展示方式</div>
            </div>
            <el-select v-model="settingsStore.notifyMethod" style="width: 160px" @change="settingsStore.saveSettings()">
              <el-option label="系统通知" value="system" />
              <el-option label="应用内通知" value="in-app" />
              <el-option label="同时使用" value="both" />
            </el-select>
          </div>
        </div>

        <!-- About -->
        <div v-show="activeSection === 'about'" class="settings-section">
          <h3 class="section-title">关于</h3>
          <div class="about-content">
            <div class="about-logo">
              <el-icon :size="48" color="#409eff"><Box /></el-icon>
              <h2>TBox</h2>
              <p class="about-version">版本 1.0.0</p>
            </div>
            <div class="setting-item">
              <div class="setting-info">
                <div class="setting-label">当前版本</div>
                <div class="setting-desc">v1.0.0 (Build 2024.03.10)</div>
              </div>
              <el-button type="primary" plain @click="checkUpdate">检查更新</el-button>
            </div>
            <el-divider />
            <div class="about-links">
              <el-link type="primary" :underline="false">
                <el-icon><Document /></el-icon> 用户协议
              </el-link>
              <el-link type="primary" :underline="false">
                <el-icon><Lock /></el-icon> 隐私政策
              </el-link>
              <el-link type="primary" :underline="false">
                <el-icon><ChatDotRound /></el-icon> 反馈与帮助
              </el-link>
            </div>
            <div class="copyright">
              <p>&copy; 2024 TBox. All rights reserved.</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useSettingsStore } from '@/stores/settings'

const settingsStore = useSettingsStore()
const activeSection = ref('general')

const navItems = [
  { key: 'general', label: '通用设置', icon: 'Setting' },
  { key: 'download', label: '下载设置', icon: 'Download' },
  { key: 'notification', label: '通知设置', icon: 'Bell' },
  { key: 'about', label: '关于', icon: 'InfoFilled' }
]

async function handleSelectPath() {
  if (window.electronAPI) {
    const path = await window.electronAPI.selectDirectory()
    if (path) {
      settingsStore.downloadPath = path
      settingsStore.saveSettings()
      ElMessage.success('下载路径已更新')
    }
  } else {
    ElMessage.info('此功能需要在 Electron 环境中使用')
  }
}

function checkUpdate() {
  ElMessage.info('正在检查更新...')
  setTimeout(() => {
    ElMessage.success('当前已是最新版本')
  }, 1500)
}
</script>

<style scoped>
.settings-page {
  max-width: 1000px;
  margin: 0 auto;
}

.page-header {
  margin-bottom: 24px;
}

.page-title {
  font-size: 22px;
  font-weight: 600;
  color: var(--text-color-primary);
}

.settings-layout {
  display: flex;
  gap: 24px;
}

.settings-nav {
  width: 200px;
  flex-shrink: 0;
  background: var(--bg-color-card);
  border-radius: var(--radius-lg);
  padding: 8px;
  border: 1px solid var(--border-color-light);
  height: fit-content;
  position: sticky;
  top: 0;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 16px;
  border-radius: var(--radius-md);
  cursor: pointer;
  font-size: 14px;
  color: var(--text-color-regular);
  transition: all 0.2s;
  margin-bottom: 2px;
}

.nav-item:hover {
  background-color: var(--bg-color);
  color: var(--text-color-primary);
}

.nav-item.active {
  background-color: #ecf5ff;
  color: var(--color-primary);
  font-weight: 500;
}

.settings-content {
  flex: 1;
  min-width: 0;
}

.settings-section {
  background: var(--bg-color-card);
  border-radius: var(--radius-lg);
  padding: 28px;
  border: 1px solid var(--border-color-light);
}

.section-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-color-primary);
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid var(--border-color-light);
}

.setting-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 0;
  border-bottom: 1px solid var(--border-color-light);
}

.setting-item:last-child {
  border-bottom: none;
}

.setting-label {
  font-size: 14px;
  font-weight: 500;
  color: var(--text-color-primary);
  margin-bottom: 4px;
}

.setting-desc {
  font-size: 12px;
  color: var(--text-color-secondary);
}

.speed-input {
  display: flex;
  align-items: center;
  gap: 8px;
}

.unit {
  font-size: 13px;
  color: var(--text-color-secondary);
}

.about-content {
  text-align: center;
}

.about-logo {
  margin-bottom: 28px;
}

.about-logo h2 {
  font-size: 24px;
  font-weight: 700;
  margin-top: 8px;
}

.about-version {
  font-size: 13px;
  color: var(--text-color-secondary);
  margin-top: 4px;
}

.about-links {
  display: flex;
  justify-content: center;
  gap: 32px;
  margin: 20px 0;
}

.about-links .el-link {
  display: flex;
  align-items: center;
  gap: 4px;
}

.copyright {
  margin-top: 24px;
  font-size: 12px;
  color: var(--text-color-placeholder);
}
</style>
