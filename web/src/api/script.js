import request from '@/utils/request'

export function searchScript(data) {
  return request({
    url: '/script/searchScript',
    method: 'post',
    data
  })
}
