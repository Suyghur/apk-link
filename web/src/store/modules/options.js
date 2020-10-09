import { getOptions } from '@/api/options'
// import {getGameGroupInfo} from "@/api/game";

const getDefaultSelectOptions = () => {
  return {
    options: {
      gameGroupOptions: [],
      channelNameOptions: [],
      pluginNameOptions: [],
      taskStatusOptions: ['未执行', '执行中', '成功', '失败']
    }
  }
}

const state = getDefaultSelectOptions()

const mutations = {
  SET_OPTIONS: (state, options) => {
    state.options.gameGroupOptions = options.game_options
    state.options.channelNameOptions = options.channel_options
    state.options.pluginNameOptions = options.plugin_options
    state.options.taskStatusOptions = ['未执行', '执行中', '成功', '失败']
  }
}
const actions = {
  fetchOptions({ commit }) {
    return new Promise((resolve, reject) => {
      getOptions().then(response => {
        if (!response) {
          return reject('数据加载异常')
        }
        commit('SET_OPTIONS', response.data.options)
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  }
}
export default {
  namespaced: true,
  state,
  mutations,
  actions
}
