import request from '@/utils/request'

// export function getOriginBagList(params) {
//   return request({
//     url: '/get_origin_bag_list',
//     method: 'get',
//     params
//   })
// }

export function searchGame(param) {
  return request({
    url: '/game/searchGame',
    method: 'get',
    params: param
  })
}

export function getGid(param) {
  return request({
    url: '/game/getGid',
    method: 'get',
    params: param
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
