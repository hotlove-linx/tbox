import request from './request'

export function getFeedbackList(params?: any) {
  return request.get('/admin/feedback', { params })
}

export function getFeedbackDetail(id: number) {
  return request.get(`/admin/feedback/${id}`)
}

export function replyFeedback(id: number, data: { content: string }) {
  return request.post(`/admin/feedback/${id}/reply`, data)
}

export function assignFeedback(id: number, data: { adminId: number }) {
  return request.put(`/admin/feedback/${id}/assign`, data)
}

export function closeFeedback(id: number) {
  return request.put(`/admin/feedback/${id}/close`)
}
