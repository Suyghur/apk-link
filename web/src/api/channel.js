import request from '@/utils/request'

export function searchChannel(param) {
  return request({
    url: '/channel/searchChannel',
    method: 'get',
    params: param
  })
}
