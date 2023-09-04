import { ref } from 'vue'
import { defineStore } from 'pinia'
import { login } from '@/services/AuthServices'
import { AuthLoginModel } from '@/Models/AuthLoginModel'
import type { ApiResponse } from '@/Interfaces/ApiResponse'
import type { LoginResponse } from '@/Interfaces/LoginResponse'

export const useAuthStore = defineStore('auth-store', () => {
  const loginData = ref(new AuthLoginModel())
  const userToken = ref<string | null>(null)

  function loginSubmitHandler(payload: AuthLoginModel) {
    return login(payload).then((data) => {
      userToken.value = data.token
      console.log('TokenData', data)
      return data
    })
  }

  return { loginData, userToken, loginSubmitHandler }
})
