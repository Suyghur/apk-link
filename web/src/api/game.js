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

export function getGids(param) {
  return request({
    url: '/gid/getGids',
    method: 'get',
    params: param
  })
}

export function getAidList(params) {
  return request({
    url: '/get_aid_list',
    method: 'get',
    params
  })
}
