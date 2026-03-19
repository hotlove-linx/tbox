import { defineStore } from 'pinia'
import { ref, watch } from 'vue'

export const useSettingsStore = defineStore('settings', () => {
  const autoStart = ref(false)
  const minimizeToTray = ref(true)
  const language = ref('zh-CN')
  const theme = ref<'light' | 'dark' | 'system'>('light')

  const downloadPath = ref('C:\\Users\\Downloads\\TBox')
  const maxConcurrent = ref(3)
  const speedLimit = ref(0)
  const autoInstall = ref(false)

  const notifyOnComplete = ref(true)
  const notifyOnUpdate = ref(true)
  const notifyMethod = ref('system')

  // Load settings from localStorage
  function loadSettings() {
    const saved = localStorage.getItem('tbox_settings')
    if (saved) {
      try {
        const data = JSON.parse(saved)
        autoStart.value = data.autoStart ?? false
        minimizeToTray.value = data.minimizeToTray ?? true
        language.value = data.language ?? 'zh-CN'
        theme.value = data.theme ?? 'light'
        downloadPath.value = data.downloadPath ?? 'C:\\Users\\Downloads\\TBox'
        maxConcurrent.value = data.maxConcurrent ?? 3
        speedLimit.value = data.speedLimit ?? 0
        autoInstall.value = data.autoInstall ?? false
        notifyOnComplete.value = data.notifyOnComplete ?? true
        notifyOnUpdate.value = data.notifyOnUpdate ?? true
        notifyMethod.value = data.notifyMethod ?? 'system'
      } catch (e) { /* ignore */ }
    }
    applyTheme()
  }

  function saveSettings() {
    localStorage.setItem('tbox_settings', JSON.stringify({
      autoStart: autoStart.value,
      minimizeToTray: minimizeToTray.value,
      language: language.value,
      theme: theme.value,
      downloadPath: downloadPath.value,
      maxConcurrent: maxConcurrent.value,
      speedLimit: speedLimit.value,
      autoInstall: autoInstall.value,
      notifyOnComplete: notifyOnComplete.value,
      notifyOnUpdate: notifyOnUpdate.value,
      notifyMethod: notifyMethod.value
    }))
  }

  function applyTheme() {
    let actualTheme = theme.value
    if (actualTheme === 'system') {
      actualTheme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
    }
    document.documentElement.setAttribute('data-theme', actualTheme)
  }

  watch(theme, () => {
    applyTheme()
    saveSettings()
  })

  // Initialize
  loadSettings()

  return {
    autoStart, minimizeToTray, language, theme,
    downloadPath, maxConcurrent, speedLimit, autoInstall,
    notifyOnComplete, notifyOnUpdate, notifyMethod,
    loadSettings, saveSettings, applyTheme
  }
})
