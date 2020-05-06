<template>
  <div>
    <v-app-bar app clipped-left>
      <router-link class="titleLink" to="/"><v-toolbar-title class="title">{{title}}</v-toolbar-title></router-link>
      <v-spacer class="spacer"></v-spacer>
      <v-btn v-show="authenticated" icon class="appbarBtn" @click="logout">
        <v-icon>mdi-logout</v-icon>
      </v-btn>
      <v-btn icon @click="toggleDark" class="appbarBtn">
        <v-icon medium>mdi-brightness-4</v-icon>
      </v-btn>
      <v-btn icon class="appbarBtn" @click="openSettingsDialog">
        <v-icon>mdi-cog-outline</v-icon>
      </v-btn>
      <v-btn icon class="appbarBtn" @click="aboutDialog = true">
        <v-icon>mdi-information-outline</v-icon>
      </v-btn>
    </v-app-bar>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'AppBar',
  computed: {
    ...mapState({
      title: state => state.core.title,
      theme: state => state.core.theme,
      authenticated: state => state.core.authenticated,
    }),
    settingsDialog: {
      get() {
        return this.$store.state.core.settingsDialog
      },
      set(value) {
        this.$store.commit('core/updateSettingsDialog', value)
      }
    },
    aboutDialog: {
      get() {
        return this.$store.state.core.aboutDialog
      },
      set(value) {
        this.$store.commit('core/updateAboutDialog', value)
      }
    },
    drawer: {
      get() {
        return this.$store.state.core.mobileDrawer
      },
      set(value) {
        this.$store.commit('core/updateMobileDrawer', value)
      }
    },
    mini: {
      get() {
        return this.$store.state.core.mobileDrawerMini
      },
      set(value) {
        this.$store.commit('core/updateMobileDrawerMini', value)
      }
    },
  },
  methods: {
    openSettingsDialog(){
      this.$store.dispatch('core/openSettingsDialog')
    },
    toggleDark(){
      this.$store.commit('core/toggleDark')
    },
    logout(){
      this.$store.dispatch('core/logout')
    }
  }
}
</script>

<style>
.titleLink {
  text-decoration: none;
  color: inherit !important;
}
@media screen and (min-width: 451px) and (min-height: 500px){
  .appbarBtn {
    margin-right:15px !important;
  }
}
/*
@media screen and (max-height: 500px) {
  .v-toolbar.v-toolbar--collapsed  {
    max-width: 125px !important;
  }
  .spacer {
    display: none;
  }
}
*/
</style>