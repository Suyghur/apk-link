import request from '@/utils/request'

export function searchOrigin(data) {
  return request({
    url: '/origin/searchOrigin',
    method: 'post',
    data
  })
}
