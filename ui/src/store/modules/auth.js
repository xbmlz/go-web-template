import { getUserInfo, login } from '@/api/methods/auth'
import { defineStore } from 'pinia'
import { useRoute, useRouter } from 'vue-router'

export const useAuthStore = defineStore('user', {
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
    async logout() {
      const router = useRouter()
      const route = useRoute()
      if (route.meta.requiresAuth) {
        router.push({
          path: '/login',
          query: { redirect: route.fullPath },
        })
      }
      this.clearToken()
    },
    async getUserInfo() {
      const data = await getUserInfo()
      this.userInfo = data
    },
  },

  persist: true,
})
