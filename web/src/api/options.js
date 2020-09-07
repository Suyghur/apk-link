import request from '@/utils/request'

export function getOptions(data) {
  return request({
    url: '/options/getOptions',
    method: 'post',
    data
  })
}

export function getOriginBags(data) {
  return request({
    url: '/origin/getOrigins',
    method: 'post',
    data
  })
}

export function getFuseSdks() {
  return request({
    url: '/sdk/fuse/getFuseSdks',
    method: 'post'
  })
}

export function getChannelSdks(data) {
  return request({
    url: '/sdk/channel/getChannelSdks',
    method: 'post',
    data
  })
}

export function getPluginSdks(data) {
  return request({
    url: '/sdk/plugin/getPluginSdks',
    method: 'post',
    data
  })
}

export function getGameKeystores(data) {
  return request({
    url: '/keystore/getKeystores',
    method: 'post',
    data
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
