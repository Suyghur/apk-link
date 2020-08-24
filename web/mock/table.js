const Mock = require('mockjs')

// Mock.setup({
//   timeout: '10-2500'
// })

const data = Mock.mock({
  'items|30': [{
    id: '@id',
    title: '@sentence(10, 20)',
    'status|1': ['published', 'draft', 'deleted'],
    author: 'name',
    display_time: '@datetime',
    pageviews: '@integer(300, 5000)'
  }]
})

const tasks = Mock.mock({
  'items|30': [{
    create_time: '@datetime',
    task_id: '@natural(1, 800)',
    'game_name|1': ['测试游戏1', '测试游戏2', '测试游戏3', '测试游戏4'],
    aid: '@integer(0, 10000)',
    'channel_name|1': ['后浪', '应用宝', 'Oppo', 'Vivo'],
    'media_name|1': ['广点通', '头条', '快手', '百度cps'],
    'status|1': ['未执行', '执行中', '已执行', '执行失败']
  }]
})
module.exports = [
  {
    url: '/table/list',
    type: 'get',
    response: config => {
      const items = data.items
      return {
        code: 200,
        msg: 'success',
        data: {
          total: items.length,
          items: items
        }
      }
    }
  },
  {
    url: '/table/tasks',
    type: 'get',
    response: config => {
      return {
        code: 200,
        msg: 'success',
        data: {
          tasks: tasks.items
        }
      }
    }
  }
]

