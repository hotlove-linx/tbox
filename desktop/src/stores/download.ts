import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface DownloadItem {
  id: number
  softwareId: number
  softwareName: string
  softwareIcon: string
  version: string
  fileSize: string
  totalBytes: number
  downloadedBytes: number
  progress: number
  speed: string
  remainingTime: string
  status: 'downloading' | 'paused' | 'completed' | 'failed'
  filePath: string
  completedAt: string
  createdAt: string
}

export const useDownloadStore = defineStore('download', () => {
  const tasks = ref<DownloadItem[]>([
    // Mock data for demonstration
    {
      id: 1, softwareId: 1, softwareName: 'Visual Studio Code', softwareIcon: '',
      version: '1.87.0', fileSize: '98.5 MB', totalBytes: 103284736, downloadedBytes: 62000000,
      progress: 60, speed: '2.5 MB/s', remainingTime: '15秒',
      status: 'downloading', filePath: '', completedAt: '', createdAt: '2024-03-10 14:30:00'
    },
    {
      id: 2, softwareId: 2, softwareName: 'Node.js', softwareIcon: '',
      version: '20.11.1', fileSize: '32.1 MB', totalBytes: 33659904, downloadedBytes: 10000000,
      progress: 30, speed: '1.8 MB/s', remainingTime: '12秒',
      status: 'downloading', filePath: '', completedAt: '', createdAt: '2024-03-10 14:32:00'
    },
    {
      id: 3, softwareId: 3, softwareName: 'Git', softwareIcon: '',
      version: '2.44.0', fileSize: '55.2 MB', totalBytes: 57881395, downloadedBytes: 57881395,
      progress: 100, speed: '', remainingTime: '',
      status: 'completed', filePath: 'C:\\Downloads\\Git-2.44.0.exe', completedAt: '2024-03-10 14:25:00', createdAt: '2024-03-10 14:20:00'
    },
    {
      id: 4, softwareId: 4, softwareName: 'Python', softwareIcon: '',
      version: '3.12.2', fileSize: '25.8 MB', totalBytes: 27053424, downloadedBytes: 27053424,
      progress: 100, speed: '', remainingTime: '',
      status: 'completed', filePath: 'C:\\Downloads\\python-3.12.2.exe', completedAt: '2024-03-10 14:15:00', createdAt: '2024-03-10 14:10:00'
    },
    {
      id: 5, softwareId: 5, softwareName: 'Docker Desktop', softwareIcon: '',
      version: '4.28.0', fileSize: '580 MB', totalBytes: 608174080, downloadedBytes: 120000000,
      progress: 20, speed: '', remainingTime: '',
      status: 'failed', filePath: '', completedAt: '', createdAt: '2024-03-10 13:50:00'
    }
  ])

  const downloadingCount = computed(() => tasks.value.filter(t => t.status === 'downloading' || t.status === 'paused').length)
  const completedCount = computed(() => tasks.value.filter(t => t.status === 'completed').length)
  const failedCount = computed(() => tasks.value.filter(t => t.status === 'failed').length)

  const downloadingTasks = computed(() => tasks.value.filter(t => t.status === 'downloading' || t.status === 'paused'))
  const completedTasks = computed(() => tasks.value.filter(t => t.status === 'completed'))
  const failedTasks = computed(() => tasks.value.filter(t => t.status === 'failed'))

  function pauseTask(id: number) {
    const task = tasks.value.find(t => t.id === id)
    if (task && task.status === 'downloading') {
      task.status = 'paused'
      task.speed = ''
      task.remainingTime = ''
    }
  }

  function resumeTask(id: number) {
    const task = tasks.value.find(t => t.id === id)
    if (task && task.status === 'paused') {
      task.status = 'downloading'
      task.speed = '2.0 MB/s'
      task.remainingTime = '计算中...'
    }
  }

  function cancelTask(id: number) {
    const idx = tasks.value.findIndex(t => t.id === id)
    if (idx !== -1) tasks.value.splice(idx, 1)
  }

  function retryTask(id: number) {
    const task = tasks.value.find(t => t.id === id)
    if (task && task.status === 'failed') {
      task.status = 'downloading'
      task.speed = '1.5 MB/s'
      task.remainingTime = '计算中...'
    }
  }

  function removeTask(id: number) {
    const idx = tasks.value.findIndex(t => t.id === id)
    if (idx !== -1) tasks.value.splice(idx, 1)
  }

  function pauseAll() {
    tasks.value.forEach(t => { if (t.status === 'downloading') { t.status = 'paused'; t.speed = ''; t.remainingTime = '' } })
  }

  function resumeAll() {
    tasks.value.forEach(t => { if (t.status === 'paused') { t.status = 'downloading'; t.speed = '2.0 MB/s'; t.remainingTime = '计算中...' } })
  }

  function clearCompleted() {
    tasks.value = tasks.value.filter(t => t.status !== 'completed')
  }

  return {
    tasks, downloadingCount, completedCount, failedCount,
    downloadingTasks, completedTasks, failedTasks,
    pauseTask, resumeTask, cancelTask, retryTask, removeTask,
    pauseAll, resumeAll, clearCompleted
  }
})
