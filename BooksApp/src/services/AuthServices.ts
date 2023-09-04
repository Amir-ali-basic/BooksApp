import { apiService } from '@/helpers/ApiService'
import type { ApiResponse } from '@/Interfaces/ApiResponse'
import type { LoginResponse } from '@/Interfaces/LoginResponse'
import type { TokenData } from '@/Interfaces/TokenData'
import type { AuthLoginModel } from '@/Models/AuthLoginModel'

export function login(payload: AuthLoginModel): Promise<TokenData> {
  return apiService.post<ApiResponse<LoginResponse>>('/users/login', payload).then((data) => {
    if (data.success) {
      // @ts-ignore
      return data.data.token
    }
  })
}
