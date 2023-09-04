import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { login } from '@/services/AuthServices'
import { AuthLoginModel } from '@/Models/AuthLoginModel'

export const useLoginStore = defineStore('login-store', () => {
  const loginData = ref(new AuthLoginModel())

  function loginSubmitHandler(payload: any) {
    return login(payload).then((data: any) => {
      return data
    })
  }

  return { loginData, loginSubmitHandler }
})
