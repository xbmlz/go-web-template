import { request } from '..'

export function login(data) {
  return request.Post('/auth/login', data)
}
