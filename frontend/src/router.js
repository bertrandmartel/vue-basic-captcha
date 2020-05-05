import Vue from 'vue'
import Router from 'vue-router'
import App from './components/app/App.vue'
/*
import Login from './components/login/Login.vue'
import Root from './components/app/Root.vue'
*/

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.VUE_APP_BASE_ROUTE || '/',
  routes: [
    { 
        path:"", 
        component: App,
        /*
        children: [
            { path: '/login', component: Login },
            { path: '/app', component: Root }
        ]
        */
    }
  ]
})