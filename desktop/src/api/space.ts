import request from './request'

export interface SpaceStats {
  softwareCount: number
  publishedCount: number
  reviewingCount: number
  privateCount: number
  totalDownloads: number
  weeklyDownloads: number
  usedSpace: string
  totalSpace: string
  usedPercent: number
}

export interface MySoftware {
  id: number
  name: string
  icon: string
  version: string
  size: string
  visibility: 'public' | 'private' | 'link'
  status: 'published' | 'reviewing' | 'rejected' | 'draft'
  downloadCount: number
  createdAt: string
}

export interface InstallRecord {
  id: number
  softwareId: number
  softwareName: string
  softwareIcon: string
  developer: string
  version: string
  size: string
  installPath: string
  installedAt: string
  status: 'installed' | 'updatable' | 'uninstalled'
  latestVersion?: string
}

export function getSpaceStats() {
  return request.get<any, { data: SpaceStats }>('/space/stats')
}

export function getMySoftware(params?: { page?: number; pageSize?: number }) {
  return request.get<any, { data: { list: MySoftware[]; total: number } }>('/space/software', { params })
}

export function deleteMySoftware(id: number) {
  return request.delete(`/space/software/${id}`)
}

export function getInstallRecords(params?: { keyword?: string; sort?: string; status?: string; page?: number; pageSize?: number }) {
  return request.get<any, { data: { list: InstallRecord[]; total: number } }>('/install-records', { params })
}

export function uninstallSoftware(id: number) {
  return request.delete(`/install-records/${id}`)
}

export function updateSoftware(id: number) {
  return request.put(`/install-records/${id}/update`)
}
