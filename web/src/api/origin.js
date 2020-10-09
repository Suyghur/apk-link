import request from '@/utils/request'

export function searchOrigin(param) {
  return request({
    url: '/origin/searchOrigin',
    method: 'get',
    params: param
  })
}
