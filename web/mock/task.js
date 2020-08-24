const Mock = require('mockjs')
const { channel, game, status, plugin } = require('./utils')

module.exports = [
  {
    url: '/get_task_list',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          tasks: Mock.mock({
            'items|30': [{
              create_time: '@datetime',
              'task_id|1-9999': 1,
              'game|1': game,
              'fuse_version': /[1-2]\.[0-9]\.[0-9]/,
              'channel|1': channel,
              'channel_version': /[1-9]\.[0-9]\.[0-9]/,
              'plugin|1': plugin,
              'aid|9000000000-9999999999': 1,
              'status|1': status
            }]
          }).items
        }
      }
    }
  }
]

