import request from '@/utils/request'

export function searchScript(param) {
  return request({
    url: '/script/searchScript',
    method: 'get',
    params: param
  })
}
