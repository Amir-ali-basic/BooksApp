import { apiService } from '@/helpers/ApiService'

interface ApiResponse<T> {
  success: boolean
  data?: T
  error?: string
}

export function login(payload: any): Promise<ApiResponse<any>> {
  return apiService.post<any>('/users/login', payload)
}
