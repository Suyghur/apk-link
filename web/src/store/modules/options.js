import { getToken } from '@/utils/auth'
import { getSelectOptions } from '@/api/task'
// import {getGameGroupInfo} from "@/api/game";

const getDefaultSelectOptions = () => {
  return {
    selectOptions: {
      gameGroupOptions: [],
      channelNamelOptions: [],
      pluginNameOptions: [],
      taskStatusOptions: []
    }
    // gameGroupInfos: {}
  }
}

const state = getDefaultSelectOptions()

const mutations = {
  SET_SELECT_OPTIONS: (state, selectOptions) => {
    state.selectOptions.gameGroupOptions = selectOptions.game_group
    state.selectOptions.channelNamelOptions = selectOptions.channel_name
    state.selectOptions.pluginNameOptions = selectOptions.plugin_name
    state.selectOptions.taskStatusOptions = selectOptions.task_status
  }
  // SET_GAME_GROUP_INFOS: (state, gameGroupInfo) => {
  //   state.gameGroupInfo = gameGroupInfo
  // }
}
const actions = {
  fetchOptions({ commit }) {
    return new Promise((resolve, reject) => {
      getSelectOptions(getToken()).then(response => {
        if (!response) {
          return reject('数据加载异常')
        }
        commit('SET_SELECT_OPTIONS', response.data)
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  }
  // fetchGameGroupInfo({commit}, gameGroup) {
  //   return new Promise((resolve, reject) => {
  //     getGameGroupInfo({game_group: gameGroup}).then(response => {
  //       if ()
  //     }).catch(error => {
  //       reject(error)
  //     })
  //   })
  // }
}
export default {
  namespaced: true,
  state,
  mutations,
  actions
}
