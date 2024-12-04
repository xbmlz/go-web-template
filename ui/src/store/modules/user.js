import { getUserInfo, login } from '@/api/methods/auth'
import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: '',
    userInfo: {},
  }),
  getters: {
    isLoggedIn: state => !!state.token,
  },
  actions: {
    setToken(token) {
      this.token = token
      localStorage.setItem('token', token)
    },
    clearToken() {
      this.token = ''
      localStorage.removeItem('token')
    },
    async login(data) {
      const { token } = await login(data)
      this.setToken(token)
    },
    async getUserInfo() {
      const data = await getUserInfo()
      this.userInfo = data
    },
  },

  persist: true,
})
