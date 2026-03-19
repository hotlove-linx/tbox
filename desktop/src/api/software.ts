import request from './request'

export interface Software {
  id: number
  name: string
  icon: string
  description: string
  shortDesc: string
  developer: string
  category: string
  categoryId: number
  version: string
  size: string
  rating: number
  downloadCount: number
  screenshots: string[]
  tags: string[]
  systemRequirement: string
  updatedAt: string
  createdAt: string
  license: string
  changelog: string
  price: number
}

export interface Category {
  id: number
  name: string
  icon: string
  count: number
}

export interface SoftwareListParams {
  page?: number
  pageSize?: number
  category?: string
  keyword?: string
  sort?: string
  priceType?: string
}

export interface PaginatedResult<T> {
  list: T[]
  total: number
  page: number
  pageSize: number
}

export function getSoftwareList(params: SoftwareListParams) {
  return request.get<any, { data: PaginatedResult<Software> }>('/software', { params })
}

export function getSoftwareDetail(id: number) {
  return request.get<any, { data: Software }>(`/software/${id}`)
}

export function getCategories() {
  return request.get<any, { data: Category[] }>('/categories')
}

export function getBanners() {
  return request.get<any, { data: Software[] }>('/software/banners')
}

export function getHotSoftware(params?: { limit?: number }) {
  return request.get<any, { data: Software[] }>('/software/hot', { params })
}

export function getRankingList(params?: { limit?: number }) {
  return request.get<any, { data: Software[] }>('/software/ranking', { params })
}

export function searchSoftware(keyword: string) {
  return request.get<any, { data: Software[] }>('/software/search', { params: { keyword } })
}

export function getSoftwareReviews(id: number, params?: { page?: number; pageSize?: number }) {
  return request.get(`/software/${id}/reviews`, { params })
}

export function addReview(id: number, data: { rating: number; content: string }) {
  return request.post(`/software/${id}/reviews`, data)
}
