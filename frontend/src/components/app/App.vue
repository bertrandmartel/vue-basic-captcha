<template>
  <v-app id="inspire">
    <AppBar v-if="$vuetify.breakpoint.height >= 500" />
    <!--<router-view></router-view>-->
    <v-content class="content">
      <LandscapeAppBar v-if="$vuetify.breakpoint.height < 500" />
      <Root v-if="this.$route.path === '/'"/>
      <Login v-if="this.$route.path === '/login'"/>
    </v-content>
    <Snackbar />
    <SettingsDialog />
    <AboutDialog />
  </v-app>
</template>

<script>
import { mapState } from 'vuex'

import AppBar from './AppBar'
import Snackbar from './Snackbar'
import Login from '../login/Login.vue'
import Root from './Root.vue'
import LandscapeAppBar from './LandscapeAppBar.vue'
import SettingsDialog from './SettingsDialog'
import AboutDialog from './AboutDialog'

export default {
  name: 'App',

  components: {
    AppBar,
    Snackbar,
    Login,
    Root,
    LandscapeAppBar,
    SettingsDialog,
    AboutDialog
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
  background-color: #121212;
  position: fixed;
  height: 100%;
  overflow: hidden;
}

body {
  width: 100vw;
  height: 100vh;
  overflow-y: auto;
  overflow-x: hidden;
  -webkit-overflow-scrolling: touch;
}

@media screen and (min-height: 500px) {
  .content {
    padding-top: 0px !important;
  }
}
.content {
overflow-x: auto;
}
</style>