<template>
  <div>
    <v-app-bar app clipped-left>
      <router-link class="titleLink" to="/"><v-toolbar-title class="title">{{title}}</v-toolbar-title></router-link>
      <v-spacer></v-spacer>
      <v-btn v-show="authenticated" icon class="appbarBtn" @click="logout">
        <v-icon>mdi-logout</v-icon>
      </v-btn>
      <v-btn icon @click="toggleDark" class="appbarBtn">
        <v-icon v-if="theme=='dark'" medium>mdi-brightness-4</v-icon>
        <v-icon v-else medium>mdi-brightness-7</v-icon>
      </v-btn>
      <v-btn icon class="appbarBtn" @click="openSettingsDialog">
        <v-icon>mdi-cog-outline</v-icon>
      </v-btn>
      <v-btn icon class="appbarBtn" @click="aboutDialog = true">
        <v-icon>mdi-information-outline</v-icon>
      </v-btn>
    </v-app-bar>
    <SettingsDialog />
    <AboutDialog />
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import SettingsDialog from './SettingsDialog'
  import AboutDialog from './AboutDialog'

  export default {
    name: 'AppBar',
    components: {
      SettingsDialog,
      AboutDialog
    },
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
.appbarBtn {
  margin-right:15px !important;
}
</style>