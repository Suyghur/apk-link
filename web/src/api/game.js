import request from '@/utils/request'

export function getOriginBagList(params) {
  return request({
    url: '/get_origin_list',
    method: 'get',
    params
  })
}

export function getLinkingBag(params) {
  return request({
    url: '/get_linking_list',
    method: 'get',
    params
  })
}

export function getKeystoreList(params) {
  return request({
    url: '/get_keystore_list',
    method: 'get',
    params
  })
}
