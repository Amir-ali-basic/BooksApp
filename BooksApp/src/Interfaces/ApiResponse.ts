export interface ApiResponse<T> {
  error?: boolean
  message?: string
  data?: T
  success: boolean
  statusCode?: number
}
