export interface ApiResponse<T> {
  success: boolean
  data?: T
  error?: ErrorDetail
}

export interface ErrorDetail {
  code: string
  message: string
  details?: Record<string, string>
}

export interface PageResponse<T> {
  content: T[]
  page: number
  size: number
  totalElements: number
  totalPages: number
  first: boolean
  last: boolean
}
