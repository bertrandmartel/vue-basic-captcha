import Vue from 'vue'
import Vuetify from 'vuetify/lib'
import colors from 'vuetify/lib/util/colors'
import '@mdi/font/css/materialdesignicons.css'

Vue.use(Vuetify);
Vue.directive('blur', {
  inserted: function (el) {
    el.onfocus = (ev) => ev.target.blur()
  }
});
export default new Vuetify({
  dark:true,
  icons: {
    iconfont: 'md',
  },
  theme: {
    themes: {
      light: {
        primary: colors.blue.lighten2,
        secondary: colors.grey.darken1,
        accent: colors.shades.black,
        error: colors.red.accent3,
      },
      dark: {
        primary: colors.blue.darken3,
      },
    },
  },
  breakpoint: {
    thresholds: {
      xs: 450,
      sm: 540,
      md: 800,
      lg: 1280,
    },
    scrollBarWidth: 24,
  },
});
