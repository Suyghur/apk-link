const Mock = require('mockjs')
const { developer, gameGroup, plugin, channel } = require('./utils')

const cesyx1 = {
  info: [
    {
      file_name: 'csyx1_1.0.0.apk',
      game_version: '1.0.0',
      fuse_version: '1.0.0',
      modify_time: '2020-08-21 09:59:26',
      url: 'https://apkfile.hihoulang.com/csyx1/csyx1_1.0.0.apk'
    }, {
      file_name: 'syx1_1.1.1.apk',
      game_version: '1.1.1',
      fuse_version: '1.0.0',
      modify_time: '2020-08-25 09:59:26',
      url: 'https://apkfile.hihoulang.com/csyx1/syx1_1.1.1.apk'
    }],
  bag: ['csyx1_1.0.0.apk', 'csyx1_1.1.1.apk'],
  keystore: ['csyx1aaa.keystore', 'csyx1bbb.keystore'],
  aid: ['9000000001', '9000000002', '9000000013', '9000000014', '9000000015', '9000000016', '9000000017', '9000000018', '9000000019']

}

const cesyx2 = {
  info: [
    {
      file_name: 'csyx2_1.0.1.apk',
      game_version: '2.0.1',
      fuse_version: '1.0.1',
      modify_time: '2020-08-21 09:59:26',
      url: 'https://apkfile.hihoulang.com/csyx2/csyx2_2.0.1.apk'
    },
    {
      file_name: 'csyx2_2.0.1.apk',
      game_version: '2.0.0',
      fuse_version: '1.0.0',
      modify_time: '2020-08-25 09:59:26',
      url: 'https://apkfile.hihoulang.com/csyx2/csyx2_2.0.0.apk'
    }
  ],
  bag: ['csyx2_1.0.1.apk', 'csyx2_2.0.1.apk'],
  keystore: ['csyx2aaa.keystore', 'csyx2bbb.keystore'],
  aid: ['9000000021', '9000000022', '9000000023', '9000000024', '9000000025', '9000000026', '9000000027', '9000000028', '9000000029', '9000000030']

}

const cesyx3 = {
  info: [
    {
      file_name: 'csyx3_4.0.0.apk',
      game_version: '4.0.0',
      fuse_version: '1.0.1',
      modify_time: '2020-08-21 09:59:26',
      url: 'https://apkfile.hihoulang.com/csyx3/csyx3_4.0.0.apk'
    }
  ],
  bag: ['csyx3_4.0.0.apk'],
  keystore: ['csyx3aaa.keystore'],
  aid: ['9000000025', '9000000026', '9000000027', '9000000028', '9000000029', '9000000030']

}

const cesyx4 = {
  info: [
    {
      file_name: 'csyx4_3.0.0.apk',
      game_version: '3.0.0',
      fuse_version: '1.0.0',
      modify_time: '2020-08-21 09:59:26',
      url: 'https://apkfile.hihoulang.com/csyx3/csyx4_3.0.0.apk'
    }
  ],
  bag: ['csyx4_3.0.0.apk'],
  keystore: ['csyx4aaa.keystore'],
  aid: ['9000000041', '9000000042', '9000000044', '9000000048', '9000000050', '9000000051']

}

function getOriginBagList(game) {
  let bag = []
  switch (game) {
    case '测试游戏1':
      bag = cesyx1.bag
      break
    case '测试游戏2':
      bag = cesyx2.bag
      break
    case '测试游戏3':
      bag = cesyx3.bag
      break
    case '测试游戏4':
      bag = cesyx4.bag
      break
    default:
      break
  }
  return bag
}

function getGameGroupInfo(game) {
  let info = []
  switch (game) {
    case '测试游戏1':
      info = cesyx1.info
      break
    case '测试游戏2':
      info = cesyx2.info
      break
    case '测试游戏3':
      info = cesyx3.info
      break
    case '测试游戏4':
      info = cesyx4.info
      break
    default:
      break
  }
  return info
}

function getGameKeystoreList(game) {
  let list = []
  switch (game) {
    case '测试游戏1':
      list = cesyx1.keystore
      break
    case '测试游戏2':
      list = cesyx2.keystore
      break
    case '测试游戏3':
      list = cesyx3.keystore
      break
    case '测试游戏4':
      list = cesyx4.keystore
      break
    default:
      break
  }
  return list
}

function getAidList(game) {
  let list = []
  switch (game) {
    case '测试游戏1':
      list = cesyx1.aid
      break
    case '测试游戏2':
      list = cesyx2.aid
      break
    case '测试游戏3':
      list = cesyx3.aid
      break
    case '测试游戏4':
      list = cesyx4.aid
      break
    default:
      break
  }
  return list
}

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
              'game_group|1': gameGroup,
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
    url: '/get_game_group_info',
    type: 'get',
    response: config => {
      const { game_group } = config.query
      return {
        code: 200,
        msg: 'success',
        data: {
          infos: getGameGroupInfo(game_group),
          gid: gameGroup.indexOf(game_group)
        }
      }
    }
  },
  {
    url: '/get_origin_bag_list',
    type: 'get',
    response: config => {
      const { game_group } = config.query
      return {
        code: 200,
        msg: 'success',
        data: {
          bag: getOriginBagList(game_group)
        }
      }
    }
  },
  {
    url: '/get_game_keystore',
    type: 'get',
    response: config => {
      const { game_group } = config.query
      return {
        code: 200,
        msg: 'success',
        data: {
          keystore: getGameKeystoreList(game_group)
        }
      }
    }
  },
  {
    url: '/get_aid_list',
    type: 'get',
    response: config => {
      const { game_group } = config.query
      return {
        code: 200,
        msg: 'success',
        data: {
          aids: getAidList(game_group)
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
              'game_group|1': gameGroup,
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
              'game_group|1': gameGroup,
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
