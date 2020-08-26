const Mock = require('mockjs')
const { channel, gameGroup, status, plugin } = require('./utils')

const List = []
const count = 100

for (let i = 0; i < count; i++) {
  List.push(Mock.mock({
    modify_time: '@datetime',
    'task_id|1-9999': 1,
    'game_group|1': gameGroup,
    'fuse_version': /[1-2]\.[0-9]\.[0-9]/,
    'channel|1': channel,
    'channel_version': /[1-9]\.[0-9]\.[0-9]/,
    'plugin|1': plugin,
    'aid|9000000000-9999999999': 1,
    'status|1': status
  }))
}
module.exports = [
  {
    url: '/get_task_list',
    type: 'get',
    response: config => {
      const { task_id, game_group, channel_name, plugin_name, task_status, page = 1, limit = 20 } = config.query

      const mockList = List.filter(item => {
        if (task_id && item.task_id.toString() !== task_id) return false
        if (game_group && item.game_group !== game_group) return false
        if (channel_name && item.channel !== channel_name) return false
        if (plugin_name && item.plugin !== plugin_name) return false
        if (task_status && item.status !== task_status) return false
        return true
      })
      const pageList = mockList.filter((item, index) => index < limit * page && index >= limit * (page - 1))
      return {
        code: 200,
        msg: 'success',
        data: {
          total: mockList.length,
          tasks: pageList
        }
      }
    }
  },
  {
    url: '/get_select_options',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          game_group: gameGroup,
          channel_name: channel,
          task_status: status,
          plugin_name: plugin
        }
      }
    }
  }
]

