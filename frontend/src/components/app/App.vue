<template>
  <v-app id="inspire">
    <AppBar />
    <!--<router-view></router-view>-->
    <Root v-if="this.$route.path === '/'"/>
    <Login v-if="this.$route.path === '/login'"/>
    <Snackbar />
  </v-app>
</template>

<script>
import { mapState } from 'vuex'

import AppBar from './AppBar'
import Snackbar from './Snackbar'
import Login from '../login/Login.vue'
import Root from './Root.vue'

export default {
  name: 'App',

  components: {
    AppBar,
    Snackbar,
    Login,
    Root
  },
  props: {
    source: String,
  },
  computed: {
    ...mapState({
    })
  },
  created () {
    this.$store.dispatch('core/init');
  }
};
</script>

<style>
html{
  overflow-y: auto !important;
}
.content {
}
</style>