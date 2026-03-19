import request from './request'

export interface UploadSoftwareParams {
  name: string
  version: string
  categoryId: number
  shortDesc: string
  description: string
  visibility: 'public' | 'private' | 'link'
  file: File
  icon?: File
  screenshots?: File[]
}

export interface UploadTask {
  id: number
  softwareName: string
  version: string
  fileSize: string
  progress: number
  status: 'uploading' | 'paused' | 'completed' | 'failed'
  createdAt: string
}

export function uploadSoftware(data: FormData) {
  return request.post('/upload/software', data, {
    headers: { 'Content-Type': 'multipart/form-data' }
  })
}

export function getUploadTasks() {
  return request.get<any, { data: UploadTask[] }>('/upload/tasks')
}

export function pauseUpload(id: number) {
  return request.put(`/upload/${id}/pause`)
}

export function resumeUpload(id: number) {
  return request.put(`/upload/${id}/resume`)
}

export function cancelUpload(id: number) {
  return request.delete(`/upload/${id}`)
}
