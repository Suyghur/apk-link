import request from '@/utils/request'

export function searchKeystore(param) {
  return request({
    url: '/keystore/searchKeystore',
    method: 'get',
    params: param
  })
}
