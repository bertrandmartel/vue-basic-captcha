<template>
  <v-dialog v-model="aboutDialog" max-width="600">
    <v-card>
      <v-card-title class="headline">About</v-card-title>
      <v-card-text class="aboutCard">
        <v-row no-gutters style='margin-top:25px;' class="aboutRow">
          <v-col cols="12" class='aboutCol'>
            <div class='aboutIcon'><v-icon large>mdi-information</v-icon></div>
            <div class='aboutInfo'>
              <div class='aboutInfoContainer'>
                <div class='aboutInfoTitle'><h3>Info</h3></div>
                <div class='aboutInfoDescription'>AWS Admin v1.0</div>
              </div>
            </div>
          </v-col>
          <v-col cols="12" class='aboutCol'>
            <div class='aboutIcon'><v-icon large>mdi-copyright</v-icon></div>
            <div class='aboutInfo'>
              <div class='aboutInfoContainer'>
                <div class='aboutInfoTitle'><h3>Copyright</h3></div>
                <div class='aboutInfoDescription'>The MIT License (MIT) Copyright (c) 2020 Bertrand Martel</div>
              </div>
            </div>
          </v-col>
          <v-col cols="12" class='aboutCol'>
            <div class='aboutIcon'><v-icon large>mdi-code-tags</v-icon></div>
            <div class='aboutInfo'>
              <div class='aboutInfoContainer'>
                <div class='aboutInfoTitle'><h3>Source code</h3></div>
                <div class='aboutInfoDescription aboutInfoLink'><a class="urlLink" target="_blank" :href="aboutLink">{{aboutLink}}</a><a class="urlTiny" target="_blank" :href="aboutLink">github.com</a></div>
              </div>
            </div>
          </v-col>
          <v-col cols="12" class='aboutCol'>
            <div class='aboutIcon'><v-icon large>mdi-apps</v-icon></div>
            <div class='aboutInfo' style='display: flex; align-items: center;'>
              <div class='aboutInfoTitle'><h3>Technologies</h3></div>
            </div>
          </v-col>
          <v-col cols="12" class='aboutCol'>
            <div class='aboutIcon'></div>
            <div class='aboutInfo'>
              <div style='display:flex;'>
                <div style='flex:30%; text-align:center;'><img class="technoImg" src="@/assets/vuejs.png"/></div>
                <div style='flex:70%; text-align:center;'><img class="technoImg" src="@/assets/go-blue.png"/></div>
              </div>
            </div>
          </v-col>
          <v-col cols="12" style='display:flex;margin-top:10px'>
            <div class='aboutIcon'></div>
            <div class='aboutInfo'>
              <div style='display:flex;'>
                <div style='flex: 30%;text-align: center'><h2>Vue.js</h2></div>
                <div style='flex: 70%;text-align: center'><h2>Golang</h2></div>
              </div>
            </div>
          </v-col>
        </v-row>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="green darken-1" text @click="aboutDialog = false">Close</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'AboutDialog',
  computed: {
    ...mapState({
      aboutLink: state => state.core.aboutLink,
    }),
    aboutDialog: {
      get() {
        return this.$store.state.core.aboutDialog
      },
      set(value) {
        this.$store.commit('core/updateAboutDialog', value)
      }
    }
  }
}
</script>

<style>
.aboutCol {
  display:flex;
  margin-top:25px
}
.aboutIcon {
  flex: 20%;
  text-align: center
}

@media screen and (min-width: 451px) {
  .urlLink {
    display: block;
  }
  .urlTiny {
    display: none;
  }
}

@media screen and (max-width: 450px) {
  .technoImg {
    display: none;
  }
  .urlLink {
    display: none;
  }
  .urlTiny {
    display: block;
  }
  .aboutInfo {
    flex: 80%;
    text-align: center;
    margin-left: 20px;
  }
}

.aboutInfo {
  flex: 80%;
  text-align: center;
}
.aboutInfoContainer{
  display: flex; 
  flex-direction: column;
}
.aboutInfoTitle {
  text-align:left;
}
.aboutInfoDescription{
  text-align:left;
}
.aboutCard {
  font-size: 1rem !important;
}
.theme--dark.v-card > .aboutCard > .aboutRow > .aboutCol > .aboutInfo > .aboutInfoContainer > .aboutInfoTitle {
  color:white !important;
}
.theme--dark.v-card > .aboutCard > .aboutRow > .aboutCol > .aboutInfo > .aboutInfoContainer > .aboutInfoLink a {
  color:#ffffffb3 !important;
}
.theme--light.v-card > .aboutCard > .aboutRow > .aboutCol > .aboutInfo > .aboutInfoContainer > .aboutInfoLink a {
  color:#00000099 !important;
}
</style>