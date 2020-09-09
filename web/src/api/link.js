import request from '@/utils/request'

export function searchLink(data) {
  return request({
    url: '/link/searchLink',
    method: 'post',
    data
  })
}
