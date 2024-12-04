import { createAlova } from 'alova'
import fetchAdapter from 'alova/fetch'
import VueHook from 'alova/vue'

export const request = createAlova({
  baseURL: 'http://localhost:8080/api',
  requestAdapter: fetchAdapter(),
  statesHook: VueHook,
  timeout: 10000,
  beforeRequest: (method) => {
    const token = localStorage.getItem('token')
    token && (method.config.headers.Authorization = `Bearer ${token}`)
  },
  responded: {
    onSuccess: async (response, _method) => {
      if (response.status >= 400) {
        throw new Error(response.statusText)
      }
      const json = await response.json()

      if (json.code !== 200) {
        throw new Error(json.msg)
      }

      return json.data
    },
    onFailure: (error, _method) => {
      console.error('error', error)
    },
    onComplete: async (_method) => {
    },
  },
})
