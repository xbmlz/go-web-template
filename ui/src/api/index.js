import { createAlova } from 'alova'
import { createServerTokenAuthentication } from 'alova/client'
import fetchAdapter from 'alova/fetch'
import VueHook from 'alova/vue'

const { onAuthRequired, onResponseRefreshToken } = createServerTokenAuthentication({
  assignToken: (method) => {
    method.config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`
  },
})

export const request = createAlova({
  baseURL: 'http://localhost:8080/api',
  requestAdapter: fetchAdapter(),
  statesHook: VueHook,
  timeout: 10000,
  beforeRequest: onAuthRequired((method) => {
    if (method.meta?.isFormPost) {
      method.config.headers['Content-Type'] = 'application/x-www-form-urlencoded'
      method.data = new URLSearchParams(method.data).toString()
    }
  }),
  responded: onResponseRefreshToken({
    onSuccess: async (response, method) => {
      const { status } = response
      if (status >= 400) {
        throw new Error(response.statusText)
      }
      // 返回blob数据
      if (method.meta?.isBlob)
        return response.blob()
        // 返回json数据
      const json = await response.json()
      if (json.code !== 200) {
        throw new Error(json.msg)
      }
      return json.data
    },
    onError: (error, method) => {
      const tip = `[${method.type}] - [${method.url}] - ${error.message}`
      console.error(tip)
    },
    onComplete: async (_method) => {},
  }),
})
