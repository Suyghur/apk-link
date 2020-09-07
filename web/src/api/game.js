import request from '@/utils/request'

// export function getOriginBagList(params) {
//   return request({
//     url: '/get_origin_bag_list',
//     method: 'get',
//     params
//   })
// }

export function getLinkingBag(params) {
  return request({
    url: '/get_linking_list',
    method: 'get',
    params
  })
}

export function getKeystoreList(params) {
  return request({
    url: '/get_game_keystore',
    method: 'get',
    params
  })
}

export function getGid(data) {
  return request({
    url: '/game/getGid',
    method: 'post',
    data
  })
}

// export function getGameKeystore(data) {
//   return request({
//     url: '/keystore/getKeystores',
//     method: 'post',
//     params
//   })
// }

export function getAidList(params) {
  return request({
    url: '/get_aid_list',
    method: 'get',
    params
  })
}
