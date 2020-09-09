import request from '@/utils/request'

export function searchChannel(data) {
  return request({
    url: '/channel/searchChannel',
    method: 'post',
    data
  })
}
