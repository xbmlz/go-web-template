import { request } from '..'

export function login(data) {
  return request.Post('/auth/login', data)
}

export function logout() {
  return request.Post('/auth/logout')
}

export function getUserInfo() {
  return request.Get('/auth/user')
}

export function getPermissions() {
  return request.Get('/auth/permissions')
}
