import Vue from 'vue'
import Vuex from 'vuex'
import createLogger from 'vuex/dist/logger'

import core from './modules/core'
import login from './modules/login'
import main from './modules/main'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
  modules: {
    core,
    login,
    main
  },
  /*
  strict: debug,
  plugins: debug ? [createLogger()] : []
  */
})