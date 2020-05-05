import Vue from 'vue'
import App from './components/app/App.vue'
import './plugins/vuetify'
import router from './router'
import vuetify from '@/plugins/vuetify'
import store from './store'
import axios from 'axios'

Vue.config.productionTip = false

axios.interceptors.request.use(function (config) {
    config.headers["Cache-Control"] = 'no-cache';
    return config;
});

new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
