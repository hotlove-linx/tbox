import request from './request'

export interface DownloadTask {
  id: number
  softwareId: number
  softwareName: string
  softwareIcon: string
  version: string
  fileSize: string
  progress: number
  speed: string
  remainingTime: string
  status: 'downloading' | 'paused' | 'completed' | 'failed'
  downloadedSize: string
  filePath: string
  completedAt: string
  createdAt: string
}

export function getDownloadTasks(params?: { status?: string }) {
  return request.get<any, { data: DownloadTask[] }>('/downloads', { params })
}

export function createDownload(softwareId: number) {
  return request.post('/downloads', { softwareId })
}

export function pauseDownload(id: number) {
  return request.put(`/downloads/${id}/pause`)
}

export function resumeDownload(id: number) {
  return request.put(`/downloads/${id}/resume`)
}

export function cancelDownload(id: number) {
  return request.delete(`/downloads/${id}`)
}

export function retryDownload(id: number) {
  return request.put(`/downloads/${id}/retry`)
}

export function clearCompleted() {
  return request.delete('/downloads/completed')
}
