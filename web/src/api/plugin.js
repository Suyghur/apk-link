import request from '@/utils/request'

export function searchPlugin(param) {
  return request({
    url: '/plugin/searchPlugin',
    method: 'get',
    params: param
  })
}
