import configuration from '../../api/configuration'

const state = {
  message: {},
  loaded: false
}

const mutations = {
  updateMessage(state, value){
    state.message = value
  },
  updateLoaded(state, value){
    state.loaded = value
  }
}

const actions = {
  async init({rootState, commit, dispatch}, noFirstSelect){
    commit("updateLoaded", false)
    await dispatch("core/getUserInfo", null, {root:true})
    if (!rootState.core.authenticated){
      document.location.href = "/login"
    }
    configuration.getConfiguration({
      serverPath: rootState.core.serverPath
    },
    (response) => {
      commit("updateMessage", response.data)
      commit("updateLoaded", true)
    },
    (err) => {
      console.log(err);
      commit("updateLoaded", true)
      if (err.response) {
        return dispatch('core/dispatchMessage', {color: 'error', message: "API error"}, {root:true})
      }
    })
  }
}

export default {
    namespaced: true,
    state,
    mutations,
    actions
}