import request from '@/utils/request'

export function searchFuseSdk(param) {
  return request({
    url: '/sdk/fuse/searchFuseSdk',
    method: 'get',
    params: param
  })
}

export function searchChannelSdk(param) {
  return request({
    url: '/sdk/channel/searchChannelSdk',
    method: 'get',
    params: param
  })
}

export function searchPluginSdk(param) {
  return request({
    url: '/sdk/plugin/searchPluginSdk',
    method: 'get',
    params: param
  })
}

