const Mock = require('mockjs')
const { channel, plugin, developer } = require('./utils')

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
  }
]

