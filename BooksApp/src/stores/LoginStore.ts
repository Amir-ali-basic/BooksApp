import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { LoginPresentable } from '@/Presentable/LoginPresetable'

export const useLoginStore = defineStore('login-store', () => {
  const loginData = ref(new LoginPresentable())

  return { loginData }
})
