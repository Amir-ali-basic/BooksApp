import { apiService } from '@/helpers/ApiService'
import type { ApiResponse } from '@/Interfaces/ApiResponse'

export function login(payload: any): Promise<ApiResponse<any>> {
  return apiService.post<any>('/users/login', payload)
}
