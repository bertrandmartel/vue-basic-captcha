<template>
  <v-navigation-drawer
    fixed
    v-model="drawer"
    :mini-variant.sync="mini"
    permanent
    left
  >
    <v-list dense>
      <v-list-item>
        <v-list-item-icon>
          <v-icon>mdi-menu</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>Menu</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list-item link @click="logout" v-show="authenticated" class="landscapeMenuItem">
        <v-list-item-icon>
          <v-icon>mdi-logout</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>Logout</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-list-item link @click="toggleDark" class="landscapeMenuItem">
        <v-list-item-icon>
          <v-icon>mdi-brightness-4</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>Theme</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-list-item link @click="openSettingsDialog" class="landscapeMenuItem">
        <v-list-item-icon>
          <v-icon>mdi-cog-outline</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>Settings</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-list-item link @click="aboutDialog = true" class="landscapeMenuItem">
        <v-list-item-icon>
          <v-icon>mdi-information-outline</v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title>About</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'LandscapeAppBar',
  computed: {
    ...mapState({
      authenticated: state => state.core.authenticated,
    }),
    drawer: {
      get() {
        return this.$store.state.core.mobileDrawer
      },
      set(value) {
      }
    },
    mini: {
      get() {
        return this.$store.state.core.mobileDrawerMini
      },
      set(value) {
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
.landscapeMenuItem {
  margin-top:10px;
}


.v-navigation-drawer--close {
  visibility: visible !important;
  transform: none !important;
}
</style>