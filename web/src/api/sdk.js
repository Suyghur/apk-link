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

export function getFuseVersionList(params) {
  return request({
    url: '/get_sdk_version_list/fuse',
    method: 'get',
    params
  })
}

export function getChannelVersionList(params) {
  return request({
    url: '/get_sdk_version_list/channel',
    method: 'get',
    params
  })
}

export function getPluginVersionList(params) {
  return request({
    url: '/get_sdk_version_list/plugin',
    method: 'get',
    params
  })
}
