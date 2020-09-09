import request from '@/utils/request'

export function searchFuseSdk(data) {
  return request({
    url: '/sdk/fuse/searchFuseSdk',
    method: 'post',
    data
  })
}

export function searchChannelSdk(data) {
  return request({
    url: '/sdk/channel/searchChannelSdk',
    method: 'post',
    data
  })
}

export function searchPluginSdk(data) {
  return request({
    url: '/sdk/plugin/searchPluginSdk',
    method: 'post',
    data
  })
}

