const Mock = require('mockjs')
const { developer, game, plugin, channel } = require('./utils')

module.exports = [
  {
    url: '/get_origin_list',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          origin_bag: Mock.mock({
            'items|4': [{
              icon: '',
              'game|1': game,
              game_version: /[1-9]\.[0-9]\.[0-9]/,
              fuse_version: /[1-2]\.[0-9]\.[0-9]/,
              file_name: '@game _xxx_@game_version .apk',
              modify_time: '@datetime',
              'modify_person|1': developer,
              url: 'http://apkfile.hihoulang.com/xxx_xxx.zpk'
            }]
          }).items
        }
      }
    }
  },
  {
    url: '/get_linking_list',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          linking_bag: Mock.mock({
            'items|30': [{
              'task_id|1-9999': 1,
              'game|1': game,
              'game_version': /[1-9]\.[0-9]\.[0-9]/,
              'fuse_version': /[1-2]\.[0-9]\.[0-9]/,
              'channel|1': channel,
              'channel_version': /[1-9]\.[0-9]\.[0-9]/,
              'plugin|1': plugin,
              'aid|9000000000-9999999999': 1,
              url: 'http://apkfile.hihoulang.com/xxx_xxx@aid .zpk',
              modify_time: '@datetime',
              'modify_person|1': developer
            }]
          }).items
        }
      }
    }
  },
  {
    url: '/get_keystore_list',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          keystore: Mock.mock({
            'items|10': [{
              'game|1': game,
              'keystore_name|': '@string("lower", 4, 4)_hl.keystore',
              'keystore_password|1': '@string("lower", 4, 4)_hl2020',
              'keystore_alias': 'alias.@keystore_password',
              'keystore_alias_password': '@keystore_alias',
              url: 'http://apkfile.hihoulang.com/@keystore_name',
              modify_time: '@datetime',
              'modify_person|1': developer
            }]
          }).items
        }
      }
    }
  }
]
