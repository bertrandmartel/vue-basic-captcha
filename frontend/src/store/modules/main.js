import configuration from '../../api/configuration'

const state = {
  message: {},
}

const mutations = {
  updateMessage(state, value){
    state.message = value
  },
}

const actions = {
  async init({rootState, commit, dispatch}, noFirstSelect){
    await dispatch("core/getUserInfo", null, {root:true})
    if (!rootState.core.authenticated){
      document.location.href = "/login"
    }
    configuration.getConfiguration({
      serverPath: rootState.core.serverPath
    },
    (response) => {
      commit("updateMessage", response.data)
    },
    (err) => {
      console.log(err);
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