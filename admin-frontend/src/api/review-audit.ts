import request from './request'

export function getReviewList(params?: any) {
  return request.get('/admin/audit/reviews', { params })
}

export function blockReview(id: number) {
  return request.put(`/admin/audit/reviews/${id}/block`)
}

export function restoreReview(id: number) {
  return request.put(`/admin/audit/reviews/${id}/restore`)
}

export function deleteReview(id: number) {
  return request.delete(`/admin/audit/reviews/${id}`)
}
