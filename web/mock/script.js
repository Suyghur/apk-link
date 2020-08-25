const Mock = require('mockjs')
const { developer, plugin, channel, gameGroup } = require('./utils')

module.exports = [
  {
    url: '/get_script_list/fuse',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          fuse_script: Mock.mock({
            'items|2': [{
              file_name: 'houlang_fuse_script.zip',
              modify_time: '@datetime',
              'modify_person|1': developer
            }]
          }).items
        }
      }
    }
  },
  {
    url: '/get_script_list/channel',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          channel_script: Mock.mock({
            'items|10': [{
              'channel|1': channel,
              file_name: '@channel _script.zip',
              modify_time: '@datetime',
              'modify_person|1': developer
            }]
          }).items
        }
      }
    }
  },
  {
    url: '/get_script_list/plugin',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          plugin_script: Mock.mock({
            'items|4': [{
              'plugin|1': plugin,
              file_name: '@plugin _script.zip',
              modify_time: '@datetime',
              'modify_person|1': developer
            }]
          }).items
        }
      }
    }
  },
  {
    url: '/get_script_list/game',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          game_script: Mock.mock({
            'items|4': [{
              'game_group|1': gameGroup,
              file_name: '@game _script.zip',
              modify_time: '@datetime',
              'modify_person|1': developer
            }]
          }).items
        }
      }
    }
  }
]
