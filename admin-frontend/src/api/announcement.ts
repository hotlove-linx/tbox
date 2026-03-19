import request from './request'

export function getAnnouncementList(params?: any) {
  return request.get('/admin/announcements', { params })
}

export function createAnnouncement(data: any) {
  return request.post('/admin/announcements', data)
}

export function updateAnnouncement(id: number, data: any) {
  return request.put(`/admin/announcements/${id}`, data)
}

export function deleteAnnouncement(id: number) {
  return request.delete(`/admin/announcements/${id}`)
}

export function publishAnnouncement(id: number) {
  return request.put(`/admin/announcements/${id}/publish`)
}
