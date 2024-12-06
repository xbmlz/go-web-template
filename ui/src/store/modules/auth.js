import { getPermissions, getUserInfo, login } from '@/api/methods/auth'
import router from '@/router'
import { defineStore } from 'pinia'
import { unref } from 'vue'

export const useAuthStore = defineStore('user', {
  state: () => ({
    token: '',
    userInfo: {},
    permissions: {},
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
      const route = unref(router.currentRoute)
      if (route.meta && route.meta.requiresAuth) {
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
    async getPermissions() {
      const data = await getPermissions()
      this.permissions = data
      console.log(data)
      const { menus } = data
      // dynamic generate routes
      const routes = []
    },
  },

  persist: true,
})
