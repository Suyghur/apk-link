import request from '@/utils/request'

export function getOptions(param) {
  return request({
    url: '/options/getOptions',
    method: 'get',
    params: param
  })
}

export function getOrigins(param) {
  return request({
    url: '/origin/getOrigins',
    method: 'get',
    params: param
  })
}

export function getFuseSdks() {
  return request({
    url: '/sdk/fuse/getFuseSdks',
    method: 'get'
  })
}

export function getChannelSdks(param) {
  return request({
    url: '/sdk/channel/getChannelSdks',
    method: 'get',
    params: param
  })
}

export function getPluginSdks(param) {
  return request({
    url: '/sdk/plugin/getPluginSdks',
    method: 'post',
    params: param
  })
}

export function getGameKeystores(param) {
  return request({
    url: '/keystore/getKeystores',
    method: 'get',
    params: param
  })
}

export function getAids(data) {
  console.log(JSON.stringify(data))
  return request({
    url: '/options/getAids',
    method: 'post',
    data
  })
}
