import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { LoginPresentable } from '@/Presentable/LoginPresetable'
import { login } from '@/services/AuthServices'

export const useLoginStore = defineStore('login-store', () => {
  const loginData = ref(new LoginPresentable())

  function loginSubmitHandler(payload: any) {
    return login(payload).then((data: any) => {
      console.log('data from store', data)
      return data
    })
  }

  return { loginData, loginSubmitHandler }
})
