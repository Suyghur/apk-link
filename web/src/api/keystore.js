import request from '@/utils/request'

export function searchKeystore(params) {
  return request({
    url: '/keystore/searchKeystore',
    method: 'get',
    params
  })
}
