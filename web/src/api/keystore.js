import request from '@/utils/request'

export function searchKeystore(data) {
  return request({
    url: '/keystore/searchKeystore',
    method: 'post',
    data
  })
}
