import request from '@/utils/request'

export function getOriginBagList(params) {
  return request({
    url: '/get_origin_bag_list',
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

export function getGameGroupInfo(params) {
  return request({
    url: '/get_game_group_info',
    method: 'get',
    params
  })
}

// export function getGameGroupInfo(params) {
//   return request({
//     url: '/get_game_group_info',
//     method: 'get',
//     params
//   })
// }
