import request from '@/utils/request'

export function searchLink(param) {
  return request({
    url: '/link/searchLink',
    method: 'get',
    params: param
  })
}
