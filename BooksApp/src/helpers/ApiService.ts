// ApiService.ts

import { RequestConfig } from './config'

export interface ApiResponse<T> {
  success: boolean
  data?: T
  error?: string
}

class ApiService {
  private async request<T>(url: string, requestOptions: RequestInit): Promise<ApiResponse<T>> {
    try {
      const response = await fetch(`${RequestConfig.basePath}${url}`, requestOptions)

      if (!response.ok) {
        throw new Error('Request failed with status ' + response.status)
      }

      const data = await response.json()

      if (data.error) {
        return { success: false, error: data.message }
      } else {
        return { success: true, data }
      }
    } catch (error: any) {
      return { success: false, error: error.message }
    }
  }

  async get<T>(url: string): Promise<ApiResponse<T>> {
    const requestOptions: RequestInit = {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    }

    return this.request<T>(url, requestOptions)
  }

  async post<T>(url: string, payload: any): Promise<ApiResponse<T>> {
    const requestOptions: RequestInit = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    }

    return this.request<T>(url, requestOptions)
  }
}

export const apiService = new ApiService()
