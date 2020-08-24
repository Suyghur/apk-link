import request from '@/utils/request'

export function getFuseScriptList(params) {
  return request({
    url: '/get_script_list/fuse',
    method: 'get',
    params
  })
}

export function getChannelScriptList(params) {
  return request({
    url: '/get_script_list/channel',
    method: 'get',
    params
  })
}

export function getPluginScriptList(params) {
  return request({
    url: '/get_script_list/plugin',
    method: 'get',
    params
  })
}
export function getGameScriptList(params) {
  return request({
    url: '/get_script_list/game',
    method: 'get',
    params
  })
}
