export function getChannelName(channelTag) {
  let name = ''
  console.log(channelTag)
  switch (channelTag) {
    case 'houlang':
      name = 'houlang（' + channelTag + ')'
      break
    case 'yyb':
      name = '应用宝（' + channelTag + ')'
      break
    case 'uc':
      name = 'UC九游（' + channelTag + ')'
      break
    case 'mi':
      name = '小米（' + channelTag + ')'
      break
    case 'meizu':
      name = '魅族（' + channelTag + ')'
      break
    case 'huawei':
      name = '华为（' + channelTag + ')'
      break
    case 'oppo':
      name = 'Oppo（' + channelTag + ')'
      break
    case 'vivo':
      name = 'Vivo（' + channelTag + ')'
      break
    default:
      break
  }
  return name
}
