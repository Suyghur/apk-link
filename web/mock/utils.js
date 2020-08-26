const channel = ['后浪（houalng）', '应用宝（yyb）', 'UC九游（uc）', '小米（mi）',
  '魅族（meizu）', '华为（huawei）', 'Oppo（oppo）', 'Vivo（vivo）', '4399（4399）', 'B站（bilibili）']

const plugin = ['广点通', '头条', '快手', '百度cps', 'UC信息流']
const gameGroup = ['测试游戏1', '测试游戏2', '测试游戏3', '测试游戏4']
const status = ['未执行', '执行中', '成功', '失败']

/**
 * @param {string} url
 * @returns {Object}
 */
function param2Obj(url) {
  const search = decodeURIComponent(url.split('?')[1]).replace(/\+/g, ' ')
  if (!search) {
    return {}
  }
  const obj = {}
  const searchArr = search.split('&')
  searchArr.forEach(v => {
    const index = v.indexOf('=')
    if (index !== -1) {
      const name = v.substring(0, index)
      const val = v.substring(index + 1, v.length)
      obj[name] = val
    }
  })
  return obj
}

const developer = ['麦锦培', '技术1', '技术2', '技术3']

module.exports = {
  param2Obj,
  plugin,
  gameGroup,
  channel,
  developer,
  status
}

