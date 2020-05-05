import Vuetify from '../../plugins/vuetify'
import login from '../../api/login'

const state = {
  authenticated: false,
  theme: "dark",
  title: "Vue Template App",
  aboutLink: "https://github.com/bertrandmartel/vue-basic-captcha",
  settingsDialog: false,
  aboutDialog: false,
  serverPath: "",
  serverPathEdit: "",
  snackbar: false,
  snackbarColor: 'error',
  snackbarMessage: ''
}

const mutations = {
  toggleDark(state){
    state.theme = Vuetify.framework.theme.dark ? "light" : "dark"
    Vuetify.framework.theme.dark = !Vuetify.framework.theme.dark
    localStorage.setItem("theme", state.theme)
  },
  updateSettingsDialog(state, value){
    state.settingsDialog = value
  },
  updateAboutDialog(state, value){
    state.aboutDialog = value
  },
  updateServerPath(state, value){
    state.serverPath = value
  },
  updateTheme(state, value){
    state.theme = value
  },
  updateServerPathEdit(state, value){
    state.serverPathEdit = value
  },
  updateSnackbar(state, value){
    state.snackbar = value
  },
  updateSnackbarColor(state, value){
    state.snackbarColor = value
  },
  updateSnackbarMessage(state, value){
    state.snackbarMessage = value
  },
  setAuthenticated(state, value){
    state.authenticated = value
  }
}

const actions = {
  init({ commit, state, dispatch}){
    commit('updateServerPath', localStorage.getItem('serverPath'))
    if (!state.serverPath){
      var path = `${location.protocol}//${location.hostname+(location.port ? ':' + location.port : '')}`;
      commit('updateServerPath', path)
      localStorage.setItem("serverPath", state.serverPath)
    }
    commit('updateTheme', localStorage.getItem('theme') || 'dark')
    Vuetify.framework.theme.dark = state.theme === 'dark'
  },
  openSettingsDialog({commit, state}){
    commit("updateServerPathEdit", state.serverPath)
    commit("updateSettingsDialog", true)
  },
  saveSettings({commit, state}){
    if (state.serverPathEdit != state.serverPath){
      commit("updateServerPath", state.serverPathEdit.replace(/\/$/, ""))
      localStorage.setItem("serverPath", state.serverPath)
      document.location.href = window.location.href
    }
    commit("updateSettingsDialog", false)
  },
  dispatchMessage({commit, state, dispatch}, data){
    commit('updateSnackbarColor', data.color)
    commit('updateSnackbarMessage', data.message)
    commit('updateSnackbar', true)
  },
  logout(){
    document.location.href = "/logout"
  },
  getUserInfo({commit, rootState}){    
    return new Promise(function(resolve, reject){
      login.userinfo({
        serverPath: rootState.core.serverPath
      },
      (response) => {
        if (response.data.authenticated){
          commit("setAuthenticated", true)
        } else {
          commit("setAuthenticated", false)
        }
        resolve(response)
      },
      (err) => {
        console.log(err);
        commit("setAuthenticated", false)
        reject(err)
      })
    });
  }
}

export default {
  namespaced: true,
  state,
  actions,
  mutations
}