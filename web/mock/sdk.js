const Mock = require('mockjs')
const { channel, plugin, developer } = require('./utils')

const fuseVersion = ['1.0.0', '1.1.1']

function getChannelVersionList(channelName) {
  let list = []
  switch (channelName) {
    case '后浪（houalng）':
      list = ['1.0.0', '1.1.1']
      break
    case '应用宝（yyb）':
      list = ['1.5.15', '1.6.3']
      break
    case 'UC九游（uc）':
      list = ['8.0.4', '9.2.0.1']
      break
    case '小米（mi）':
      list = ['3.2.1A', '3.2.1C']
      break
    case '魅族（meizu）':
      list = ['5.2.7.P', '5.4.2_P']
      break
    case '华为（huawei）':
      list = ['2.6.3.306', '4.0.3.302']
      break
    case 'Oppo（oppo）':
      list = ['2.0.2.306', '2.0.4.009']
      break
    case 'Vivo（vivo）':
      list = ['4.4.2.0', '4.5.0.0']
      break
    case '4399（4399）':
      list = ['2.30.0.33', '2.32.1.1']
      break
    case 'B站（bilibili）':
      list = ['2.6.0', '2.7.0']
      break
    default:
      break
  }
  return list
}

function getPluginVersionList(pluginName) {
  let list = []
  switch (pluginName) {
    case '广点通':
      list = ['1.5.0']
      break
    case '头条':
      list = ['1.0.0', '3.1.12']
      break
    case '快手':
      list = ['1.0.0']
      break
    case '百度cps':
      list = ['1.0.0']
      break
    case 'UC信息流':
      list = ['1.0']
      break
  }
  return list
}

module.exports = [
  {
    url: '/get_sdk_list/fuse',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          fuse_sdk: Mock.mock({
            'items|3': [{
              sdk_version: /[1-2]\.[0-9]\.[0-9]/,
              file_name: 'houlang_fuse_sdk_@sdk_version',
              modify_time: '@datetime',
              'modify_person|1': developer
            }]
          }).items
        }
      }
    }
  },
  {
    url: '/get_sdk_list/channel',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          channel_sdk: Mock.mock({
            'items|10': [{
              'channel|1': channel,
              sdk_version: /[1-2]\.[0-9]\.[0-9]/,
              file_name: '@channel' + ' _' + '@sdk_version',
              modify_time: '@datetime',
              'modify_person|1': developer
            }]
          }).items
        }
      }
    }
  },
  {
    url: '/get_sdk_list/plugin',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          plugin_sdk: Mock.mock({
            'items|10': [{
              'plugin|1': plugin,
              sdk_version: /[1-2]\.[0-9]\.[0-9]/,
              file_name: '@plugin' + ' _' + '@sdk_version',
              modify_time: '@datetime',
              'modify_person|1': developer
            }]
          }).items
        }
      }
    }
  },
  {
    url: '/get_sdk_version_list/fuse',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          version_list: fuseVersion
        }
      }
    }
  },
  {
    url: '/get_sdk_version_list/channel',
    type: 'get',
    response: config => {
      const { channel_name } = config.query
      return {
        code: 200,
        msg: 'success',
        data: {
          version_list: getChannelVersionList(channel_name)
        }
      }
    }
  },
  {
    url: '/get_sdk_version_list/plugin',
    type: 'get',
    response: config => {
      const { plugin_name } = config.query
      return {
        code: 200,
        msg: 'success',
        data: {
          version_list: getPluginVersionList(plugin_name)
        }
      }
    }
  }
]

