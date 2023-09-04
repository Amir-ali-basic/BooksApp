import type { ApiResponse } from '@/Interfaces/ApiResponse'
import { RequestConfig } from './config'

class ApiService {
  private async request<T>(url: string, requestOptions: RequestInit): Promise<ApiResponse<T>> {
    try {
      const response = await Promise.race([
        fetch(`${RequestConfig.basePath}${url}`, requestOptions),
        new Promise<Response>(
          (_, reject) => setTimeout(() => reject(new Error('Request timeout')), 10000) // Add a timeout
        )
      ])

      const data = await response.json()

      if (!response.ok) {
        console.log('reponse', response)
        console.log('reponse.status', response.status)

        return { success: false, error: data.message, statusCode: response.status }
      } else {
        return { success: true, data }
      }
    } catch (error: any) {
      return { success: false, error: error.message }
    }
  }

  async get<T>(url: string, headers: Record<string, string> = {}): Promise<ApiResponse<T>> {
    const requestOptions: RequestInit = {
      method: 'GET',
      headers: {
        ...headers,
        'Content-Type': 'application/json'
      }
    }

    return this.request<T>(url, requestOptions)
  }

  async post<T>(
    url: string,
    payload: any,
    headers: Record<string, string> = {}
  ): Promise<ApiResponse<T>> {
    const requestOptions: RequestInit = {
      method: 'POST',
      headers: {
        ...headers,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    }

    return this.request<T>(url, requestOptions)
  }
}

export const apiService = new ApiService()
