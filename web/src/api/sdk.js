import request from '@/utils/request'

export function getFuseSdkList(params) {
  return request({
    url: '/get_sdk_list/fuse',
    method: 'get',
    params
  })
}

export function getChannelSdkList(params) {
  return request({
    url: '/get_sdk_list/channel',
    method: 'get',
    params
  })
}

export function getPluginSdkList(params) {
  return request({
    url: '/get_sdk_list/plugin',
    method: 'get',
    params
  })
}

