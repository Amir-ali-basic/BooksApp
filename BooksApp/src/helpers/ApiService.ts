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

      if (!response.ok) {
        const errorData = await response.json()
        console.error('Request error:', errorData)
        return { success: false, error: errorData.message, statusCode: response.status }
      }

      const responseData = await response.json()
      const processedData = responseData.data || responseData

      return { success: true, data: processedData }
    } catch (error: any) {
      console.error('Request error:', error.message)
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

  async put<T>(
    url: string,
    payload: any,
    headers: Record<string, string> = {}
  ): Promise<ApiResponse<T>> {
    const requestOptions: RequestInit = {
      method: 'PUT',
      headers: {
        ...headers,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    }

    return this.request<T>(url, requestOptions)
  }

  async delete<T>(url: string, headers: Record<string, string> = {}): Promise<ApiResponse<T>> {
    const requestOptions: RequestInit = {
      method: 'DELETE',
      headers: {
        ...headers,
        'Content-Type': 'application/json'
      }
    }

    return this.request<T>(url, requestOptions)
  }

  async patch<T>(
    url: string,
    payload: any,
    headers: Record<string, string> = {}
  ): Promise<ApiResponse<T>> {
    const requestOptions: RequestInit = {
      method: 'PATCH',
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
