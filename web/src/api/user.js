import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/base/login',
    method: 'post',
    data
  })
}

export function getInfo(param) {
  return request({
    url: '/user/getUserInfo',
    method: 'get',
    params: param
  })
}

export function logout() {
  return request({
    url: '/user/logout',
    method: 'post'
  })
}
