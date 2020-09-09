import request from '@/utils/request'

export function searchPlugin(data) {
  return request({
    url: '/plugin/searchPlugin',
    method: 'post',
    data
  })
}
